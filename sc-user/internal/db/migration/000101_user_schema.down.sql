-- removing a trigger
DROP TRIGGER IF EXISTS update_users_timestamp ON users;
DROP FUNCTION IF EXISTS update_timestamp;

-- removing tables
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS users;