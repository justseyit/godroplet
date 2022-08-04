package utils

import (
	"fmt"
	"log"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErrorAsResponse(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
}
