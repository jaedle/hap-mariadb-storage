package storelib

func New() *MariaDbStore {
	return &MariaDbStore{}
}

type MariaDbStore struct {
}

func (m *MariaDbStore) Set(key string, value []byte) error {
	panic("implement me")
}

func (m *MariaDbStore) Get(key string) ([]byte, error) {
	panic("implement me")
}

func (m *MariaDbStore) Delete(key string) error {
	panic("implement me")
}

func (m *MariaDbStore) KeysWithSuffix(suffix string) ([]string, error) {
	panic("implement me")
}
