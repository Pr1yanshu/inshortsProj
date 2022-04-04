package utility

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"inshortsProj/constant"
	"inshortsProj/database_manager"
	"inshortsProj/logger"
	"inshortsProj/models"
)

func GetDataForRegion(ctx *gin.Context, regionCode string) (models.Finalresponse, error) {
	var finalResponse models.Finalresponse

	if val, err := database_manager.GetFromRedis(ctx, regionCode); err == nil {
		err = json.Unmarshal([]byte(val.(string)), &finalResponse)
		if err != nil {
			fmt.Println("Error in unmarshalling cache : " + err.Error())
			logger.LogErrorForScalyr(err.Error(), "GetDataForRegion", constant.GET_COVID_DATA_VERTICAL, "")
			return finalResponse, err
		}
		return finalResponse, nil
	}

	mongoData, err := database_manager.GetStateCollectionData(ctx)
	if err != nil {
		logger.LogErrorForScalyr(err.Error(), "GetDataForRegion", constant.GET_COVID_DATA_VERTICAL, "")
		return finalResponse, err
	}

	if val, ok := mongoData.Data[regionCode]; ok {
		finalResponse.LastUpdated = val.Meta.LastUpdated
		finalResponse.Confirmed = val.Total.Confirmed
		finalResponse.Deceased = val.Total.Deceased
		finalResponse.Recovered = val.Total.Recovered
		finalResponse.RegionCode = regionCode

	} else {
		return finalResponse, errors.New("no data found in database for this regionCode :" + regionCode)
	}
	marshalledCache, _ := json.Marshal(finalResponse)

	//caching result for this session//////////////
	go cacheSession(ctx, regionCode, marshalledCache)

	return finalResponse, nil
}

func cacheSession(ctx *gin.Context, regionCode string, marshalledCache []byte) {
	Error := database_manager.SetinRedis(ctx, regionCode, string(marshalledCache))
	if Error != nil {
		fmt.Println("Error in setting cache for this session : " + Error.Error())
		logger.LogErrorForScalyr(Error.Error(), "GetDataForRegion", constant.GET_COVID_DATA_VERTICAL, "")

	}
}
