package handler

import (
	"net/http"

	"github.com/RealSnowKid/ResIT/logic"
	"github.com/gin-gonic/gin"
)

var _TSlogic logic.ITimeSlotLogic

type TShandler struct{}

func NewTimeSlotHandler(logic logic.ITimeSlotLogic) *handler {
	_TSlogic = logic
	return &handler{}
}

// @Description Get all timeslots
// @Accept  json
// @Produce  json
// @Success 200 {array} []model.TimeSlot "ok"
// @Router /timeslots [get]
func (*handler) GetAllTimeSlots(c *gin.Context) {
	var timeSlot = _TSlogic.GetAllTimeSlots()
	c.JSON(http.StatusOK, timeSlot)
}
