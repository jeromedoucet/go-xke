package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/vil-coyote-acme/go-concurrency/commons"
)

const bartenderPath string = "/orders"

type Server struct {
	playerId     string
	bartenderUrl string
	mux          *http.ServeMux
	orderChan    chan commons.Order
}


/*
 * ########################################################################################
 * ########################################################################################
 * #########  			     THE CODE YOU MUST CARE 		        ###########
 * ########################################################################################
 * ########################################################################################
 */

func NewServer(playerId string, bartenderUrl string) (s *Server) {
	s = new(Server)
	s.bartenderUrl = bartenderUrl
	s.playerId = playerId
	/*
	 * TODO : initialisez votre channel ici !
	 */
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/orders", s.handleOrder)
	return s
}

func (s *Server) Start(url string) {
	/*
	 * TODO : vous devez demarrer votre consommateur de commande ici !
	 */
	err := http.ListenAndServe(url, s.mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) consumeOrder(){
	for {
		/*
		 * TODO : c'est ici que les commandes doivent etre traite maintenant !
		 */
	}
}

func (s *Server) handleOrder(w http.ResponseWriter, r *http.Request) {
	// first step : unmarshal the incoming order
	var order commons.Order
	unMarshallErr := unmarshalOrderFromHttp(r, &order)
	if unMarshallErr != nil {
		log.Println(unMarshallErr.Error())
		return
	}
	/*
	 * TODO : Plutot que de continuer executer le code qui suit, vous devrier deleguer a un consommateur de commande maintenat!
	 */
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
	getDataFromCallback(order)
	w.WriteHeader(200)
}




/*
 * ########################################################################################
 * ########################################################################################
 * #########  			SOME CODE YOU DON'T NEED TO CHANGE		###########
 * ########################################################################################
 * ########################################################################################
 */




func (s *Server) postOrder(order commons.Order) (r *http.Response, err error) {
	order.PlayerId = s.playerId
	buf, marshalErr := json.Marshal(order)
	if marshalErr != nil {
		log.Println(marshalErr.Error())
		return
	}
	bartenderUrl := s.bartenderUrl + bartenderPath
	return http.Post(bartenderUrl, "application/json", bytes.NewBuffer(buf))
}

func unmarshalOrderFromHttp(r *http.Request, order *commons.Order) (err error) {
	buf := make([]byte, r.ContentLength)
	// a. use io library to read message from the body of the request
	io.ReadFull(r.Body, buf)
	// b. use json library to unmarsall the message in a "commons.Order"
	err = json.Unmarshal(buf, &order)
	// c. return
	return
}

func getDataFromCallback(order commons.Order) (err error) {
	paymentRes, paymentErr := http.Get(order.CallBackUrl)
	if paymentErr != nil {
		log.Printf("get an error when calling payment api : %s", paymentErr.Error())
		return paymentErr
	}
	if paymentRes != nil && paymentRes.StatusCode != 200 {
		log.Printf("get a non 200 response when calling payment api : %s", paymentRes.Status)
		return fmt.Errorf("get a non 200 response when calling payment api : %s", paymentRes.Status)
	}
	return nil
}



