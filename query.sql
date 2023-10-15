CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    nik VARCHAR(255) NOT NULL UNIQUE,
    user_name VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    legal_name VARCHAR(255) NOT NULL,
    tempat_lahir VARCHAR(255) NOT NULL,
    tanggal_lahir VARCHAR(255) NOT NULL,
    gaji INT NOT NULL,
    foto_ktp VARCHAR(255) NOT NULL,
    foto_selfie VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL DEFAULT 'user'
);



CREATE TABLE transactions (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id integer REFERENCES users(id) NOT NULL,
    no_kontrak character varying(255) NOT NULL,
    otr integer NOT NULL,
    admin_fee integer NOT NULL,
    cicilan integer NOT NULL,
    bunga integer NOT NULL,
    nama_asset character varying(255) NOT NULL,
    status boolean NOT NULL DEFAULT false
);
