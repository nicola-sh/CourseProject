package hikedb

import (
	"context"
	"database/sql"
	"fmt"
)

// Provides all functions to execute db queries and transaction
type Store struct {
	*Queries
	db *sql.DB
}

// Create a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a functions within database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// // contains the ipnut parametres of the transfer transaction
// type TransferTxParams struct {
// 	FromUserID int64
// 	ToUserId   int64
// 	amount     int64
// }

// // is the result of the transfer transaction
// type TransferTxResult struct {
// 	Transfer    Transfer
// 	FromAccount Account
// 	ToAccount   Account
// 	FromEntry   Entry
// 	ToEntry     Entry
// }

// // from 1 user to 2 user
// func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {

// }
