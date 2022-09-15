package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContinentsModel = (*customContinentsModel)(nil)

type (
	// ContinentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContinentsModel.
	ContinentsModel interface {
		continentsModel
		Continents(ctx context.Context) (continents []*Continents, err error)
	}

	customContinentsModel struct {
		*defaultContinentsModel
	}
)

// NewContinentsModel returns a model for the database table.
func NewContinentsModel(conn sqlx.SqlConn) ContinentsModel {
	return &customContinentsModel{
		defaultContinentsModel: newContinentsModel(conn),
	}
}

func (m *defaultContinentsModel) Continents(ctx context.Context) (continents []*Continents, err error) {
	query := fmt.Sprintf("select * from %s", m.table)
	err = m.conn.QueryRowsCtx(ctx, &continents, query)
	switch err {
	case nil:
		return continents, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
