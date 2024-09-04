CREATE OR REPLACE FUNCTION change_selected_avatar()
RETURNS TRIGGER AS $$
BEGIN
    -- Check if the last deleted avatar was selected
    IF OLD.is_selected = true THEN
        -- find the second avatar after the selected one
        WITH last_avatar AS (
            SELECT avatar_id 
            FROM profile_avatars 
            WHERE profile_id = OLD.profile_id
            ORDER BY avatar_id DESC 
            LIMIT 1
        )
        -- Update last avatar by making it selected
        UPDATE profile_avatars
        SET is_selected = true
        WHERE avatar_id = (SELECT avatar_id FROM last_avatar);
    END IF;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_profile_avatar
AFTER DELETE ON profile_avatars
FOR EACH ROW
EXECUTE FUNCTION change_selected_avatar();
