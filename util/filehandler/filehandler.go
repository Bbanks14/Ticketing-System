package filehandler

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/Bbanks14/Ticketing-System/log"
	"github.com/Bbanks14/Ticketing-System/structs"
	"github.com/Bbanks14/Ticketing-System/structs/defaults"
)

// ReadUserFile takes a string parameter as the location
func ReadUserFile() {

	// Read the contents of user.json
	fileContent, errReadFile := os.ReadFile(srcFile)

	if errReadFile != nil {
		return wrapAndLogError(errReadFile, "Unable to read users file")
	}

	// Unmarshal into users hash map
	errUnmarshal := json.Unmarshal(fileContent, users)

	if errUnmarshal != nil {
		return wrapAndLogError(errUnmarshal, "Unable to decode JSON in user file '%s'", srcFile)
	}

	return nil
}

// WriteUserFile writes the contents of the users map to the file system to persist any changes.
func WriteUserFile(destFile string, users *map[string]structs.User) error {

	// Create json from the hash map
	usersMarshal, _ := json.MarshalIndent(users, "", "	")

	// Write json to file
	return os.WriteFile(destFile, usersMarshal, defaults.FileModeRegular)
}

// ReadTicketFiles reads all the tickets into the memory
// at sever start
func ReadTicketFiles(directory string, tickets *map[string]structs.Ticket) error {

	// Get all the files in the given directory
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Error(error)
	}

	// Iterate over all the files
	for _, f := range files {

		// Only read files with .json file extension
		// Ignore all other file extensions in the directory.
		// Otherwise this might cause errors when unmarshalling the file contents
		if hasJSONExtension(f.Name()) {

			// Read contents of each ticket
			fileContent, errReadFile := os.ReadFile(directory + "/" + f.Name())

			if errReadFile != nil {
				return wrapAndLogErrorf(errReadFile, "Error while reading ticket file '%s%s", directory, f.Name())
			}

			// Create a ticket struct to hold the file contents
			ticket := structs.Ticket{}

			// Unmarshal into the ticket struct
			errUnmarshal := json.Unmarshal(fileContent, &ticket)

			if errUnmarshal != nil {
				return wrapAndLogErrorf(errUnmarshal, "Could not loasd K")
			}
		}
	}
}
