package telemetry

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

// DefaultOpenSearchURI is the URI of the OpenSearch instance.
// TODO: Fix to actual URI
const DefaultOpenSearchURI = "https://localhost:9200"

// DefaultIndexName is the name of the index to which telemetry events are sent.
const DefaultIndexName = "conduit-telemetry"

// DefaultTelemetryUserName is the username for the OpenSearch instance.
// We intentionally store credentials in the source code
// to report telemetry to a write-only database.
// TODO: Fix to actual username for algorand
const DefaultTelemetryUserName = "admin"

// DefaultTelemetryPassword is the password for the OpenSearch instance.
const DefaultTelemetryPassword = "admin"

// MakeTelemetryConfig initializes a new TelemetryConfig.
func MakeTelemetryConfig() Config {
	return Config{
		Enable:   true,
		URI:      DefaultOpenSearchURI,
		GUID:     uuid.New().String(), // Use Google UUID instead of go-algorand utils
		Index:    DefaultIndexName,
		UserName: DefaultTelemetryUserName,
		Password: DefaultTelemetryPassword,
	}
}

// initializeOpenSearchClient creates a new OpenSearch client.
func initializeOpenSearchClient(cfg Config) (*opensearch.Client, error) {
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{cfg.URI},
		// These credentials are here intentionally. Not a bug.
		Username: cfg.UserName,
		Password: cfg.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create new OpenSearch client with URI %s: %w", cfg.URI, err)
	}
	return client, nil
}

// MakeOpenSearchClient initializes a new TelemetryState.
func MakeOpenSearchClient(cfg Config) (*OpenSearchClient, error) {
	client, err := initializeOpenSearchClient(cfg)
	if err != nil {
		return nil, err
	}

	telemetryState := &OpenSearchClient{
		Client:          client,
		TelemetryConfig: cfg,
	}
	return telemetryState, nil
}

// MakeTelemetryStartupEvent sends a startup event when the pipeline is initialized.
func (t *OpenSearchClient) MakeTelemetryStartupEvent() Event {
	return Event{
		Message: "starting conduit",
		GUID:    t.TelemetryConfig.GUID,
		Time:    time.Now(),
	}
}

// SendEvent sends a TelemetryEvent to OpenSearch.
func (t *OpenSearchClient) SendEvent(event Event) error {
	data, _ := json.Marshal(event)
	req := opensearchapi.IndexRequest{
		Index: t.TelemetryConfig.Index,
		Body:  bytes.NewReader(data),
	}
	_, err := req.Do(context.Background(), t.Client)
	if err != nil {
		return fmt.Errorf("failed to insert event: %w", err)
	}
	return nil
}