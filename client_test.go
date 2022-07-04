package discordwebhook_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	discord "github.com/SebastiaanPasterkamp/go-discordwebhook"
)

func TestClientFromEnv(t *testing.T) {
	const environment = "DISCORD_INIT_TEST"

	testCases := []struct {
		name          string
		value         string
		expectedError error
	}{
		{"Success", "http://example.com", nil},
		{"Empty or missing", "", discord.ErrMissingWebhookURL},
		{"Malformed", "not a url", discord.ErrInvalidURL},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(environment, tt.value)
			defer os.Unsetenv(environment)

			_, err := discord.NewFromEnv(environment)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("Unexpected error. Expected %d, got %d.",
					tt.expectedError, err)
			}
		})
	}
}

func TestSend(t *testing.T) {
	testCases := []struct {
		name             string
		message          discord.Message
		fauxURL          string
		statusCode       int
		responseMessage  *discord.Message
		responseAPIError *discord.APIError
		expectedMessage  *discord.Message
		expectedAPIError *discord.APIError
		expectedError    error
	}{
		{"Success", discord.Message{Content: "Hello"},
			"", http.StatusOK, &discord.Message{Content: "Hello"}, nil,
			&discord.Message{Content: "Hello"}, nil, nil},
		{"Success without a body", discord.Message{},
			"", http.StatusOK, nil, nil,
			&discord.Message{}, nil, nil},
		{"No content", discord.Message{},
			"", http.StatusNoContent, nil, nil,
			nil, nil, nil},
		{"Invalid URL", discord.Message{},
			"http://localhost:1", 0, nil, nil,
			nil, nil, discord.ErrUnexpectedResponse},
		{"Bad request", discord.Message{},
			"", http.StatusBadRequest, nil, &discord.APIError{Code: 50035, Message: "Invalid Form Body"},
			nil, &discord.APIError{Code: 50035, Message: "Invalid Form Body"}, discord.ErrBadRequest},
		{"Invalid response", discord.Message{},
			"", http.StatusBadGateway, nil, nil,
			nil, nil, discord.ErrUnexpectedResponse},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ct := r.Header.Get("content-type")
				if ct != "application/json" {
					t.Errorf("Unexpected content-type. Expected %q, got %q.",
						"application/json", ct)
				}

				w.Header().Add("content-type", "application/json")
				w.WriteHeader(tt.statusCode)
				if tt.responseMessage != nil {
					json.NewEncoder(w).Encode(tt.responseMessage)
				} else if tt.responseAPIError != nil {
					json.NewEncoder(w).Encode(tt.responseAPIError)
				}
			}))
			defer ts.Close()

			dwc, err := discord.New(ts.URL)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if tt.fauxURL != "" {
				dwc, err = discord.New(tt.fauxURL)
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			}

			message, apiError, err := dwc.Send(tt.message, true)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("Unexpected error. Expected %d, got %d.",
					tt.expectedError, err)
			}

			if !reflect.DeepEqual(message, tt.expectedMessage) {
				t.Errorf("Unexpected message. Expected %v, got %v.",
					tt.expectedMessage, message)
			}

			if !reflect.DeepEqual(apiError, tt.expectedAPIError) {
				t.Errorf("Unexpected api error. Expected %v, got %v.",
					tt.expectedAPIError, apiError)
			}
		})
	}
}
