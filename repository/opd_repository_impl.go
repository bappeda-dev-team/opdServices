package repository

import (
	"context"
	"database/sql"
	"opdServices/helper"
	"opdServices/model/domain"
)

type OpdRepositoryImpl struct {
}

func NewOpdRepositoryImpl() *OpdRepositoryImpl {
	return &OpdRepositoryImpl{}
}

func (repository *OpdRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, opd domain.Opd) (domain.Opd, error) {
	script := "INSERT INTO tb_operasional_daerah (id, nama_opd, kode_opd, singkatan, alamat, telepon, fax, email, website, nip_kepala_opd, nama_kepala_opd, pangkat_kepala) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, script, opd.Id, opd.NamaOpd, opd.KodeOpd, opd.Singkatan, opd.Alamat, opd.Telepon, opd.Fax, opd.Email, opd.Website, opd.NipKepalaOpd, opd.NamaKepalaOpd, opd.PangkatKepala)
	if err != nil {
		return domain.Opd{}, err
	}
	return opd, nil
}

func (repository *OpdRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, opd domain.Opd) domain.Opd {
	script := "UPDATE tb_operasional_daerah SET nama_opd = ?, kode_opd = ?, singkatan = ?, alamat = ?, telepon = ?, fax = ?, email = ?, website = ?, nip_kepala_opd = ?, nama_kepala_opd = ?, pangkat_kepala = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, opd.NamaOpd, opd.KodeOpd, opd.Singkatan, opd.Alamat, opd.Telepon, opd.Fax, opd.Email, opd.Website, opd.NipKepalaOpd, opd.NamaKepalaOpd, opd.PangkatKepala, opd.Id)
	helper.PanicIfError(err)
	return opd
}

func (repository *OpdRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	script := "UPDATE tb_operasional_daerah SET deleted_at = NOW() WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, id)
	helper.PanicIfError(err)
	return nil
}

func (repository *OpdRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (domain.Opd, error) {
	script := "SELECT id, nama_opd, kode_opd, singkatan, alamat, telepon, fax, email, website, nip_kepala_opd, nama_kepala_opd, pangkat_kepala FROM tb_operasional_daerah WHERE id = ?"
	rows, err := tx.QueryContext(ctx, script, id)
	helper.PanicIfError(err)
	defer rows.Close()

	opd := domain.Opd{}
	if rows.Next() {
		err = rows.Scan(&opd.Id, &opd.NamaOpd, &opd.KodeOpd, &opd.Singkatan, &opd.Alamat, &opd.Telepon, &opd.Fax, &opd.Email, &opd.Website, &opd.NipKepalaOpd, &opd.NamaKepalaOpd, &opd.PangkatKepala)
		helper.PanicIfError(err)
	}
	return opd, nil
}

func (repository *OpdRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Opd, error) {
	script := "SELECT id, nama_opd, kode_opd, singkatan, alamat, telepon, fax, email, website, nip_kepala_opd, nama_kepala_opd, pangkat_kepala, created_at, updated_at FROM tb_operasional_daerah WHERE deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	opds := []domain.Opd{}
	for rows.Next() {
		opd := domain.Opd{}
		err = rows.Scan(&opd.Id, &opd.NamaOpd, &opd.KodeOpd, &opd.Singkatan, &opd.Alamat, &opd.Telepon, &opd.Fax, &opd.Email, &opd.Website, &opd.NipKepalaOpd, &opd.NamaKepalaOpd, &opd.PangkatKepala, &opd.CreatedAt, &opd.UpdatedAt)
		helper.PanicIfError(err)
		opds = append(opds, opd)
	}
	return opds, nil
}

func (repository *OpdRepositoryImpl) FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) (domain.Opd, error) {
	script := "SELECT id, nama_opd, kode_opd, singkatan, alamat, telepon, fax, email, website, nip_kepala_opd, nama_kepala_opd, pangkat_kepala FROM tb_operasional_daerah WHERE kode_opd = ? AND deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, script, kodeOpd)
	helper.PanicIfError(err)
	defer rows.Close()

	opd := domain.Opd{}
	if rows.Next() {
		err = rows.Scan(&opd.Id, &opd.NamaOpd, &opd.KodeOpd, &opd.Singkatan, &opd.Alamat, &opd.Telepon, &opd.Fax, &opd.Email, &opd.Website, &opd.NipKepalaOpd, &opd.NamaKepalaOpd, &opd.PangkatKepala)
		helper.PanicIfError(err)
	}
	return opd, nil
}
