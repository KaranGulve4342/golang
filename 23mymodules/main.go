package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in Golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving home page")
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}

// go mod tidy: Cleans up the go.mod file by removing unnecessary dependencies.
// go mod verify: Verifies the dependencies of the module.
// go list: Lists the packages in the module.
// go list all: Lists all packages available in the Go environment.
// go list -m all: Lists all modules used in the build.
// go list -m -versions github.com/gorilla/mux: Lists all available versions of the github.com/gorilla/mux module.
// go mod why: Explains why packages or modules are needed.
// go mod why github.com/gorilla/mux: Explains why the github.com/gorilla/mux module is needed.
// go mod graph: Displays the module dependency graph.
// go mod edit -go 1.16: Sets the Go version in the go.mod file to 1.16.
// go mod edit -go 1.23.3: Sets the Go version in the go.mod file to 1.23.3.
// go mod vendor: Creates a vendor directory with all dependencies.
