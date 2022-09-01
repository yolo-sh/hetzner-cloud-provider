package infrastructure

import (
	"crypto/sha1"
	"encoding/hex"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

const YoloConfigFileName = "config.json"

func CreateYoloConfigStorage(
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

	return os.MkdirAll(
		filepath.Dir(configFilePath),
		fs.FileMode(0700),
	)
}

func computeYoloConfigStorageDirName(
	apiToken string,
	region string,
) (string, error) {

	sha1Computer := sha1.New()
	_, err := sha1Computer.Write([]byte("hetzner." + apiToken + "." + region))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(sha1Computer.Sum(nil)), nil
}

func getYoloConfigFilePath(
	yoloConfigDir string,
	apiToken string,
	region string,
) (string, error) {

	storageDirName, err := computeYoloConfigStorageDirName(apiToken, region)

	if err != nil {
		return "", err
	}

	return path.Join(
		yoloConfigDir,
		storageDirName,
		YoloConfigFileName,
	), nil
}
