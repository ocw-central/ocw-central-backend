BEGIN;

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
    
    INDEX(subject_numbering) #FIXME
);

COMMIT;
