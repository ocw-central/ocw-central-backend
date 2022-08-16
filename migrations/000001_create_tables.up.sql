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

CREATE TABLE IF NOT EXISTS syllabuses(
    id BINARY(16) NOT NULL PRIMARY KEY,
    language VARCHAR(100),
    subject_numbering VARCHAR(100),
    academic_year INT,
    semester VARCHAR(100),
    num_credit TINYINT,
    course_format VARCHAR(100),
    assigned_grade VARCHAR(100),
    targeted_audience VARCHAR(100),
    day_of_week ENUM('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'),
    course_period VARCHAR(100),
    outline TEXT,
    objective TEXT,
    lesson_plan TEXT,
    grading_method TEXT,
    course_requirement TEXT,
    outclass_learning TEXT,
    reference TEXT,
    remark TEXT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    
);

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
);

CREATE TABLE IF NOT EXISTS faculties(
    id BINARY(16) NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    department VARCHAR(100),
    job_title VARCHAR(50),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,    
    INDEX(name)
);

CREATE TABLE IF NOT EXISTS pdfs(
    id BINARY(16) NOT NULL PRIMARY KEY,
    subject_id BINARY(16) NOT NULL,
    link VARCHAR(200) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (subject_id)
        REFERENCES subjects(id)
        ON DELETE RESTRICT,
);


CREATE TABLE IF NOT EXISTS subpages(
    id BINARY(16) NOT NULL PRIMARY KEY,
    subject_id BINARY(16) NOT NULL,
    link VARCHAR(200) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,    
    FOREIGN KEY (subject_id)
        REFERENCES subjects(id)
        ON DELETE RESTRICT,
);

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
);

CREATE TABLE IF NOT EXISTS subject_faculties(
    subject_id BINARY(16) NOT NULL,
    faculty_id BINARY(16) NOT NULL,
    PRIMARY KEY (subject_id, faculty_id),
    FOREIGN KEY (subject_id)
        REFERENCES subjects(id)
        ON DELETE RESTRICT,
    Foreign Key (faculty_id)
        REFERENCES faculties(id)
        ON DELETE RESTRICT,    
);

CREATE TABLE IF NOT EXISTS syllabus_faculties(
    syllabus_id BINARY(16) NOT NULL,
    faculty_id BINARY(16) NOT NULL,
    PRIMARY KEY (syllabus_id, faculty_id),    
    FOREIGN KEY (syllabus_id)
        REFERENCES syllabuses(id)
        ON DELETE RESTRICT,
    FOREIGN Key (faculty_id)
        REFERENCES faculties(id)
        ON DELETE RESTRICT,
);

CREATE TABLE IF NOT EXISTS video_faculties(
    video_id BINARY(16) NOT NULL,
    faculty_id BINARY(16) NOT NULL,
    PRIMARY KEY (video_id, faculty_id),    
    FOREIGN KEY (video_id)
        REFERENCES videos(id)
        ON DELETE RESTRICT,
    Foreign Key (faculty_id)
        REFERENCES faculties(id)
        ON DELETE RESTRICT,
);

COMMIT;
