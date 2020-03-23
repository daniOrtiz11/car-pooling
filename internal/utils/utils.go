package utils

import (
	"net/http"
	"os"
)

/*
Config environment struct

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}


Cfg environment var

var Cfg = getConfig("config.yml")

func getConfig(configFile string) Config {
	f, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

*/

/*
GetContentType is a
*/
func GetContentType(r *http.Request) string {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		//set default value
		contentType = "application/octet-stream"
	}
	return contentType
}

/*
GetAccept is a
*/
func GetAccept(r *http.Request) string {
	contentType := r.Header.Get("Accept")
	if contentType == "" {
		//set default value
		contentType = "application/octet-stream"
	}
	return contentType
}

/*
GetEnv is a
*/
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
