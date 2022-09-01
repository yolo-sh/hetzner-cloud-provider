package service

import (
	"encoding/json"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/yolo-sh/hetzner-cloud-provider/infrastructure"
	"github.com/yolo-sh/yolo/entities"
	"github.com/yolo-sh/yolo/stepper"
)

func (h *Hetzner) ClosePort(
	stepper stepper.Stepper,
	config *entities.Config,
	cluster *entities.Cluster,
	env *entities.Env,
	portToClose string,
) error {

	var envInfra *EnvInfrastructure
	err := json.Unmarshal([]byte(env.InfrastructureJSON), &envInfra)

	if err != nil {
		return err
	}

	hetznerClient := hcloud.NewClient(hcloud.WithToken(h.config.Credentials.APIToken))

	err = infrastructure.CloseServerPort(
		hetznerClient,
		envInfra.Firewall.ID,
		portToClose,
	)

	if err != nil {
		return err
	}

	delete(env.OpenedPorts, portToClose)

	return nil
}
