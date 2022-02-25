package ocp

import (
	"net/http"
)

// a HTTP health check handler in short form
func healthCheckShort(resp http.ResponseWriter, _ *http.Request) {
	resp.WriteHeader(http.StatusNoContent)
}

func healthCheckShortUsage() {
	http.Handle("/health", http.HandlerFunc(healthCheckShort))
}


type aStruct struct {}

func (*aStruct) ServeHTTP(response http.ResponseWriter, r *http.Request) {

}

func anyName(response http.ResponseWriter, r *http.Request) {

}

func launch() {
	http.Handle("/", http.HandlerFunc(anyName))
}