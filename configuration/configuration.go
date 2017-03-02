package configuration

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/DaveBlooman/awsr/output"
	"github.com/go-ini/ini"
	homedir "github.com/mitchellh/go-homedir"
)

type Config map[string]map[string]string

func Load() (*ini.File, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return nil, errors.New("Couldn't determine home directory")
	}
	path := filepath.Join(homeDir, ".aws/credentials")

	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		output.Error("opening config file" + err.Error())
	}

	cfg, err := ini.Load(configFile)

	if err != nil {
		return nil, err
	}

	return cfg, nil

}
