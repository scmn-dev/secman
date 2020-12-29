package factory

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cli/cli/git"
	"github.com/abdfnx/secman/v3/api/config"
	"github.com/abdfnx/secman/v3/api/common"
	"github.com/abdfnx/secman/v3/api/repox"
	"github.com/cli/cli/pkg/cmdutil"
	"github.com/cli/cli/pkg/iostreams"
)

func New(appVersion string) *cmdutil.Factory {
	io := iostreams.System()

	var cachedConfig config.Config
	var configError error
	configFunc := func() (config.Config, error) {
		if cachedConfig != nil || configError != nil {
			return cachedConfig, configError
		}
		cachedConfig, configError = config.ParseDefaultConfig()
		if errors.Is(configError, os.ErrNotExist) {
			cachedConfig = config.NewBlankConfig()
			configError = nil
		}
		cachedConfig = config.InheritEnv(cachedConfig)
		return cachedConfig, configError
	}

	hostOverride := ""
	if !strings.EqualFold(common.Default(), common.OverridableDefault()) {
		hostOverride = common.OverridableDefault()
	}

	rr := &remoteResolver{
		readRemotes: git.Remotes,
		getConfig:   configFunc,
	}
	remotesFunc := rr.Resolver(hostOverride)

	return &cmdutil.Factory{
		IOStreams: io,
		Config:    configFunc,
		Remotes:   remotesFunc,
		HttpClient: func() (*http.Client, error) {
			cfg, err := configFunc()
			if err != nil {
				return nil, err
			}

			return NewHTTPClient(io, cfg, appVersion, true), nil
		},
		BaseRepo: func() (repox.Interface, error) {
			remotes, err := remotesFunc()
			if err != nil {
				return nil, err
			}
			return remotes[0], nil
		},
		Branch: func() (string, error) {
			currentBranch, err := git.CurrentBranch()
			if err != nil {
				return "", fmt.Errorf("could not determine current branch: %w", err)
			}
			return currentBranch, nil
		},
	}
}
