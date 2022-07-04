package discordwebhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	// ErrMissingWebhookURL is the error returned when the environment variable
	// for the discord Client is missing or empty.
	ErrMissingWebhookURL = fmt.Errorf("no discord URL environment variable")
	// ErrInvalidURL is the error returned when the client URL cannot be parsed.
	ErrInvalidURL = fmt.Errorf("invalid url")
	// ErrMalformedMessage is the error returned when the message to be sent
	// cannot be converted to json.
	ErrMalformedMessage = fmt.Errorf("malformed message")
	// ErrBadRequest is the error returned when discord rejects the message,
	// possibly with an error message.
	ErrBadRequest = fmt.Errorf("bad request")
	// ErrUnexpectedResponse is the error returned when the response does not
	// look like a discord success or error message.
	ErrUnexpectedResponse = fmt.Errorf("unexpected response")
)

// Client is an object that enables sending messages to a discord webhook
// configured on the client. The webhook URL should be treated as a secret, so
// a constructor to initialize a client using an environment variable is
// provided.
type Client struct {
	webhook *url.URL
}

// apiResponse is a combination of a Message and APIError struct so a single
// response body can be parsed into either format.
type apiResponse struct {
	*Message
	*APIError
}

// New creates a discord webhook Client using the provided webhook URL.
func New(raw string) (*Client, error) {
	webhook, err := url.Parse(raw)
	if err != nil || webhook.Scheme == "" || webhook.Fragment != "" {
		return nil, fmt.Errorf("%w: %v: %q", ErrInvalidURL, err, raw)
	}

	return &Client{webhook: webhook}, nil
}

// NewFromEnv creates a discord webhook Client using an environment variable
// for the secret webhook URL.
func NewFromEnv(env string) (*Client, error) {
	raw := os.Getenv(env)
	if raw == "" {
		return nil, fmt.Errorf("%w: %q", ErrMissingWebhookURL, env)
	}

	return New(raw)
}

// Send allows posting a Message to the configured Client webhook URL.
func (c Client) Send(m Message, wait bool) (*Message, *APIError, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(m)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v: %v", ErrMalformedMessage, err, m)
	}

	webhook := *c.webhook
	if wait {
		values := webhook.Query()
		values.Add("wait", "true")
		webhook.RawQuery = values.Encode()
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, webhook.String(), &buf)
	req.Header.Add("content-type", "application/json")
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v: %v", ErrBadRequest, err, m)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrUnexpectedResponse, err)
	}
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	result := Message{}
	apiError := &APIError{}

	if resp.Body != nil {
		body := apiResponse{
			Message:  &result,
			APIError: apiError,
		}
		err = json.NewDecoder(resp.Body).Decode(&body)
		switch err {
		case io.EOF:
		case nil:
		default:
			return nil, nil, fmt.Errorf("%w: message %v", ErrUnexpectedResponse, err)
		}
	}

	if apiError.Code == 0 {
		apiError = nil
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return &result, apiError, nil
	case http.StatusNoContent:
		// No response body expected
		return nil, apiError, nil
	case http.StatusBadRequest:
		return nil, apiError, fmt.Errorf("%w", ErrBadRequest)
	default:
	}

	return nil, apiError, fmt.Errorf("%w: Received status code %d",
		ErrUnexpectedResponse, resp.StatusCode)
}
