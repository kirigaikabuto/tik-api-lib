package tik_api_lib

import tik_lib "github.com/kirigaikabuto/tik-lib"

type LoginCommand struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (cmd *LoginCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).Login(cmd)
}

type RegisterCommand struct {
	tik_lib.User
}

func (cmd *RegisterCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).Register(cmd)
}

type CreateFileCommand struct {
	tik_lib.File
	UserId string `json:"-"`
}

func (cmd *CreateFileCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).CreateFile(cmd)
}

type UpdateFileCommand struct {
	tik_lib.FileUpdate
	UserId string `json:"-"`
}

func (cmd *UpdateFileCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).UpdateFile(cmd)
}

type GetFileByIdCommand struct {
	Id     string `json:"id"`
	UserId string `json:"-"`
}

func (cmd *GetFileByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).GetFileById(cmd)
}

type ListFilesCommand struct {
	UserId string `json:"-"`
}

func (cmd *ListFilesCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).ListFiles(cmd)
}

type DeleteFileCommand struct {
	Id     string `json:"id"`
	UserId string `json:"-"`
}

func (cmd *DeleteFileCommand) Exec(svc interface{}) (interface{}, error) {
	return nil, svc.(Service).DeleteFile(cmd)
}
