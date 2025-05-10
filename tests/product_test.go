package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:1323/products"

type Product struct {
	ID    uint   `json:"ID"` // <-- zmiana: uint + wielka litera (dla JSON z GORM.Model)
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func TestProductEndpoints(t *testing.T) {
	var createdID uint

	// CREATE
	t.Run("Create Product", func(t *testing.T) {
		body := Product{
			Name:  "Test Product",
			Price: 999,
		}
		jsonBody, _ := json.Marshal(body)

		resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		defer func() {
			_ = resp.Body.Close()
		}()
		var created Product
		if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
			t.Fatalf("❌ Failed to decode response: %v", err)
		}
		assert.Equal(t, body.Name, created.Name)
		assert.Equal(t, body.Price, created.Price)

		createdID = created.ID
		assert.NotZero(t, createdID)
	})

	// t.Run("Create Product 2", func(t *testing.T) {
	// 	body := Product{
	// 		Name:  "Test Product 2",
	// 		Price: 222,
	// 	}
	// 	jsonBody, _ := json.Marshal(body)

	// 	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(jsonBody))
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// 	defer resp.Body.Close()
	// 	var created Product
	// 	json.NewDecoder(resp.Body).Decode(&created)
	// 	assert.Equal(t, body.Name, created.Name)
	// 	assert.Equal(t, body.Price, created.Price)

	// 	createdID = created.ID
	// 	assert.NotZero(t, createdID)
	// })

	// GET ALL
	t.Run("Get Products", func(t *testing.T) {
		resp, err := http.Get(baseURL)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		defer func() {
			_ = resp.Body.Close()
		}()

		body, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(body), "Test Product")
	})

	// GET ONE
	t.Run("Get Product By ID", func(t *testing.T) {
		url := fmt.Sprintf("%s/%d", baseURL, createdID)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		defer func() {
			_ = resp.Body.Close()
		}()

		var p Product
		if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
			t.Fatalf("❌ Failed to decode response: %v", err)
		}

		assert.Equal(t, createdID, p.ID)
		assert.Equal(t, "Test Product", p.Name)
	})

	// UPDATE
	t.Run("Update Product", func(t *testing.T) {
		url := fmt.Sprintf("%s/%d", baseURL, createdID)
		body := Product{
			Name:  "Updated Product",
			Price: 123,
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		defer func() {
			_ = resp.Body.Close()
		}()
		var updated Product
		if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
			t.Fatalf("❌ Failed to decode response: %v", err)
		}

		assert.Equal(t, "Updated Product", updated.Name)
		assert.Equal(t, uint(123), updated.Price)
	})

	// DELETE
	t.Run("Delete Product", func(t *testing.T) {
		url := fmt.Sprintf("%s/%d", baseURL, createdID)
		req, _ := http.NewRequest(http.MethodDelete, url, nil)

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Sprawdź, czy już go nie ma
		resp, _ = http.Get(url)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
