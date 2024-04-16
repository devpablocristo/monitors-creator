package furyapp

import (
	"github.com/mercadolibre/fury_go-seq-client"
	"github.com/mercadolibre/go-meli-toolkit/godsclient"
	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
	"github.com/mercadolibre/go-meli-toolkit/goosclient"
	"users_example/internal/platform/emails"
	"users_example/internal/platform/environment"
	"users_example/internal/platform/localdb"
	"users_example/internal/supervisor/developer"
	"users_example/internal/supervisor/developer/task"
	"users_example/internal/supervisor/report"
)

type Dependencies struct {
	DeveloperRepository developer.Repository
	ReportRepository    report.Repository
	ReportNotifier      report.Notifier
	TaskPublisher       task.Publisher
}

func BuildDependencies(env environment.Environment) (*Dependencies, error) {
	switch env {
	case environment.Production:
		devDS := godsclient.NewEntityClient(godsclient.NewDsClientConfig().WithServiceName("prod-dev-ds"))
		devKVS := gokvsclient.MakeKvsClient("prod-dev-kvs", gokvsclient.MakeKvsConfig())
		sequenceService := gosequence.MakeSequenceClient("prod-sequence", 1000)
		reportsOS := goosclient.MakeOsClient("prod-reports-os", goosclient.MakeOSClientConfigRead(), goosclient.MakeOSClientConfigWrite())
		emailsClient := emails.NewClient()

		// infra adapters
		reportNotifier := report.NewEmailNotifier(emailsClient)
		reportRepo := report.NewOSRepo(reportsOS)
		devRepo := developer.NewFuryRepo(devKVS, sequenceService, devDS)
		taskPublisher := task.NewFakePublisher()

		return &Dependencies{
			DeveloperRepository: devRepo,
			ReportRepository:    reportRepo,
			ReportNotifier:      reportNotifier,
			TaskPublisher:       taskPublisher,
		}, nil
	case environment.Beta:
		devDS := godsclient.NewEntityClient(godsclient.NewDsClientConfig().WithServiceName("beta-dev-ds"))
		devKVS := gokvsclient.MakeKvsClient("beta-dev-kvs", gokvsclient.MakeKvsConfig())
		sequenceService := gosequence.MakeSequenceClient("beta-sequence", 1000)
		reportsOS := goosclient.MakeOsClient("beta-reports-os", goosclient.MakeOSClientConfigRead(), goosclient.MakeOSClientConfigWrite())
		emailsClient := emails.NewClient()

		// infra adapters
		reportNotifier := report.NewEmailNotifier(emailsClient)
		reportRepo := report.NewOSRepo(reportsOS)
		devRepo := developer.NewFuryRepo(devKVS, sequenceService, devDS)
		taskPublisher := task.NewFakePublisher()

		return &Dependencies{
			DeveloperRepository: devRepo,
			ReportRepository:    reportRepo,
			ReportNotifier:      reportNotifier,
			TaskPublisher:       taskPublisher,
		}, nil
	case environment.Development:
		localDb := localdb.New()

		// infra adapters
		reportNotifier := report.NewFakeNotifier()
		reportRepo := report.NewLocalRepo(localDb)
		devRepo := developer.NewLocalRepo(localDb)
		taskPublisher := task.NewFakePublisher()

		return &Dependencies{
			DeveloperRepository: devRepo,
			ReportRepository:    reportRepo,
			ReportNotifier:      reportNotifier,
			TaskPublisher:       taskPublisher,
		}, nil
	}

	return nil, nil
}
