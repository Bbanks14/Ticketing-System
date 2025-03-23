package api_in

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	//"github.com/pkg/errors"
)

// answerSubjectRegex is a regular expression defining the syntax of a subject
// that creates a new answer to an existing ticket
var answerSubjectRegex = regexp.MustCompile(`\[Ticket "([A-Za-z0-9}+)"\].*`)

// emailRegex defines the syntax of valid email addresses
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-z0-9-]{0,61}[a-zA-Z0-9])?+$")

// stringType is the type for string in order to
// check the request parameter's type for validity
var stringType = reflect.TypeOf("")

// parameterMap is used as type for the expected API parameters
// and their types to check for all existing of all required and additional parameters
type parameterMap map[string]reflect.Type

//
