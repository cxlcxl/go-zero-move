package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PositionBasePricesModel = (*customPositionBasePricesModel)(nil)

type (
	// PositionBasePricesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionBasePricesModel.
	PositionBasePricesModel interface {
		positionBasePricesModel
		BatchInsert(ctx context.Context, basePrices []*PositionBasePrices) (err error)
	}

	customPositionBasePricesModel struct {
		*defaultPositionBasePricesModel
	}
)

// NewPositionBasePricesModel returns a model for the database table.
func NewPositionBasePricesModel(conn sqlx.SqlConn) PositionBasePricesModel {
	return &customPositionBasePricesModel{
		defaultPositionBasePricesModel: newPositionBasePricesModel(conn),
	}
}

func (m *defaultPositionBasePricesModel) BatchInsert(ctx context.Context, basePrices []*PositionBasePrices) (err error) {
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, positionBasePricesRowsExpectAutoSet)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		if _, err = session.ExecCtx(ctx, "TRUNCATE "+m.table); err != nil {
			return err
		}
		for i := 0; i < len(basePrices); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?)")
			values = append(values, basePrices[i].CreativeSizeId, basePrices[i].PriceType, basePrices[i].BasePrice)
			// 达到了 300 条数据，或最后一条了
			if chunk == 800 || i == len(basePrices)-1 {
				// 写入库
				insertSQL := query + strings.Join(valueStatement, ",")
				if _, err = session.ExecCtx(ctx, insertSQL, values...); err != nil {
					return err
				}
				// 重置
				values, valueStatement = make([]interface{}, 0), make([]string, 0)
				chunk = 0
			}
			chunk++
		}
		return nil
	})
}
