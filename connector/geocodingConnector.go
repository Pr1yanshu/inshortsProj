package connector

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inshortsProj/constant"
	"inshortsProj/http_call"
	"inshortsProj/logger"
)

func ReverseGeoCode(ctx *gin.Context, lat string, long string) ([]byte, error) {
	url := fmt.Sprintf(constant.REVERSE_GEO_CODE_API_URL+"&query=%s,%s", lat, long)
	header := make(map[string]interface{})
	header["Content-Type"] = "application/json"
	header["Accept"] = "application/json"
	data, err := http_call.MakeGetHttpCall(url, header, constant.REVERSE_GEO_CODE_API_TIMEOUT)
	if err != nil {
		logger.LogErrorForScalyr(err.Error(), "reverseGeoCode", constant.GET_COVID_DATA_VERTICAL, "")
		return data, err
	}
	return data, nil

}
