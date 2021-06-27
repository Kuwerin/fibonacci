package server

import (
	"encoding/json"
	"net/http"

	"github.com/Kuwerin/fibonacci/internal/model"
	"github.com/Kuwerin/fibonacci/internal/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)


type RESTServer struct {
	// httpServer *http.Server
	router *mux.Router
	services *service.Service 
	// handler *handler.Handler

}

func NewHttpServer(services *service.Service) *RESTServer{
	s := &RESTServer{
			router: mux.NewRouter(),
			services: services,
	}
	s.configureRouter()
	return s

}

func (s *RESTServer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	s.router.ServeHTTP(w,r)
}

// appendJsonHeader
func (s *RESTServer) appendJsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}


func (s *RESTServer) calculateFibonacci(w http.ResponseWriter, r *http.Request)  {
	req := &model.Fibonacci{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		s.newError(w, r, http.StatusBadRequest, err)	
		return
	}
	res := s.services.GetSlice(*req) 
	s.respond(w, r, http.StatusOK, res)
}

func (s *RESTServer) newError(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *RESTServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}){
	w.WriteHeader(code)
	if data != nil{
		json.NewEncoder(w).Encode(data)
	}
}


func (s *RESTServer) configureRouter()  {
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.Use(s.appendJsonHeader)
	s.router.HandleFunc("/fibonacci", s.calculateFibonacci).Methods("POST")
}

