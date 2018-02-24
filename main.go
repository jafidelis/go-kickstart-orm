package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-kickstart-orm/business"
	"github.com/go-kickstart-orm/server"
)

func main() {
	us := business.UserService{}
	us.CreateDefaultUser()
	server := server.NewServer()
	fmt.Println("Server responding in localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
