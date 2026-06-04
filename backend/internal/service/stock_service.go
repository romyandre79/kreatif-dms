package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type StockService struct {
	repo repository.Querier
}

func NewStockService(repo repository.Querier) *StockService {
	return &StockService{repo: repo}
}

type StockStats struct {
	PendingMissions int32   `json:"pending_missions"`
	ActiveMissions  int32   `json:"active_missions"`
	AccuracyRate    float64 `json:"accuracy_rate"`
}

func (s *StockService) GetStats(ctx context.Context) (StockStats, error) {
	stats, err := s.repo.GetStockMissionStats(ctx)
	if err != nil {
		return StockStats{}, err
	}
	return StockStats{
		PendingMissions: stats.PendingMissions,
		ActiveMissions:  stats.ActiveMissions,
		AccuracyRate:    stats.AccuracyRate,
	}, nil
}

func (s *StockService) GetMissions(ctx context.Context, tab string, page, limit int) ([]repository.ListStockMissionsRow, int64, error) {
	offset := (page - 1) * limit
	missions, err := s.repo.ListStockMissions(ctx, repository.ListStockMissionsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, 0, err
	}

	var filtered []repository.ListStockMissionsRow
	for _, m := range missions {
		status := strings.ToLower(m.Status.String)
		if tab == "scheduled" {
			if status == "pending" || status == "in_progress" {
				filtered = append(filtered, m)
			}
		} else if tab == "completed" {
			if status == "completed" {
				filtered = append(filtered, m)
			}
		} else {
			filtered = append(filtered, m)
		}
	}

	return filtered, int64(len(filtered)), nil
}

func (s *StockService) CreateMission(ctx context.Context, title, description, assignedTo string, targetArea string) (repository.StockOpnameMission, error) {
	opID, err := uuid.Parse(assignedTo)
	if err != nil {
		return repository.StockOpnameMission{}, fmt.Errorf("invalid operator ID: %w", err)
	}

	// 1. Generate session number (e.g. REC-YYYYMMDD-XXXX)
	rand.Seed(time.Now().UnixNano())
	sessionNo := fmt.Sprintf("REC-%s-%04d", time.Now().Format("20060102"), rand.Intn(10000))

	// 2. Create Stock Session
	session, err := s.repo.CreateStockSession(ctx, repository.CreateStockSessionParams{
		SessionNo:   sessionNo,
		ConductedBy: opID,
	})
	if err != nil {
		return repository.StockOpnameMission{}, fmt.Errorf("failed to create stock session: %w", err)
	}

	// 3. Create Stock Mission
	mission, err := s.repo.CreateStockMission(ctx, repository.CreateStockMissionParams{
		SessionID:   pgtype.UUID{Bytes: session.ID, Valid: true},
		Title:       title,
		Description: pgtype.Text{String: description, Valid: true},
		AssignedTo:  opID,
		TargetArea:  pgtype.Text{String: targetArea, Valid: true},
	})
	if err != nil {
		return repository.StockOpnameMission{}, fmt.Errorf("failed to create stock mission: %w", err)
	}

	// 4. Query documents in target area
	searchPattern := "%" + targetArea + "%"
	if targetArea == "" || strings.ToLower(targetArea) == "all" {
		searchPattern = "all"
	}
	docs, err := s.repo.GetDocumentsByTargetArea(ctx, searchPattern)
	if err != nil {
		return repository.StockOpnameMission{}, fmt.Errorf("failed to query target documents: %w", err)
	}

	// 5. Populate stock items
	for _, d := range docs {
		locName := d.RackName.String
		if locName == "" {
			locName = "Unassigned Rack"
		}
		subLoc := d.BoxName.String
		if subLoc == "" {
			subLoc = "Unassigned Box"
		}
		sku := d.FileName.String
		if sku == "" {
			sku = d.ID.String()
		}

		_, err = s.repo.CreateStockItem(ctx, repository.CreateStockItemParams{
			SessionID:    session.ID,
			DocumentID:   pgtype.UUID{Bytes: d.ID, Valid: true},
			SkuCode:      pgtype.Text{String: sku, Valid: true},
			ItemName:     pgtype.Text{String: d.Title, Valid: true},
			SystemQty:    1,
			PhysicalQty:  0,
			Status:       "MISSING",
			LocationName: pgtype.Text{String: locName, Valid: true},
			SubLocation:  pgtype.Text{String: subLoc, Valid: true},
		})
		if err != nil {
			return repository.StockOpnameMission{}, fmt.Errorf("failed to populate stock item: %w", err)
		}
	}

	// Update session counts initially
	s.RecalculateSessionSummary(ctx, session.ID)

	return mission, nil
}

