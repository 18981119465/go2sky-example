package data

import (
	"context"

	"github.com/18981119465/go2sky-example/model"
)

// IData is interface for accessing storage.
type IData interface {
	OpenDb() error
	QueryAccount(ctx context.Context, accountID string) (model.Account, error)
	Seed()
	Check() bool
}
