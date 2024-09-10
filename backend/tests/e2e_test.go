package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"ideahive/backend/internal/models"
	"ideahive/backend/tests/testutil"
)

func TestCreateIdea(t *testing.T) {
	// Ensure test environment is set up
	if err := testutil.SetupTestEnvironment(); err != nil {
		t.Fatalf("Failed to setup test environment: %v", err)
	}
	defer testutil.TeardownTestEnvironment()

	// Create a new idea
	newIdea := models.Idea{
		Title:       "Test Idea",
		Description: "This is a test idea",
	}

	// Convert idea to JSON
	payload, err := json.Marshal(newIdea)
	if err != nil {
		t.Fatalf("Failed to marshal idea: %v", err)
	}

	// Create a new request
	req, err := http.NewRequest("POST", testutil.TestServer.URL+"/api/ideas",
		bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	// Check the response
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status Created; got %v", resp.Status)
	}

	// Decode the response
	var createdIdea models.Idea
	if err := json.NewDecoder(resp.Body).Decode(&createdIdea); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check the created idea
	if createdIdea.Title != newIdea.Title {
		t.Errorf("Expected title %v; got %v", newIdea.Title, createdIdea.Title)
	}
	if createdIdea.Description != newIdea.Description {
		t.Errorf("Expected description %v; got %v", newIdea.Description, createdIdea.Description)
	}
	if createdIdea.ID == 0 {
		t.Error("Expected non-zero ID")
	}

	// Verify the idea was actually saved in the database
	var savedIdea models.Idea
	result := testutil.TestDB.First(&savedIdea, createdIdea.ID)
	if result.Error != nil {
		t.Fatalf("Failed to retrieve idea from database: %v", result.Error)
	}
	if savedIdea.Title != newIdea.Title || savedIdea.Description != newIdea.Description {
		t.Errorf("Database record doesn't match created idea")
	}
}
