CREATE TABLE profiles (
    profile_id        BIGSERIAL PRIMARY KEY,
    username          VARCHAR(16) NOT NULL,
    description       VARCHAR(64) NOT NULL DEFAULT '',
    number_of_friends INT NOT NULL DEFAULT 0,
    online            BOOLEAN NOT NULL DEFAULT false,
    last_seen         TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at        TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE TABLE profile_avatars (
    avatar_id BIGSERIAL PRIMARY KEY,
    profile_id BIGINT NOT NULL,
    orig_url VARCHAR(128) NOT NULL,
    added_at TIMESTAMP NOT NULL DEFAULT NOW(),
    is_selected BOOLEAN NOT NULL DEFAULT false,
    FOREIGN KEY (profile_id) REFERENCES profiles(profile_id) 
);


CREATE TABLE  profile_biographies (
    biography_id BIGSERIAL PRIMARY KEY,
    profile_id   BIGINT  UNIQUE NOT NULL,
    first_name   VARCHAR(16) NOT NULL,
    second_name  VARCHAR(16) NOT NULL,
    birthday     TIMESTAMP,

    CONSTRAINT fk_profile_id 
    FOREIGN KEY (profile_id) 
    REFERENCES profiles(profile_id) ON DELETE CASCADE
);

-- function for a trigger that fires before each insertion of a new record
-- sets the username field to “200”+“${profile_id}”.
CREATE OR REPLACE FUNCTION set_default_username()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.username IS NULL THEN
        NEW.username := '200' || NEW.profile_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER before_insert_profiles
BEFORE INSERT ON profiles
FOR EACH ROW
EXECUTE FUNCTION set_default_username();
