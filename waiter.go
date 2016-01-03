package main
import (
	"log"
	"flag"
	"sync"
	"github.com/nsqio/go-nsq"
	mes "github.com/jeromedoucet/go-concurrency/messages"
	"bytes"
	"net/http"
	"strconv"
	"io"
)

var (
	host = "51.254.216.243"
	lookupaddr string = host + ":4161"
	bartenderAddr string = host + ":3000"
	deliverAddr string = host + ":3002"

	// TODO replace by your player id
	playerId string = "foo"
)


func main() {
	topic := flag.String("topic", "orders#ephemeral", "the topic to subscribe on")
	channel := flag.String("channel", "chan#ephemeral", "the channel to use to consume topic message")// to do remove and make it empty
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	initListener(*topic, *channel)
	wg.Wait()
}

type Handler struct {
}

func initListener(topic, channel string) {
	conf := nsq.NewConfig()
	cons, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		log.Panicf("error when trying to create a consumer for topic : %v and channel : %v", topic, channel)
	}
	cons.AddConcurrentHandlers(new(Handler), 5)
	cons.ConnectToNSQLookupd(lookupaddr)
}


func (* Handler) HandleMessage(message *nsq.Message) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
			return
		}
	}()
	log.Printf("receive a message %v", message)
	order := unmarshallMes(message)
	resB := askBartender(askBartenderUrl(bartenderAddr, &order))
	log.Printf("receive a response from bartender %f", resB)
	if resB == 200 {
		deliver(deliverUrl(deliverAddr), createDeliverBody(&order))
	}
	return
}

func unmarshallMes (message *nsq.Message) mes.Order {
	var order mes.Order
	// TODO use the json package to get a decoder in order to unmarshall the message body. You will need a Buffer from bytes package too.
	// tips : the decoder has a decode function that will need a pointer on order
	return order
}

func askBartender(url string) (statusCode int) {
	resp, err := http.Post(url, "text/plain", bytes.NewBufferString(""))
	resp.Body.Close()
	if err != nil {
		log.Panicf("error when trying to send post on %v ", url)
	} else {
		statusCode = resp.StatusCode
	}
	return
}

func askBartenderUrl(host string, order *mes.Order) string {
	// TODO return the expected value
	return "" + strconv.Itoa(int(order.Id))
}

func createDeliverBody(order *mes.Order) []byte {
	orderCheck := mes.NewOrderCheck(order.Id, playerId)
	// TODO use the json package in order to marshall the orderCheck
	var b []byte
	var err error
	if err != nil {
		log.Panicf("error when trying serialise %s ", orderCheck)
	}
	return b
}

func deliverUrl(host string) string {
	// TODO return the expected value
	return ""
}

func deliver(url string, body []byte) {
	// TODO use the Buffer type in the bytes package in order to use it in rest post
	// tips : the Buffer type implements io.reader
	var buffer io.Reader
	resp, err := http.Post(url, "application/json", buffer)
	resp.Body.Close()
	if err != nil {
		log.Panicf("error when trying post on %v ", url)
	}
}