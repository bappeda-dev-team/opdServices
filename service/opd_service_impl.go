package service

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"opdServices/helper"
	"opdServices/model/domain"
	"opdServices/model/web"
	"opdServices/repository"
	"regexp"

	"github.com/go-playground/validator/v10"
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

	if request.KodeOpd == "" {
		return web.OpdResponse{}, fmt.Errorf("kode OPD wajib terisi")
	}

	regex := `^\d\.\d{2}\.\d\.\d{2}\.\d\.\d{2}\.\d{2}\.\d{4}$`
	matched, err := regexp.MatchString(regex, request.KodeOpd)
	if err != nil {
		return web.OpdResponse{}, fmt.Errorf("error validasi format kode: %v", err)
	}
	if !matched {
		return web.OpdResponse{}, fmt.Errorf("format kode tidak valid")
	}

	idRandom := rand.Intn(100000)

	opd, err := service.opdRepository.Create(ctx, tx, domain.Opd{
		Id:            idRandom,
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

	if request.KodeOpd == "" {
		return web.OpdResponse{}, fmt.Errorf("kode OPD wajib terisi")
	}

	regex := `^\d\.\d{2}\.\d\.\d{2}\.\d\.\d{2}\.\d{2}\.\d{4}$`
	matched, err := regexp.MatchString(regex, request.KodeOpd)
	if err != nil {
		return web.OpdResponse{}, fmt.Errorf("error validasi format kode: %v", err)
	}
	if !matched {
		return web.OpdResponse{}, fmt.Errorf("format kode tidak valid")
	}

	opd := service.opdRepository.Update(ctx, tx, domain.Opd{
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
			CreatedAt:     opd.CreatedAt,
			UpdatedAt:     opd.UpdatedAt,
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
