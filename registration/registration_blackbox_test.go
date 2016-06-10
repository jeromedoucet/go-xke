package registration_test

import (
	"testing"
	"sync"
	"net/http/httptest"
	"net/http"
	"github.com/vil-coyote-acme/go-concurrency/commons"
	"github.com/stretchr/testify/assert"
	"github.com/vil-coyote-acme/go-xke/registration"
	"time"
)

func Test_register_should_return_no_error_on_200_response(t *testing.T) {
	// given
	wg := new(sync.WaitGroup)
	wg.Add(1)
	playerId := "test"
	ip := "127.0.0.1"

	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		wg.Done()
		var reg commons.Registration
		commons.UnmarshalRegistrationFromHttp(rq, &reg)
		assert.Equal(t, playerId, reg.PlayerId)
		assert.Equal(t, ip, reg.Ip)
		assert.Equal(t, "/registration", rq.URL.Path)
		assert.Equal(t, http.MethodPost, rq.Method)
	}))
	defer srv.Close()
	// when
	err := registration.Register(srv.URL, ip, playerId)
	assert.Nil(t, err)
	timeOut := commons.WaitTimeout(wg, time.Second * 5)
	assert.False(t, timeOut)
}

func Test_register_should_return_no_error_on_non_200_response(t *testing.T) {
	// given
	wg := new(sync.WaitGroup)
	wg.Add(1)
	playerId := "test"
	ip := "127.0.0.1"

	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		wg.Done()
		var reg commons.Registration
		commons.UnmarshalRegistrationFromHttp(rq, &reg)
		assert.Equal(t, playerId, reg.PlayerId)
		assert.Equal(t, ip, reg.Ip)
		assert.Equal(t, "/registration", rq.URL.Path)
		assert.Equal(t, http.MethodPost, rq.Method)
		rw.WriteHeader(403)
	}))
	defer srv.Close()
	// when
	err := registration.Register(srv.URL, ip, playerId)
	assert.NotNil(t, err)
	timeOut := commons.WaitTimeout(wg, time.Second * 5)
	assert.False(t, timeOut)
}

func Test_register_should_return_error_when_connection_issue(t *testing.T) {
	// given
	playerId := "test"
	ip := "127.0.0.1"
	// when
	err := registration.Register("http://toto_titi_tata", ip, playerId)
	assert.NotNil(t, err)
}