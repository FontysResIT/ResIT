package logic

import (
	"github.com/FontysResIT/config_api/model"
	"github.com/FontysResIT/config_api/repository"
	"github.com/FontysResIT/config_api/util"
)

type IRestaurantLogic interface {
	GetAllRestaurants() []model.Restaurant
	CreateRestaurant(model.Restaurant) (model.Restaurant, error)
}

var _repositoryRestaurant repository.IRestaurant

type Restaurantlogic struct{}

func NewRestaurantLogic(repository repository.IRestaurant) *Restaurantlogic {
	_repositoryRestaurant = repository
	return &Restaurantlogic{}
}

func (*Restaurantlogic) GetAllRestaurants() []model.Restaurant {
	return _repositoryRestaurant.All()
}

func (*Restaurantlogic) CreateRestaurant(restaurant model.Restaurant) (model.Restaurant, error) {
	restaurant, err := _repositoryRestaurant.Create(restaurant)
	if err == nil {
		go func() {
			util.CreateRestaurant(restaurant)
		}()
	}
	return restaurant, err
}
