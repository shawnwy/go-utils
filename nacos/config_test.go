package nacos

import (
	"fmt"
	"testing"
	"time"
)

var (
	testuri = "localhost:8848"
	usr     = "nacos"
	pwd     = "nacos"
)

func testInitConfigCLI(t *testing.T) {
	err := InitConfigCLI(testuri, usr, pwd)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestInitConfigCLI(t *testing.T) {
	testInitConfigCLI(t)
}

func TestReadConfig(t *testing.T) {
	testInitConfigCLI(t)
	data, err := ReadConfig("test", "DEFAULT_GROUP")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println(data)
}

func TestPublishConfig(t *testing.T) {
	testInitConfigCLI(t)
	err := PublishConfig("test", "DEFAULT_GROUP", "test1")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	data, err := ReadConfig("test", "DEFAULT_GROUP")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if data != "test1" {
		t.FailNow()
	}
}

func TestWatchConfig(t *testing.T) {
	testInitConfigCLI(t)
	err := PublishConfig("test", "DEFAULT_GROUP", "test")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	var cfg string
	if err := WatchConfig("test", "DEFAULT_GROUP", func(data string) error {
		cfg = data
		fmt.Println("assigning")
		return nil
	}); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if cfg != "test" {
		t.FailNow()
	}
	time.Sleep(5 * time.Second)
	if cfg != "test1" {
		t.FailNow()
	}
}
