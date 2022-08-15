BEGIN;

CREATE TABLE IF NOT EXISTS subpages(
    id BINARY(16) NOT NULL PRIMARY KEY,
    subject_id BINARY(16) NOT NULL,
    link VARCHAR(200) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    
    FOREIGN KEY (subject_id)
        REFERENCES subjects(id)
        ON DELETE RESTRICT,
    
    INDEX(link) #TODO: need confirmation
);

COMMIT;
