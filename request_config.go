package main

import (
	"errors"
	"github.com/deis/deis/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	ResponseParserJSON = "json"
	ResponseParserTEXT = "text"
)

type RequestConfig struct {
	URL         string            `json:"yaml"`
	URI         string            `yaml:"uri"`
	Method      string            `yaml:"method"`
	ContentType string            `yaml:"content_type"`
	Headers     map[string]string `json:"headers"`
	Payload     string            `yaml:"payload"`
}

type Config struct {
	Mainly         *RequestConfig `yaml:"mainly"`
	Secondly       *RequestConfig `yaml:"secondly"`
	ResponseParser string         `yaml:"response_parser"`
}

func NewConfigFromFile(path string) (*Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("can't read config file:" + err.Error())
	}
	return NewConfigFromContent(content)
}

func NewConfigFromContent(content []byte) (*Config, error) {
	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, errors.New("Can't parse config file:" + err.Error())
	}
	return &config, nil
}
