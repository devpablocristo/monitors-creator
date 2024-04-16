package supervisor

import (
	"github.com/google/uuid"
	"time"
	"users_example/internal/supervisor/developer"
	"users_example/internal/supervisor/developer/task"
	"users_example/internal/supervisor/report"
)

type ReporterUseCase interface {
	LazyReport() (*report.Report, error)
}

type useCaseLazyReporter struct {
	reportRepo     report.Repository
	reportNotifier report.Notifier
	devRepo        developer.Repository
}

func NewLazyReporterUseCase(reportRepo report.Repository, reportNotifier report.Notifier, devRepo developer.Repository) *useCaseLazyReporter {
	return &useCaseLazyReporter{reportRepo: reportRepo, reportNotifier: reportNotifier, devRepo: devRepo}
}

func (u useCaseLazyReporter) LazyReport() (*report.Report, error) {
	lazyDevs, err := u.devRepo.SearchByStatus(task.StatusNone)
	if err != nil {
		return nil, err
	}

	var r report.Report
	if len(lazyDevs) == 0 {
		r = report.Report{
			Id:             uuid.New().String(),
			CreatedAt:      time.Now(),
			LazyDevelopers: 0,
			Details:        nil,
		}
	} else {
		r = report.Report{
			Id:             uuid.New().String(),
			CreatedAt:      time.Now(),
			LazyDevelopers: len(lazyDevs),
		}

		var details []report.Detail
		for _, dev := range lazyDevs {
			details = append(details, report.Detail{
				DeveloperName:  dev.Name,
				Seniority:      dev.Seniority.String(),
				TaskStatus:     dev.Task.Status.String(),
				LastActivityAt: dev.Task.UpdatedAt,
			})
		}

		r.Details = details
	}

	if err := u.reportRepo.Save(r.Id, r); err != nil {
		return nil, err
	}

	if err := u.reportNotifier.Notify(r); err != nil {
		return nil, err
	}

	return &r, nil
}
