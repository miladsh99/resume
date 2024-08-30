package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"project1/dto"
)

func ReadConfig(fileName string) (*dto.ConfigDB, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var config dto.ConfigDB
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func GenerateConfig(config *dto.ConfigDB) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=%t",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
		config.Database.ParseTime,
	)
}
