package rewrite

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/FontysResIT/ResIT/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	CustomerId primitive.ObjectID `json:"customer_id"`
}

type SharableReservation struct {
	Id             primitive.ObjectID `json:"id"`
	RestaurantId   primitive.ObjectID `json:"restaurant_id"`
	CustomerId     primitive.ObjectID `json:"customer_id"`
	GroupSize      int                `json:"groupSize"`
	Group          []Group            `json:"group"`
	TableNumber    string             `json:"tableNumber"`
	CreatedAt      time.Time          `json:"created_at"`
	State          string             `json:"state"`
	Comment        string             `json:"comment"`
	SingleHoushold bool               `json:"single_household"`
}

type DietaryRestriction struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type SherableCustomer struct {
	Id                  primitive.ObjectID   `json:"id"`
	Name                string               `json:"name"`
	Email               string               `json:"email"`
	DietaryRestrictions []DietaryRestriction `json:"dietaryRestrictions"`
}

func ReservationToModels(reservation model.ReservationReadDTO) (SharableReservation, []SherableCustomer) {
	creatorId := primitive.NewObjectID()
	customers := []SherableCustomer{}
	customerIds := []Group{}
	length := len(reservation.GuestPersona)
	for i := 0; i < length; i++ {
		g := reservation.GuestPersona[i]
		dr := []DietaryRestriction{}
		var c = SherableCustomer{}
		for _, dreq := range g.DietaryRequirements {
			dr = append(dr, DietaryRestriction{Name: dreq, Comment: "No " + strings.ToLower(dreq)})
		}
		for _, fp := range g.FoodPreferences {
			dr = append(dr, DietaryRestriction{Name: fp, Comment: "Include " + strings.ToLower(fp)})
		}
		if i == 0 {
			c = SherableCustomer{Id: creatorId, Name: g.GuestName, Email: "", DietaryRestrictions: dr}
		} else {
			c = SherableCustomer{Id: primitive.NewObjectID(), Name: g.GuestName, Email: "", DietaryRestrictions: dr}
		}
		customers = append(customers, c)
		customerIds = append(customerIds, Group{c.Id})
	}
	var reservationStatus string
	if reservation.IsCancelled {
		reservationStatus = "CANCELED_BY_CUSTOMER"
	} else {
		reservationStatus = "RESERVED"
	}
	sr := SharableReservation{Id: reservation.Id, RestaurantId: reservation.RestaurantId, CustomerId: creatorId,
		GroupSize: reservation.GuestCount, Group: customerIds, TableNumber: strconv.Itoa(rand.Intn(100)),
		CreatedAt: reservation.DateTimeSlot.Date, State: reservationStatus, Comment: reservation.Remark, SingleHoushold: false}
	return sr, customers
}

func StructToJson(res model.ReservationReadDTO) ([]byte, []byte) {
	sr, sc := ReservationToModels(res)
	final_res, err := json.Marshal(sr)
	if err != nil {
		return nil, nil
	}
	final_custs, err := json.Marshal(sc)
	if err != nil {
		return nil, nil
	}

	return final_res, final_custs
}
