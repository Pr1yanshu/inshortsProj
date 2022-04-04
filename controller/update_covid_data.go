package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"inshortsProj/connector"
	"inshortsProj/constant"
	"inshortsProj/database_manager"
	"inshortsProj/logger"
	"inshortsProj/models"
	"net/http"
	"runtime/debug"
)

// UpdateCovidData godoc
// @Summary      update covid data
// @Description  fetch covid data and pers ist in mongo
// @Tags         CovidAPIs
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]models.RegionData
// @Failure      500  {string}  Internal Server Error
// @Router       /api/covid/updateCovidData/ [post]
func UpdateCovidData(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic in UpdateCovidData ,General Error: " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			fmt.Println(ErrorString)
			panic(ErrorString)
		}
	}()
	//header := ctx.Request.Header
	var response map[string]models.RegionData
	data, err := connector.GetData(ctx)
	if err != nil {
		logger.LogErrorForScalyr(err.Error(), "UpdateCovidData", constant.UPDATE_COVID_DATA_VERTICAL, "")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error, state API failing!!")
		return
	}
	error := json.Unmarshal(data, &response)
	if error != nil {
		logger.LogErrorForScalyr(err.Error(), "UpdateCovidData", constant.UPDATE_COVID_DATA_VERTICAL, "")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error, Unmarshaling covid state data failing!!")
		return
	}

	if err := database_manager.SetStateCollectionData(ctx, response); err != nil {
		logger.LogErrorForScalyr(err.Error(), "UpdateCovidData", constant.UPDATE_COVID_DATA_VERTICAL, "")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error,update data failing!!")
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}
