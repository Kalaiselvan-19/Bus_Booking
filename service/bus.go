package service

import (
	"Bus_Booking/graph/model"
	"Bus_Booking/initializers"
	"context"
)

func BusCreate(ctx context.Context, input model.NewBus) (*model.Buses, error) {
	bus := model.Buses{
		BusName:   input.BusName,
		BusNumber: input.BusNumber,
		TotalSeat: input.TotalSeat,
	}

	if err := initializers.DB.Create(&bus).Error; err != nil {
		return nil, err
	}
	return &bus, nil
}

func BusGetByNum(ctx context.Context, Busnum string) (*model.Buses, error) {

	var bus model.Buses

	if err := initializers.DB.Model(&bus).Where("Bus_Number LIKE ?", Busnum).Take(&bus).Error; err != nil {
		return nil, err
	}

	return &bus, nil
}
