package main

import (
	monitor "github.com/lambda-zhang/systemmonitor-web/cron"
	"github.com/lambda-zhang/systemmonitor-web/routes"

	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func atexit() {
	log.Println(">>> exit")
	monitor.SM.Stop()
}

func printlistenaddr(port string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Printf("http://%s%s", ipnet.IP.String(), port)
			}
		}
	}
}

func main() {
	r := routes.Router()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	endPoint := fmt.Sprintf(":%s", port)

	defer atexit()

	printlistenaddr(endPoint)
	if err := http.ListenAndServe(endPoint, r); err != nil {
		log.Fatal(err)
	}
}
