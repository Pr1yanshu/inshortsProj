package utility

import (
	"errors"
	"github.com/gin-gonic/gin"
	"inshortsProj/models"
)

func GetRegionCode(ctx *gin.Context, geoCodingresponse models.GeoCodingResponse) (string, error) {
	if len(geoCodingresponse.Data) == 0 {
		return "", errors.New("empty data recived from geo coding API")
	}
	node := geoCodingresponse.Data[0]
	return node.RegionCode, nil
}
