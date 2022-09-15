package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hendralatumeten/vehicles_rental/src/routers"
)

func main() {
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Serve run on port 8080")
	http.ListenAndServe(":8080", mainRoute)

}
