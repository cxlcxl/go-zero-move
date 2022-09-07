package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ JobsModel = (*customJobsModel)(nil)

type (
	// JobsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJobsModel.
	JobsModel interface {
		jobsModel
		GetJobs(ctx context.Context) (jobs []*Jobs, err error)
		UpdateJobDayByModule(ctx context.Context, module, day string) error
	}

	customJobsModel struct {
		*defaultJobsModel
	}
)

// NewJobsModel returns a model for the database table.
func NewJobsModel(conn sqlx.SqlConn) JobsModel {
	return &customJobsModel{
		defaultJobsModel: newJobsModel(conn),
	}
}

func (m *defaultJobsModel) GetJobs(ctx context.Context) (jobs []*Jobs, err error) {
	sql, _, err := squirrel.Select(jobsRows).From(m.table).ToSql()
	if err != nil {
		return nil, err
	}
	err = m.conn.QueryRowsCtx(ctx, &jobs, sql)
	switch err {
	case nil:
		return jobs, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJobsModel) UpdateJobDayByModule(ctx context.Context, module, day string) error {
	query := fmt.Sprintf("update %s set stat_day = ? where `api_module` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, day, module)
	return err
}
