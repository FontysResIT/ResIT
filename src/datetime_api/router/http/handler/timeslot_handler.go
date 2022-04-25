package handler

import (
	"net/http"

	"github.com/FontysResIT/datetime_api/logic"
	"github.com/gin-gonic/gin"
)

var _TSlogic logic.ITimeSlotLogic

type TShandler struct{}

func NewTimeSlotHandler(logic logic.ITimeSlotLogic) *TShandler {
	_TSlogic = logic
	return &TShandler{}
}

// @Description Get all timeslots
// @Accept  json
// @Produce  json
// @Success 200 {array} []model.TimeSlot "ok"
// @Router /timeslots [get]
func (*TShandler) GetAllTimeSlots(c *gin.Context) {
	var timeSlot = _TSlogic.GetAllTimeSlots()
	c.JSON(http.StatusOK, timeSlot)
}
