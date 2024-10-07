-- removing triggers
DROP TRIGGER IF EXISTS update_posts_timestamp ON posts;
DROP TRIGGER IF EXISTS update_comments_timestamp ON comments;

-- removing tables
DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS labels;