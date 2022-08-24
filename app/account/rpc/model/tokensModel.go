package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TokensModel = (*customTokensModel)(nil)

type (
	// TokensModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTokensModel.
	TokensModel interface {
		tokensModel
	}

	customTokensModel struct {
		*defaultTokensModel
	}
)

// NewTokensModel returns a model for the database table.
func NewTokensModel(conn sqlx.SqlConn) TokensModel {
	return &customTokensModel{
		defaultTokensModel: newTokensModel(conn),
	}
}
