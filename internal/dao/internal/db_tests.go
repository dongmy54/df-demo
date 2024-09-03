// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DbTestsDao is the data access object for table db_tests.
type DbTestsDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns DbTestsColumns // columns contains all the column names of Table for convenient usage.
}

// DbTestsColumns defines and stores column names for table db_tests.
type DbTestsColumns struct {
	A string //
	B string //
	C string //
}

// dbTestsColumns holds the columns for table db_tests.
var dbTestsColumns = DbTestsColumns{
	A: "a",
	B: "b",
	C: "c",
}

// NewDbTestsDao creates and returns a new DAO object for table data access.
func NewDbTestsDao() *DbTestsDao {
	return &DbTestsDao{
		group:   "default",
		table:   "db_tests",
		columns: dbTestsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DbTestsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DbTestsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DbTestsDao) Columns() DbTestsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DbTestsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DbTestsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DbTestsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
