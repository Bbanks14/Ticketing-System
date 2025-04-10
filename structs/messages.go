package structs

// CliMessage defines different standard messages used
// by the command-line interface that are printed to the console
type CliMessage string

// Messages in the command-line tool
const (
	// RequestCommandInput is the message for requesting a command input
	RequestCommandInput CliMessage = "\nTo Fetch Mails from the server, type '0'\n" +
		"To send mail to the server type '1'\n" +
		"To exit this program type '2'\n" +
		"Command: "

	// CommandNotAccepted is the error message if a command is not accepted
	CommandNotAccepted CliMessage = "Input not accepted, error: "

	// RequestEmailAddress is the message to prompt the user for a valid e-mail address
	RequestEmailAddress CliMessage = "Please enter your valid email address.\nEmail Address:"

	// RequestSubject is the message prompt to ask the user for a ticket subject
	RequestSubject CliMessage = "Please enter the subject of the message.\nSubject: "

	// RequestMessage is the message to prompt the user for a ticket message
	RequestMessage CliMessage = "Please enter the body of the message.\nMessage: "

	// RequestTicketID is the message to prompt the user for an optional ticket id.
	RequestTicketID CliMessage = "Please enter the ticket id if applicable." +
		" If left empty a new ticket wil be created.\nTicket ID (optional): "

	// TO is the string for the output of the recipient of an e-mail
	To CliMessage = "To: "

	// Subject is the string for the output of the subject of a ticket
	Subject CliMessage = "Subject: "
)

// CliErrMessage defines error messages printed
// to the console when the command-line tool faces an error
type CliErrMessage string

// Various error messages for wrong user inputs
const (
	TooManyInputs CliErrMessage = "Too many wrong user inputs. Arborting program execution...\n"
	NoValidOption CliErrMessage = "Not within the range of valid options"
	EmptyString   CliErrMessage = "String is Empty"
	InvalidEmail  CliErrMessage = "Email Address is invalid"
)
