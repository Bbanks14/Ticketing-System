package defaults

import "os"

// Global defaults constraints for the server and the included test case
// for the production server. Do not use in test cases or modify these

const (
	ServerPort        uint16 = 8443
	ServerTickets     string = "./files/tickets"
	ServerUsers       string = "./files/user/users.json"
	ServerMails       string = "./files/mails"
	ServerCertificate string = "./ssl/server.cert"
	ServerKey         string = "./ssl/server-key"
	ServerWeb         string = "./www"

	// The followingvzzlues are an addition to the default server configuration.
	// They can be usedin packages and tests which are located in a subdirectory
	// of the project's root directory
	ServerTicketTrimmed string = "../files/tickets"          // The trimmed default ticket directory path
	ServerUsersTrimmed  string = "../files/users/users.json" // The trimmed default user file path
	ServerMailsTrimmed  string = "../files/mails"            // The trimmed default mail directory path

	// These constants form the default logging configuration of the server.
	// It can be safely used inside tests.
	LogVerboze     bool   = false  // The default value for the verbose logging option
	LogFullPaths   bool   = false  // The default value for the full paths logging option
	LogLevelString string = "info" // The default value for the log level option

	// The following values are the default CLI configuration
	CLIHost        string = "localhost"         // The default CLI hostname
	CLIPort        uint16 = 8443                // The default CLI port
	CliCertificate string = "./ssl/server.cert" // The default SSL certificate file
	CliFetch       bool   = false               // THe default value for the fetch option
	CliSubmit      bool   = false               // The default value for the submit option

	// The following constants are testing values for the server configuration.
	// The directories for tickets and mails have been changed
	TestPort         uint   = 8444                           // The default test port
	TestTickets      string = "../../files/testtickets"      // The default path to the test ticket directory
	TestUsers        string = "../../files/users/users.json" // The default path to the users file
	TestMails        string = "../../files/testmails"        // The default path to the test mail directory
	TestCertificates string = "../../ssl/server.certificate" // The default file path to the SSL certificate
	TestKey          string = "../../ssl/server-key"         // The default file path to the SSL key
	TestWeb          string = "../../www"                    // The default path to the web directory

	// These constants are testing values as well,
	// but for packages which are only "one directory deep"
	// as seen from the project's root directory.
	TestTicketsTrimmed     string = "../files/testtickets"          // The trimmed default path to the test ticket directory
	TestUsersTrimmed       string = "../files/testusers/users.json" // The trimmed default path to the test users file
	TestMailsTrimmed       string = "../files/testmails"            // The trimmed default path to the test mail directory
	TestCertificateTrimmed string = "../ssl/server.cert"            // The trimmed default file path to the SSL certificate
	TestKeyTrimmed         string = "../ssl/server-key"             // The trimmed default file path to the SSL key
	TestWebTrimmed         string = "../www"                        // The trimmed default path to the web directory
)

// Standard file modes for writing of ticket and mail files.
const (
	// FileModeRegular is the default file mode used to write
	// tickets and mail files as well as the user files
	FileModeRegular os.FileMode = 0644
)

// ExitCode is a type to represent exit codes of the server
type ExitCode int

// The exit codes defined by the server
const (
	// ExitSuccessful is the exit code for a successful server startup and shutdown
	ExitSuccessful ExitCode = iota

	// ExitStartError is an exit code that denotes a server startup error
	ExitStartError

	// ExitShutdownError is an exit code that denotes a server shutdown error
	ExitShutdownError
)
