package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Root     string
	LogLevel string `yaml:"log_level"`
}

func LoadYaml(root string, path string, env string) (cfg *Config, err error) {
	root, err = filepath.Abs(root)
	if err != nil {
		return
	}

	fullpath := strings.Join([]string{root, path}, "/")
	var data []byte
	data, err = ioutil.ReadFile(fullpath)
	if err != nil {
		return
	}

	fullCfg := map[string]*Config{}
	err = yaml.Unmarshal([]byte(data), &fullCfg)
	if err != nil {
		return
	}
	cfg = fullCfg[env]
	cfg.Root = root

	return
}
