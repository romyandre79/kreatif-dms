package repository

import (
	"context"

	"github.com/google/uuid"
)




func (q *Queries) CreateDocumentFile(ctx context.Context, documentID uuid.UUID, fileName, filePath string, fileSize int64, mimeType string, sortOrder int32) (DocumentFile, error) {
	const query = `
		INSERT INTO document_files (document_id, file_name, file_path, file_size, mime_type, sort_order)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, document_id, file_name, file_path, file_size, mime_type, sort_order, created_at
	`
	var f DocumentFile
	row := q.db.QueryRow(ctx, query, documentID, fileName, filePath, fileSize, mimeType, sortOrder)
	err := row.Scan(&f.ID, &f.DocumentID, &f.FileName, &f.FilePath, &f.FileSize, &f.MimeType, &f.SortOrder, &f.CreatedAt)
	return f, err
}

func (q *Queries) GetDocumentFileByID(ctx context.Context, fileID uuid.UUID) (DocumentFile, error) {
	const query = `
		SELECT id, document_id, file_name, file_path, file_size, mime_type, sort_order, created_at
		FROM document_files
		WHERE id = $1
	`
	var f DocumentFile
	err := q.db.QueryRow(ctx, query, fileID).Scan(
		&f.ID, &f.DocumentID, &f.FileName, &f.FilePath, &f.FileSize, &f.MimeType, &f.SortOrder, &f.CreatedAt,
	)
	return f, err
}

func (q *Queries) ListDocumentFiles(ctx context.Context, documentID uuid.UUID) ([]DocumentFile, error) {
	const query = `
		SELECT id, document_id, file_name, file_path, file_size, mime_type, sort_order, created_at
		FROM document_files
		WHERE document_id = $1
		ORDER BY sort_order ASC, created_at ASC
	`
	rows, err := q.db.Query(ctx, query, documentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []DocumentFile
	for rows.Next() {
		var f DocumentFile
		if err := rows.Scan(&f.ID, &f.DocumentID, &f.FileName, &f.FilePath, &f.FileSize, &f.MimeType, &f.SortOrder, &f.CreatedAt); err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	return files, rows.Err()
}
