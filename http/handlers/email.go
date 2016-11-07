package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pbk/kit-api/email"
	"github.com/pbk/kit-api/models"
)

// SendEmail ...
func SendEmail(w http.ResponseWriter, r *http.Request) {
	v := models.Event{}
	json.NewDecoder(r.Body).Decode(&v)

	log.Printf("SendEmail: %+v", v)

	// TODO: move the message body into template.

	// Create message string.
	conv := "Follow the conversation <br/><br/>"

	for _, m := range v.Messages {
		conv += fmt.Sprintf("%s: %s<br/>", m.CreatedByUser.Email, m.Message)
	}

	subj := fmt.Sprintf("Your friend %s has invited you to an event!", v.CreatedByUser.Email)

	for _, i := range v.Invitees {
		curr := fmt.Sprintf(
			"<a href=\"%s/event/%s?curr_user_uuid=%s\">Your Event Link</a> <br/><br/>",
			os.Getenv("KIT_WEB_URL"),
			v.UUID,
			i.UUID,
		)
		curr += conv
		email.Send(subj, curr, i.Email)
	}

	writeJSON(w, "Success")
}
