package main

import "net/http"
import "fmt"
import "encoding/json"
import "io/ioutil"
import "log"

type IssueField struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type IssueAssignee struct {
	*IssueField
	Key          string `json:"key"`
	EmailAddress string `json:"emailAddress"`
	AccountId    string `json:"accountId"`
	DisplayName  string `json:"displayName"`
}

type Issue struct {
	Id     string `json:"id"`
	Fields struct {
		StoryPoints float64       `json:"customfield_10020"`
		IssueType   IssueField    `json:"issuetype"`
		Resolution  IssueField    `json:"resolution"`
		Assignee    IssueAssignee `json:"assignee"`
		Project     IssueField    `json:"project"`
	}
}

func getIssue(body []byte) (*Issue, error) {
	var issue = new(Issue)
	err := json.Unmarshal(body, issue)
	return issue, err
}

func main() {
	var BASE_ISSUE_DETAIL_URL = "https://aplijobs.atlassian.net/rest/api/3/issue"
	var BASIC_TOKEN = "aGVub2NkekBnbWFpbC5jb206ZFpyMkd3aE92U3M0ZldYNFdydGFBREMy"
	var issueKey = "EVA-668"
	var issueUrl = fmt.Sprintf("%s/%s", BASE_ISSUE_DETAIL_URL, issueKey)

	client := &http.Client{}
	req, err := http.NewRequest("GET", issueUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	var authorizationHeader = fmt.Sprintf("Basic %s", BASIC_TOKEN)
	req.Header.Add("Authorization", authorizationHeader)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	issue, err := getIssue([]byte(body))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(issue.Fields.StoryPoints)
}
