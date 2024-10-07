-- Drop triggers and functions

-- Drop trigger for logging successful login attempts
DROP TRIGGER IF EXISTS log_successful_login ON auth;
DROP FUNCTION IF EXISTS log_login_attempt;

-- Drop trigger for updating the last login timestamp
DROP TRIGGER IF EXISTS update_login_timestamp ON auth;
DROP FUNCTION IF EXISTS update_last_login;

-- Drop trigger for deleting old sessions
DROP TRIGGER IF EXISTS clear_old_sessions ON auth;
DROP FUNCTION IF EXISTS delete_old_sessions;

-- Drop tables

-- Drop the sessions table
DROP TABLE IF EXISTS sessions;

-- Drop the login_attempts table
DROP TABLE IF EXISTS login_attempts;

-- Drop the auth table
DROP TABLE IF EXISTS auth;