package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	fmt.Println("Jajanan Suci is packing")

	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Printf("Listening to port %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func Run() {
	var server = Server{}

	server.Initialize()
	server.Run(":9000")
}
