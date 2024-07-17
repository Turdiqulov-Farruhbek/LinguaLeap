CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    native_language VARCHAR(5),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE user_languages (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    language_code VARCHAR(5) NOT NULL,
    proficiency_level VARCHAR(20) NOT NULL,
    started_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- O'zgartirilgan
    UNIQUE (user_id, language_code)
);

-- Tillarning asosiy jadvali
CREATE TABLE languages (
    id UUID PRIMARY KEY, -- UUID turi, birlamchi kalit
    language_code VARCHAR(10) UNIQUE NOT NULL, -- Til kodi, masalan, 'en', 'uz'
    name VARCHAR(50) NOT NULL, -- Til nomi, masalan, 'Ingliz tili', 'O'zbek tili'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Yaratilgan vaqti
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- O'zgartirilgan
);

-- Darslar jadvali
CREATE TABLE lessons (
    id UUID PRIMARY KEY, -- UUID turi, birlamchi kalit
    language_code VARCHAR(5) NOT NULL, -- Til kodi, masalan, 'en', 'uz'
    title VARCHAR(100) NOT NULL, -- Dars nomi
    level VARCHAR(20) NOT NULL, -- Dars darajasi, masalan, 'boshlang'ich', 'o'rta', 'yuqori'
    orderNum INT NOT NULL, 
    content JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Dars yaratilgan vaqti
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- O'zgartirilgan
    UNIQUE (language_code, orderNum) -- Kombinatsiyalangan UNIQUE cheklovi
);


CREATE TABLE user_lessons (
    id SERIAL PRIMARY KEY,
    user_id UUID ,
    lesson_id UUID REFERENCES lessons(id),
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,-- Yaratilgan vaqti
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,-- O'zgartirilgan
    UNIQUE (user_id, lesson_id)
);



-- Leksika (so'z boyligi) jadvali
CREATE TABLE vocabulary (
    id UUID PRIMARY KEY, -- UUID turi, birlamchi kalit
    language_code VARCHAR(5) NOT NULL, -- Til kodi, masalan, 'en', 'uz'
    word VARCHAR(100) NOT NULL, -- So'z
    translation VARCHAR(100) NOT NULL, -- So'z tarjimasi
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Yaratilgan vaqti
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP-- Yaratilgan vaqti
);

-- Foydalanuvchi taraqqiyoti jadvali
CREATE TABLE user_progress (
    user_id UUID NOT NULL, -- Foydalanuvchi ID'si
    language_id UUID REFERENCES languages(id), -- Til jadvalidan til ID'si
    lesson_id UUID REFERENCES lessons(id), -- Dars jadvalidan dars ID'si
    exercise_id UUID UNIQUE NOT NULL, -- Mashq jadvalidan mashq ID'si
    vocabulary_id UUID REFERENCES vocabulary(id), -- Leksika jadvalidan so'z ID'si
    completed_at TIMESTAMP, -- Tugatilgan vaqti
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,-- Yaratilgan vaqti
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,-- O'zgartirilgan
    PRIMARY KEY (user_id, language_id, lesson_id, exercise_id, vocabulary_id) -- Birlamchi kalitlar to'plami
);

