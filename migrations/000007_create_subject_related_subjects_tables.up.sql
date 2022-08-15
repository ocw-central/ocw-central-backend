BEGIN;

CREATE TABLE IF NOT EXISTS subject_related_subjects(
    subject_id BINARY(16) NOT NULL,
    related_subject_id BINARY(16) NOT NULL,
    PRIMARY KEY (subject_id, related_subject_id),

    FOREIGN KEY (subject_id)
        REFERENCES subjects(id)
        ON DELETE RESTRICT,
    
    FOREIGN KEY (related_subject_id)
        REFERENCES subjects(id)
        ON DELETE RESTRICT,
    
    INDEX(subject_id) #TODO: need checking
);

COMMIT;
