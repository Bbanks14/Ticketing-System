package globals

import (
	"github.com/Bbanks14/ticketing-system/structs"
)

// Tickets holds all the created tickets.
var Tickets = make(map[string]structs.Ticket)

// Mails hold all currently cached mails.
var Mails = make(map[string]structs.Mail)

// ServerConfig holds the given server config for access to backend systems.
var ServerConfig *structs.ServerConfig

// LogConfig contains the global logging configuration
var LogConfig *structs.LogConfig

// Sessions holds all the sessions for the users.
var Sessions = make(map[string]structs.SessionManager)
