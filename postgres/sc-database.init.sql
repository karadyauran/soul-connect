\c sc_db
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Creating a trigger to update the time (updated_at) when a record changes
CREATE OR REPLACE FUNCTION update_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;