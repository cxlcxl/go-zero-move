package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TokensModel = (*customTokensModel)(nil)

type (
	// TokensModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTokensModel.
	TokensModel interface {
		tokensModel
		GetAccessTokenList(ctx context.Context) ([]*Tokens, error)
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
func (m *defaultTokensModel) GetAccessTokenList(ctx context.Context) (tokens []*Tokens, err error) {
	query := fmt.Sprintf("select %s from %s", tokensRows, m.table)
	err = m.conn.QueryRowsCtx(ctx, &tokens, query)
	return
}
