package report

import (
	"time"
)

type Repository interface {
	Save(id string, report Report) error
}

type Notifier interface {
	Notify(report Report) error
}

type Report struct {
	Id             string
	CreatedAt      time.Time
	LazyDevelopers int
	Details        []Detail
}

type Detail struct {
	DeveloperName  string
	Seniority      string
	TaskStatus     string
	LastActivityAt time.Time
}
