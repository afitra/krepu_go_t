CREATE TABLE IF NOT EXISTS users (
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



CREATE TABLE IF NOT EXISTS transactions (
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
<<<<<<< Updated upstream
=======



INSERT INTO "users" ( "nik", "user_name", "password", "full_name", "legal_name", "tempat_lahir", "tanggal_lahir", "gaji", "foto_ktp", "foto_selfie", "role", "tenor_satu", "tenor_dua", "tenor_tiga", "tenor_empat") VALUES
('3493493049d3049',	'apitoong',	'$2a$10$pOJ7JS/Uqqk68VZDi0ZimeGwDpxAmZ8s5oo0H2mnG6fhT4rAD/N/u',	'Afitra Mamor Bikhoir',	'Afitra Mamor Bikhoir',	'kediri',	'03-03-1996',	10000000,	'https://img.inews.co.id/media/600/files/networks/2022/06/20/ffb4c_anya-geraldine.jpg',	'https://img.inews.co.id/media/600/files/networks/2022/06/20/ffb4c_anya-geraldine.jpg',	'user',	1200000,	1800000,	2000000,	2000000),
('34934930493049',	'admin123',	'$2a$10$Efh/oCrUYazpvci7Gh5P7Or/zA7vkeMuQfQRyv0ABFlrMaPhBx2GW',	'Afitra Mamor Bikhoir',	'Afitra Mamor Bikhoir',	'kediri',	'03-03-1996',	10000000,	'https://img.inews.co.id/media/600/files/networks/2022/06/20/ffb4c_anya-geraldine.jpg',	'https://img.inews.co.id/media/600/files/networks/2022/06/20/ffb4c_anya-geraldine.jpg',	'admin',	600000,	1000000,	1800000,	1800000);
>>>>>>> Stashed changes
