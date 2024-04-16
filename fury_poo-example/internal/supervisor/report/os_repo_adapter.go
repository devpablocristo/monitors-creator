package report

import (
	"encoding/json"
	"github.com/mercadolibre/go-meli-toolkit/goosclient"
)

type osRepo struct {
	os goosclient.Client
}

func NewOSRepo(os goosclient.Client) *osRepo {
	return &osRepo{os}
}

func (repo osRepo) Save(id string, report Report) error {
	bb, err := json.Marshal(report)
	if err != nil {
		return err
	}

	if err := repo.os.Put(id, bb); err != nil {
		return err
	}

	return nil
}
