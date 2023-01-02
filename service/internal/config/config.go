package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	c                = config{}
	GoogleMapsAPIKey = c.GoogleMapsAPIKey
	MySQL            = c.MySQL
	Redis            = c.Redis
)

func Configure() {
	configYaml, err := os.ReadFile(findConfig("test"))
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(configYaml, &c); err != nil {
		panic(err)
	}
	GoogleMapsAPIKey = c.GoogleMapsAPIKey
	MySQL = c.MySQL
	Redis = c.Redis
}

// 現在いる階層以上のディレクトリにあるconfig/xxx.ymlを探してパスを返す
func findConfig(env string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir := filepath.Clean(wd)
	for {
		configPath := filepath.Join(dir, "config", fmt.Sprintf("%s.yml", env))
		if fi, err := os.Stat(configPath); err == nil && !fi.IsDir() {
			return configPath
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}

		dir = parentDir
	}
	return ""
}

type config struct {
	GoogleMapsAPIKey string `yaml:"GoogleMapsAPIKey"`
	MySQL            struct {
		User   string `yaml:"User"`
		Passwd string `yaml:"Passwd"`
		Host   string `yaml:"Host"`
		Port   string `yaml:"Port"`
	} `yaml:"MySQL"`
	Redis struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
	} `yaml:"Redis"`
}
