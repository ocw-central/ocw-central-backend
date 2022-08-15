BEGIN;

CREATE TABLE IF NOT EXISTS subject_faculty(
    subject_id BINARY(16) NOT NULL,
    faculty_id BINARY(16) NOT NULL,
    PRIMARY KEY (subject_id, faculty_id),

    FOREIGN KEY (subject_id)
        REFERENCES subjects(id)
        ON DELETE RESTRICT,
    
    Foreign Key (faculty_id)
        REFERENCES faculties(id)
        ON DELETE RESTRICT,
    
    INDEX(subject_id) #TODO: need confirmation
);

COMMIT;
