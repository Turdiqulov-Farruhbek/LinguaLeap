CREATE TABLE users (
    id UUID PRIMARY KEY, -- Foydalanuvchi ID'si
    username VARCHAR(50) UNIQUE NOT NULL, -- Foydalanuvchi ismi (login)
    password VARCHAR(255) NOT NULL, -- Parol hash'i
    email VARCHAR(100) UNIQUE NOT NULL, -- Foydalanuvchi email manzili
    native_language VARCHAR(5),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Yaratilgan sana
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Yangilangan sana
    deleted_at BIGINT DEFAULT 0
);



CREATE TABLE user_profiles (
    user_id UUID PRIMARY KEY REFERENCES users(id), -- Foydalanuvchi ID'si
    first_name VARCHAR(50) NOT NULL, -- Ism
    last_name VARCHAR(50) NOT NULL, -- Familiya
    date_of_birth DATE, -- Tug'ilgan sana
    gender VARCHAR(10), -- Jinsi
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Yaratilgan sana
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Yangilangan sana
    deleted_at BIGINT DEFAULT 0
);


CREATE TABLE password_reset_tokens (
    id UUID PRIMARY KEY, -- Token ID'si
    user_id UUID REFERENCES users(id), -- Foydalanuvchi ID'si
    token VARCHAR(255) NOT NULL, -- Parol tiklash tokeni
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL, -- Token amal qilish muddati
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Yaratilgan sana
);
