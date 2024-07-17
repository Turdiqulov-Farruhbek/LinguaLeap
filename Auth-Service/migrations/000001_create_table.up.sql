
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    role VARCHAR(50) DEFAULT 'user',
    avatar_url VARCHAR(255),
    level INTEGER DEFAULT 1,
    xp INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);


CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY,
    token VARCHAR(400) UNIQUE NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

