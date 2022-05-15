package storelib

import "database/sql"

func New(*sql.DB, string) *MariaDbStore {
	return &MariaDbStore{}
}

type MariaDbStore struct {
}

func (m *MariaDbStore) Set(key string, value []byte) error {
	return nil
}

func (m *MariaDbStore) Get(key string) ([]byte, error) {
	return nil, nil
}

func (m *MariaDbStore) Delete(key string) error {
	panic("implement me")
}

func (m *MariaDbStore) KeysWithSuffix(suffix string) ([]string, error) {
	panic("implement me")
}

func (m *MariaDbStore) Init() error {
	return nil
}
