package service

import (
	"context"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/config"
	"time"

	"errors"
)



type CartService struct {
	conf *config.Config
	CartRepository domain.CartRepository
	ProductRepository domain.ProductRepository

}


func NewCart(cnf *config.Config, cartRepository domain.CartRepository, productRepository domain.ProductRepository) domain.CartService {
	return &CartService{
		conf:           cnf,
		CartRepository: cartRepository,
		ProductRepository: productRepository,
	}
}

func (c CartService) AddOrUpdate(ctx context.Context, req dto.AddCartReq) error {

	
	for _, item := range req.Items {

		product, err := c.ProductRepository.FindById(ctx, item.ProductId.String())
		if err != nil {
			return err
		}

		if product.Stock < item.Quantity {
			return errors.New("stock not enough")
		}
		
		err = c.CartRepository.Insert(ctx, req.UserId, item.ProductId, item.Quantity)
		
		if err != nil {
			return err
		}


	}

	return nil
}
	
func (c CartService) GetCartByUserId(ctx context.Context, userId string) (dto.CartFullRes , error) {
	carts , err  := c.CartRepository.GetByCartId(ctx, userId)
	if err != nil {
		return dto.CartFullRes{}, err
	}
	var items []dto.CartItemRes
	var total int

	for _, cart := range carts {
		subtotal := cart.Quantity * cart.Product.Price
		total += subtotal
		items = append(items, dto.CartItemRes{
			 Id:          cart.Id,
        	ProductId:   cart.ProductId,
        	ProductName: cart.Product.Name,
       	 	Price:       cart.Product.Price,
        	Quantity:    cart.Quantity,
        	Subtotal:    subtotal,
        	Tax:         "0%",
        	Discount:    "0%",
			CreatedAt:   time.Now().Format(time.RFC3339),
			UpdatedAt:   time.Now().Format(time.RFC3339),
		})
	}


	return dto.CartFullRes{
        UserId: userId,
        Items:  items,
        Total:  total,
    }, nil
}




func (c CartService) GetAll(ctx context.Context) ([]domain.Cart, error) {
	return c.CartRepository.GetAll(ctx)
}



func (c CartService) DeleteCartById(ctx context.Context, id string) error {
	err := c.CartRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}