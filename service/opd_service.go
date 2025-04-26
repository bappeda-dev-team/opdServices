package service

import (
	"context"
	"opdServices/model/web"
)

type OpdService interface {
	Create(ctx context.Context, request web.OpdCreateRequest) (web.OpdResponse, error)
	Update(ctx context.Context, request web.OpdUpdateRequest) (web.OpdResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (web.OpdResponse, error)
	FindAll(ctx context.Context) ([]web.OpdResponse, error)
	FindAllOnlyOpd(ctx context.Context) ([]web.OpdOnlyResponse, error)
	FindByKodeOpd(ctx context.Context, kodeOpd string) (web.OpdResponse, error)
}
