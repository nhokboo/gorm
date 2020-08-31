package tests

import (
	"github.com/nhokboo/gorm"
	"github.com/nhokboo/gorm/clause"
	"github.com/nhokboo/gorm/logger"
	"github.com/nhokboo/gorm/schema"
)

type DummyDialector struct {
}

func (DummyDialector) Initialize(*gorm.DB) error {
	return nil
}

func (DummyDialector) Migrator(*gorm.DB) gorm.Migrator {
	return nil
}

func (DummyDialector) BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{}) {
	writer.WriteByte('?')
}

func (DummyDialector) QuoteTo(writer clause.Writer, str string) {
	writer.WriteByte('`')
	writer.WriteString(str)
	writer.WriteByte('`')
}

func (DummyDialector) Explain(sql string, vars ...interface{}) string {
	return logger.ExplainSQL(sql, nil, `"`, vars...)
}

func (DummyDialector) DataTypeOf(*schema.Field) string {
	return ""
}
