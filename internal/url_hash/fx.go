package url_hash

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("url_hash", fx.Provide(NewURLHash))
}
