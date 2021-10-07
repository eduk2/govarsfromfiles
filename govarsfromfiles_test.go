package govarsfromfiles

import (
	"testing"
)

func TestSetPath(t *testing.T) {

	if File.SetPath("./configuration_file_example.php") != true {
		t.Fatal("configuration_file_example.php exists and it cannot be initialized")
	}

}

func TestGetValues(t *testing.T) {

	testVarsSearched := []struct {
		searching string
		expected  string
	}{
		{"$db_host", "localhost"},
		{"$db_password", "this is the password"},
	}

	for _, values := range testVarsSearched {

		File.SetPath("./configuration_file_example.php")
		myValues := File.GetValues(values.searching)
		if myValues[values.searching] != values.expected {
			t.Errorf("%+v did not equal to %+v", values.searching, values.expected)
		}

	}

}
