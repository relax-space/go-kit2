package echomiddleware

import (
	"context"
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

const ContextDBName = "DB"

func ContextDB(db *xorm.Engine) echo.MiddlewareFunc {
	db.ShowExecTime()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			switch req.Method {
			case "POST", "PUT", "DELETE":
				session := db.NewSession()
				defer session.Close()

				c.SetRequest(req.WithContext(context.WithValue(req.Context(), ContextDBName, session)))

				if err := session.Begin(); err != nil {
					log.Println(err)
				}
				if err := next(c); err != nil {
					session.Rollback()
					return err
				}
				if c.Response().Status >= 500 {
					session.Rollback()
					return nil
				}
				if err := session.Commit(); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
			default:
				c.SetRequest(req.WithContext(context.WithValue(req.Context(), ContextDBName, db)))
				return next(c)
			}

			return nil
		}
	}
}
