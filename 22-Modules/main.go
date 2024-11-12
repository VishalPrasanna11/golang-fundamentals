package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, from Vishal Prasanna's GoLang Playground!")
	greeter()
	fmt.Println("Server started at port 8080")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r), "Server started at port 8080")

}

func greeter() {
	fmt.Println("Welcome to Functions in GoLang")
	fmt.Println("go mod tidy is used to clean up the go.mod file")
	fmt.Println("go mod init is used to initialize the go.mod file")
	fmt.Println("go mod vendor is used to create a vendor folder")
	fmt.Println("go mod download is used to download the dependencies")
	fmt.Println("go mod verify is used to verify the dependencies")
	fmt.Println("go mod graph is used to display the dependency graph")
	fmt.Println("go mod edit is used to edit the go.mod file")
	fmt.Println("go mod why is used to explain why a package is needed")
	fmt.Println("go list is used to list the packages")
	fmt.Println("go get is used to download the packages")
	fmt.Println("go run is used to run the go files")
	fmt.Println("go build is used to build the go files")
	fmt.Println("go test is used to test the go files")
	fmt.Println("go fmt is used to format the go files")
	fmt.Println("go doc is used to get the documentation of the go files")
	fmt.Println("go env is used to get the environment variables")
	fmt.Println("go version is used to get the version of the go")
	fmt.Println("go help is used to get the help on go commands")
	fmt.Println("go clean is used to clean the go files")
	fmt.Println("go install is used to install the go files")
	fmt.Println("go generate is used to generate the go files")
	fmt.Println("go list -m all is used to list all the modules")

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to GoLang Web Development"))
}
