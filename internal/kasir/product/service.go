package product

import (
	"context"

	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/config"
)

type Service struct {
	conf *config.Config
	repo domain.ProductRepository
}

func NewService(cnf *config.Config, repo domain.ProductRepository) *Service {
	return &Service{conf: cnf, repo: repo}
}

// ===== use‑case =====
func (s *Service) GetList(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]dto.ProductResponse, 0, len(products))
	for _, v := range products {
		resp = append(resp, dto.ProductResponse{
			Id:    v.Id.String(),
			Name:  v.Name,
			Sku:   v.Sku,
			Price: float64(v.Price),
			Stock: v.Stock,
		})
	}
	return resp, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (dto.ProductData, error) {
	p, err := s.repo.FindById(ctx, id)
	if err != nil {
		return dto.ProductData{}, err
	}

	return dto.ProductData{
		Id:    p.Id.String(),
		Name:  p.Name,
		Sku:   p.Sku,
		Price: p.Price,
		Stock: p.Stock,
	}, nil
}

func (s *Service) Create(ctx context.Context, req dto.ProductRequest) error {
	return s.repo.Insert(ctx, domain.Product{
		Name:  req.Name,
		Sku:   req.Sku,
		Price: int(req.Price),
		Stock: req.Stock,
	})
}

func (s *Service) Delete(ctx context.Context, id string) error {
	// opsional: validasi id di sini
	return s.repo.Delete(ctx, id)
}
