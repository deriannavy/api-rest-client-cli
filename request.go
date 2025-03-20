package main

import (
	"io"
	"net/http"
	"time"

	"github.com/deriannavy/api-rest-client-cli/ui"
)

type Response struct {
	StatusCode string
}

func MakeRequest(item ui.Item) string {
	var (
		url    = item.UrlFormat()
		method = item.Request.Method
	)

	//  body := []byte(`{"key": "value"}`)
	// req, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	// Crea una nueva solicitud
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err.Error()
	}

	// Establece los encabezados necesarios
	req.Header.Set("Content-Type", "application/json")

	// Enviar la solicitud utilizando un cliente HTTP
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	// Leer y procesar el body de la respuesta
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	return res.Status + string(resBody)
}

func HandleBody(body ui.Body) string {
	// REFAC form y raw only supported
	return ""
}
