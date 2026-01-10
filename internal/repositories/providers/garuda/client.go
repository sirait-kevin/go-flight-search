package garuda

import (
	"encoding/json"
	"go-flight-search/internal/domain"
	"go-flight-search/pkg/helper"
	"go-flight-search/pkg/logger"
	"os"
	"path/filepath"
	"time"
)

type Client struct {
	dataPath  string
	baseDelay int64
	maxDelay  int64
}

func New(dataPath string, baseDelay, maxDelay int64) *Client {
	return &Client{
		dataPath:  dataPath,
		baseDelay: baseDelay,
		maxDelay:  maxDelay,
	}
}

func (c *Client) Name() string {
	return "Garuda"
}

func (c *Client) Search(q domain.SearchQuery) (*[]domain.Flight, error) {
	// Simulate Garuda latency: 50–100ms
	helper.SimulateDelay(50, 100)

	// Load mock JSON
	logger.Info(filepath.Clean(c.dataPath))
	raw, err := os.ReadFile(filepath.Clean(c.dataPath))
	if err != nil {
		return nil, err
	}
	logger.Info(string(raw))

	var resp GarudaSearchResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	// Map to domain
	flights := make([]domain.Flight, len(resp.Flights))
	for i, f := range resp.Flights {
		flights[i] = MapToDomain(f)
	}

	return &flights, nil
}

func (c *Client) BaseDelay() time.Duration {
	return time.Duration(c.baseDelay) * time.Millisecond
}

func (c *Client) MaxDelay() time.Duration {
	return time.Duration(c.maxDelay) * time.Millisecond
}
