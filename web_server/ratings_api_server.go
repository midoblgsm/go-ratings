package web_server

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/midoblgsm/go-ratings/resources"
	"github.com/midoblgsm/go-ratings/utils"

	"github.com/gorilla/mux"
)

type RatingsApiServer struct {
	port int
}

func NewRatingsApiServer(port int) RatingsApiServer {
	return RatingsApiServer{port: port}
}

func (r *RatingsApiServer) InitializeHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/v1/ratings/{productId}", r.GetRating()).Methods("GET")
	return router
}

func (r *RatingsApiServer) Start() error {
	router := r.InitializeHandler()
	http.Handle("/", router)

	log.Printf("Starting Ratings API server on port %d ....", r.port)
	log.Println("CTL-C to exit/stop Ratings API server")

	return http.ListenAndServe(fmt.Sprintf(":%d", r.port), nil)

}

func (r *RatingsApiServer) GetRating() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		id := utils.ExtractVarsFromRequest(req, "productId")
		min := 1
		max := 5
		stars := rand.Intn(max-min) + min

		rating := resources.Rating{Id: id, Stars: stars, Color: "red"}
		fmt.Printf("rating %#v", rating)
		utils.WriteResponse(w, http.StatusOK, rating)
	}
}

func (r *RatingsApiServer) EnableCORS(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

	return
}
