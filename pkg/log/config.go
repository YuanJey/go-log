package log

import "github.com/YuanJey/env-config/pkg/load"

var Config config

type config struct {
	Log struct {
		StorageLocation       string   `env:"STORAGE_LOCATION" def:"./logs/"`
		RotationTime          int      `env:"ROTATION_TIME" def:"24"`
		RemainRotationCount   int      `env:"REMAIN_ROTATION_COUNT" def:"2"`
		RemainLogLevel        int      `env:"REMAIN_LOG_LEVEL" def:"6"`
		ElasticSearchSwitch   bool     `env:"ELASTIC_SEARCH_SWITCH" def:"false"`
		ElasticSearchAddr     []string `env:"ELASTIC_SEARCH_ADDR" def:"http://127.0.0.1:9200/"`
		ElasticSearchUser     string   `env:"ELASTIC_SEARCH_USER" def:""`
		ElasticSearchPassword string   `env:"ELASTIC_SEARCH_PASSWORD" def:""`
	} `yaml:"log"`
}

func init() {
	err := load.LoadEnv(&Config)
	if err != nil {
		panic(err)
	}
}
