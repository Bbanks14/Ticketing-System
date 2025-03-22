package client

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mortenterhart/trivial-tickets/structs"
)

// cliConfig holds the currenr command line tool
var cliConfig structs.CLIConfig

// client is an instance of the http.Client
// Makes GET and POST requests to the server
var client = http.Client{}

// clientConfigured indicates whether the HTTP client has already been configured i.e. if it is already initialized
var clientConfigured bool

// get and post are replaceacble functions for the http.Client.Get and http.Client.Post functions
var get = makeGETRequest
var post = makePOSTRequest

// FetchEmails sends a GET request to the path `api/fetchMails` and returns the response
func FetchEmails() (mails map[string]structs.Mail, err error) {
	response, err := get("api/fetchMails")
	if err != nil {
		err = fmt.Errorf("Failed to fetch emails: %v", err)
		return
	}

	err = json.Unmarshal([]byte(response), &mails)
	if err != nil {
		err = fmt.Errorf("Error occured while unmarshaling the JSON: %v", err)
		return
	}

	return
}

// AcknowledgeEmailReception sends a POST request with an `ID` to the received email of the server
func AcknowledgeEmailReception(mail structs.Mail) (err error) {
	jsonID := `{"id":"` + mail.ID + `"}`
	_, err = post(jsonID, "api/verifyMail")

	if err != nil {
		err = fmt.Errorf("email acknowldgement failed: %v", err)
	}

	return
}

// SubmitEmail takes a structs.Mail and sends it to the server as JSON per POST request
func SubmitEmail(mail string) (err error) {
	resp, err := post(mail, "api/receive")
	fmt.Println(resp)

	return

}

// makeGETRequest requests a path string
// and sends a GET request to the "path" on the server specified in the CLIConfig
func makeGETRequest(path string) (response string, err error) {
	if !clientConfigured {
		initializeClient()
	}

	url := "https://" + cliConfig.Host + ":" + strconv.Itoa(int(cliConfig.Port)) + "/" + path

	resp, err := client.Get(url)

	if err != nil {
		err = fmt.Errorf("Error sending get request: %v", err)
		return
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("Error status code: %v %s", resp.Status, string(responseData))
		return
	}

	return string(responseData), nil
}

// makePOSTRequest takes a payload and a path string and sends a
// POST request to the "path" on the server specified in the CLIConfig
func makePOSTRequest(payload string, path string) (response string, err error) {
	if !clientConfigured {
		initializeClient()
	}

	reader := strings.NewReader(payload)
	url := fmt.Sprintf("https://%s:%d/%s", cliConfig.Host, cliConfig.Port, path)

	resp, err := client.Post(url, "application/json", reader)

	if err != nil {
		return "", fmt.Errorf("Error sending post request: %v", err)
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error status code: %v %s", resp.Status, string(responseData))
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error with https request. Status code: %v %s", resp.Status, string(responseData))
	}
	response = string(responseData)

	return
}

// initializeClient initializes the HTTP client with the server's certificate when using tls
func initializeClient() {
	configFilePath, _ := filepath.Abs(cliConfig.Cert)
	cert, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)

	client = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: certPool,
			},
		},
		Timeout: 5 * time.Second,
	}
	clientConfigured = true
}

// SetCLIConfig sets the current CLIConfig
func SetCLIConfig(config structs.CLIConfig) {
	cliConfig = config
}
