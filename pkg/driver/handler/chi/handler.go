package chi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

type ctx struct {
	ctx     context.Context
	request *http.Request
	writer  http.ResponseWriter
}

func (c *ctx) WriteContextValue(key, val interface{}) {
	c.ctx = context.WithValue(c.ctx, key, val)
}

func (c *ctx) ContextValue(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func (c *ctx) Context() context.Context {
	return c.ctx
}

func (c *ctx) JSON(status int, v interface{}) {
	c.writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.writer.WriteHeader(status)
	if err := json.NewEncoder(c.writer).Encode(v); err != nil {
		http.Error(c.writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *ctx) Param(key string) string {
	return chi.URLParam(c.request, key)
}

func (c *ctx) Query(key string) string {
	return c.request.FormValue(key)
}

func (c *ctx) BodyToJSON(v interface{}) error {
	if err := json.NewDecoder(c.request.Body).Decode(v); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
