package config_test

import (
	"gotemplate/pkg/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead_BadConfig_ReturnErr(t *testing.T) {
	assert.NotNil(t, config.Read("badconfig", "test", nil))
}

func TestRead_GoodConfig_FillStruct(t *testing.T) {
	type Test struct {
		Appname string
		Host    string
		Port    int
	}
	var actual Test
	err := config.Read("../../testdata/test.yml", "", &actual)

	assert.Nil(t, err)
	assert.Equal(t, Test{
		Appname: "test-app",
		Host:    "localhost",
		Port:    8000,
	}, actual)
}

func TestRead_InvalidConfigDataToFill_UnmarshalReturnErr(t *testing.T) {
	var invalidConf int

	assert.NotNil(t, config.Read("../../testdata/test.yml", "", &invalidConf))
}
