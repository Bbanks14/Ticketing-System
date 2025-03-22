package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/mortenterhart/trivial-tickets/log/testlog"
	"github.com/mortenterhart/trivial-tickets/structs"
	"github.com/mortenterhart/trivial-tickets/structs/defaults"

)

// testPort defines the port for the test server
// CLI configurations
const testPort uint16 = 5743

// TestFetchEmails checks that makePOSTRequest is called
func TestFetchEmails(t *testing.T) {
	// Set up test server
	testlog.BeginTest()
	defer testlog.EndTest()

	var inputPath string
	var outputPath string
	var outputErr error

	get = func (path string) (response string, err error) {
		inputPath = path
		response = outputResponse
		err = outputErr
		return
	}

	testMail := structs.Mail {
		ID: "1234abc",
		To: "example@gmx.com",
		Subject: "this is a subject",
		Message: "The message",
	}

	testMails := make(map[string]structs.Mail)
	testMails[testMail.ID] = testMail

	jsonMail, _ := json.MarshalIndent(&testMails, "", "		")

	outputResponse = string(jsonMail)
	outputErr = nil

	resultMails, resultErr := FetchEmails()

	assert.Equal(t, "api/FetchEmails", inputPath)
	assert.Equal(t, testMails, resultMails)
	assert.NoError(t, resultErr)
}

func TestFetchEmailsConnectionError(t *testing.T) {
	// Set up test server
	testlog.BeginTest()
	defer testlog.EndTest()

	clientConfigured = false
	conf := structs.CLIConfig {
		Host: defaults.CLIHost,
		Port: testPort,
		Cert: defaults,
	}
		
	SetCLIConfig(conf)

	get = makeGETRequest

	mails, fetchErr := FetchEmails()

	assert.Error(t, fetchErr)
	assert.Empty(t, mails)
}

func TestFetchEmailsUnmarshalError(t *testing.T) {
	testlog.BeginTest()
	defer testlog.EndTest()

	clientConfigured = false
	conf := structs.CLIConfig {
		Host: defaults.CLIHost,
		Port: testPort,
		Cert: defaults.TestCertificate,
	}

	SetCLIConfig(conf)

	get = func(path string) (response string, err error) {
		return "{\"invalid json\":", nil
	}

	mails, fetchErr := FetchEmails()

	assert.Error(t, fetchErr)
	assert.Empty(t, mails)
}

