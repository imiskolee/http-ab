package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestConfigParser(t *testing.T) {

	content := `common: &common
  uri: /a/b?c=d&e=f#g=1
  method: GET
  content_type: application/json
  headers:
    FOO: BAR
    
mainly:
  <<: *common
  url: https://for.bar

secondly:
  <<: *common
  url: https://foo.bar
  
response_parser: json  
  
  `

	config, err := NewConfigFromContent([]byte(content))

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, config.Mainly.URL, "https://for.bar")
	assert.Equal(t, config.Secondly.URL, "https://foo.bar")
	assert.Equal(t, config.Mainly.Method, "GET")
	assert.Equal(t, config.Mainly.ContentType, "application/json")
	assert.Equal(t, config.Mainly.Headers, map[string]string{"FOO": "BAR"})

}
