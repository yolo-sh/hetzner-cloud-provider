package service

import (
	"github.com/yolo-sh/hetzner-cloud-provider/userconfig"
	"github.com/yolo-sh/yolo/entities"
)

//go:generate mockgen -destination ../mocks/user_config_resolver.go -package mocks -mock_names UserConfigResolver=UserConfigResolver github.com/yolo-sh/hetzner-cloud-provider/service UserConfigResolver
type UserConfigResolver interface {
	Resolve() (*userconfig.Config, error)
}

type UserConfigValidator interface {
	Validate(userConfig *userconfig.Config) error
}

type BuilderOpts struct {
	YoloConfigDir string
}

type Builder struct {
	opts                BuilderOpts
	userConfigResolver  UserConfigResolver
	userConfigValidator UserConfigValidator
}

func NewBuilder(
	opts BuilderOpts,
	userConfigResolver UserConfigResolver,
	userConfigValidator UserConfigValidator,
) Builder {

	return Builder{
		opts:                opts,
		userConfigResolver:  userConfigResolver,
		userConfigValidator: userConfigValidator,
	}
}

func (b Builder) Build() (entities.CloudService, error) {
	userConfig, err := b.userConfigResolver.Resolve()

	if err != nil {
		return nil, err
	}

	if err := b.userConfigValidator.Validate(userConfig); err != nil {
		return nil, err
	}

	userConfig.YoloConfigDir = b.opts.YoloConfigDir

	return NewHetzner(userConfig), nil
}
