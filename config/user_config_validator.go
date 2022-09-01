package config

import (
	"github.com/yolo-sh/hetzner-cloud-provider/userconfig"
)

var validRegions = []string{
	"fsn1",
	"nbg1",
	"hel1",
	"ash",
}

type UserConfigValidator struct{}

func NewUserConfigValidator() UserConfigValidator {
	return UserConfigValidator{}
}

func (u UserConfigValidator) Validate(userConfig *userconfig.Config) error {
	region := userConfig.Region

	if err := u.validateRegion(region); err != nil {
		return err
	}

	creds := userConfig.Credentials
	apiToken := creds.APIToken

	if err := u.validateAPIToken(apiToken); err != nil {
		return err
	}

	return nil
}

func (UserConfigValidator) validateRegion(region string) error {
	for _, validRegion := range validRegions {
		if validRegion == region {
			return nil
		}
	}

	return ErrInvalidRegion{
		Region: region,
	}
}

func (UserConfigValidator) validateAPIToken(apiToken string) error {
	if len(apiToken) != 64 {
		return ErrInvalidAPIToken{
			APIToken: apiToken,
		}
	}

	return nil
}
