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
	ErrPhoneNumberNotProvided         = com.NewMiddleError(errors.New("phone number not provided"), 400, 5)
	ErrPasswordNotProvided            = com.NewMiddleError(errors.New("password not provided"), 400, 6)
	ErrUserAlreadyExistByPhone        = com.NewMiddleError(errors.New("user with that phone number already exist"), 400, 7)

)
