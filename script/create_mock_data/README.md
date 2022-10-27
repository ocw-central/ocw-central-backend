The script to create [mock data](../../mock_data.md).


A subject has its related subjects, but due to foreign key constraints, the related subjects also have to be included in the subjects in mock data. This means selecting closed set of subjects is needed. To this end, BFS is used to find the closed set in `create_mock.py` and execute SQLs in `create_mock.sql` with found subjects.
