package dto

type ConfigDB struct {
	Database struct {
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		DBName    string `yaml:"dbname"`
		ParseTime bool   `yaml:"parseTime"`
	} `yaml:"database"`
}
