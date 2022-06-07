package http

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Context context.Context
}

func NewAPI(ctx context.Context) *API {
	return &API{
		Context: ctx,
	}
}

func (api *API) SetupServer() {
	fmt.Println("my first api in golang", api.Context.Value("key"))

	r := mux.NewRouter()

	r.HandleFunc("/movies", create).Methods("POST")
	r.HandleFunc("/movies", findAll).Methods("GET")
	r.HandleFunc("/movies/{id}", findById).Methods("GET")
	r.HandleFunc("/movies/{id}", remove).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

func create(w http.ResponseWriter, r *http.Request) {
	json, _ := ioutil.ReadAll(r.Body)

	fmt.Println("create method", string(json))
}

func findAll(http.ResponseWriter, *http.Request) {
	fmt.Println("create findAll")
}

func findById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("create findById", params["id"])
}

func remove(http.ResponseWriter, *http.Request) {
	fmt.Println("create remove")
}
