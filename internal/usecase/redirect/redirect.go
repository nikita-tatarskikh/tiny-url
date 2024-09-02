package redirect

import (
	"github.com/valyala/fasthttp"
	"tiny-url/internal/storage"
)

type RedirectSvc struct {
	storage storage.Storage
}

func NewRedirectSvc(storage storage.Storage) *RedirectSvc {
	return &RedirectSvc{storage: storage}
}

func (h RedirectSvc) Redirect(ctx *fasthttp.RequestCtx) {
	longURL, err := h.storage.Get(ctx, string(ctx.QueryArgs().Peek("short-url")))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	if longURL == "" {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.Redirect(longURL, fasthttp.StatusMovedPermanently)
}
