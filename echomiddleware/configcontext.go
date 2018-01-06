package echomiddleware

import (
	"context"

	"github.com/labstack/echo"
)

const (
	ContextConfigName = "config_context"
)

func Config(ctx context.Context) *map[string]interface{} {
	v := ctx.Value(ContextConfigName)
	if v == nil {
		panic("config context is not exist")
	}
	if customConfig, ok := v.(*map[string]interface{}); ok {
		return customConfig
	}
	panic("config context is not exist")
}

func ContextConfig(d map[string]interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			c.SetRequest(req.WithContext(context.WithValue(req.Context(), ContextConfigName, &d)))
			return next(c)
		}
	}
}
