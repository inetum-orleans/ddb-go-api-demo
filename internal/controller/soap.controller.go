package controller

import (
	"ddb-go-demo/internal/dto"
	"ddb-go-demo/internal/repository"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	XMLName xml.Name `xml:"Body"`
	User    dto.User `xml:"User"`
}

func (c *BaseController) SoapController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var envelope Envelope
	err = xml.Unmarshal(body, &envelope)
	if err != nil {
		http.Error(w, "Failed to unmarshal XML", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Received user email: %s\n", envelope.Body.User.Email)

	err = repository.CreateUser(c.ctx.Db, &envelope.Body.User)
	if err != nil {
		http.Error(w, "Failed to save User: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "User Saved successfully")
}
