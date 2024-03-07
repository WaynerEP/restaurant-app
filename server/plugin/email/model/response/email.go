package response

type Email struct {
	To      string `json:"to"`      // Recipient's email address
	Subject string `json:"subject"` // Email subject
	Body    string `json:"body"`    // Email content
}
