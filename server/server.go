package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vil-coyote-acme/go-concurrency/commons"
)

const bartenderPath string = "/orders"

type Server struct {
	playerId     string
	bartenderUrl string
	mux          *http.ServeMux
}

func (s *Server) handleOrder(w http.ResponseWriter, r *http.Request) {
	// first step : unmarshal the incoming order
	var order commons.Order
	/** TODO 1. complete the method unmarshalOrderFromHttp() (at the end), to unmarshall
	  the http entry of message**/
	unMarshallErr := unmarshalOrderFromHttp(r, &order)
	if unMarshallErr != nil {
		log.Println(unMarshallErr.Error())
		return
	}
	log.Printf("receive one order : %s", order)

	// second step, send the order to the bartender
	/** TODO 2. assign the playerId from the server to the order and marshall order to send
	to the bartender. complete the method postOrder**/
	res, err := s.postOrder(order)

	if err != nil {
		log.Printf("error when calling bartender api : %s", err)
		return
	}
	if res.StatusCode != 200 {
		log.Printf("get a non 200 response when calling bartender api : %s", res.Status)
		return
	}

	// third step, if all is right, get your money back !
	/**TODO 3. Get payment information from the CallBackUrl in the order
	Complete the method*/
	getDataFromCallback(order)
	w.WriteHeader(200)
}

func (s *Server) Start(url string) {
	err := http.ListenAndServe(url, s.mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func NewServer(playerId string, bartenderUrl string) (s *Server) {
	s = new(Server)
	s.bartenderUrl = bartenderUrl
	s.playerId = playerId
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/orders", s.handleOrder)
	return s
}

func unmarshalOrderFromHttp(r *http.Request, order *commons.Order) (err error) {
	/** TODO  1.*/
	// a. create a variable of type []byte to contain the body of the request.
	// Hint: buf := make([]byte, r.ContentLength)
	// b. use io library to read the message from the body of the request
	// and save it in the variable of step a. buf

	// c. use json library to unmarsall the message in the order

	// d. return
	return
}

func (s *Server) postOrder(order commons.Order) (r *http.Response, err error) {
	/** TODO 2.*/
	//a. assing playerId to the order

	//b. import and use json library to marshall order in a variable.

	//c. if error in marshalling, log it and return

	bartenderUrl := s.bartenderUrl + bartenderPath
	fmt.Println(bartenderUrl)
	// d. use http post to send to the bartenderUrl, as application/json, the marshalled order
	// you may want to explore la function bytes.NewBuffer. Return the response and the error if it exists
	return nil, nil
}

func getDataFromCallback(order commons.Order) (err error) {
	/**TODO 3.*/
	//a. use the method Get from http to get my payment status

	//b. if error, log it and return

	//c. if StatusCode from the response of the server != 200, log it  and return an error
	// hint: explore fmt.Errorf

	return nil
}
