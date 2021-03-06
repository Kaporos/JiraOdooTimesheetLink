package teamsnotifier

import (
	"bytes"
	"encoding/json"
	"jira-timesheet/jira"
	"net/http"
	"os"
)

type SectionFact struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type NotificationSection struct {
	Title string `json:"activityTitle"`
	Subtitle string `json:"activitySubtitle"`
	Image string `json:"activityImage"`
	Facts []SectionFact `json:"facts"`
	Markdown bool `json:"markdown"`
}

type Notification struct {
	Summary string `json:"summary"`
	Type string `json:"@type"`
	Context string `json:"@context"`
	Color string `json:"themecolor"`
	Sections []NotificationSection `json:"sections"`
}

func NewOdooError(request jira.Request, err error) Notification{
	lastWorklog := request.GetLastWorklog()
	return Notification{
		Summary: err.Error(),
		Type: "MessageCard",
		Context: "http://schema.org/extensions",
		Color: "EE0000",
		Sections: []NotificationSection{
			{
				Title: "Odoo Status",
				Subtitle: err.Error(),
				Image: "https://repository-images.githubusercontent.com/202264544/3ce58c00-19ab-11ea-8a01-81d62334b3ed",
				Facts: []SectionFact{
					{
						Name: "Issue Key",
						Value: request.Issue.Key,
					},
					{
						Name: "Employee",
						Value: request.User.Email,
					},
					{
						Name: "Worklog Time",
						Value: lastWorklog.TimeSpent,
					},
					{
						Name:"Timesheet Code",
						Value: request.Issue.Fields.TimesheetCode,
					},
				},
				Markdown: true,

			},
		},
	}
}


func NewJiraError(err error) Notification{
	return Notification{
		Summary: err.Error(),
		Type: "MessageCard",
		Context: "http://schema.org/extensions",
		Color: "EE0000",
		Sections: []NotificationSection{
			{
				Title: "Jira Status",
				Subtitle: err.Error(),
				Image: "https://repository-images.githubusercontent.com/202264544/3ce58c00-19ab-11ea-8a01-81d62334b3ed",
				Markdown: true,

			},
		},
	}
}

func NewSuccess(request jira.Request) Notification {
	var lastWorklog = request.GetLastWorklog()
	return Notification{
		Summary: "success",
		Type: "MessageCard",
		Context: "http://schema.org/extensions",
		Color: "00EE00",
		Sections: []NotificationSection{
			{
				Title: "Process Status",
				Subtitle: "Done",
				Image: "https://repository-images.githubusercontent.com/202264544/3ce58c00-19ab-11ea-8a01-81d62334b3ed",
				Facts: []SectionFact{
					{
						Name: "Issue Key",
						Value: request.Issue.Key,
					},
					{
						Name: "Employee",
						Value: request.User.Email,
					},
					{
						Name: "Worklog Time",
						Value: lastWorklog.TimeSpent,
					},
					{
						Name:"Timesheet Code",
						Value: request.Issue.Fields.TimesheetCode,
					},
				},
				Markdown: true,

			},
		},
	}
}



func Notify(notification Notification) {
	data, _ := json.Marshal(notification)
	http.Post(os.Getenv("TEAMS_WEBHOOK_URL"),"application/json",bytes.NewReader(data))
}
