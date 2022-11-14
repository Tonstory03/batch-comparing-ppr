package config

type Config struct {
	Application Application `yaml:"application"`
	Server      Server      `yaml:"server"`
	Elastic     Elastic     `yaml:"elastic"`
	Cronjob     []Cronjob   `yaml:"cronjob"`
}

type Application struct {
	Name    string `yaml:"name"`
	Profile string `yaml:"profile"`
}

type Server struct {
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}

type Cronjob struct {
	Name       string `yaml:"name"`
	Expression string `yaml:"expression"`
	Enable     bool   `yaml:"enable"`
}

type Elastic struct {
	Endpoint   string `yaml:"endpoint"`
	EnableAuth bool   `yaml:"enableAuth"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}
