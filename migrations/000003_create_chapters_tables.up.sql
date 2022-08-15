BEGIN;

CREATE TABLE IF NOT EXISTS chapters(
    id BINARY(16) NOT NULL PRIMARY KEY,
    video_id BINARY(16) NOT NULL,
    start_at INT NOT NULL,
    topic VARCHAR(100) NOT NULL,
    thumbnail_link VARCHAR(200) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    
    FOREIGN KEY (video_id)
        REFERENCES videos(id)
        ON DELETE RESTRICT,

    INDEX(topic) #TODO: need confirmation
);

COMMIT;
