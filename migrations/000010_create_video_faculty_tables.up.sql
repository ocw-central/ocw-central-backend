BEGIN;

CREATE TABLE IF NOT EXISTS video_faculty(
    video_id BINARY(16) NOT NULL,
    faculty_id BINARY(16) NOT NULL,
    PRIMARY KEY (video_id, faculty_id),
    
    FOREIGN KEY (video_id)
        REFERENCES videos(id)
        ON DELETE RESTRICT,

    Foreign Key (faculty_id)
        REFERENCES faculties(id)
        ON DELETE RESTRICT,
    
    INDEX(video_id) #TODO: need confirmation
);

COMMIT;
