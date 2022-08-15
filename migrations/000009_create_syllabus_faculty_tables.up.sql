BEGIN;

CREATE TABLE IF NOT EXISTS syllabus_faculty(
    syllabus_id BINARY(16) NOT NULL,
    faculty_id BINARY(16) NOT NULL,
    PRIMARY KEY (syllabus_id, faculty_id),
    
    FOREIGN KEY (syllabus_id)
        REFERENCES syllabuses(id)
        ON DELETE RESTRICT,

    FOREIGN Key (faculty_id)
        REFERENCES faculties(id)
        ON DELETE RESTRICT,
    
    INDEX(syllabus_id) #TODO: need confirmation
);

COMMIT;
