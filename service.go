package tik_api_lib

import (
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"github.com/kirigaikabuto/tik-api-lib/auth"
	tik_lib "github.com/kirigaikabuto/tik-lib"
)

type service struct {
	amqpRequests AmqpRequests
	tokenStore   auth.TokenStore
}

type Service interface {
	Login(cmd *LoginCommand) (*auth.TokenDetails, error)
}

func NewService(amqpRequests AmqpRequests, tknStore auth.TokenStore) Service {
	return &service{amqpRequests: amqpRequests, tokenStore: tknStore}
}

func (s *service) Login(cmd *LoginCommand) (*auth.TokenDetails, error) {
	user, err := s.amqpRequests.GetUserByPhoneNumber(&tik_lib.GetUserByPhoneNumberCommand{PhoneNumber: cmd.PhoneNumber})
	if err != nil {
		return nil, err
	}
	if !setdata_common.CheckPasswordHash(cmd.Password, user.Password) {
		return nil, ErrNoUserByPhoneNumberAndPassword
	}
	token, err := s.tokenStore.CreateToken(user.Id)
	if err != nil {
		return nil, err
	}
	return token, nil
}
