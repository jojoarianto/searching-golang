package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// app has router using mux and db instance
type App struct {
	Router *mux.Router
}

// init app with predefined app configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling products
	a.Router.HandleFunc("/", a.indexPage).Methods("GET")
}

// handler for index page
func (a *App) indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Printf("Start a server on http://127.0.0.1%s", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
