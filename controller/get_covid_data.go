package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"inshortsProj/connector"
	"inshortsProj/constant"
	"inshortsProj/logger"
	"inshortsProj/models"
	"inshortsProj/utility"
	"net/http"
	"runtime/debug"
)

// GetCovidData godoc
// @Summary      get covid data
// @Description  get covid data by lat,long
// @Tags         CovidAPIs
// @Accept       json
// @Produce      json
// @Param        lat    query     string  false  "latitude of the user"
// @Param        long    query     string  false  "longitude of the user"
// @Success      200  {object}  models.Finalresponse
// @Failure      400  {string}  bad request
// @Failure      500  {string}  Internal Server Error
// @Router       /api/covid/getCovidData/ [get]
func GetCovidData(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic in GetCovidData ,General Error: " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			fmt.Println(ErrorString)
			panic(ErrorString)
		}
	}()

	lat := ctx.Request.URL.Query().Get("lat")
	long := ctx.Request.URL.Query().Get("long")
	if len(lat) == 0 || len(long) == 0 {
		ctx.JSON(http.StatusBadRequest, "please pass valid latitude/longitude")
		return
	}

	response, err := connector.ReverseGeoCode(ctx, lat, long)
	if err != nil {
		logger.LogErrorForScalyr(err.Error(), "GetCovidData", constant.GET_COVID_DATA_VERTICAL, "")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error, geoCode API failing!!")
		return
	}
	var finalResponse models.GeoCodingResponse
	err = json.Unmarshal(response, &finalResponse)
	if err != nil {
		logger.LogErrorForScalyr(err.Error(), "GetCovidData", constant.GET_COVID_DATA_VERTICAL, "")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error, unmarshaling failing!!")
		return
	}

	regionCode, err := utility.GetRegionCode(ctx, finalResponse)
	if err != nil {
		logger.LogErrorForScalyr(err.Error(), "GetCovidData", constant.GET_COVID_DATA_VERTICAL, "")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error, no region code found for the given lat long!!")
		return
	}

	apiResponse, err := utility.GetDataForRegion(ctx, regionCode)
	ctx.JSON(http.StatusOK, apiResponse)
	//ctx.JSON(http.StatusOK, byte)
}
