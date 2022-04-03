package utility

import (
	"errors"
	"github.com/gin-gonic/gin"
	"inshortsProj/database_manager"
	"inshortsProj/logger"
	"inshortsProj/models"
)

func GetDataForRegion(ctx *gin.Context, regionCode string) (models.Finalresponse, error) {
	var finalResponse models.Finalresponse
	mongoData, err := database_manager.GetStateCollectionData(ctx)
	if err != nil {
		logger.LogErrorForScalyr(err.Error(), "GetDataForRegion", "api/getCovidData", "")
		return finalResponse, err
	}

	if val, ok := mongoData.Data[regionCode]; ok {
		finalResponse.LastUpdated = val.Meta.LastUpdated
		finalResponse.Confirmed = val.Total.Confirmed
		finalResponse.Deceased = val.Total.Deceased
		finalResponse.Recovered = val.Total.Recovered

	} else {
		return finalResponse, errors.New("no data found in database for this regionCode :" + regionCode)
	}

	return finalResponse, nil
}
