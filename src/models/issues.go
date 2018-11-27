package models

type Issue struct {
	Id           string
	IssueType    string
	AssigneeId   string
	ProjectId    string
	Status       string
	StoryPoints  float64
	ProjectName  string
	AssigneeName string
}

func GetAllIssues() ([]*Issue, error) {
	rows, err := db.Query(`
		SELECT I.id AS issue_id, I.issue_type, I.assignee_id, I.project_id, I.status, I.story_points, P.name AS project_name, A.name AS assignee_name FROM issue I
		LEFT JOIN project P ON I.project_id = P.id
		LEFT JOIN assignee A ON I.assignee_id = A.id
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	issues := make([]*Issue, 0)
	for rows.Next() {
		issue := new(Issue)
		err := rows.Scan(&issue.Id, &issue.IssueType, &issue.AssigneeId, &issue.ProjectId, &issue.Status, &issue.StoryPoints, &issue.ProjectName, &issue.AssigneeName)
		if err != nil {
			return nil, err
		}
		issues = append(issues, issue)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return issues, nil
}
