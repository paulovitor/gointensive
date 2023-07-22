package usecase

import (
	"github.com/paulovitor/gointensive/internal/entity"
	"github.com/paulovitor/gointensive/internal/infra/database"
)

type OrderInput struct {
	ID    string  `json:"id"` // tag
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

// {"id": "1", "price": 10.0, "tax": 0.1}

type OrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPrice(orderRepository *database.OrderRepository) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		OrderRepository: orderRepository,
	}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