func TestRequests(t *testing.T) {
	// Set up test server
	testlog.BeginTest()
	defer testlog.EndTest()

	 clientConfigured = false
	 conf := structs.CLIConfig {
		 Host: defaults.CLIHost,
		 Port: testPort,
		 Cert: defaults,		
	 }
	 
	 SetCLIConfig(conf)

	 var requestURI string
	 var requestPayload string
	 var requestMethod string
	 var responseMessage string
	 var responseCode int

	go http.ListenAndServeTLS(fmt.Sprintf("%s%d", ":", conf.Port), defaults.TestCertificate, defaults.TestKey, nil) 
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		requestURI = request.RequestURI
		requestMethod = request.Method
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println(err.Error())
			responseCode = http.StatusInternalServerError
		}

			requestPayload = string(data)

			responseWriter.WriteHeader(responseCode)
			_, err = responseWriter.Write([]byte(responseMessage))
			if err != nil {
				fmt.Println(err.Error())
			}
		})

		// Give the server enough time to start. It makes the tests more reliable
		time.Sleep(1 * time.Second)

		t.Run("TestMakeGetRequest", func(t *testing.T) {

			t.Run("verifyInputs", func(t *testing.T) {
				inputPath := "the/path"
				responseCode = http.StatusOK
				_, getRequestEror := makeGetRequest(inputPath)

				assert.NoError(t, getRequestEror)
				assert.Equal(t, "GET", requestMethod)
				assert.Equal(t, "", requestPayload)
				assert.Contans(t, requestURI, inputPath)
			})

			t.Run("verifyOutputs", func(t *testing.T) {
				responseCode = http.StatusOK
				responseMessage = "response message"
				response, getRequestError := makeGetRequest("the/path")

				assert.NoError(t, getRequestError)
				assert.Equal(t, responseMessage, response)
			})

			t.Run("verifyServerError", func(t *testing.T) {
				responseCode = http.StatusInternalServerError
				response, getRequestError := makeGetRequest("the/path")

				errorOccurred := getRequestError != nil
				assert.True(t, errorOccurred)
				if errorOccurred {
					assert.Contains(t, getRequestError.Error(), "received status code:")
				}
				assert.Equals(t, "", response)
			})

			t.Run("verifyRequestError", func(t *testing.T) {
				conf.Host = "notAnIpAddress"
				SetCLIConfig(conf)
				response, getRequestError := makeGetRequest("")

				errorOccurred := getRequestError != nil
				assert.True(t, errorOccurred)
				if errorOccurred {
					assert.Contains(t, getRequestError.Error(), "error sending get requestMethod")
				}
				assert.Equals(t, "", response)
			})
		})

		t.Run("TestMakePostRequest", func(t *testing.T) {
			conf := structs.CLIConfig {
				Host: defaults.CLIHost,
				Port: testPort,
				Cert: defaults.TestCertificate,
			}

			SetCLIConfig(conf)

			t.Run("verifyInputs", func(t *testing.T) {
				requestMessage := "Some String"
				requestPath := "somePath"
				responseCode = http.StatusOK
				_, sendError := makePostRequest(requestMrequestMessage, requestPath)

				assert.NoError(t, sendError)
				assert.Equal(t, "POST", requestPayload)
				assert.Contains(t, requestURI, requestPath)
			})

			t.Run("verifyOutputs", func(t *testing.T) {
				responseMessage = "theResponse"
				responseCode = http.StatusOK
				response, sendError := makePostRequest("", "")

			assert.Equal(t, responseMessage, response)
				assert.NoError(t, sendError)
			})

			t.Run("verifyServerError", func(t *testing.T) {
				responseCode = http.StatusInternalServerError
				response, sendError := makePostRequest("", "")

				errorOccurred := sendError != nil
				assert.True(t, errorOccurred)
				if errorOccurred {
					assert.Contains(t, sendError.Error(), "error with http request. Status code:")
				}
				assert.Equal(t, "", response)
			})

			t.Run("verifyPostError", func(t *testing.T) {
				conf.Host = "notAnIpAddress"
				SetCLIConfig(conf)
				response, sendError := makePostRequest("", "")

				assert.Equal(t, "", response)
				errorOccurred := sendError != nil
				assert.True(t, errorOccurred)
				if errorOccurred {
					assert.Contains(t, sendError.Error(), "error sending post request")
				}
			})
		})
	}

	func TestSubmitEmail()  {
		testlog.BeginTest()
		defer testlog.EndTest()
		clientConfigured = false

		var inputPayload string
		var inputPath string
		var outputResponse string
		var outputErr error

		post = func(payload string, path string) (response string, err error) {
			inputPayload = payload
			inputPath = path
			response = outputResponse
			err = outputErr
			return
		}

		testMail := `{"from":"example@gmx.com", "subject":"This is a test", "message":"The message"}`
		resultErr := SubmitEmail(testMail)

		assert.NoError(t, resultErr)
		assert.Equal(t, tstMail, inputPayload)
		assert.Equal(t, "api/receive", inputPath)
	}

	func TestSubmitEmailConnectionError(t *testing.T) {
		testlog.BeginTest()
		defer testlog.EndTest()

		clientConfigured = false
		conf := structs.CLIConfig {
			Host: defaults.CLIHost,
			Port: testPort,
			Cert: defaults.TestCertificate,
		}

		SetCLIConfig(conf)

		post = makePostRequest

		testMail := `{"from":"example@gmx.com", "subject":"This is a test", "message":"The message"}`
		submitErr := SubmitEmail(testMail)

		assert.Error(t, submitErr)
	}

	func TestServerConfig(t, *testing.T) {
		testlog.BeginTest()
		defer testlog.EndTest()

		conf := structs.CLIConfig {
			Host: "127.0.0.1",
			Port: 433,
		}

		SetCLIConfig(conf)

		assert.Equal(t, conf, cliConfig)

		conf = strings.CLIConfig {
			Host: "10.168.0.1",
			Port: 1010,
		}

		SetCLIConfig(conf)

		assert.Equal(t, conf, cliConfig)
	}

	func TestInitializeCLient(t *testing){
		testlog.BeginTest()
		defer testlog.EndTest()

		cliConfig.Cert = defaults.TestCertificate
		clientConfigured = false

		InitializeClient()

		assert.True(t, clientConfigured)
		assert.Equal(t, 5*time.Second, client.Timeout)
		assert.NotEqual(t, http.Transport{}, client.Transport)
	}

	func TestAcknowledgeEmailReception(t *testing.T) {
		testlog.BeginTest()
		defer testlog.EndTest()

		clientConfigured = false
		testMail := structs.Mail {
			ID: "IDString",
			To: "example@gmail.com",
			Subject: "Example",
			Message: "An example message",
		}

		var inputPayload string
		var inputPath string

		post = func(payload string, path string) (response string, err error) {
			inputPayload = payload
			inputPath = path
			return
		}

		acknowledgementError := AcknowledgeEmailReception(testMail)

		assert.Equal(t, `{id":"'+testMail.ID+'"}`, inputPayload)
		assert.NoError(t, acknowledgementError)
		assert.Equal(t, "api/verifyMail", inputPath)
	}

	func TestAcknowledgeEmailReceptionConnectionError(t *testing.T) {
		testlog.BeginTest()
		defer testlog.EndTest()

		clientConfigured = false
		conf := structs.CLIConfig {
			Host: defaults.CLIHost,
			Port: testPort,
			Cert: defaults.TestCertificate,
		}

		SetCLIConfig(conf)

		post = makePOSTRequest

		testMail := structs.Mail {
			ID: "IDString",
			To: "example@gmail.com",
			Subject: "Example",
			Message: "An example message",
		}

		acknowledgementError := AcknowledgeEmailReception(testMail)

		assert.Error(t, acknowledgementError)
	}
}

