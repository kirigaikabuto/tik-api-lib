package tik_api_lib

import (
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"github.com/kirigaikabuto/tik-api-lib/auth"
	tik_lib "github.com/kirigaikabuto/tik-lib"
	"io/ioutil"
	"os"
)

type service struct {
	amqpRequests AmqpRequests
	tokenStore   auth.TokenStore
	s3           S3Uploader
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
	UploadFile(cmd *UploadFileCommand) (*tik_lib.File, error)
}

func NewService(amqpRequests AmqpRequests, tknStore auth.TokenStore, s3 S3Uploader) Service {
	return &service{amqpRequests: amqpRequests, tokenStore: tknStore, s3: s3}
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
	if err != nil && err.Error() != tik_lib.ErrUserNotFound.Error() {
		return nil, err
	} else if err == nil {
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

func (s *service) UploadFile(cmd *UploadFileCommand) (*tik_lib.File, error) {
	folderCreateDir := "./videos/"
	file, err := s.amqpRequests.GetFileById(&tik_lib.GetFileByIdCommand{Id: cmd.Id})
	if err != nil {
		return nil, err
	}
	videoFolderName := "video_" + file.Id + "/"
	videoFullPath := folderCreateDir + videoFolderName
	err = os.Mkdir(videoFullPath, 0700)
	if err != nil {
		return nil, err
	}
	hlsFolder := videoFullPath + "hls/"
	err = os.Mkdir(hlsFolder, 0700)
	if err != nil {
		return nil, err
	}
	filePath := videoFolderName + cmd.Name + "." + cmd.Type
	err = ioutil.WriteFile(filePath, cmd.File.Bytes(), 0700)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
