CREATE TABLE IF NOT EXISTS photos(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(50),
    caption VARCHAR(150),
    photo_url VARCHAR(255) UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);