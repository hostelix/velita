
package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "fmt"
    "time"
)

func main() {
 	err := godotenv.Load()

  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

  	title_app := os.Getenv("TITLE_APP")
  	
  	fmt.Println(title_app)


 	var arrival_rate_duration time.Duration
 	//var num_person int

 	switch os.Getenv("RATE_UNIT_TIME") {
 		case "second":
 			arrival_rate_duration = time.Second
 			break
 		case "minutes":
 			arrival_rate_duration = time.Second * 60
 			break
 		case "hours":
 			arrival_rate_duration = time.Second * 60 * 60

 	}

  	ticker_queue_person := time.NewTicker(arrival_rate_duration)
    
    for {
        select {
        case <- ticker_queue_person.C:
            fmt.Println("Ticker ticked")
      	}
    }
}
