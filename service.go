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
	//user
	Login(cmd *LoginCommand) (*auth.TokenDetails, error)
	Register(cmd *RegisterCommand) (*tik_lib.User, error)
	//files
	CreateFile(cmd *CreateFileCommand) (*tik_lib.File, error)
	UpdateFile(cmd *UpdateFileCommand) (*tik_lib.File, error)
	GetFileById(cmd *GetFileByIdCommand) (*tik_lib.File, error)
	ListFiles(cmd *ListFilesCommand) ([]tik_lib.File, error)
	DeleteFile(cmd *DeleteFileCommand) error
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

func (s *service) Register(cmd *RegisterCommand) (*tik_lib.User, error) {
	if cmd.PhoneNumber == "" {
		return nil, ErrPhoneNumberNotProvided
	} else if cmd.Password == "" {
		return nil, ErrPasswordNotProvided
	}

	_, err := s.amqpRequests.GetUserByPhoneNumber(&tik_lib.GetUserByPhoneNumberCommand{PhoneNumber: cmd.PhoneNumber})
	if err == nil {
		return nil, ErrUserAlreadyExistByPhone
	}
	cmd.Password, err = setdata_common.HashPassword(cmd.Password)
	if err != nil {
		return nil, err
	}
	return s.amqpRequests.CreateUser(&tik_lib.CreateUserCommand{User: &cmd.User})
}

func (s *service) CreateFile(cmd *CreateFileCommand) (*tik_lib.File, error) {
	if cmd.FileType != "" {
		if !tik_lib.IsFileTypeExist(cmd.FileType.ToString()) {
			return nil, ErrFileTypeError
		}
	} else {

	}
	return s.amqpRequests.CreateFile(&tik_lib.CreateFileCommand{File: &cmd.File})
}

func (s *service) UpdateFile(cmd *UpdateFileCommand) (*tik_lib.File, error) {
	if *cmd.FileType != "" {
		if !tik_lib.IsFileTypeExist(cmd.FileType.ToString()) {
			return nil, ErrFileTypeError
		}
	} else {

	}
	return s.amqpRequests.UpdateFile(&tik_lib.UpdateFileCommand{FileUpdate: &cmd.FileUpdate})
}

func (s *service) GetFileById(cmd *GetFileByIdCommand) (*tik_lib.File, error) {
	return s.amqpRequests.GetFileById(&tik_lib.GetFileByIdCommand{Id: cmd.Id})
}

func (s *service) ListFiles(cmd *ListFilesCommand) ([]tik_lib.File, error) {
	return s.amqpRequests.ListFiles(&tik_lib.ListFilesCommand{})
}

func (s *service) DeleteFile(cmd *DeleteFileCommand) error {
	return s.amqpRequests.DeleteFile(&tik_lib.DeleteFileCommand{Id: cmd.Id})
}
