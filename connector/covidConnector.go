package connector

import (
	"github.com/gin-gonic/gin"
	"inshortsProj/constant"
	"inshortsProj/http_call"
	"inshortsProj/logger"
)

func GetData(ctx *gin.Context) ([]byte, error) {
	url := constant.COVID_STATE_API_URL
	header := make(map[string]interface{})
	header["Content-Type"] = "application/json"
	header["Accept"] = "application/json"
	data, err := http_call.MakeGetHttpCall(url, header, constant.COVID_STATE_API_TIMEOUT)
	if err != nil {
		logger.LogErrorForScalyr(err.Error(), "GetData", constant.UPDATE_COVID_DATA_VERTICAL, "")
		return data, err
	}
	return data, nil

}
