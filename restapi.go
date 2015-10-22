package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostRequest struct {
	Name string `json:"name"`
}
type PostResponse struct {
	Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func helloPost(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	decoder := json.NewDecoder(req.Body)
	var t PostRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic("Some error in decoding the JSON")
	}

	var postResponse PostResponse
	postResponse.Greeting = "Hello, " + t.Name
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
	outputJson, err := json.Marshal(postResponse)
	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`
{
    "error": "Unable to marshal response."
}`))
		panic(err.Error())
	}
	rw.Write(outputJson)

}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hellopost/", helloPost)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
