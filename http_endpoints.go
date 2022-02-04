package tik_api_lib

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"net/http"
)

type HttpEndpoints interface {
	MakeLoginEndpoint() gin.HandlerFunc
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
