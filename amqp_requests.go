package tik_api_lib

import (
	"encoding/json"
	"fmt"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	tik_lib "github.com/kirigaikabuto/tik-lib"
)

const (
	getUserByPhoneNumber = "user.getByPhoneNumber"
	createUser           = "user.create"

	createFile  = "file.create"
	getFileByID = "file.getById"
	updateFile  = "file.update"
	deleteFile  = "file.delete"
	listFiles   = "file.list"
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
	fmt.Println(err.(setdata_common.MiddleError))
	if err == nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *AmqpRequests) CreateUser(cmd *tik_lib.CreateUserCommand) (*tik_lib.User, error) {
	user := &tik_lib.User{}
	jsonData, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	resp, err := a.clt.Call(createUser, amqp.Message{Body: jsonData})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *AmqpRequests) CreateFile(cmd *tik_lib.CreateFileCommand) (*tik_lib.File, error) {
	file := &tik_lib.File{}
	jsonData, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	resp, err := a.clt.Call(createFile, amqp.Message{Body: jsonData})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body, &file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (a *AmqpRequests) UpdateFile(cmd *tik_lib.UpdateFileCommand) (*tik_lib.File, error) {
	file := &tik_lib.File{}
	jsonData, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	resp, err := a.clt.Call(updateFile, amqp.Message{Body: jsonData})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body, &file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (a *AmqpRequests) GetFileById(cmd *tik_lib.GetFileByIdCommand) (*tik_lib.File, error) {
	file := &tik_lib.File{}
	jsonData, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	resp, err := a.clt.Call(getFileByID, amqp.Message{Body: jsonData})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body, &file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (a *AmqpRequests) ListFiles(cmd *tik_lib.ListFilesCommand) ([]tik_lib.File, error) {
	var files []tik_lib.File
	jsonData, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	resp, err := a.clt.Call(listFiles, amqp.Message{Body: jsonData})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body, &files)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (a *AmqpRequests) DeleteFile(cmd *tik_lib.DeleteFileCommand) error {
	jsonData, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	_, err = a.clt.Call(deleteFile, amqp.Message{Body: jsonData})
	if err != nil {
		return err
	}
	return nil
}
