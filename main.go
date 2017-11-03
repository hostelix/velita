
package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "fmt"
)

func main() {
 	err := godotenv.Load()

  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

  	title_app := os.Getenv("TITLE_APP")
  	
  	fmt.Println(title_app)
}