-- Tillarning asosiy jadvali
CREATE TABLE languages (
    id UUID PRIMARY KEY, -- UUID turi, birlamchi kalit
    language_code VARCHAR(10) UNIQUE NOT NULL, -- Til kodi, masalan, 'en', 'uz'
    name VARCHAR(50) NOT NULL -- Til nomi, masalan, 'Ingliz tili', 'O'zbek tili'
);

-- Darslar jadvali
CREATE TABLE lessons (
    id UUID PRIMARY KEY, -- UUID turi, birlamchi kalit
    language_code VARCHAR(5) NOT NULL, -- Til kodi, masalan, 'en', 'uz'
    title VARCHAR(100) NOT NULL, -- Dars nomi
    level VARCHAR(20) NOT NULL, -- Dars darajasi, masalan, 'boshlang'ich', 'o'rta', 'yuqori'
    content JSONB NOT NULL, -- Dars mazmuni JSON formatida
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Dars yaratilgan vaqti
);

-- Mashqlar jadvali
CREATE TABLE exercises (
    id UUID PRIMARY KEY, -- UUID turi, birlamchi kalit
    language_id UUID REFERENCES languages(id), -- Til jadvalidan til ID'si
    type VARCHAR(50) NOT NULL, -- Mashq turi, masalan, 'tinglash', 'talaffuz'
    title VARCHAR(100) NOT NULL, -- Mashq nomi
    content TEXT NOT NULL -- Mashq mazmuni
);

-- Leksika (so'z boyligi) jadvali
CREATE TABLE vocabulary (
    id UUID PRIMARY KEY, -- UUID turi, birlamchi kalit
    language_code VARCHAR(5) NOT NULL, -- Til kodi, masalan, 'en', 'uz'
    word VARCHAR(100) NOT NULL, -- So'z
    translation VARCHAR(100) NOT NULL, -- So'z tarjimasi
    part_of_speech VARCHAR(20), -- So'zning qismlari (masalan, fe'l, ot)
    example_sentence TEXT, -- Misol jumla
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Yaratilgan vaqti
);

-- Foydalanuvchi taraqqiyoti jadvali
CREATE TABLE user_progress (
    user_id UUID NOT NULL, -- Foydalanuvchi ID'si
    language_id UUID REFERENCES languages(id), -- Til jadvalidan til ID'si
    lesson_id UUID REFERENCES lessons(id), -- Dars jadvalidan dars ID'si
    exercise_id UUID REFERENCES exercises(id), -- Mashq jadvalidan mashq ID'si
    vocabulary_id UUID REFERENCES vocabulary(id), -- Leksika jadvalidan so'z ID'si
    completed_at TIMESTAMP, -- Tugatilgan vaqti
    PRIMARY KEY (user_id, language_id, lesson_id, exercise_id, vocabulary_id) -- Birlamchi kalitlar to'plami
);
