package service

import (
	"encoding/json"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/yolo-sh/hetzner-cloud-provider/infrastructure"
	"github.com/yolo-sh/yolo/entities"
	"github.com/yolo-sh/yolo/stepper"
)

func (h *Hetzner) OpenPort(
	stepper stepper.Stepper,
	config *entities.Config,
	cluster *entities.Cluster,
	env *entities.Env,
	portToOpen string,
) error {

	var envInfra *EnvInfrastructure
	err := json.Unmarshal([]byte(env.InfrastructureJSON), &envInfra)

	if err != nil {
		return err
	}

	hetznerClient := hcloud.NewClient(hcloud.WithToken(h.config.Credentials.APIToken))

	err = infrastructure.OpenServerPort(
		hetznerClient,
		envInfra.Firewall.ID,
		portToOpen,
	)

	if err != nil {
		return err
	}

	env.OpenedPorts[portToOpen] = true

	return nil
}
