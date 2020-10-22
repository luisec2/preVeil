package main

import 	(
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, Preveil!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut Preveil!")
}

func latinHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hola, Preveil!")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, Preveil!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/latin", latinHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	log.Fatal(err)
}

