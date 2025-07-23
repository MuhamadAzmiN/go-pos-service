package transaction

import (
	"context"
	"errors"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/config"
	"my-golang-service-pos/internal/utils"
	"time"

	"github.com/google/uuid"
)


 type Service struct {
	conf *config.Config
	transactionRepository domain.TransactionRepository
	cartRepository domain.CartRepository
	productRepository domain.ProductRepository

 }

 

 func NewService(config *config.Config, transactionRepository domain.TransactionRepository, cartRepository domain.CartRepository, productRepository domain.ProductRepository) *Service {
	return &Service{
		conf: config,
		transactionRepository: transactionRepository,
		cartRepository: cartRepository,
		productRepository: productRepository,
	}
 }




 func (s *Service) GetList(ctx context.Context) ([]domain.Transaction, error) {
	data, err := s.transactionRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
 }


 func (c *Service) CreateTransaction(ctx context.Context, req dto.TransactionRequest) (domain.Transaction, error) {
	cartItems, err := c.cartRepository.GetByCartId(ctx, req.UserId.String())
	if err != nil {
		return domain.Transaction{}, errors.New("failed to get cart items")
	}


	if len(cartItems) == 0 {
		return domain.Transaction{}, errors.New("cart is empty")
	}

	var total float64
	for _, item := range cartItems {
		total += float64(item.Product.Price) * float64(item.Quantity)
	}

	if req.PaidAmount < total {
		return domain.Transaction{}, errors.New("paid amount is not enough")
	}


	trx := domain.Transaction{
		Id:            uuid.New().String(),
		InvoiceNumber:  utils.GenerateInvoiceNumber(), // kamu bisa bikin logic-nya
		UserId:        req.UserId.String(),
		TotalPrice:    total,
		PaidAmount:    req.PaidAmount,
		PaymentMethod: req.PaymentMethod,
		ChangeAmount:  req.PaidAmount - total,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}


	var items []domain.TransactionItem
	for _, item := range cartItems {
		items = append(items, domain.TransactionItem{
			TransactionId: trx.Id,
			ProductId:     item.ProductId.String(),
			Quantity:      item.Quantity,
			Price:         float64(item.Product.Price),
			Subtotal:      float64(item.Product.Price) * float64(item.Quantity),
		})
	}


	


	savedTrx, err := c.transactionRepository.Create(ctx, trx, items)
	if err != nil {
		return domain.Transaction{}, err
	}


	for _, item := range items {
		productUUID, err := uuid.Parse(item.ProductId)
		if err != nil {
			return domain.Transaction{}, err
		}
		err = c.productRepository.DecreaseStock(ctx, productUUID, item.Quantity)
		if err != nil {
			return domain.Transaction{}, err
		}

		err = c.cartRepository.ClearCart(ctx, req.UserId.String())
		if err != nil {
			return domain.Transaction{}, err
		}
	}

	savedTrx.Items = items

	return savedTrx, nil


}
