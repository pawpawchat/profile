CREATE TABLE profiles (
    profile_id        BIGSERIAL PRIMARY KEY,
    username          VARCHAR(16) NOT NULL,
    avatar_id         BIGINT,
    description       VARCHAR(64) NOT NULL DEFAULT '',
    number_of_friends INT NOT NULL DEFAULT 0,
    online            BOOLEAN NOT NULL DEFAULT false,
    last_seen         TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at        TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE profile_avatars (
    avatar_id       BIGINT PRIMARY KEY,
    profile_id      BIGINT NOT NULL,
    original_source VARCHAR(256) NOT NULL,
    small_source    VARCHAR(256) NOT NULL,
    added_at        TIMESTAMP NOT NULL DEFAULT NOW()
);

-- внешний ключ на id профилей для таблицы с аватарами
ALTER TABLE profile_avatars ADD CONSTRAINT fk_profile_id
FOREIGN KEY (profile_id) REFERENCES profiles(profile_id) ON DELETE CASCADE;

CREATE TABLE  profile_biographies (
    biography_id BIGSERIAL PRIMARY KEY,
    profile_id   BIGINT NOT NULL,
    first_name   VARCHAR(16) NOT NULL,
    second_name  VARCHAR(16) NOT NULL,
    birthday     TIMESTAMP,

    CONSTRAINT fk_profile_id FOREIGN KEY (profile_id) REFERENCES profiles(profile_id)
);

-- внешний ключ на id аватарки для таблицы profiles
ALTER TABLE profiles ADD CONSTRAINT fk_avatar_id 
FOREIGN KEY (avatar_id) REFERENCES profile_avatars(avatar_id) ON DELETE CASCADE;

-- функция для триггера, который срабатывает перед каждым инсертом новой записи
-- устанавливает в поле username значение типа "200"+"${profile_id}"
CREATE OR REPLACE FUNCTION set_default_username()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.username IS NULL THEN
        NEW.username := '200' || NEW.profile_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Тригер для вставки в 
CREATE TRIGGER before_insert_profiles
BEFORE INSERT ON profiles
FOR EACH ROW
EXECUTE FUNCTION set_default_username();
