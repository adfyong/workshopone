// Command server provides a basic RESTesque API for participants to practice
// writing simple Go programs against.
//
// Basic HTTP API sketch
//   GET  /task/                - Provides documentation
//   GET  /task/?name=[name]    - Gets a new task token
//   POST /task/[token]         - Post form data in reply
//   GET  /task/square/         - Gets documentation
//   GET  /task/square/[token]  - Gets a number to square
//   POST /task/square/[token]  - Post your squared number
//   GET  /task/frequency/[token]  - Gets a paragraph to analyze
//   POST /task/frequency/[token]  - Post your top three words
//
//   GET  /status         - HTML showing status
//   GET  /status/updates - Websocket providing updates to status
//
package main

import (
	"fmt"
	"net/http"
)

func writeDocumentation(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(wr, "Documentation!")
}
