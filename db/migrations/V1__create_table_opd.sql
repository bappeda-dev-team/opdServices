CREATE TABLE tb_operasional_daerah(
    id BIGSERIAL PRIMARY KEY,
    nama_opd VARCHAR(255) NOT NULL,
    kode_opd VARCHAR(255) NOT NULL,
    singkatan VARCHAR(255) NOT NULL,
    alamat TEXT NOT NULL,
    telepon VARCHAR(255) NOT NULL,
    fax VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    website VARCHAR(255) NOT NULL,
    nip_kepala_opd VARCHAR(255) NOT NULL,
    nama_kepala_opd TEXT,
    pangkat_kepala VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);


CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_tb_operasional_daerah_timestamp
BEFORE UPDATE ON tb_operasional_daerah
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();