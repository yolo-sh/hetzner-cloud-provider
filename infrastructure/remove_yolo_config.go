package infrastructure

import (
	"os"
	"path/filepath"
)

func RemoveYoloConfig(
	yoloConfigDir string,
	apiToken string,
	region string,
) error {

	configFilePath, err := getYoloConfigFilePath(
		yoloConfigDir,
		apiToken,
		region,
	)

	if err != nil {
		return err
	}

	err = os.Remove(configFilePath)

	if err != nil {
		return err
	}

	configDirName := filepath.Dir(
		configFilePath,
	)

	return os.Remove(configDirName)
}
