package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		port := os.Args[1]

		configFolder := os.Getenv("HOME") + "/.servedir"

		serverCertFile := configFolder + "/server.crt"
		serverKeyFile := configFolder + "/server.key"

		directory, err := os.Getwd()
		if err == nil {
			mux := http.NewServeMux()
			mux.Handle("/", corsHeader(http.FileServer(http.Dir(directory))))

			localIP := getLocalIP()

			if !fileExists(serverCertFile) || !fileExists(serverKeyFile) {
				log.Println("Starting web server at http://" + localIP + ":" + port)
				panic(http.ListenAndServe(":"+port, mux))
			} else {
				log.Println("Starting web server at https://" + localIP + ":" + port)
				panic(http.ListenAndServeTLS(":"+port, serverCertFile, serverKeyFile, mux))
			}
		} else {
			log.Println("couldn't get current working directory", err)
		}
	} else {
		fmt.Println("Serve Static files easily\n")
		fmt.Println("Usage:")
		fmt.Println("	serveDir <port>")
		fmt.Println("Example:")
		fmt.Println("	serveDir 3000")
	}
}

// add CORS header to http requests
func corsHeader(handler http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("serving file", request.URL.Path, "to", request.RemoteAddr)
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Credentials", "true")
		writer.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		handler.ServeHTTP(writer, request)
	}
}

// GetLocalIP returns the non loopback local IP of the host
func getLocalIP() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addresses {
		ip := address.(*net.IPNet)
		if ip.IP.To4() != nil && !ip.IP.IsLoopback() {
			return ip.IP.String()
		}
	}
	return "127.0.0.1"
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}
