package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ TargetingDictionariesModel = (*customTargetingDictionariesModel)(nil)

type (
	// TargetingDictionariesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTargetingDictionariesModel.
	TargetingDictionariesModel interface {
		BatchInsert(ctx context.Context, dictionaries []*TargetingDictionaries) error
	}

	customTargetingDictionariesModel struct {
		*defaultTargetingDictionariesModel
	}
)

// NewTargetingDictionariesModel returns a model for the database table.
func NewTargetingDictionariesModel(conn sqlx.SqlConn) TargetingDictionariesModel {
	return &customTargetingDictionariesModel{
		defaultTargetingDictionariesModel: newTargetingDictionariesModel(conn),
	}
}

func (m *defaultTargetingDictionariesModel) BatchInsert(ctx context.Context, dictionaries []*TargetingDictionaries) (err error) {
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, targetingDictionariesRows)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	err = m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		if _, err = session.ExecCtx(ctx, "TRUNCATE "+m.table); err != nil {
			return err
		}
		for i := 0; i < len(dictionaries); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?)")
			values = append(values,
				dictionaries[i].DictKey,
				dictionaries[i].Id,
				dictionaries[i].Pid,
				dictionaries[i].Label,
				dictionaries[i].Value,
				dictionaries[i].Code,
				dictionaries[i].Seq,
				dictionaries[i].DataStruct,
			)
			// 达到了 500 条数据，或最后一条了
			if chunk == 500 || i == len(dictionaries)-1 {
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
	return err
}
