package db

import (
	"context"
	"database/sql"
	"fmt"
)

// store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new object of Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

var txKey = struct{}{}

// execTx executes the transaction . incase of error it rolls back, otherwise it commits the transaction
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		rbackErr := tx.Rollback()
		if rbackErr != nil {
			return fmt.Errorf("the error is %v %v", rbackErr, err)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
	Transfer    Transfer `json:"transfer"`
}

// TransferTx executes the money transaction . sends money from one account to another
// it creates transfer record, add 2 account entries(1 for from_account & 1 for to_account), and
// update account's balance within a single db transaction
func (s *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	txName := ctx.Value(txKey)
	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		// create transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "created transfer")
		// add acount entry for from_account
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "created entry1")
		// add acount entry for to_account
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "created entry2")
		// TODO: update account as it includes db lock

		// update the amount in fromAccount - debited
		fmt.Println(txName, "get account for update1")

		// To avoid deadlock, query statements order is also important
		// for ex: if we update account 1, do the operation on account1 again,
		// then do the operations on account2

		// TODO: put below repeated code in a single function
		if arg.FromAccountID < arg.ToAccountID {
			err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.FromAccountID,
				Amount: -arg.Amount,
			})
			if err != nil {
				return err
			}
			fmt.Println(txName, "update account1")
			// update the amount in ToAccount - credited
			err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.ToAccountID,
				Amount: arg.Amount,
			})
			if err != nil {
				return err
			}
			fmt.Println(txName, "update account2")
			return nil
		} else {
			err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.ToAccountID,
				Amount: arg.Amount,
			})
			if err != nil {
				return err
			}
			fmt.Println(txName, "update account1")
			// update the amount in ToAccount - credited
			err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.FromAccountID,
				Amount: -arg.Amount,
			})
			if err != nil {
				return err
			}
			fmt.Println(txName, "update account2")
			return nil

		}
	})
	return result, err

}
