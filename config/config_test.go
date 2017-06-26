package config

import (
	"fmt"
	"testing"
)

func TestLoadYaml(t *testing.T) {
	var (
		cfg *Config
		err error
	)

	cfg, err = LoadYaml(".", "testdata/not_exist.yaml", "development")
	if err == nil {
		fmt.Println("when file is not exist, should err is not nil")
		t.Errorf(err.Error())
	}

	cfg, err = LoadYaml(".", "testdata/sample.yaml", "development")
	if err != nil {
		fmt.Println("when file is exist, should err is nil")
		t.Errorf(err.Error())
	}

	cfg, err = LoadYaml(".", "testdata/sample.yaml", "development")
	if cfg.LogLevel != "debug" {
		t.Errorf(`cfg.LogLevel is not equal:
		expect: debug
		got:    %s`, cfg.LogLevel)
	}

	cfg, err = LoadYaml(".", "testdata/sample.yaml", "production")
	if cfg.LogLevel != "warn" {
		t.Errorf(`cfg.LogLevel is not equal:
		expect: warn
		got:    %s`, cfg.LogLevel)
	}
}
