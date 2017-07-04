package io

import (
	"testing"
	"path/filepath"
)

func TestConfig(t *testing.T) {

	testFileFormat(t, "testConfig.json")
	testFileFormat(t, "legacyConfig.json")
}

func testFileFormat(t *testing.T, testFile string) {

	config := new(Config)
	expectedTarget := "http://some.site:8081"

	path, err := filepath.Abs(testFile)
	if err != nil {
		t.Error(err)
	}
	config.FilePath = path
	config.read()
	if config.Map["target"] != expectedTarget {
		t.Errorf("target != %s: %s", expectedTarget, config.Map["target"])
	}
	_, username, password, err := config.GetNetworkCredentials()
	assertUserPassword(err, t, username, "user1", password, "password1")

	username, password, err = config.GetNetworkCredentialsForTarget("http://another.one:8081")
	assertUserPassword(err, t, username, "user2", password, "password2")
}

func assertUserPassword(err error, t *testing.T, username string, expectedUser string, password string, expectedPassword string) {
	if err != nil {
		t.Error(err)
	}
	if username != expectedUser {
		t.Errorf("username != %s: %s", expectedUser, username)
	}
	if password != expectedPassword {
		t.Errorf("password != %s: %s", expectedPassword, username)
	}
}