func (s *StockService) StartMission(ctx context.Context, missionID uuid.UUID) error {
	mission, err := s.repo.GetStockMission(ctx, missionID)
	if err != nil {
		return err
	}

	err = s.repo.UpdateStockMissionStatus(ctx, repository.UpdateStockMissionStatusParams{
		ID:     missionID,
		Status: pgtype.Text{String: "in_progress", Valid: true},
	})
	if err != nil {
		return err
	}

	if mission.SessionID.Valid {
		err = s.repo.UpdateStockSessionStatus(ctx, repository.UpdateStockSessionStatusParams{
			ID:     mission.SessionID.Bytes,
			Status: "in_progress",
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StockService) GetSessionItems(ctx context.Context, sessionID uuid.UUID) ([]repository.StockOpnameItem, error) {
	return s.repo.ListStockItemsBySession(ctx, sessionID)
}

type ScanResult struct {
	Item     repository.StockOpnameItem `json:"item"`
	IsMatch  bool                       `json:"is_match"`
	IsExtra  bool                       `json:"is_extra"`
	IsMissed bool                       `json:"is_missed"`
	Message  string                     `json:"message"`
}

func (s *StockService) ScanItem(ctx context.Context, sessionID uuid.UUID, skuCode string, currentRack string) (ScanResult, error) {
	// 1. Check if item exists in expected list for this session
	item, err := s.repo.GetStockItemBySku(ctx, repository.GetStockItemBySkuParams{
		SessionID: sessionID,
		SkuCode:   pgtype.Text{String: skuCode, Valid: true},
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Item is EXTRA (not expected in this audit zone)
			// Let's look up document title and location in global documents
			var docID uuid.UUID
			var docTitle, docRack, docBox string
			_, err := s.repo.GetDocumentsByTargetArea(ctx, "all") // query all
			if err == nil {
				// We can find matching document in db
				// Quick scan in database
				dbURL := "postgresql://postgres:mysecretpassword@mifsaka.com:8200/kreatifdms?sslmode=disable"
				conn, err := pgx.Connect(ctx, dbURL)
				if err == nil {
					defer conn.Close(ctx)
					conn.QueryRow(ctx, `
						SELECT d.id, d.title, r.name, b.name 
						FROM documents d 
						LEFT JOIN racks r ON d.rack_id = r.id 
						LEFT JOIN boxes b ON d.box_id = b.id 
						WHERE d.file_name = $1 OR d.id::text = $1
					`, skuCode).Scan(&docID, &docTitle, &docRack, &docBox)
				}
			}

			if docTitle == "" {
				docTitle = "Unknown Document (" + skuCode + ")"
			}

			// Create EXTRA item record
			newItem, err := s.repo.CreateStockItem(ctx, repository.CreateStockItemParams{
				SessionID:    sessionID,
				DocumentID:   pgtype.UUID{Bytes: docID, Valid: docID != uuid.Nil},
				SkuCode:      pgtype.Text{String: skuCode, Valid: true},
				ItemName:     pgtype.Text{String: docTitle, Valid: true},
				SystemQty:    0,
				PhysicalQty:  1,
				Status:       "EXTRA",
				LocationName: pgtype.Text{String: docRack, Valid: docRack != ""},
				SubLocation:  pgtype.Text{String: docBox, Valid: docBox != ""},
			})
			if err != nil {
				return ScanResult{}, err
			}

			s.RecalculateSessionSummary(ctx, sessionID)

			return ScanResult{
				Item:     newItem,
				IsMatch:  false,
				IsExtra:  true,
				IsMissed: false,
				Message:  fmt.Sprintf("Extra item found! System expected it at %s.", docRack),
			}, nil
		}
		return ScanResult{}, err
	}

	// 2. Item exists. Check location match.
	expectedRack := item.LocationName.String
	isMatch := strings.EqualFold(expectedRack, currentRack) || currentRack == ""

	status := "MATCH"
	if !isMatch {
		status = "EXTRA" // physically here but expected elsewhere
	}

	err = s.repo.UpdateStockItemQuantity(ctx, repository.UpdateStockItemQuantityParams{
		ID:          item.ID,
		PhysicalQty: 1,
		Status:      status,
	})
	if err != nil {
		return ScanResult{}, err
	}

	// Reload item
	item, _ = s.repo.GetStockItemBySku(ctx, repository.GetStockItemBySkuParams{
		SessionID: sessionID,
		SkuCode:   pgtype.Text{String: skuCode, Valid: true},
	})

	s.RecalculateSessionSummary(ctx, sessionID)

	if !isMatch {
		return ScanResult{
			Item:     item,
			IsMatch:  false,
			IsExtra:  true,
			IsMissed: false,
			Message:  fmt.Sprintf("Discrepancy alert! Expected at %s but scanned here.", expectedRack),
		}, nil
	}

	return ScanResult{
		Item:     item,
		IsMatch:  true,
		IsExtra:  false,
		IsMissed: false,
		Message:  "Scan match verified.",
	}, nil
}

func (s *StockService) ResolveItem(ctx context.Context, itemID uuid.UUID, resolution string, note string) error {
	return s.repo.UpdateStockItemResolution(ctx, repository.UpdateStockItemResolutionParams{
		ID:             itemID,
		Resolution:     pgtype.Text{String: resolution, Valid: true},
		ResolutionNote: pgtype.Text{String: note, Valid: true},
	})
}

func (s *StockService) ApproveSession(ctx context.Context, sessionID uuid.UUID, approverID uuid.UUID, pin string, resolutionNote string) error {
	// 1. Verify PIN
	user, err := s.repo.GetUserByID(ctx, approverID)
	if err != nil {
		return fmt.Errorf("approver not found: %w", err)
	}

	if !user.Pin.Valid {
		return errors.New("approver PIN not configured")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pin.String), []byte(pin)); err != nil {
		return errors.New("invalid authorization PIN")
	}

	// 2. Fetch all items in the session to commit updates to the inventory
	items, err := s.repo.ListStockItemsBySession(ctx, sessionID)
	if err != nil {
		return err
	}

	// 3. Process each discrepancy and apply to main inventory
	for _, item := range items {
		if item.DocumentID.Valid {
			docID := item.DocumentID.Bytes
			res := strings.ToLower(item.Resolution.String)

			if item.Status == "MISSING" && res == "confirmed_lost" {
				// Write off missing document -> Mark status as 'lost'
				dbURL := "postgresql://postgres:mysecretpassword@mifsaka.com:8200/kreatifdms?sslmode=disable"
				conn, err := pgx.Connect(ctx, dbURL)
				if err == nil {
					defer conn.Close(ctx)
					conn.Exec(ctx, "UPDATE documents SET physical_status = 'lost', status = 'archived' WHERE id = $1", docID)
				}
			} else if item.Status == "EXTRA" && res == "relocated" {
				// Relocate extra items -> Update rack/box location in main documents table
				// Find rack and box id based on sublocation text
				dbURL := "postgresql://postgres:mysecretpassword@mifsaka.com:8200/kreatifdms?sslmode=disable"
				conn, err := pgx.Connect(ctx, dbURL)
				if err == nil {
					defer conn.Close(ctx)
					// Simple relocation update: find rack matching item.LocationName
					var targetRackID, targetBoxID uuid.UUID
					conn.QueryRow(ctx, "SELECT id FROM racks WHERE name = $1 LIMIT 1", item.LocationName.String).Scan(&targetRackID)
					conn.QueryRow(ctx, "SELECT id FROM boxes WHERE name = $1 LIMIT 1", item.SubLocation.String).Scan(&targetBoxID)
					
					if targetRackID != uuid.Nil {
						conn.Exec(ctx, "UPDATE documents SET rack_id = $1, box_id = $2, physical_status = 'archived' WHERE id = $3", 
							targetRackID, targetBoxID, docID)
					}
				}
			}
		}
	}

	// 4. Update session status to approved
	err = s.repo.UpdateStockSessionStatus(ctx, repository.UpdateStockSessionStatusParams{
		ID:             sessionID,
		Status:         "approved",
		ApprovedBy:     pgtype.UUID{Bytes: approverID, Valid: true},
		ResolutionNote: pgtype.Text{String: resolutionNote, Valid: true},
	})
	if err != nil {
		return err
	}

	// 5. Update stock mission status to completed
	// Find mission linked to this session
	dbURL := "postgresql://postgres:mysecretpassword@mifsaka.com:8200/kreatifdms?sslmode=disable"
	conn, err := pgx.Connect(ctx, dbURL)
	if err == nil {
		defer conn.Close(ctx)
		var missionID uuid.UUID
		conn.QueryRow(ctx, "SELECT id FROM stock_opname_missions WHERE session_id = $1 LIMIT 1", sessionID).Scan(&missionID)
		if missionID != uuid.Nil {
			s.repo.UpdateStockMissionStatus(ctx, repository.UpdateStockMissionStatusParams{
				ID:     missionID,
				Status: pgtype.Text{String: "completed", Valid: true},
			})
		}
	}

	return nil
}

func (s *StockService) RecalculateSessionSummary(ctx context.Context, sessionID uuid.UUID) {
	items, err := s.repo.ListStockItemsBySession(ctx, sessionID)
	if err != nil {
		return
	}

	var matched, loan, missing, extra int32
	for _, item := range items {
		switch item.Status {
		case "MATCH":
			matched++
		case "ON LOAN":
			loan++
		case "MISSING":
			missing++
		case "EXTRA":
			extra++
		}
	}

	s.repo.UpdateStockSessionSummary(ctx, repository.UpdateStockSessionSummaryParams{
		ID:           sessionID,
		TotalMatched: pgtype.Int4{Int32: matched, Valid: true},
		TotalOnLoan:  pgtype.Int4{Int32: loan, Valid: true},
		TotalMissing: pgtype.Int4{Int32: missing, Valid: true},
		TotalExtra:   pgtype.Int4{Int32: extra, Valid: true},
	})
}
