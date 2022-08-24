package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AccountsModel = (*customAccountsModel)(nil)

type (
	// AccountsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountsModel.
	AccountsModel interface {
		accountsModel
	}

	customAccountsModel struct {
		*defaultAccountsModel
	}
)

// NewAccountsModel returns a model for the database table.
func NewAccountsModel(conn sqlx.SqlConn) AccountsModel {
	return &customAccountsModel{
		defaultAccountsModel: newAccountsModel(conn),
	}
}
