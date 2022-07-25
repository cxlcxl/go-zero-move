package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ RolesModel = (*customRolesModel)(nil)

type (
	// RolesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRolesModel.
	RolesModel interface {
		rolesModel
		RoleList(ctx context.Context, roleName string, state int64, ids []int64) ([]*Roles, error)
	}

	customRolesModel struct {
		*defaultRolesModel
	}
)

// NewRolesModel returns a model for the database table.
func NewRolesModel(conn sqlx.SqlConn, c cache.CacheConf) RolesModel {
	return &customRolesModel{
		defaultRolesModel: newRolesModel(conn, c),
	}
}

func (m *defaultRolesModel) RoleList(ctx context.Context, roleName string, state int64, ids []int64) ([]*Roles, error) {
	var roles []*Roles
	sb := squirrel.Select(rolesRows).From(m.table)
	if state >= 0 {
		sb = sb.Where("state = ?", state)
	}
	if len(ids) > 0 {
		s, v := whereIntIn(ids)
		sb = sb.Where("id in "+s, v...)
	}
	if len(roleName) > 0 {
		sb = sb.Where("role_name like ?", "%"+roleName+"%")
	}
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	err = m.QueryRowsNoCacheCtx(ctx, &roles, query, args...)
	switch err {
	case nil:
		return roles, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func whereIntIn(ids []int64) (w string, args []interface{}) {
	query := make([]string, 0)
	for i := range ids {
		query = append(query, "?")
		args = append(args, ids[i])
	}

	return "(" + strings.Join(query, ",") + ")", args
}
