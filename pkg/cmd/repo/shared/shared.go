package shared

import (
	"github.com/cli/cli/api"
	"github.com/cli/cli/internal/config"
	"github.com/cli/cli/internal/ghrepo"
	"github.com/cli/cli/pkg/cmdutil"
)

// shared.NewRepo is a small wrapper around cmdutil.NewRepo
// providing sane fallback functions that are commonly needed in the repo commands
func NewRepo(nwo string, config func() (config.Config, error), client *api.Client) (ghrepo.Interface, error) {
	cfg, err := config()
	if err != nil {
		return nil, err
	}

	host, err := cfg.DefaultHost()
	if err != nil {
		return nil, err
	}

	hostFallbackFunc := func() (string, error) {
		return host, nil
	}

	ownerFallbackFunc := func() (string, error) {
		if client != nil {
			return api.CurrentLoginName(client, host)
		}
		return "", nil
	}

	return cmdutil.NewRepo(nwo, hostFallbackFunc, ownerFallbackFunc)
}