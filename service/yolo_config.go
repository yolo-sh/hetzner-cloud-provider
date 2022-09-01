package service

import (
	"encoding/json"
	"errors"

	"github.com/yolo-sh/hetzner-cloud-provider/infrastructure"
	"github.com/yolo-sh/yolo/entities"
	"github.com/yolo-sh/yolo/stepper"
)

func (h *Hetzner) CreateYoloConfigStorage(
	stepper stepper.Stepper,
) error {

	stepper.StartTemporaryStep("Creating Yolo's config storage")

	return infrastructure.CreateYoloConfigStorage(
		h.config.YoloConfigDir,
		h.config.Credentials.APIToken,
		h.config.Region,
	)
}

func (h *Hetzner) LookupYoloConfig(
	stepper stepper.Stepper,
) (*entities.Config, error) {

	configJSON, err := infrastructure.LookupYoloConfig(
		h.config.YoloConfigDir,
		h.config.Credentials.APIToken,
		h.config.Region,
	)

	if err != nil {

		if errors.Is(err, infrastructure.ErrNoYoloConfigFound) {
			return nil, entities.ErrYoloNotInstalled
		}

		return nil, err
	}

	var yoloConfig *entities.Config
	err = json.Unmarshal([]byte(configJSON), &yoloConfig)

	if err != nil {
		return nil, err
	}

	return yoloConfig, nil
}

func (h *Hetzner) SaveYoloConfig(
	stepper stepper.Stepper,
	config *entities.Config,
) error {

	configJSON, err := json.Marshal(config)

	if err != nil {
		return err
	}

	return infrastructure.SaveYoloConfig(
		h.config.YoloConfigDir,
		h.config.Credentials.APIToken,
		h.config.Region,
		configJSON,
	)
}

func (h *Hetzner) RemoveYoloConfigStorage(
	stepper stepper.Stepper,
) error {

	stepper.StartTemporaryStep("Removing Yolo's config storage")

	return infrastructure.RemoveYoloConfig(
		h.config.YoloConfigDir,
		h.config.Credentials.APIToken,
		h.config.Region,
	)
}
