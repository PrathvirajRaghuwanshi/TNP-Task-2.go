package handlers

import (
	"log"
	"net/http"

	"TnpTask2/database"
	"TnpTask2/utils"

	"gopkg.in/gomail.v2"
)

// SendBulkMessages sends bulk emails efficiently
func SendBulkMessages(w http.ResponseWriter, r *http.Request) {
	// Use the global DB connection
	rows, err := database.DB.Query("SELECT email FROM certificates")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var emails []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		emails = append(emails, email)
	}

	// Configure SMTP server
	dialer := gomail.NewDialer("smtp.example.com", 587, "your_email@example.com", "your_password")

	// Establish a connection
	sender, err := dialer.Dial()
	if err != nil {
		http.Error(w, "Failed to connect to mail server", http.StatusInternalServerError)
		return
	}
	defer sender.Close()

	// Send emails efficiently using the same connection
	for _, email := range emails {
		m := gomail.NewMessage()
		m.SetHeader("From", "your_email@example.com")
		m.SetHeader("To", email)
		m.SetHeader("Subject", "Bulk Notification")
		m.SetBody("text/plain", "Hello, this is a bulk message!")

		if err := gomail.Send(sender, m); err != nil {
			log.Printf("Failed to send email to %s: %v", email, err)
		} else {
			log.Printf("Email sent to %s", email)
		}
	}

	utils.JSONResponse(w, map[string]string{"message": "Bulk messages sent successfully"})
}
