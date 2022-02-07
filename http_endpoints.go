package tik_api_lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"net/http"
)

type HttpEndpoints interface {
	MakeLoginEndpoint() gin.HandlerFunc

	MakeCreateFileEndpoint() gin.HandlerFunc
	MakeListFilesEndpoint() gin.HandlerFunc
	MakeGetFileByIdEndpoint() gin.HandlerFunc
	MakeUpdateFileEndpoint() gin.HandlerFunc
	MakeDeleteFileEndpoint() gin.HandlerFunc
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeLoginEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &LoginCommand{}
		err := context.ShouldBindJSON(&cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeCreateFileEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &CreateFileCommand{}
		userIdVal, ok := context.Get("user_id")
		if !ok {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(errors.New("no user id in context")))
			return
		}
		cmd.UserId = userIdVal.(string)
		fmt.Println(cmd.UserId)
		err := context.ShouldBindJSON(&cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusCreated, resp)
	}
}

func (h *httpEndpoints) MakeListFilesEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &ListFilesCommand{}
		userIdVal, ok := context.Get("user_id")
		if !ok {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(errors.New("no user id in context")))
			return
		}
		cmd.UserId = userIdVal.(string)
		fmt.Println(cmd.UserId)
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeGetFileByIdEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &GetFileByIdCommand{}
		userIdVal, ok := context.Get("user_id")
		if !ok {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(errors.New("no user id in context")))
			return
		}
		cmd.UserId = userIdVal.(string)
		fileId := context.Request.URL.Query().Get("id")
		if fileId == ""{
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(ErrFileIdNotProvided))
			return
		}
		cmd.Id = fileId
		fmt.Println(cmd.UserId)
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeUpdateFileEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &UpdateFileCommand{}
		userIdVal, ok := context.Get("user_id")
		if !ok {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(errors.New("no user id in context")))
			return
		}
		cmd.UserId = userIdVal.(string)
		fileId := context.Request.URL.Query().Get("id")
		if fileId == ""{
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(ErrFileIdNotProvided))
			return
		}
		cmd.Id = fileId
		fmt.Println(cmd.UserId)
		err := context.ShouldBindJSON(&cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeDeleteFileEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &DeleteFileCommand{}
		userIdVal, ok := context.Get("user_id")
		if !ok {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(errors.New("no user id in context")))
			return
		}
		cmd.UserId = userIdVal.(string)
		fileId := context.Request.URL.Query().Get("id")
		if fileId == ""{
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(ErrFileIdNotProvided))
			return
		}
		cmd.Id = fileId

		fmt.Println(cmd.UserId)
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusOK, resp)
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
