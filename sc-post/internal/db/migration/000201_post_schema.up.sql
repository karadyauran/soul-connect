SET TIMEZONE = 'UTC';

-- creating a post table
CREATE TABLE IF NOT EXISTS  posts (
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
    user_id UUID NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    likes_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

--creating a comment table
CREATE TABLE IF NOT EXISTS  comments (
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
    post_id UUID NOT NULL REFERENCES posts(id),
    user_id UUID NOT NULL REFERENCES users(id),
    content TEXT NOT NULL,
    likes_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Creating a label table with a limit of 5 predefined values
CREATE TABLE IF NOT EXISTS  labels (
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
    name VARCHAR(50) UNIQUE NOT NULL
);

-- insert 5 predefined labels by mood
INSERT INTO labels (name) VALUES
('Happy'),
('Sad'),
('Angry'),
('Excited'),
('Calm');

CREATE TABLE IF NOT EXISTS  labels_posts (
    label_id UUID NOT NULL REFERENCES labels(id),
    post_id UUID NOT NULL REFERENCES posts(id)
);

-- creating a table of likes
CREATE TABLE IF NOT EXISTS  likes (
   id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
   post_id UUID REFERENCES posts(id) ON DELETE CASCADE,
   comment_id UUID REFERENCES comments(id) ON DELETE CASCADE,
   user_id UUID NOT NULL REFERENCES users(id),
   created_at TIMESTAMP DEFAULT NOW(),
   CONSTRAINT unique_post_comment_user_like UNIQUE (post_id, comment_id, user_id)
);

-- Creating a trigger to update the time (updated_at) when a record changes
CREATE OR REPLACE FUNCTION update_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- applying a post trigger
CREATE OR REPLACE TRIGGER update_posts_timestamp
    BEFORE UPDATE ON posts
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- applying a comment trigger
CREATE TRIGGER update_comments_timestamp
    BEFORE UPDATE ON comments
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- trigger for adding a like to a post
CREATE OR REPLACE FUNCTION increment_post_likes()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE posts
    SET likes_count = likes_count + 1
    WHERE id = NEW.post_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE  TRIGGER after_like_post
    AFTER INSERT ON likes
    FOR EACH ROW
    WHEN (NEW.post_id IS NOT NULL)
EXECUTE FUNCTION increment_post_likes();

-- trigger for removing a like from a post
CREATE OR REPLACE FUNCTION decrement_post_likes()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE posts
    SET likes_count = likes_count - 1
    WHERE id = OLD.post_id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE  TRIGGER after_unlike_post
    AFTER DELETE ON likes
    FOR EACH ROW
    WHEN (OLD.post_id IS NOT NULL)
EXECUTE FUNCTION decrement_post_likes();

-- trigger for adding a like to a comment
CREATE OR REPLACE FUNCTION increment_comment_likes()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE comments
    SET likes_count = likes_count + 1
    WHERE id = NEW.comment_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE  TRIGGER after_like_comment
    AFTER INSERT ON likes
    FOR EACH ROW
    WHEN (NEW.comment_id IS NOT NULL)
EXECUTE FUNCTION increment_comment_likes();

-- trigger for removing a like from a comment
CREATE OR REPLACE FUNCTION decrement_comment_likes()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE comments
    SET likes_count = likes_count - 1
    WHERE id = OLD.comment_id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE  TRIGGER after_unlike_comment
    AFTER DELETE ON likes
    FOR EACH ROW
    WHEN (OLD.comment_id IS NOT NULL)
EXECUTE FUNCTION decrement_comment_likes();