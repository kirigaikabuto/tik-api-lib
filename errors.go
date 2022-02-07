package tik_api_lib

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrNoUserByPhoneNumber            = com.NewMiddleError(errors.New("no user by this phone number"), 400, 1)
	ErrNoUserByPhoneNumberAndPassword = com.NewMiddleError(errors.New("no user by this phone number and password"), 400, 2)
	ErrFileTypeError                  = com.NewMiddleError(errors.New("not correct file type, it should be video or image"), 400, 3)
	ErrFileIdNotProvided              = com.NewMiddleError(errors.New("file id not provided"), 400, 4)
)
