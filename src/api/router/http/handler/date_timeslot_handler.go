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

// @Description Get all reservations
// @Accept  json
// @Produce  json
// @Success 200 {array} []model.Reservation	"ok"
// @Router /reservation [get]
func (*handler) GetAllDateTimeslots(c *gin.Context) {
	var dateTimeSlot = _DTSlogic.GetAllDateTimeslots()
	c.JSON(http.StatusOK, dateTimeSlot)
}
