# CleanURIAPI
 Shorten URL API

This is a command-line tool written in Go that takes a URL as input and returns a shortened version of that URL using the CleanURI API.

## Installation

To install and use this tool, follow these steps:

1. Make sure you have Go installed on your machine.
2. Clone this repository or download the `main.go` file.
3. Open a terminal or command prompt and navigate to the directory containing the `main.go` file.
4. Run the following command to build the executable:


   go build -o shorten-url main.go

The above command will create an executable file named shorten-url in the current directory.
Usage

After building the executable, you can use it as follows:

./shorten-url [-u <url>]

The tool accepts one optional command-line flag:

-u, --url: Specifies the URL to be shortened. If not provided, the default URL is set to "https://google.com".

Example usage:

./shorten-url -u https://example.com

Dependencies

This tool uses the following Go packages as dependencies:

bytes
flag
io/ioutil
log
net/http
os
strings
These packages are part of the Go standard library and do not require any additional installation.

License

This project is licensed under the MIT License.
