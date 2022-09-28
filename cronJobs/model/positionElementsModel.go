package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PositionElementsModel = (*customPositionElementsModel)(nil)

type (
	// PositionElementsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionElementsModel.
	PositionElementsModel interface {
		positionElementsModel
		BatchInsert(ctx context.Context, elements []*PositionElements) (err error)
	}

	customPositionElementsModel struct {
		*defaultPositionElementsModel
	}
)

// NewPositionElementsModel returns a model for the database table.
func NewPositionElementsModel(conn sqlx.SqlConn) PositionElementsModel {
	return &customPositionElementsModel{
		defaultPositionElementsModel: newPositionElementsModel(conn),
	}
}

func (m *defaultPositionElementsModel) BatchInsert(ctx context.Context, elements []*PositionElements) (err error) {
	query := fmt.Sprintf("insert ignore into %s (%s) values ", m.table, positionElementsRowsExpectAutoSet)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		if _, err = session.ExecCtx(ctx, "TRUNCATE "+m.table); err != nil {
			return err
		}
		for i := 0; i < len(elements); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			values = append(values,
				elements[i].CreativeSizeId, elements[i].SubType, elements[i].GroupNumber, elements[i].ElementId, elements[i].ElementName,
				elements[i].ElementTitle, elements[i].ElementCaption, elements[i].Width, elements[i].Height,
				elements[i].MinWidth, elements[i].MinHeight, elements[i].MinLength, elements[i].MaxLength,
				elements[i].FileSizeKbLimit, elements[i].GifSizeKbLimit, elements[i].FileFormat, elements[i].Pattern,
				elements[i].Duration, elements[i].MinOccurs, elements[i].MaxOccurs,
			)
			// 达到了 200 条数据，或最后一条了
			if chunk == 200 || i == len(elements)-1 {
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
