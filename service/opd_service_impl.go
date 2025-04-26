package service

import (
	"context"
	"database/sql"
	"fmt"
	"opdServices/helper"
	"opdServices/model/domain"
	"opdServices/model/web"
	"opdServices/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type OpdServiceImpl struct {
	opdRepository repository.OpdRepository
	db            *sql.DB
	validate      *validator.Validate
}

func NewOpdServiceImpl(opdRepository repository.OpdRepository, db *sql.DB, validate *validator.Validate) *OpdServiceImpl {
	return &OpdServiceImpl{
		opdRepository: opdRepository,
		db:            db,
		validate:      validate,
	}
}

func (service *OpdServiceImpl) Create(ctx context.Context, request web.OpdCreateRequest) (web.OpdResponse, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return web.OpdResponse{}, err
	}

	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.OpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	ts := time.Now().Format("060102150405")
	currentYear := time.Now().Year()
	u := uuid.New()
	randomPart := u.String()[0:6]
	id := fmt.Sprintf("OPD-%v-%v-%v", currentYear, ts, randomPart)

	opd, err := service.opdRepository.Create(ctx, tx, domain.Opd{
		Id:            id,
		NamaOpd:       request.NamaOpd,
		KodeOpd:       request.KodeOpd,
		Singkatan:     request.Singkatan,
		Alamat:        request.Alamat,
		Telepon:       helper.EmptyStringIfNull(request.Telepon),
		Fax:           helper.EmptyStringIfNull(request.Fax),
		Email:         helper.EmptyStringIfNull(request.Email),
		Website:       helper.EmptyStringIfNull(request.Website),
		NipKepalaOpd:  request.NipKepalaOpd,
		NamaKepalaOpd: request.NamaKepalaOpd,
		PangkatKepala: helper.EmptyStringIfNull(request.PangkatKepala),
	})
	if err != nil {
		return web.OpdResponse{}, err
	}

	return web.OpdResponse{
		Id:            opd.Id,
		NamaOpd:       opd.NamaOpd,
		KodeOpd:       opd.KodeOpd,
		Singkatan:     opd.Singkatan,
		Alamat:        opd.Alamat,
		Telepon:       opd.Telepon,
		Fax:           opd.Fax,
		Email:         opd.Email,
		Website:       opd.Website,
		NipKepalaOpd:  opd.NipKepalaOpd,
		NamaKepalaOpd: opd.NamaKepalaOpd,
		PangkatKepala: opd.PangkatKepala,
	}, nil
}

func (service *OpdServiceImpl) Update(ctx context.Context, request web.OpdUpdateRequest) (web.OpdResponse, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return web.OpdResponse{}, err
	}

	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.OpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	opd := service.opdRepository.Update(ctx, tx, domain.Opd{
		Id:            request.Id,
		NamaOpd:       request.NamaOpd,
		KodeOpd:       request.KodeOpd,
		Singkatan:     request.Singkatan,
		Alamat:        request.Alamat,
		Telepon:       helper.EmptyStringIfNull(request.Telepon),
		Fax:           helper.EmptyStringIfNull(request.Fax),
		Email:         helper.EmptyStringIfNull(request.Email),
		Website:       helper.EmptyStringIfNull(request.Website),
		NipKepalaOpd:  request.NipKepalaOpd,
		NamaKepalaOpd: request.NamaKepalaOpd,
		PangkatKepala: helper.EmptyStringIfNull(request.PangkatKepala),
	})

	return web.OpdResponse{
		Id:            opd.Id,
		NamaOpd:       opd.NamaOpd,
		KodeOpd:       opd.KodeOpd,
		Singkatan:     opd.Singkatan,
		Alamat:        opd.Alamat,
		Telepon:       opd.Telepon,
		Fax:           opd.Fax,
		Email:         opd.Email,
		Website:       opd.Website,
		NipKepalaOpd:  opd.NipKepalaOpd,
		NamaKepalaOpd: opd.NamaKepalaOpd,
		PangkatKepala: opd.PangkatKepala,
	}, nil
}

func (service *OpdServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.opdRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *OpdServiceImpl) FindById(ctx context.Context, id string) (web.OpdResponse, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.OpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	opd, err := service.opdRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.OpdResponse{}, err
	}

	return web.OpdResponse{
		Id:            opd.Id,
		NamaOpd:       opd.NamaOpd,
		KodeOpd:       opd.KodeOpd,
		Singkatan:     opd.Singkatan,
		Alamat:        opd.Alamat,
		Telepon:       opd.Telepon,
		Fax:           opd.Fax,
		Email:         opd.Email,
		Website:       opd.Website,
		NipKepalaOpd:  opd.NipKepalaOpd,
		NamaKepalaOpd: opd.NamaKepalaOpd,
		PangkatKepala: opd.PangkatKepala,
	}, nil
}

func (service *OpdServiceImpl) FindAll(ctx context.Context) ([]web.OpdResponse, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return []web.OpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	opds, err := service.opdRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.OpdResponse{}, err
	}

	opdResponses := []web.OpdResponse{}
	for _, opd := range opds {
		opdResponses = append(opdResponses, web.OpdResponse{
			Id:            opd.Id,
			NamaOpd:       opd.NamaOpd,
			KodeOpd:       opd.KodeOpd,
			Singkatan:     opd.Singkatan,
			Alamat:        opd.Alamat,
			Telepon:       opd.Telepon,
			Fax:           opd.Fax,
			Email:         opd.Email,
			Website:       opd.Website,
			NipKepalaOpd:  opd.NipKepalaOpd,
			NamaKepalaOpd: opd.NamaKepalaOpd,
			PangkatKepala: opd.PangkatKepala,
		})
	}

	return opdResponses, nil
}

func (service *OpdServiceImpl) FindAllOnlyOpd(ctx context.Context) ([]web.OpdOnlyResponse, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return []web.OpdOnlyResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	opds, err := service.opdRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.OpdOnlyResponse{}, err
	}

	opdResponses := []web.OpdOnlyResponse{}
	for _, opd := range opds {
		opdResponses = append(opdResponses, web.OpdOnlyResponse{
			NamaOpd: opd.NamaOpd,
			KodeOpd: opd.KodeOpd,
		})
	}

	return opdResponses, nil
}

func (service *OpdServiceImpl) FindByKodeOpd(ctx context.Context, kodeOpd string) (web.OpdResponse, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.OpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	opd, err := service.opdRepository.FindByKodeOpd(ctx, tx, kodeOpd)
	if err != nil {
		return web.OpdResponse{}, err
	}

	return web.OpdResponse{
		NamaOpd: opd.NamaOpd,
		KodeOpd: opd.KodeOpd,
	}, nil
}
