package make_short_url

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
	"net/http"
	"tiny-url/internal/dto"
	"tiny-url/internal/short_url"
)

type ShortUrlHandlerParams struct {
	fx.In

	ShortURL short_url.ShortUrl
}

type shortURLHandler struct {
	shortUrl short_url.ShortUrl
}

func NewShortURLHandler(params ShortUrlHandlerParams) fasthttp.RequestHandler {
	return shortURLHandler{shortUrl: params.ShortURL}.GetShortURL
}

func (h shortURLHandler) GetShortURL(ctx *fasthttp.RequestCtx) {
	var request dto.ShortURLRequest

	err := json.Unmarshal(ctx.PostBody(), &request)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	url, err := h.shortUrl.MakeShortURL(ctx, request.LongURL)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Redirect(url, http.StatusMovedPermanently)
}
