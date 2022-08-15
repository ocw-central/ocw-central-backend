BEGIN;

CREATE TABLE IF NOT EXISTS faculties(
    id BINARY(16) NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    department VARCHAR(100),
    job_position VARCHAR(50),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    
    INDEX(name) #TODO: need confirmation
);

COMMIT;
