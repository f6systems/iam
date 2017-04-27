//filename=iam.go
//see=hopley@ipcloud.net
package main

import (
	"fmt"
	"log"
	       "net/http"
		  
)

func init() {
	log.Printf("init()\n")
}

func main() {
	log.Printf("Starting.\n")
	        http.Handle("/", http.HandlerFunc(sayHello))
        http.Handle("/ip",http.HandlerFunc(ip))
        http.Handle("/dude", http.HandlerFunc(Dude))
        //err := http.ListenAndServe("0.0.0.0:80", nil)
        //err := http.ListenAndServe("10.12.1.50:80", nil)
        err := http.ListenAndServe("10.12.1.239:80", nil)
        err = http.ListenAndServe(":80", nil)
        // err := http.ListenAndServe("10.9.8.78:8080", nil)
        if err != nil {
        //fmt.Printf("ListenAndServe Error :" + err.String())
        fmt.Printf("ListenAndServe Error. (err=%s)\n",err)
        }


}

