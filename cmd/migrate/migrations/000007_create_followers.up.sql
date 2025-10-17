CREATE TABLE followers (
    user_id BIGSERIAL REFERENCES users(id) ON DELETE CASCADE,
    follower_id BIGSERIAL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, follower_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE
);
