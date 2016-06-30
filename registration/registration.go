package registration

import (
	"fmt"
	"github.com/vil-coyote-acme/go-concurrency/commons"
	"log"
)

func Register(clientUrl string, ourIp string, playerId string) error {
	/**TODO 1. use json library to marshall an object of type registration*/
	//a. create an object type registration
	reg := commons.Registration{}
	log.Println(fmt.Sprintf("server | will do a registration: %s", reg))

	//b. use json library to marshall the object

	//c. if error (in marshall) return the error

	/**TODO 2. send a post to the clientRegistrationURL with the marshalled registration, to register the client*/
	clientRegistrationURL := clientUrl + "/registration"
	fmt.Println(clientRegistrationURL)
	// a. use http library to send a post of type application/json with the response of the previous step

	// b.if error (in post) , return the error, example:
	//if httpErr != nil {
	//	return httpErr
	//}
	// c. if the StatusCode of the response is != 200 return new error
	// hint: use error library combined with fmt to create a new error and return it

	return nil
}
