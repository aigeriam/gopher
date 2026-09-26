ALTER TABLE posts
    DROP CONSTRAINT IF EXISTS posts_user_id_fkey,
    DROP COLUMN IF EXISTS tags,
    DROP COLUMN IF EXISTS updated_at;