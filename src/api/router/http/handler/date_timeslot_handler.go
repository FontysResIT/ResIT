package handler

import (
	"net/http"

	"github.com/RealSnowKid/ResIT/logic"
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
// @Router /dateTimeSlots [get]
func (*handler) GetAllDateTimeslots(c *gin.Context) {
	var dateTimeSlot = _DTSlogic.GetAllDateTimeslots()
	c.JSON(http.StatusOK, dateTimeSlot)
}

func (*handler) GetDateTimeslotByDate(c *gin.Context) {
	var dateTimeSlots = _DTSlogic.GetDateTimeslotByDate()
	c.JSON(http.StatusOK, dateTimeSlots)
}
