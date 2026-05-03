package infra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/uuid"
)

type SearchService struct {
	es *elasticsearch.Client
}

func NewSearchService(es *elasticsearch.Client) *SearchService {
	return &SearchService{es: es}
}

type DocumentIndex struct {
	ID           uuid.UUID              `json:"id"`
	Title        string                 `json:"title"`
	Content      string                 `json:"content"`
	DepartmentID uuid.UUID              `json:"department_id"`
	Tags         []string               `json:"tags"`
	Metadata     map[string]interface{} `json:"metadata"`
	CreatedAt    string                 `json:"created_at"`
}

func (s *SearchService) IndexDocument(ctx context.Context, doc DocumentIndex) error {
	log.Printf("[SearchService] Indexing document: %s (Title: %s)", doc.ID, doc.Title)
	data, err := json.Marshal(doc)
	if err != nil {
		log.Printf("[SearchService] Error marshaling document %s: %v", doc.ID, err)
		return err
	}

	res, err := s.es.Index(
		"documents",
		bytes.NewReader(data),
		s.es.Index.WithDocumentID(doc.ID.String()),
		s.es.Index.WithContext(ctx),
	)
	if err != nil {
		log.Printf("[SearchService] Error calling ES Index for %s: %v", doc.ID, err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[SearchService] ES returned error for %s: %s", doc.ID, res.String())
		return fmt.Errorf("error indexing document: %s", res.String())
	}

	log.Printf("[SearchService] Document indexed successfully: %s", doc.ID)
	return nil
}

func (s *SearchService) Search(ctx context.Context, query string, deptID uuid.UUID) ([]uuid.UUID, error) {
	log.Printf("[SearchService] Searching documents for: '%s' (Dept: %s)", query, deptID)
	var buf bytes.Buffer
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"multi_match": map[string]interface{}{
							"query":  query,
							"fields": []string{"title", "content", "tags", "metadata.*"},
						},
					},
				},
			},
		},
	}

	// Add department filter if specified
	if deptID != uuid.Nil {
		searchQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"] = []map[string]interface{}{
			{
				"term": map[string]interface{}{
					"department_id": deptID.String(),
				},
			},
		}
	}

	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		log.Printf("[SearchService] Error encoding search query: %v", err)
		return nil, err
	}

	res, err := s.es.Search(
		s.es.Search.WithContext(ctx),
		s.es.Search.WithIndex("documents"),
		s.es.Search.WithBody(&buf),
		s.es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		log.Printf("[SearchService] Error calling ES Search: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[SearchService] ES search returned error: %s", res.String())
		return nil, fmt.Errorf("error searching documents: %s", res.String())
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	var ids []uuid.UUID
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		idStr := hit.(map[string]interface{})["_id"].(string)
		id, _ := uuid.Parse(idStr)
		ids = append(ids, id)
	}

	return ids, nil
}
func (s *SearchService) TestConnection(ctx context.Context, addresses []string, username, password, apiKey string) (map[string]interface{}, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
		APIKey:    apiKey,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create ES client: %v", err)
	}

	res, err := client.Info(client.Info.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ES: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("ES error: %s", res.String())
	}

	var info map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
		return nil, fmt.Errorf("failed to decode ES info: %v", err)
	}

	return info, nil
}
