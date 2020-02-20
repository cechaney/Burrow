package core

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	InitConfig()
	ConfigureLogger(Config)
	ConfigureStatic(Static)

	os.Exit(m.Run())
}

func TestInitConfig(t *testing.T) {

	if nil == Config {
		t.Errorf("core.Config was not initilized")
	}

}

func TestPortConfigured(t *testing.T) {

	if nil != Config && Config.GetString("port") == "" {
		t.Errorf("core.Config.port is not set")
	}
}
