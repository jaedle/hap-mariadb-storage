package storelib

import (
	"database/sql"
	"fmt"
)

func New(db *sql.DB, table string) *MariaDbStore {
	return &MariaDbStore{
		table: table,
		db:    db,
	}
}

type MariaDbStore struct {
	table string
	db    *sql.DB
}

func (m *MariaDbStore) Init() error {
	_, err := m.db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (`key` varchar(255), `value` MEDIUMBLOB, CONSTRAINT PK PRIMARY KEY (`key`)) ;", m.table))
	return err
}

func (m *MariaDbStore) Set(key string, value []byte) error {
	_, err := m.db.Exec(fmt.Sprintf("INSERT INTO `%s` (`key`, `value`) VALUES (?,?) ON DUPLICATE KEY UPDATE `value`=?", m.table), key, value, value)
	return err
}

func (m *MariaDbStore) Get(key string) ([]byte, error) {
	row := m.db.QueryRow(fmt.Sprintf("SELECT `value` FROM `%s` WHERE `key` = ?;", m.table), key)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var result []byte
	if err := row.Scan(&result); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (m *MariaDbStore) Delete(key string) error {
	panic("implement me")
}

func (m *MariaDbStore) KeysWithSuffix(suffix string) ([]string, error) {
	panic("implement me")
}
