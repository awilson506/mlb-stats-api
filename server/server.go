package server

import (
	"encoding/json"
	"net/http"

	"github.com/awilson506/mlb-stats-api/api"
)

type Server struct {
	client api.Client
	server *http.Server
	mux    *http.ServeMux
}

type handler struct {
	pattern string
	handler http.HandlerFunc
}

// New - get a new instance of the server
func New(addr string, client api.Client) *Server {
	s := &Server{
		client: client,
		mux:    http.NewServeMux(),
	}

	handlers := []handler{
		{pattern: "/v1/stats", handler: s.statsHandler},
	}

	for _, h := range handlers {
		s.mux.HandleFunc(h.pattern, h.handler)
	}

	s.server = &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}
	return s
}

// Start - start the server
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// statsHandler - handle the get stats request
func (s *Server) statsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	strTeamId := r.URL.Query().Get("team-id")
	date := r.URL.Query().Get("date")
	teamId, msg, err := api.ValidateContentGetRequest(strTeamId, date)

	if err {
		s.WriteErrorResponse(w, msg.Errors)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.client.GetFavoriteMLBStats(int(teamId), date))
}

// WriteErrorResponse - handle writing error responses
func (s *Server) WriteErrorResponse(w http.ResponseWriter, errors map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errors)
}
