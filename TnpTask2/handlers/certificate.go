package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gopkg.in/gomail.v2"

	"TnpTask2/database"
	"TnpTask2/utils"
)

var validate = validator.New()

// Certificate struct with validation
type Certificate struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Course   string `json:"course" validate:"required"`
	FilePath string `json:"file_path" validate:"required,url"`
}

// GetCertificates fetches all certificates
func GetCertificates(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, email, course, file_path FROM certificates")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var certificates []Certificate
	for rows.Next() {
		var c Certificate
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Course, &c.FilePath); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		certificates = append(certificates, c)
	}

	utils.JSONResponse(w, certificates)
}

func SendCertificate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var cert Certificate
	err := database.DB.QueryRow("SELECT name, email, course, file_path FROM certificates WHERE id = $1", id).
		Scan(&cert.Name, &cert.Email, &cert.Course, &cert.FilePath)
	if err != nil {
		http.Error(w, "Certificate not found", http.StatusNotFound)
		return
	}

	// Validate certificate
	if err := validate.Struct(cert); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	d := gomail.NewDialer("smtp.example.com", 587, "your_email@example.com", "your_password")
	m := gomail.NewMessage()
	m.SetHeader("From", "your_email@example.com")
	m.SetHeader("To", cert.Email)
	m.SetHeader("Subject", "Your Certificate")
	m.SetBody("text/plain", fmt.Sprintf("Hello %s,\n\nHere is your certificate.\n\nBest Regards", cert.Name))
	m.Attach(cert.FilePath)

	if err := d.DialAndSend(m); err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, map[string]string{"message": "Email sent successfully"})
}
