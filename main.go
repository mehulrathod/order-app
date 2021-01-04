package main

import (
	"log"
	"os"
	Routers "updated_structure/orderapp/routers"

	"github.com/joho/godotenv"
)

//Execution starts from main function
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR : ", err)
	}
	r := Routers.SetupRouter()

	port := os.Getenv("port")

	// For run on requested port
	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080" //localhost
	}
	type Job interface {
		Run()
	}

	err = r.Run(":" + port)
	if err != nil {
		log.Println("ERROR")
	}



}