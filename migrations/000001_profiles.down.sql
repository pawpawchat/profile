ALTER TABLE IF EXISTS profiles DROP CONSTRAINT IF EXISTS fk_avatar_id;

DROP TABLE IF EXISTS profile_avatars;
DROP TABLE IF EXISTS profile_biographies;
DROP TABLE IF EXISTS profiles;
