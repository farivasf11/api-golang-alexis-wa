package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func AnimeQuotes(rw http.ResponseWriter, r *http.Request) {
	clienteHttp := &http.Client{}

	url := "https://animechan.vercel.app/api/quotes"
	peticion, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)

	}
	peticion.Header.Add("Content-Type", "application/json")
	peticion.Header.Add("X-Hola-Mundo", "Ejemplo")
	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
	}
	defer respuesta.Body.Close()

	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}

	respuestaString := string(cuerpoRespuesta)
	log.Printf("Response Code: %d", respuesta.StatusCode)
	log.Printf("Headers: '%q'", respuesta.Header)
	contentType := respuesta.Header.Get("Content-Type")
	log.Printf("Content type: '%s'", contentType)
	log.Printf("Server response: '%s'", respuestaString)

	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(respuestaString))
}
