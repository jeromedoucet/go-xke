package registration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vil-coyote-acme/go-concurrency/commons"
	"net/http"
)

func Register(clientUrl string, ourIp string, playerId string) error {
	body, marshErr := json.Marshal(commons.Registration{PlayerId: playerId, Ip: ourIp})
	if marshErr != nil {
		return marshErr
	}
	res, httpErr := http.Post(clientUrl+"/registration", "application/json", bytes.NewBuffer(body))
	if httpErr != nil {
		return httpErr
	}
	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("bad http code. Expected 200, got %d", res.StatusCode))
	}
	return nil
}
