CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    nik VARCHAR(255)  NOT NULL UNIQUE,
    user_name VARCHAR(255)  NOT NULL UNIQUE,
    password VARCHAR(255)  NOT NULL,
    full_name VARCHAR(255)  NOT NULL,
    legal_name VARCHAR(255)  NOT NULL,
    tempat_lahir VARCHAR(255)  NOT NULL,
    tanggal_lahir VARCHAR(255)  NOT NULL,
    gaji INT  NOT NULL,
    foto_ktp VARCHAR(255)  NOT NULL,
    foto_selfie VARCHAR(255)  NOT NULL
);