package service

import (
	"errors"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/yolo-sh/hetzner-cloud-provider/infrastructure"
	"github.com/yolo-sh/yolo/stepper"
)

type ErrInvalidInstanceType struct {
	InstanceType string
	Region       string
}

func (ErrInvalidInstanceType) Error() string {
	return "ErrInvalidInstanceType"
}

func (h *Hetzner) CheckInstanceTypeValidity(
	stepper stepper.Stepper,
	instanceType string,
) error {

	hetznerClient := hcloud.NewClient(hcloud.WithToken(h.config.Credentials.APIToken))

	_, err := infrastructure.LookupServerTypeInfos(
		hetznerClient,
		instanceType,
		h.config.Region,
	)

	if err != nil {

		if errors.Is(err, infrastructure.ErrInvalidServerType) {
			return ErrInvalidInstanceType{
				InstanceType: instanceType,
				Region:       h.config.Region,
			}
		}

		return err
	}

	return nil
}
