package greyhound

import "encoding/json"
import "testing"

var testConfig = []byte(`
{"Projects": {
    "code": {
        "Root": "/home/myroot/code/",
        "Exclusions": [".*\\.class"]
    },
    "statics": {
        "Root": "/home/myroot/code/statics/"}
    }
}
`)

func TestConfigLoadFromString(t *testing.T) {
	var config GreyhoundConfig
	err:= json.Unmarshal(testConfig, &config)
	if err != nil {
		t.Error("Error while deserializing json!", err)
	}
	// test the demarshalling
	codeProject, ok := config.Projects["code"]
	if !ok {
		t.Error("Code project does not exist!")
	}
	if codeProject.Root != "/home/myroot/code/" {
		t.Errorf("root for project code is not as expected! got %s", codeProject.Root)
	}
	if codeProject.Exclusions[0] != ".*\\.class" {
		t.Error("exclusions for project code is not as expected! got: ", codeProject.Exclusions)
	}
}
