package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ TargetingDictionariesModel = (*customTargetingDictionariesModel)(nil)

type (
	// TargetingDictionariesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTargetingDictionariesModel.
	TargetingDictionariesModel interface {
		FindDictionaries(ctx context.Context, keys []string) (dictionaries []*TargetingDictionaries, err error)
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

func (m *defaultTargetingDictionariesModel) FindDictionaries(ctx context.Context, keys []string) (dictionaries []*TargetingDictionaries, err error) {
	sb := squirrel.Select(targetingDictionariesRows).From(m.table)
	if len(keys) > 0 {
		s, v := whereIn(keys)
		sb = sb.Where("dict_key in "+s, v...)
	}
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	err = m.conn.QueryRowsCtx(ctx, &dictionaries, query, args...)
	switch err {
	case nil:
		return dictionaries, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// where in 条件组合
func whereIn(s []string) (w string, args []interface{}) {
	query := make([]string, 0)
	for i := range s {
		query = append(query, "?")
		args = append(args, s[i])
	}

	return "(" + strings.Join(query, ",") + ")", args
}
