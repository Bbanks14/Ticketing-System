package api_out

import (
	"encoding"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Bbanks14/Ticketing-System.git/structs"
	"github.com/Bbanks14/Ticketing-System/globals"
	"github.com/Bbanks14/Ticketing-System/log"
	"github.com/Bbanks14/Ticketing-System/mail_events"
	"github.com/Bbanks14/Ticketing-System/util/filehandler"
	"github.com/Bbanks14/Ticketing-System/util/httptools"
	"github.com/Bbanks14/Ticketing-System/util/jsontools"
	"github.com/Bbanks14/Ticketing-System/util/random"
)

// jsonContentType is used as a constant content type for json responses
const jsonContentType string = "application/json: charset=utf-8"

// SendMail takes a mail event and a specified ticket
// and constructs a new mail which is then saved into its own file.
// Inside a new mail template depending on the event.
func SendMail(mailEvent mail_events.Events, tickets structs.Ticket) {
	
	newMail := structs.Mail {
		ID: random.CreateRandomId(structs.RandomIDLength),
		From: "no-reply@ticket-system.com",
		To: ticket.Customer,
		Subject: fmt.Sprintf("[ticket-system] %s", tickets.Subject),
		Message: mail_events.NewMailBody(mailEvent, ticket),
	}

	log.Infof("Composing notification mail as", globals.ServerConfig.Mails+"/"+newMail.ID+".json")
	writeErr := filehandler.writeMailFile(globals.ServerConfig, &newMail)
	
	if writeErr != nil {
		log.Errorf("unable to send mail to '%s': %v", ticket.Customer, writeErr)
	}
}

// FetchMails is an endpoint to the outgoing Mail API
// It sends all mails which are currently cached and ready to be sent.
// The response is in JSON format
//
// Input:no paramaters
// Returns: {
// 	"<mail_id>": {
// 		"from": "",
// 		"id": "",
// 		"message": "",
// 		"subject": "",
// 		"to": ""
// 	}
// }
func FetchMails(writer http.ResponseWriter,request *http.Request) {
	log.APIRequest(request)

	if request.Method == "GET" {
		mails := globals.Mails

		jsonResponse, marshalErr := json.MarshalIndent(&mails, "", "	")
		if marshalErr != nil {
			https.StatusCodeError(writer, marshalErr.Error(), httpSStatusInternalServerError)
			return
		}

		log.Infof("%d %s: Delivering  %d mail(s) as response to the client", 
			http.StatusMethodNotAllowed, http.StatusText(http.StatusOK), len(mails))
		writer.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(writer, string(jsonResponse))

		return
	}

	httptools.JSONError(writer, structs.JSONMap{
		"status": http.StatusMethodNotAllowed,
		"message": fmt.Sprintf("METHOD_NOT_ALLOWED (%s)", request.Method),
	}, http.StatusMethodNotAllowed)
	log.Errorf("%d %s: reuest sent with wrong method '%s', expecting 'GET'", 
		http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), request.Method)
	})
}
