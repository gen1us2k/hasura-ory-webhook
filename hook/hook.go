package hook

import (
	"context"
	"encoding/json"
	"net/http"
	"webhook/config"

	"github.com/akrylysov/algnhsa"
	client "github.com/ory/client-go"
)

type (
	// Hook ...
	Hook struct {
		config *config.HookConfig
		client *client.APIClient
	}

	// Response struct returns the following object
	// back to Hasura
	//
	// Hasura decides what to do with the request
	Response struct {
		HasuraUserID string `json:"X-Hasura-User-Id"`
		HasuraRole   string `json:"X-Hasura-Role"`
	}
	// Request represents data sent by Hasura
	// to a webhook
	Request struct {
		Headers map[string]string `json:"headers"`
	}
)

// NewHook creates a new Hasura Webhook
func NewHook(c *config.HookConfig) *Hook {
	h := &Hook{config: c}
	conf := client.NewConfiguration()
	conf.Servers = client.ServerConfigurations{
		{
			URL: c.OrySDKURL,
		},
	}
	h.client = client.NewAPIClient(conf)
	h.init()
	return h
}

func (h *Hook) init() {
	http.HandleFunc("/", h.handler)
}

func (h *Hook) handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Not implemented"))
		return
	}
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to decode json"))
		return
	}
	sess, err := h.toSession(r.Context(), req.Headers["cookie"])
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Response{
		HasuraUserID: sess.Identity.Id,
		HasuraRole:   "user",
	})

}

func (h *Hook) toSession(ctx context.Context, cookie string) (*client.Session, error) {
	sess, _, err := h.client.V0alpha2Api.
		ToSession(ctx).
		Cookie(cookie).
		Execute()
	return sess, err
}

// Start starts a web server
func (h *Hook) Start() error {
	if h.config.Environment == config.Production {
		algnhsa.ListenAndServe(http.DefaultServeMux, nil)
	}
	return http.ListenAndServe(":8090", nil)
}
