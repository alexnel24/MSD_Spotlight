package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// "os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Hello Alex")

	port := os.Getenv("PORT")
	fmt.Println("Port is: ", port)

	router := httprouter.New()
	router.GET("/healthcheck", Healthcheck)
	router.GET("/student", StudentHello)

	log.Fatal(http.ListenAndServe(":" + port, router))
}

func Healthcheck(out http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	log.Println("Healthcheck hit")
	resp := "Hello MSD!"
	jsResp, err := json.Marshal(resp)

	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	out.WriteHeader(statusCode)
	out.Write(jsResp)
}

func StudentHello(out http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	log.Println("StudentHello hit")

	queryParam := req.URL.Query()
	studentName := queryParam.Get("name")
	if studentName == "" {
		log.Println("No name given")
		Healthcheck(out, req, httprouter.Params{})
		return
	}

	stringToReturn := "Hello " + studentName + "!"
	jsResp, err := json.Marshal(stringToReturn)

	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	out.WriteHeader(statusCode)
	out.Write(jsResp)
}
