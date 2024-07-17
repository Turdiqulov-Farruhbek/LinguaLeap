-- -- Maqsadlar jadvali
-- CREATE TABLE goals (
--     id UUID PRIMARY KEY, -- Maqsad ID'si
--     user_id UUID NOT NULL, -- Foydalanuvchi ID'si
--     language_code VARCHAR(10) NOT NULL, -- Til kodi (masalan, 'fr', 'es', va hokazo)
--     goal_type VARCHAR(50) NOT NULL, -- Maqsad turi (masalan, 'daily_xp', 'weekly_lessons', va hokazo)
--     target_value INT NOT NULL, -- Maqsad uchun kerakli qiymat
--     current_value INT DEFAULT 0, -- Hozirgi qiymat
--     start_date TIMESTAMP WITH TIME ZONE NOT NULL, -- Maqsad boshlanish sanasi
--     end_date TIMESTAMP WITH TIME ZONE NOT NULL, -- Maqsad tugash sanasi
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Maqsad yaratilgan sana
--     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Maqsad yangilangan sana
-- );


-- CREATE TABLE progress (
--     id UUID PRIMARY KEY, -- Progress ID'si
--     user_id UUID NOT NULL, -- Foydalanuvchi ID'si
--     language_code VARCHAR(10) NOT NULL, -- Til kodi
--     date DATE NOT NULL, -- Progress sanasi
--     period_type VARCHAR(10) NOT NULL, -- Period turi (daily, weekly, monthly)
--     xp_earned INT DEFAULT 0, -- XP miqdori
--     lessons_completed INT DEFAULT 0, -- Bajarilgan darslar soni
--     new_words_learned INT DEFAULT 0, -- O'rganilgan yangi so'zlar soni
--     minutes_practiced INT DEFAULT 0, -- Mashq qilingan umumiy daqiqalar
--     streak_days INT DEFAULT 0, -- Davomiylik kunlari (faqat daily uchun)
--     most_active_day DATE, -- Eng faol kun (faqat weekly uchun)
--     most_improved_skill VARCHAR(50), -- Eng ko'p rivojlangan ko'nikma (faqat monthly uchun)
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Yaratilgan sana
--     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Yangilangan sana
-- );


-- -- Ko'nikmalar jadvali
-- CREATE TABLE skills (
--     id UUID PRIMARY KEY, -- Ko'nikma ID'si
--     user_id UUID NOT NULL, -- Foydalanuvchi ID'si
--     language_code VARCHAR(10) NOT NULL, -- Til kodi
--     name VARCHAR(50) NOT NULL, -- Ko'nikma nomi (masalan, 'Vocabulary', 'Grammar', 'Listening', va hokazo)
--     level INT NOT NULL, -- Ko'nikma darajasi
--     progress_to_next_level INT NOT NULL, -- Keyingi darajaga o'tish uchun progress
--     words_learned INT NOT NULL, -- O'rganilgan so'zlar soni
--     rules_mastered INT NOT NULL, -- Egallangan qoidalar soni
--     comprehension_rate INT NOT NULL, -- Tinglab tushunish darajasi
--     fluency_rate INT NOT NULL, -- So'zlashuv ravonligi darajasi
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Ko'nikma yaratilgan sana
-- );


-- -- Foydalanuvchi statistikasi jadvali
-- CREATE TABLE user_statistics (
--     id UUID PRIMARY KEY, -- Statistik ma'lumot ID'si
--     user_id UUID NOT NULL, -- Foydalanuvchi ID'si
--     language_code VARCHAR(10) NOT NULL, -- Til kodi (masalan, 'fr', 'es', va hokazo)
--     total_xp_earned INT NOT NULL, -- Umumiy olingan XP miqdori
--     lessons_completed INT NOT NULL, -- Bajarilgan darslar soni
--     new_words_learned INT NOT NULL, -- O'rganilgan yangi so'zlar soni
--     total_minutes_practiced INT NOT NULL, -- Mashq qilingan umumiy daqiqalar
--     comprehension_rate INT NOT NULL, -- Tinglab tushunish darajasi
--     fluency_rate INT NOT NULL, -- So'zlashuv ravonligi darajasi
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Statistik ma'lumot yaratilgan sana
--     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Statistik ma'lumot yangilangan sana
-- );

