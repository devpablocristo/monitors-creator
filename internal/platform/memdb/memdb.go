package memdb

type MemDB map[string]any

func NewDB() *MemDB {
	db := make(MemDB)
	return &db
}
