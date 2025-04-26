package repository

import (
	"context"
	"database/sql"
	"opdServices/model/domain"
)

type OpdRepository interface {
	Create(ctx context.Context, tx *sql.Tx, opd domain.Opd) (domain.Opd, error)
	Update(ctx context.Context, tx *sql.Tx, opd domain.Opd) domain.Opd
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.Opd, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Opd, error)
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) (domain.Opd, error)
}
