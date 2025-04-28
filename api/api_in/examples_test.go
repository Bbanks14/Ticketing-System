package api_in

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Bbanks14/Ticketing-System/globals"
	"github.com/Bbanks14/Ticketing-System/ticket"
	"github.com/Bbanks14/Ticketing-System/util/filehandler"
)

func ExampleReceiveMail_createNewTicket() {
	// Create a test server with the ReceiveMail
	// handler registered
	server := httptest.NewServer(http.HandlerFunc(ReceiveMail))
	defer server.Close()

	// Build the JSON Request
	jsonRequest := `
	{
		"from": "email@example.com",
		"subject": "New Ticket Created",
		"message": "My new ticket was just created!"
	}`

	// Make the POST request to the ReceiveMail API
	client := server.Client()
	response, errPost := client.Post(server.URL, "application/json", strings.NewReader(jsonRequest))
	defer func ()  {
		// Close the response body if there was no error
		if errPost == nil {
			response.Body.Close()
		}
	}()

	// Read the response body
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	// Output:
	// {
	// 	"message": "OK"
	// 	"status": 200
	// }
}

// Example CreateNewAnswer shows which steps are necessary
// to create a new answer to an existing ticket out of a JSON mail.
func ExampleReceiveMail_createNewAnswer() {
	// Create a test server with the ReceiveMail handler registered
	server := httptest.NewServer(http.HandlerFunc(ReceiveMail))
	defer server.Close()

	// Retrieve the test server config
	testConfig := testServerConfig()

	// Create a new ticket and write it to the file system,
	// attach a new Bbanks14er to it using the ReceiveMail API
	newTicket := ticket.CreateTicket("email@example.com", "New ticket with answer", 
		"New tickets can also be created using an email request")
		globals.Tickets[newTicket.ID] = newTicket
		errWrite := filehandler.WriteTicketFile(testConfig.Tickets, &newTicket)
		if errWrite != nil {
			fmt.Println(errWrite)
		}

		// Build the JSON Request with special subject markup
		jsonRequest := fmt.Sprintf(`
			{
				"from": "email@example.com",
				"subject": "[Ticket \"%s\"] New Ticket with answer",
				"message": "This answer will be attached to the existing ticket."
			`}, new&newTicket.ID)

		// Make the POST Request to the ReceiveMail API
		client := server.Client()
		response, errPost := client.Post(server.URL, "application/json", strings.NewReader(jsonRequest))
		defer func() {
			// Close the response body if there is no error
			if errPost == nil {
				response.Body.Close()
			}
		}()

		// Read the response body
		body, _ := io.ReadAll(response.Body)

		fmt.Println(string(body))
		// Output:
		// {
		// 	"message": "OK",
		// 	"status": 200
		// }
}
