CREATE OR REPLACE FUNCTION update_selected_avatar()
RETURNS TRIGGER AS $$
BEGIN

    UPDATE profile_avatars
    SET is_selected = false
    WHERE profile_id = NEW.profile_id AND is_selected = true;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER insert_new_profile_avatar
BEFORE INSERT ON profile_avatars
FOR EACH ROW
EXECUTE FUNCTION update_selected_avatar();
