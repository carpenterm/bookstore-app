package repository

import (
	"DemoApp/internal/models"
	"os"
	"testing"
)

// TestSearchCluster runs against a real Elasticsearch or OpenSearch named by
// SEARCH_TEST_URL (with SEARCH_TEST_USERNAME/PASSWORD), e.g.
// SEARCH_TEST_URL=https://localhost:9200 against a local OpenSearch.
func TestSearchCluster(t *testing.T) {
	url := os.Getenv("SEARCH_TEST_URL")
	if url == "" {
		t.Skip("SEARCH_TEST_URL not set")
	}
	repo, err := NewElasticsearchRepository(SearchConfig{
		URL: url, Username: os.Getenv("SEARCH_TEST_USERNAME"), Password: os.Getenv("SEARCH_TEST_PASSWORD"), InsecureTLS: true,
	})
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	str := func(s string) *string { return &s }
	num := func(n int) *int { return &n }
	products := []models.Product{
		{ID: 1, Name: "Pride and Prejudice", Author: str("Jane Austen"), Description: "A novel of manners", CategoryID: num(2)},
		{ID: 2, Name: "Moby Dick", Author: str("Herman Melville"), Description: "The whale", CategoryID: num(3)},
	}
	if err := repo.IndexProducts(products); err != nil {
		t.Fatalf("index: %v", err)
	}
	ids, err := repo.SearchProducts("austen", 0)
	if err != nil || len(ids) != 1 || ids[0] != 1 {
		t.Fatalf("search austen = %v, %v", ids, err)
	}
	ids, err = repo.SearchProducts("", 3)
	if err != nil || len(ids) != 1 || ids[0] != 2 {
		t.Fatalf("search category 3 = %v, %v", ids, err)
	}
	if err := repo.DeleteProduct(2); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if err := repo.IndexProduct(products[1]); err != nil {
		t.Fatalf("index one: %v", err)
	}
}
