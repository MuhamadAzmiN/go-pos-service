package service

import (
	"context"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/config"
)

type productService struct {
	conf              *config.Config
	productRepository domain.ProductRepository
}

func NewProduct(cnf *config.Config, productRepository domain.ProductRepository) domain.ProductService {
	return &productService{
		conf:              cnf,
		productRepository: productRepository,
	}
}

func (p *productService) GetProductList(ctx context.Context) ([]dto.ProductResponse, error) {
	product, err := p.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.ProductResponse
	for _, v := range product {
		customerData = append(customerData, dto.ProductResponse{
			Id:    v.Id.String(),
			Name:  v.Name,
			Sku:   v.Sku,
			Price: float64(v.Price),
			Stock: v.Stock,
			
		})
	}

	return customerData, nil
}

func (p *productService) GetProductById(ctx context.Context, id string) (dto.ProductData, error) {
	product, err := p.productRepository.FindById(ctx, id)
	if err != nil {
		return dto.ProductData{}, err
	}

	return dto.ProductData{
		Id:    product.Id.String(),
		Name:  product.Name,
		Sku:   product.Sku,
		Price: product.Price,
		Stock: product.Stock,
	}, nil
}

func (p *productService) CreateProduct(ctx context.Context, data dto.ProductRequest) error {
	newProduct := domain.Product{
		Name:  data.Name,
		Sku:   data.Sku,
		Price: int(data.Price),
		Stock: data.Stock,
	}

	err := p.productRepository.Insert(ctx, newProduct)
	if err != nil {
		return err
	}
	return nil
}



func (p *productService) DeleteProduct(ctx context.Context, id string) error {
	checkProduct, err := p.productRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if checkProduct.Id == (domain.Product{}).Id {
		return nil // No product found with the given ID
	}
	
	err = p.productRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

