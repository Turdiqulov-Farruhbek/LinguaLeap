-- Maqsadlar jadvali
CREATE TABLE goals (
    id UUID PRIMARY KEY, -- Maqsad ID'si
    user_id UUID NOT NULL, -- Foydalanuvchi ID'si
    language_code VARCHAR(10) NOT NULL, -- Til kodi (masalan, 'fr', 'es', va hokazo)
    goal_type VARCHAR(50) NOT NULL, -- Maqsad turi (masalan, 'daily_xp', 'weekly_lessons', va hokazo)
    target_value INT NOT NULL, -- Maqsad uchun kerakli qiymat
    current_value INT DEFAULT 0, -- Hozirgi qiymat
    start_date TIMESTAMP WITH TIME ZONE NOT NULL, -- Maqsad boshlanish sanasi
    end_date TIMESTAMP WITH TIME ZONE NOT NULL, -- Maqsad tugash sanasi
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Maqsad yaratilgan sana
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Maqsad yangilangan sana
);

-- Kunlik progress jadvali
CREATE TABLE daily_progress (
    id UUID PRIMARY KEY, -- Kunlik progress ID'si
    user_id UUID NOT NULL, -- Foydalanuvchi ID'si
    date DATE NOT NULL, -- Sana
    xp_earned INT NOT NULL, -- Kunlik XP miqdori
    lessons_completed INT NOT NULL, -- Bajarilgan darslar soni
    new_words_learned INT NOT NULL, -- O'rganilgan yangi so'zlar soni
    minutes_practiced INT NOT NULL, -- Mashq qilingan daqiqalar
    streak_days INT NOT NULL, -- Kunlik davomiylik kunlari
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Progress yaratilgan sana
);

-- Haftalik progress jadvali
CREATE TABLE weekly_progress (
    id UUID PRIMARY KEY, -- Haftalik progress ID'si
    user_id UUID NOT NULL, -- Foydalanuvchi ID'si
    week_start DATE NOT NULL, -- Haftaning boshlanish sanasi
    week_end DATE NOT NULL, -- Haftaning tugash sanasi
    total_xp_earned INT NOT NULL, -- Haftalik umumiy XP miqdori
    lessons_completed INT NOT NULL, -- Bajarilgan darslar soni
    new_words_learned INT NOT NULL, -- O'rganilgan yangi so'zlar soni
    total_minutes_practiced INT NOT NULL, -- Mashq qilingan umumiy daqiqalar
    most_active_day DATE NOT NULL, -- Eng faol kun
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Progress yaratilgan sana
);

-- Oylik progress jadvali
CREATE TABLE monthly_progress (
    id UUID PRIMARY KEY, -- Oylik progress ID'si
    user_id UUID NOT NULL, -- Foydalanuvchi ID'si
    month DATE NOT NULL, -- Oy
    total_xp_earned INT NOT NULL, -- Oylik umumiy XP miqdori
    lessons_completed INT NOT NULL, -- Bajarilgan darslar soni
    new_words_learned INT NOT NULL, -- O'rganilgan yangi so'zlar soni
    total_minutes_practiced INT NOT NULL, -- Mashq qilingan umumiy daqiqalar
    most_improved_skill VARCHAR(50) NOT NULL, -- Eng ko'p rivojlangan ko'nikma
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Progress yaratilgan sana
);

-- Yutuqlar jadvali
CREATE TABLE achievements (
    id UUID PRIMARY KEY, -- Yutuq ID'si
    user_id UUID NOT NULL, -- Foydalanuvchi ID'si
    title VARCHAR(100) NOT NULL, -- Yutuq nomi
    description TEXT NOT NULL, -- Yutuq ta'rifi
    earned_at TIMESTAMP WITH TIME ZONE NOT NULL, -- Yutuq erishilgan sana
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Yutuq yaratilgan sana
);

-- Liderlar jadvali
CREATE TABLE leaderboards (
    id UUID PRIMARY KEY, -- Liderlar jadvali ID'si
    user_id UUID NOT NULL, -- Foydalanuvchi ID'si
    language_code VARCHAR(10) NOT NULL, -- Til kodi
    rank INT NOT NULL, -- O'rin
    xp INT NOT NULL, -- XP miqdori
    level VARCHAR(20) NOT NULL, -- Daraja
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Jadval yaratilgan sana
);

-- Ko'nikmalar jadvali
CREATE TABLE skills (
    id UUID PRIMARY KEY, -- Ko'nikma ID'si
    user_id UUID NOT NULL, -- Foydalanuvchi ID'si
    language_code VARCHAR(10) NOT NULL, -- Til kodi
    name VARCHAR(50) NOT NULL, -- Ko'nikma nomi (masalan, 'Vocabulary', 'Grammar', 'Listening', va hokazo)
    level INT NOT NULL, -- Ko'nikma darajasi
    progress_to_next_level INT NOT NULL, -- Keyingi darajaga o'tish uchun progress
    words_learned INT NOT NULL, -- O'rganilgan so'zlar soni
    rules_mastered INT NOT NULL, -- Egallangan qoidalar soni
    comprehension_rate INT NOT NULL, -- Tinglab tushunish darajasi
    fluency_rate INT NOT NULL, -- So'zlashuv ravonligi darajasi
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Ko'nikma yaratilgan sana
);

-- Statistika jadvali
CREATE TABLE statistics (
    id UUID PRIMARY KEY, -- Statistika ID'si
    user_id UUID NOT NULL, -- Foydalanuvchi ID'si
    language_code VARCHAR(10) NOT NULL, -- Til kodi
    total_learning_time INT NOT NULL, -- Umumiy o'qish vaqti (daqiqalarda)
    average_daily_time INT NOT NULL, -- O'rtacha kunlik o'qish vaqti (daqiqalarda)
    total_lessons_completed INT NOT NULL, -- Bajarilgan umumiy darslar soni
    total_exercises_completed INT NOT NULL, -- Bajarilgan umumiy mashqlar soni
    accuracy_rate INT NOT NULL, -- Aniqlik darajasi (%)
    most_challenging_skill VARCHAR(50) NOT NULL, -- Eng qiyin ko'nikma
    strongest_skill VARCHAR(50) NOT NULL, -- Eng kuchli ko'nikma
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Statistika yaratilgan sana
);
