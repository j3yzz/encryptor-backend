package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var environments = map[string]string{
	"prod":    "./settings/prod.json",
	"develop": "./settings/develop.json",
}

type Settings struct {
	PrivateKeyPath string
	PublicKeyPath  string
	JWTExpiration  int
}

var settings Settings = Settings{}
var env = "develop"

func Init() {
	env = os.Getenv("GO_ENV")

	if env == "" {
		env = "develop"
	}

	LoadEnvSettings(env)
}

func LoadEnvSettings(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("error while loading environment file.")
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing file", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}
