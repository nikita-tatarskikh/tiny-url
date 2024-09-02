package url_generator

type Generator interface {
	GenerateShortURL(url string) string
}
