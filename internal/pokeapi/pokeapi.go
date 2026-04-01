package pokeapi

import (
	"net/http"
	"time"

	"github.com/thegreatestgiant/Go-Pokemon/internal/pokecache"
	"github.com/thegreatestgiant/Go-Pokemon/internal/theme"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
	theme      *theme.CLITheme
	themeFunc  *theme.CLIThemeFunc
	Debug      bool
}

func NewClient(cacheInterval time.Duration, isDebug bool) Client {
	appTheme := theme.LoadTheme()
	appThemeFunc := theme.LoadThemeFunc()

	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache:     pokecache.NewCache(cacheInterval),
		theme:     appTheme,
		themeFunc: appThemeFunc,
		Debug:     isDebug,
	}
}
