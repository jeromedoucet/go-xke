package registration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/vil-coyote-acme/go-concurrency/commons"
)

func Register(clientUrl string, ourIp string, playerId string) error {
	/**TODO 1. use json library to marshall an object of type registration*/
	//a. create an object type registration
	registration := commons.Registration{PlayerId: playerId, Ip: ourIp}
	//b. use json library to marshall the object
	body, marshErr := json.Marshal(registration)
	//c. if error (in marshall) return the error
	if marshErr != nil {
		return marshErr
	}
	/**TODO 2. send a post to the clientRegistrationURL with the marshalled registration, to register the client*/
	clientRegistrationURL := clientUrl + "/registration"
	// a. use http library to send a post of type application/json with the response of the previous step
	res, httpErr := http.Post(clientRegistrationURL, "application/json", bytes.NewBuffer(body))
	// b.if error (in post) , return the error
	if httpErr != nil {
		return httpErr
	}
	// c. if the StatusCode of the response is != 200 return new error
	// hint: use error library combined with fmt to create a new error and return it
	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("bad http code. Expected 200, got %d", res.StatusCode))
	}
	return nil
}
