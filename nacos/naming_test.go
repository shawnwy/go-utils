package nacos

import "testing"

func testInitNamingCLI(t *testing.T) {
	err := InitNamingCLI(t)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestInitNamingCLI(t *testing.T) {
	testInitNamingCLI(t)
}

func TestRegisterService(t *testing.T) {
	testInitNamingCLI(t)
	err := RegisterService("testSVC", "TEST_GROUP", "1.1.1.1", 33033)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestDeregisterService(t *testing.T) {
	testInitNamingCLI(t)
	err := DeregisterService("testSVC", "TEST_GROUP", "1.1.1.1", 33033)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
