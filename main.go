package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gitgithub.com/VarunPandya07/credit-card-validation/internal/api"
)

func main() {

	port := "3000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	http.HandleFunc("/validate", api.CardValidationHandler)
	fmt.Printf("server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
