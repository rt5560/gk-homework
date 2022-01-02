package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main()  {
	for{
		response,err:= http.Get("http://localhost:8080/ping")
		if err != nil {
			log.Fatal(err)
		}else {
			body, _ := io.ReadAll(response.Body)
			fmt.Printf("status: %d,message: %s\n",response.StatusCode,body )
		}
		time.Sleep(time.Millisecond*100)
	}
}