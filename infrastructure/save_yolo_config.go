package infrastructure

import (
	"os"
)

func SaveYoloConfig(
	yoloConfigDir string,
	apiToken string,
	region string,
	configJSON []byte,
) error {

	configFilePath, err := getYoloConfigFilePath(
		yoloConfigDir,
		apiToken,
		region,
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		configFilePath,
		configJSON,
		os.FileMode(0600),
	)
}
