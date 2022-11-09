delete from syllabuses;
delete from resources;
delete from subpages;
delete from subject_related_subjects;
delete from chapters;
delete from videos;
delete from subjects;


insert into `ocw-central-mock`.subjects
select * from `ocw-central`.subjects
where id in (
0x0182c9d30e1e008a640778c32f551522,
0x0182c9d30e211b82fe06bbcd18d34e61,
0x0182c9d30e253e1721b0210e69e71100,
0x0182c9d30e1e6cfd5a8177a879954bd6,
0x0182c9d30e20fee8bfa48f6c059dfb84,
0x0182c9d30e1e750a87edd9dbac80917b,
0x0182c9d30e25de094541a1466a05c5c1,
0x0182c9d30e1ebab8d2e3457474bc4faa,
0x0182c9d30e1edf7a7eea966633c65ce9
);

insert into videos
select
videos.id,
videos.subject_id,
videos.title,
videos.faculty,
videos.ordering,
videos.link,
videos.lectured_on,
videos.video_length,
videos.language,
videos.created_at,
videos.updated_at,
videos.transcription
from `ocw-central`.videos
join subjects on subjects.id = videos.subject_id;

insert into chapters
select
chapters.id,
chapters.video_id,
chapters.start_at,
chapters.topic,
chapters.thumbnail_link,
chapters.created_at,
chapters.updated_at
from `ocw-central`.chapters
join videos on videos.id = chapters.video_id;

insert into resources
select
resources.id,
resources.subject_id,
resources.title,
resources.description,
resources.ordering,
resources.link,
resources.created_at,
resources.updated_at
from `ocw-central`.resources
join subjects on subjects.id = resources.subject_id;

insert into syllabuses
select
syllabuses.id,
syllabuses.subject_id,
syllabuses.faculty,
syllabuses.language,
syllabuses.subject_numbering,
syllabuses.academic_year,
syllabuses.semester,
syllabuses.num_credit,
syllabuses.course_format,
syllabuses.assigned_grade,
syllabuses.targeted_audience,
syllabuses.course_day_period,
syllabuses.outline,
syllabuses.objective,
syllabuses.lesson_plan,
syllabuses.grading_method,
syllabuses.course_requirement,
syllabuses.outclass_learning,
syllabuses.reference,
syllabuses.remark,
syllabuses.created_at,
syllabuses.updated_at
from `ocw-central`.syllabuses
join subjects on subjects.id = syllabuses.subject_id;

insert into subpages
select
subpages.id,
subpages.subject_id,
subpages.content,
subpages.created_at,
subpages.updated_at
from `ocw-central`.subpages
join subjects on subjects.id = subpages.subject_id;

insert into subject_related_subjects
select
subject_related_subjects.subject_id,
subject_related_subjects.related_subject_id,
subject_related_subjects.created_at,
subject_related_subjects.updated_at
from `ocw-central`.subject_related_subjects
join subjects on subjects.id = subject_related_subjects.subject_id;
