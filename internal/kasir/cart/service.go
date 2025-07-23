package cart

import (
	"context"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/config"
	"time"

	"errors"
)

type Service struct {
	conf *config.Config
	cartRepository domain.CartRepository
	productRepository domain.ProductRepository

}


func NewService(config *config.Config, cartRepository domain.CartRepository, productRepository domain.ProductRepository) *Service {
	return &Service{
		conf: config,
		cartRepository: cartRepository,
		productRepository: productRepository,
	}
}



func (c *Service) AddOrUpdate(ctx context.Context, req dto.AddCartReq) error {

	
	for _, item := range req.Items {

		product, err := c.productRepository.FindById(ctx, item.ProductId.String())
		if err != nil {
			return err
		}

		if product.Stock < item.Quantity {
			return errors.New("stock not enough")
		}
		
		err = c.cartRepository.Insert(ctx, req.UserId, item.ProductId, item.Quantity)
		
		if err != nil {
			return err
		}


	}

	return nil
}
	
func (c *Service) GetCartByUserId(ctx context.Context, userId string) (dto.CartFullRes , error) {
	carts , err  := c.cartRepository.GetByCartId(ctx, userId)
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




func (c *Service) GetAll(ctx context.Context) ([]dto.CartFullRes, error) {
	carts, err := c.cartRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var cartFullResList []dto.CartFullRes
	for _, cart := range carts {
		var items []dto.CartItemRes
		var total int

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

		cartFullResList = append(cartFullResList, dto.CartFullRes{
			UserId: cart.UserId,
			Items:  items,
			Total:  total,
		})
	}

	return cartFullResList, nil
}



func (c *Service) DeleteCartById(ctx context.Context, id string) error {
	err := c.cartRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}