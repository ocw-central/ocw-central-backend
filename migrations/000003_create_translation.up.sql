BEGIN;

CREATE TABLE IF NOT EXISTS translations(
    id BINARY(16) NOT NULL PRIMARY KEY,
    video_id BINARY(16) NOT NULL,
    language_code VARCHAR(5) NOT NULL,
    translation MEDIUMTEXT NOT NULL,

    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,

    FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE RESTRICT
);

COMMIT;
