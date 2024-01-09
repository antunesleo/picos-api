package persistence

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	SQLXTx *sql.Tx
}

func (t *Transaction) GetSQLXTransaction() (*sql.Tx, error) {
	if t.SQLXTx != nil {
		return t.SQLXTx, nil
	}
	return nil, errors.New("No current transaction")
}

type TransactionManager interface {
	Begin() *Transaction
	Commit(t *Transaction) error
	Rollback(t *Transaction) error
}

type SQLXTransactionManager struct {
	DBConnection *sqlx.DB
}

func (tm *SQLXTransactionManager) Begin() (*Transaction, error) {
	tx, err := tm.DBConnection.Begin()
	if err != nil {
		return &Transaction{}, err
	}
	return &Transaction{tx}, err
}

func (tm *SQLXTransactionManager) Commit(t *Transaction) error {
	sqlxTx, err := t.GetSQLXTransaction()
	if err != nil {
		return err
	}
	sqlxTx.Commit()
	return nil
}

func (tm *SQLXTransactionManager) Rollback(t *Transaction) error {
	sqlxTx, err := t.GetSQLXTransaction()
	if err != nil {
		return err
	}
	sqlxTx.Rollback()
	return nil
}
