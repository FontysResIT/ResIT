package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/FontysResIT/ResIT/logic"
	"github.com/gin-gonic/gin"
)

var _DTSlogic logic.IDateTimeslotLogic

type DTShandler struct{}

func NewDateTimeslotHandler(logic logic.IDateTimeslotLogic) *handler {
	_DTSlogic = logic
	return &handler{}
}

// @Description Get all date time slots
// @Accept  json
// @Produce  json
// @Success 200 {array} []model.DateTimeSlot "ok"
// @Router /datetimeslots [get]
func (*handler) GetAllDateTimeslots(c *gin.Context) {
	var dateTimeSlot = _DTSlogic.GetAllDateTimeslots()
	c.JSON(http.StatusOK, dateTimeSlot)
}

func (*handler) GetDateTimeslotById(c *gin.Context) {
	var dateTimeSlot = _DTSlogic.GetDateTimeslotById(c.Param("id"))
	c.JSON(http.StatusOK, dateTimeSlot)
}

// @Description Get all date time slots by a parameter
// @Accept  json
// @Produce  json
// @Param	query	path	string	false	"search date time slots by a query (date, dateId)"
// @Param	param	path	string	false	"search date time slots by a paramteter (e.g. 2022-04-07)"
// @Success 200 {array} []model.DateTimeSlot "ok"
// @Success	200	{array}	int	"ok"
// @Router /datetimeslots/{query}/{param} [get]
func (*handler) GetDateTimeslotByParam(c *gin.Context) {
	query := c.Param("query")
	param := c.Param("param")
	switch query {
	case "date":
		var params = strings.Split(strings.Split(param, "/")[1], "-")
		var dates = []int{}

		for _, v := range params {
			j, err := strconv.Atoi(v)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			dates = append(dates, j)
		}
		date := time.Date(dates[0], time.Month(dates[1]), dates[2], 0, 0, 0, 0, time.FixedZone("CEST", 2*60*60))
		dateTimeSlots := _DTSlogic.GetDateTimeslotsByDate(date)
		c.JSON(http.StatusOK, dateTimeSlots)
	case "dateId":
		var params = strings.Split(strings.Split(param, "/")[1], "-")
		var dates = []int{}

		for _, v := range params {
			j, err := strconv.Atoi(v)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			dates = append(dates, j)
		}
		date := time.Date(dates[0], time.Month(dates[1]), dates[2], 0, 0, 0, 0, time.FixedZone("CEST", 2*60*60))
		dateTimeSlot := _DTSlogic.GetDateTimeslotsIdByDate(date)
		c.JSON(http.StatusOK, dateTimeSlot)
	}
}
