package report

import "users_example/internal/platform/localdb"

type localRepo struct {
	db *localdb.LocalDB
}

func NewLocalRepo(db *localdb.LocalDB) *localRepo {
	return &localRepo{db: db}
}

func (l localRepo) Save(id string, report Report) error {
	l.db.SaveItem(id, report)
	return nil
}
