CREATE TABLE assignee (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(250) UNIQUE NOT NULL,
    profile_picture_url TEXT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE project (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE issue (
    id VARCHAR(50),
    issue_type VARCHAR(50),
    assignee_id  VARCHAR(50) NOT NULL REFERENCES assignee(id) ON DELETE RESTRICT,
    project_id   VARCHAR(50) NOT NULL REFERENCES project(id) ON DELETE RESTRICT,
    status VARCHAR(20),
    story_points FLOAT DEFAULT 0
);
