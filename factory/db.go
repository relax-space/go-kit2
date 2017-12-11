package factory

import (
	"context"

	"github.com/relax-space/go-kitt/echomiddleware"

	"github.com/go-xorm/xorm"
)

func DB(ctx context.Context) *xorm.Session {
	v := ctx.Value(echomiddleware.ContextDBName)
	if v == nil {
		panic("DB is not exist")
	}
	if db, ok := v.(*xorm.Session); ok {
		return db
	}
	if db, ok := v.(*xorm.Engine); ok {
		return db.NewSession()
	}
	panic("DB is not exist")
}
