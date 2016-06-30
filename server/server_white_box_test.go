package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vil-coyote-acme/go-concurrency/commons"
	"net/http/httptest"
)


/*
 * TESTS THAT MUST BECOME GREEN ^^
 */


func Test_postOrder_should_fail(t *testing.T) {
	// given
	var order commons.Order
	server := new(Server)
	server.playerId = "new id"
	// when
	_, err := server.postOrder(order)
	// then
	assert.NotNil(t, err)
}

func Test_postOrder_should_do_without_error(t *testing.T) {
	// given
	server := new(Server)
	server.playerId = "new id"
	order := commons.Order{Id: 1, Quantity: 4, CallBackUrl: "111.111.111.111"}
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		var recievedOrder commons.Order
		commons.UnmarshalOrderFromHttp(rq, &recievedOrder)
		assert.Equal(t, order.Id, recievedOrder.Id)
		assert.Equal(t, server.playerId, recievedOrder.PlayerId)
		assert.Equal(t, order.CallBackUrl, recievedOrder.CallBackUrl)
		assert.Equal(t, order.Quantity, recievedOrder.Quantity)
		rw.WriteHeader(200)
	}))
	server.bartenderUrl = srv.URL
	// when
	resp, err := server.postOrder(order)
	// then
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func Test_getDataFromCallback_should_fail_with_error_in_url(t *testing.T) {
	// given
	var order commons.Order
	// when
	err := getDataFromCallback(order)
	// then
	assert.NotNil(t, err)
}

func Test_getDataFromCallback_should_not_fail(t *testing.T) {
	// given
	var order commons.Order
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		assert.Equal(t, http.MethodGet, rq.Method)
		rw.WriteHeader(200)
	}))
	order.CallBackUrl = srv.URL
	// when
	err := getDataFromCallback(order)
	// then
	assert.Nil(t, err)
}










/*
 * OTHERS TESTS
 */

func Test_unmarshallOrder_should_unmarshal_without_error(t *testing.T) {
	// given
	expectedOrder := commons.Order{Id: 1, Quantity: 5, Type: commons.Beer, CallBackUrl: "http://callback.com/money"}
	order := new(commons.Order)
	body, _ := json.Marshal(expectedOrder)
	var req http.Request
	req.Body = nopCloser{bytes.NewBuffer(body)}
	req.ContentLength = int64(len(body))
	// when
	err := unmarshalOrderFromHttp(&req, order)
	// then
	assert.Nil(t, err)
	assert.Equal(t, expectedOrder, *order)
}

func Test_unmarshallOrder_should_unmarshal_with_error(t *testing.T) {
	// given
	order := new(commons.Order)
	var req http.Request
	req.Body = nopCloser{bytes.NewBuffer(make([]byte, 0))}
	req.ContentLength = int64(0)
	// when
	err := unmarshalOrderFromHttp(&req, order)
	// then
	assert.NotNil(t, err)
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}
