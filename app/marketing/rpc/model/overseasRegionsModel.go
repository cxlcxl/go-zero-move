package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OverseasRegionsModel = (*customOverseasRegionsModel)(nil)

type (
	// OverseasRegionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOverseasRegionsModel.
	OverseasRegionsModel interface {
		overseasRegionsModel
		GetCountries(ctx context.Context) (countries []*OverseasRegions, err error)
	}

	customOverseasRegionsModel struct {
		*defaultOverseasRegionsModel
	}
)

// NewOverseasRegionsModel returns a model for the database table.
func NewOverseasRegionsModel(conn sqlx.SqlConn) OverseasRegionsModel {
	return &customOverseasRegionsModel{
		defaultOverseasRegionsModel: newOverseasRegionsModel(conn),
	}
}

func (m *defaultOverseasRegionsModel) GetCountries(ctx context.Context) (countries []*OverseasRegions, err error) {
	query := fmt.Sprintf("select %s from %s where `level` = 0 order by c_name asc", overseasRegionsRows, m.table)
	err = m.conn.QueryRowsCtx(ctx, &countries, query)
	switch err {
	case nil:
		return countries, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
