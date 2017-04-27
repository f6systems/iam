//filename=iam.go
//see=hopley@ipcloud.net
package main

import (
	"fmt"
	"log"
	       "net/http"
		   "time"
)

func init() {
	log.Printf("init()\n")
}

func main() {
	log.Printf("Starting.\n")
	        http.Handle("/", http.HandlerFunc(sayHello))
        http.Handle("/ip",http.HandlerFunc(ip))
        http.Handle("/dude", http.HandlerFunc(dude))
        //err := http.ListenAndServe("0.0.0.0:80", nil)
        //err := http.ListenAndServe("10.12.1.50:80", nil)
        err := http.ListenAndServe("10.12.1.239:80", nil)
        // err := http.ListenAndServe("10.9.8.78:8080", nil)
        if err != nil {
        //fmt.Printf("ListenAndServe Error :" + err.String())
        fmt.Printf("ListenAndServe Error. (err=%s)\n",err)
        }


}

//
func ip(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        w.Write([]byte(ip))

}//end ip()

// func sayHello(c *http.Conn, req *http.Request) {
func sayHello(w http.ResponseWriter, r *http.Request) {
        //fmt.Printf("New Request... \n")
        log.Printf("[%v] - %s ; %s ; / \n",time.Now(),r.RemoteAddr,r.Host)
        //c.Write([]byte("<h1>Go Say's Hello</h1><h2>(Via http)</h2>"))
       w.Header().Set("Content-Type", "text/plain")
        fmt.Fprintf(w,"\nIPCloud - see hopley@ipcloud.net  on %s\n\n",r.Host)
        fmt.Fprintf(w, "Method: %s\n", r.Method)
        fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
        fmt.Fprintf(w, "Host: %s\n", r.Host)
        fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
        fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
        fmt.Fprintf(w, "URL: %#v\n", r.URL)
        fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
        fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
        //fmt.Fprintf(w, "TLS: %#v\n", r.TLS)                                                                                                                    
        fmt.Fprintf(w, "\nHeaders:\n")
        //dh - Print request Headers
        for k, v := range r.Header {
                fmt.Fprintf(w, " [%s] = %s\n",k,v)
        }
        //c.Flush()
}

func dude(c http.ResponseWriter, req *http.Request){

 fmt.Println("~Got a dude request")

}//end dude()
