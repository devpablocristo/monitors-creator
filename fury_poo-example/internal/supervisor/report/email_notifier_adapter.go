package report

import (
	"encoding/json"
	"users_example/internal/platform/emails"
)

type emailNotifier struct {
	emails *emails.Client
}

func NewEmailNotifier(client *emails.Client) *emailNotifier {
	return &emailNotifier{client}
}

func (n emailNotifier) Notify(report Report) error {
	bb, err := json.Marshal(report)
	if err != nil {
		return err
	}

	if err := n.emails.Send(bb, "supervisor@mercadolibre.com"); err != nil {
		return err
	}

	return nil
}
