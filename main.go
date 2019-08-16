package main

import (
	_ "expvar"
	"golang-service/transport"
	"kit/log"
	"net/http"
	"os"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	transport.RegisterHttpsServicesAndStartListener()
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "8080"
	//}
	port := "8888"
	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
		//fmt.Println("Error")
	}
}
