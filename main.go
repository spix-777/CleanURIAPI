package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Loggers
var (
	BannerLogger *log.Logger
	InfoLogger   *log.Logger
	WarnLogger   *log.Logger
)

// Initialize loggers to stdout
func init() {
	BannerLogger = log.New(os.Stdout, "", 0)
	InfoLogger = log.New(os.Stdout, " [ + ] ", 0)
	WarnLogger = log.New(os.Stdout, " [ ! ] ", 0)

}

func main() {
	var urlStr string
	flag.StringVar(&urlStr, "u", "https://google.com", "Input URL")
	flag.Parse()

	banner()

	// Check if the URL starts with "https"
	if strings.Index(urlStr, "https") == -1 {
		urlStr = "https://" + urlStr
	}

	// Define the URL you want to send the POST request to
	url := "https://cleanuri.com/api/v1/shorten"

	// Create a JSON payload to send in the request body
	urlInput := urlStr
	payload := []byte(`{"url":"` + urlInput + `"}`)

	// Create a new HTTP request with the POST method and the payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		WarnLogger.Fatalln("Error creating request:", err)
		return
	}

	// Set the Content-Type header for the request
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		WarnLogger.Fatalln("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		WarnLogger.Fatalln("Error reading response:", err)
		return
	}

	// Check if the response body contains "error"
	if strings.Contains(string(body), "error") {
		WarnLogger.Fatalln("Error:", string(body))
		return
	}

	// Print the response body to human readable format
	out := string(body)
	index := strings.Index(out, "https")
	out = out[index:]
	out = strings.Replace(out, "\"}", "", -1)
	out = strings.Replace(out, "\\", "", -1)

	// Print the response body
	InfoLogger.Println("Response:", out)
}

func banner() {
	banner := `                                             
     _                     _    _____ _____ _____ 
 ___| |___ ___ ___ _ _ ___|_|  |  _  |  _  |     |
|  _| | -_| .'|   | | |  _| |  |     |   __|-   -|
|___|_|___|__,|_|_|___|_| |_|  |__|__|__|  |_____|
													 
	`
	BannerLogger.Println(banner)
}
