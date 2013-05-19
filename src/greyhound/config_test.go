package greyhound

var testConfig = `
{"Projects": {
    "code": {
        "Root": "/home/p4/code/comp_hi/",
        "Exclusions": [".*\.class"]
    },
    "statics": {
        "Root": "/home/p4/code/comp_hi/ZillowWeb/webapp.war/static/"}
    }
}
`

func TestConfigLoadFromString(t *testing.T) {
	gs := NewGreyhoundSearch()
	gs.LoadConfigFromString()
}
