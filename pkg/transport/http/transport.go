package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/Kuwerin/fibonacci/pkg/service"
	"github.com/Kuwerin/fibonacci/pkg/transport/grpc/fibonaccipb"
)

type HTTPServer struct {
	router   *mux.Router
	services *service.Service
	http.Server
}

func NewHTTPServer(services *service.Service) *HTTPServer {
	s := &HTTPServer{
		router:   mux.NewRouter(),
		services: services,
	}
	s.configureRouter()

	return s
}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *HTTPServer) calculateFibonacci(w http.ResponseWriter, r *http.Request) {
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

func (s *HTTPServer) newErrorResponse(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *HTTPServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	if data != nil {
		w.WriteHeader(code)

		if err := json.NewEncoder(w).Encode(data); err != nil {
			return
		}

		return
	}
}

func (s *HTTPServer) appendJsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func (s *HTTPServer) configureRouter() {
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.Use(s.appendJsonHeader)
	s.router.HandleFunc("/fibonacci", s.calculateFibonacci).Methods("POST")
}
