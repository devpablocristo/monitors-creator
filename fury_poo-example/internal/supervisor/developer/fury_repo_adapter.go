package developer

import (
	"github.com/mercadolibre/fury_go-seq-client"
	"github.com/mercadolibre/go-meli-toolkit/godsclient"
	"github.com/mercadolibre/go-meli-toolkit/godsclient/querybuilders"
	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
	"users_example/internal/supervisor/developer/task"
)

type furyRepo struct {
	kvs      gokvsclient.Client
	ds       godsclient.Client
	sequence gosequence.Client
}

func NewFuryRepo(kvs gokvsclient.Client, sequence gosequence.Client, ds godsclient.Client) *furyRepo {
	return &furyRepo{kvs, ds, sequence}
}

func (repo furyRepo) Save(d *Developer) error {
	id, err := repo.sequence.GetID()
	if err != nil {
		return err
	}

	item := gokvsclient.MakeItem(id.String(), *d)
	if err := repo.kvs.Save(item); err != nil {
		return err
	}

	return nil
}

func (repo furyRepo) Get(id string) (*Developer, error) {
	item, err := repo.kvs.Get(id)
	if err != nil {
		return nil, err
	}

	var dev Developer
	if err := item.GetValue(&dev); err != nil {
		return nil, err
	}

	return &dev, nil
}

func (repo furyRepo) Delete(id string) error {
	return nil
}

func (repo furyRepo) SearchByStatus(status task.Status) ([]Developer, error) {
	queryBuilder := querybuilders.And(
		querybuilders.Eq("status", status.String()),
	)
	response, err := repo.ds.SearchBuilder().WithQuery(queryBuilder).Execute()
	if err != nil {
		return nil, err
	}

	var developers []Developer
	for _, dsItem := range response.Documents {
		var developer Developer
		err := dsItem.To(&developer)
		if err != nil {
			return nil, err
		}

		developers = append(developers, developer)
	}
	return developers, nil
}
