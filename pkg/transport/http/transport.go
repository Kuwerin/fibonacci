package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/Kuwerin/fibonacci/pkg/service"
	"github.com/Kuwerin/fibonacci/pkg/transport/grpc/fibonaccipb"
)

type httpServer struct {
	router   *mux.Router
	services *service.Service
	logger   log.Logger
	http.Server
}

func NewHTTPServer(port int, logger log.Logger, services *service.Service) *httpServer {
	s := &httpServer{
		router:   mux.NewRouter(),
		services: services,
		logger:   logger,
	}
	s.configureRouter()

	s.Server.Addr = fmt.Sprintf(":%d", port)
	s.Server.Handler = s.router

	return s
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *httpServer) calculateFibonacci(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func(begin time.Time) {
		s.logger.Log(
			"entity", "transport",
			"type", "http",
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	req := new(fibonaccipb.GetFibonacciSliceRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.newErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	res, err := s.services.GetSlice(req)
	if err != nil {
		s.newErrorResponse(w, r, http.StatusInternalServerError, err)
	}

	s.respond(w, r, http.StatusOK, res)
}

func (s *httpServer) newErrorResponse(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *httpServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	if data != nil {
		w.WriteHeader(code)

		if err := json.NewEncoder(w).Encode(data); err != nil {
			return
		}

		return
	}
}

func (s *httpServer) appendJsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func (s *httpServer) configureRouter() {
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.Use(s.appendJsonHeader)
	s.router.HandleFunc("/fibonacci", s.calculateFibonacci).Methods("POST")
}
