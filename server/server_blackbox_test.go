package server_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/vil-coyote-acme/go-concurrency/commons"
	"github.com/vil-coyote-acme/go-xke/server"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
	"time"
)

var (
	playerId string = "player"
	brtPath  string = "/orders/"
	orderId  int    = 1
	cbkPath  string = "/" + playerId + "/bill/" + strconv.Itoa(orderId)
)

// nominal functional test
func Test_server_should_handle_new_order_call_bartender_and_answer_on_callback_addr(t *testing.T) {
	// given
	// wait group used to wait for async call
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// mock client callback
	clientCallBack := mockClientCallback(t, wg)
	defer clientCallBack.Close()

	// create & marshall order
	order := commons.Order{Id: orderId, Quantity: 5, Type: commons.Beer, CallBackUrl: clientCallBack.URL + cbkPath}
	body, _ := json.Marshal(order)

	// mock bartender api
	bartender := mockBartender(t, wg, &order)
	defer bartender.Close()
	srv := server.NewServer(playerId, bartender.URL)

	// when
	startHttpServeAsync(srv)
	time.Sleep(time.Millisecond * 100)
	resp, err := http.Post("http://127.0.0.1:4242/orders", "application/json", bytes.NewBuffer(body))

	// then
	assert.Nil(t, err)
	if err == nil {
		assert.Equal(t, resp.StatusCode, 200)
		assert.False(t, commons.WaitTimeout(wg, time.Second*5))
	}
}

func mockClientCallback(t *testing.T, wg *sync.WaitGroup) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		wg.Done()
		assert.Equal(t, http.MethodGet, rq.Method)
		assert.Equal(t, cbkPath, rq.URL.Path)
	}))
}

func mockBartender(t *testing.T, wg *sync.WaitGroup, order *commons.Order) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		wg.Done()
		assertOnBartenderCall(rw, rq, order, t)
		rw.WriteHeader(200)
	}))
}

func assertOnBartenderCall(rw http.ResponseWriter, rq *http.Request, order *commons.Order, t *testing.T) {
	var brtOrder commons.Order
	buf := make([]byte, rq.ContentLength)
	io.ReadFull(rq.Body, buf)
	json.Unmarshal(buf, &brtOrder)
	// assert on query
	assert.Equal(t, http.MethodPost, rq.Method, "http method assert on bartender")
	assert.Equal(t, brtPath, rq.URL.Path, "path assert on bartender")
	// assert on body
	assert.Equal(t, order.Id, brtOrder.Id, "order id assert on bartender")
	assert.Equal(t, order.Quantity, brtOrder.Quantity, "quantity assert on bartender")
	assert.Equal(t, order.Type, brtOrder.Type, "beverage type on bartender")
}

func Test_server_should_answer_200_to_health_check(t *testing.T) {
	// given
	srv := server.NewServer(playerId, "http:localhost:1234")
	// when
	startHttpServeAsync(srv)
	time.Sleep(time.Millisecond * 100)
	resp, err := http.Get("http://127.0.0.1:4242/status")
	// then
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

func startHttpServeAsync(srv *server.Server) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		wg.Done()
		srv.Start()
	}()
	wg.Wait()
}
