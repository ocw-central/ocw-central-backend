BEGIN;

CREATE TABLE IF NOT EXISTS subjects(
    id BINARY(16) NOT NULL PRIMARY KEY,
    category VARCHAR(100),
    academic_field VARCHAR(100),
    title VARCHAR(191) NOT NULL,
    location VARCHAR(100),
    department VARCHAR(100),
    language VARCHAR(100),
    first_held_on DATE,
    free_description VARCHAR(5000),
    series VARCHAR(191),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,

    INDEX(title)
);


CREATE TABLE IF NOT EXISTS videos(
    id BINARY(16) NOT NULL PRIMARY KEY,
    subject_id BINARY(16) NOT NULL,
    title VARCHAR(191) NOT NULL,
    link VARCHAR(200) NOT NULL,
    lectured_on DATE,
    video_length INT,
    language VARCHAR(100),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (subject_id)
        REFERENCES subjects(id)
        ON DELETE RESTRICT,
    
    INDEX(title)
);

COMMIT;
