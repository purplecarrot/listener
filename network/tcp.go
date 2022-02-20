package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
)

func processPOSTdata(contentType string, postdata string) (response string, err error) {

	switch contentType {
	case "application/json":
		var result map[string]interface{}
		e := json.Unmarshal([]byte(postdata), &result)
		if e != nil {
			return "", fmt.Errorf("invalid JSON")
		}
		return postdata, nil
	default:
		if len(postdata) > 0 {
			response = fmt.Sprintf("<p>You said:<br>\n%s</p>", postdata)
		} else {
			response = fmt.Sprintf("<p>Listener here, nothing posted</p>")
		}
	}
	return response, nil
}

func makeResponse(contentType string, data string) (response string, err error) {

	switch contentType {
	case "application/json":
		json_string := `{"name": "mark", "message": "test message"}`

		var result map[string]interface{}
		e := json.Unmarshal([]byte(json_string), &result)
		if e != nil {
			return "", fmt.Errorf("can't marshal json %s: %v", response, e)
		}
		response = fmt.Sprintf("{\"received\": %s}\n", json_string)
	default:
		response = fmt.Sprintf("<head><title>Listener</title></head><body>%s</body>\n", data)
	}

	return response, nil
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-type")
	// for our purposes we always use contentType
	// accept := r.Header.Get("Accept")

	_, port, _ := net.SplitHostPort(r.Context().Value(http.LocalAddrContextKey).(net.Addr).String())
	connID := fmt.Sprintf("%s/tcp", port)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("bad request from %s", "test")
		http.Error(w, "invalid body", http.StatusBadRequest)
	} else {
		log.Printf("[%s] >%s %s %s %s (%s) %s", connID, r.RemoteAddr, r.Method, r.URL, r.Proto, contentType, string(body))
	}

	postdata, err := processPOSTdata(contentType, string(body))
	if err != nil {
		log.Printf("[%s] <%s Bad Request - %s", connID, r.RemoteAddr, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := makeResponse(contentType, postdata)
	if err != nil {
		response = "Could not process post data"
	}

	log.Printf("[%s] <%s responding (%s): %s", connID, r.RemoteAddr, contentType, response)
	w.Header().Set("Server", "Listener")
	w.Write([]byte(response))

}

func TCPListen(port int) {

	log.Printf("[%d/tcp] listening on TCP port %d for HTTP connections", port, port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Printf("[%d/tcp] error binding to port %d - %v ", port, port, err)
	}

}
