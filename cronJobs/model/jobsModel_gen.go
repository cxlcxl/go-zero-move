// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	jobsFieldNames          = builder.RawFieldNames(&Jobs{})
	jobsRows                = strings.Join(jobsFieldNames, ",")
	jobsRowsExpectAutoSet   = strings.Join(stringx.Remove(jobsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	jobsRowsWithPlaceHolder = strings.Join(stringx.Remove(jobsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	jobsModel interface {
		Insert(ctx context.Context, data *Jobs) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Jobs, error)
		FindOneByApiModule(ctx context.Context, apiModule string) (*Jobs, error)
		Update(ctx context.Context, newData *Jobs) error
		Delete(ctx context.Context, id int64) error
	}

	defaultJobsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Jobs struct {
		Id          int64     `db:"id"`
		StatDay     time.Time `db:"stat_day"`     // 数据日期
		ApiModule   string    `db:"api_module"`   // 数据模块
		TotalPage   int64     `db:"total_page"`   // 总页数
		TotalNum    int64     `db:"total_num"`    // 总数
		JobSchedule string    `db:"job_schedule"` // cron 调度
		PauseRule   int64     `db:"pause_rule"`   // 调度截止规则：0 调度到当天；-1 停止调度此任务；> 0 为当前日期减{pause_rule}天
		Version     int64     `db:"version"`      // 版本：每次有规则或调度修改 +1
	}
)

func newJobsModel(conn sqlx.SqlConn) *defaultJobsModel {
	return &defaultJobsModel{
		conn:  conn,
		table: "`jobs`",
	}
}

func (m *defaultJobsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultJobsModel) FindOne(ctx context.Context, id int64) (*Jobs, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", jobsRows, m.table)
	var resp Jobs
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJobsModel) FindOneByApiModule(ctx context.Context, apiModule string) (*Jobs, error) {
	var resp Jobs
	query := fmt.Sprintf("select %s from %s where `api_module` = ? limit 1", jobsRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, apiModule)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultJobsModel) Insert(ctx context.Context, data *Jobs) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, jobsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.StatDay, data.ApiModule, data.TotalPage, data.TotalNum, data.JobSchedule, data.PauseRule, data.Version)
	return ret, err
}

func (m *defaultJobsModel) Update(ctx context.Context, newData *Jobs) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, jobsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.StatDay, newData.ApiModule, newData.TotalPage, newData.TotalNum, newData.JobSchedule, newData.PauseRule, newData.Version, newData.Id)
	return err
}

func (m *defaultJobsModel) tableName() string {
	return m.table
}
