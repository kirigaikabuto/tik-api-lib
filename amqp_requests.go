package tik_api_lib

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	tik_lib "github.com/kirigaikabuto/tik-lib"
)

const (
	getUserByPhoneNumber = "user.getByPhoneNumber"
)

type AmqpRequests struct {
	clt amqp.Client
}

func NewAmqpRequests(clt amqp.Client) AmqpRequests {
	return AmqpRequests{clt: clt}
}

func (a *AmqpRequests) GetUserByPhoneNumber(cmd *tik_lib.GetUserByPhoneNumberCommand) (*tik_lib.User, error) {
	user := &tik_lib.User{}
	jsonData, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	resp, err := a.clt.Call(getUserByPhoneNumber, amqp.Message{Body: jsonData})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
