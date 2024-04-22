package monitor

import "monitors-creator/internal/platform/memdb"

func NewMemoryRepo(db *memdb.MemDB) MonitorRepositoryPort {
	return &MemRepo{
		db: *db,
	}
}

type MemRepo struct {
	db memdb.MemDB
}

func (m *MemRepo) Create(monitor *Monitor) error {
	m.db[monitor.ID.String()] = monitor
	return nil
}

// func (l localRepo) Save(id string, report Report) error {
// 	l.db.SaveItem(id, report)
// 	return nil
// }

// func (db LocalDB) SaveItem(id string, item interface{}) {
// 	db.storage[id] = item
// }

// func (db LocalDB) GetItem(id string) interface{} {
// 	item, ok := db.storage[id]
// 	if !ok {
// 		return nil
// 	}

// 	return item
// }

// func (db LocalDB) DeleteItem(id string) {
// 	delete(db.storage, id)
// }

// func (db LocalDB) Dump() []interface{} {
// 	var items []interface{}
// 	for _, value := range db.storage {
// 		items = append(items, value)
// 	}

// 	return items
// }
