CREATE TABLE tb_operasional_daerah (
    id INT AUTO_INCREMENT PRIMARY KEY,
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
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
