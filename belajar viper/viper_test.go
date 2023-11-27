package belajar_viper

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	var config *viper.Viper = viper.New()

	assert.NotNil(t, config)
}

func TestJson(t *testing.T) {
	var config *viper.Viper = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Fuad", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestYaml(t *testing.T) {
	var config *viper.Viper = viper.New()
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Fuad", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestEnv(t *testing.T) {
	var config *viper.Viper = viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "belajar-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Fuad", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
}
