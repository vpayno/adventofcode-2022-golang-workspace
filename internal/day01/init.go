// Package day01 is the module with the cli logic for the cmd application.
package day01

// Config holds the application's configuration.
type Config struct {
	appName       string
	inputFileName string
}

// Setup creates the applications configuration object.
func Setup(appName string) Config {

	conf := Config{
		appName:       appName,
		inputFileName: "./data/" + appName + "/" + appName + "-input.txt",
	}

	return conf
}
