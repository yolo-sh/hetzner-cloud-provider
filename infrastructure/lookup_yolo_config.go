package infrastructure

import (
	"errors"
	"io/fs"
	"os"
)

var (
	ErrNoYoloConfigFound = errors.New("ErrNoYoloConfigFound")
)

func LookupYoloConfig(
	yoloConfigDir string,
	apiToken string,
	region string,
) (string, error) {

	configFilePath, err := getYoloConfigFilePath(
		yoloConfigDir,
		apiToken,
		region,
	)

	if err != nil {
		return "", err
	}

	configFileContent, err := os.ReadFile(configFilePath)

	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return "", ErrNoYoloConfigFound
	}

	if err != nil {
		return "", err
	}

	return string(configFileContent), nil
}
