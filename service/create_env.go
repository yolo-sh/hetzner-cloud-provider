package service

import (
	"encoding/json"
	"net"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/yolo-sh/agent/constants"
	"github.com/yolo-sh/hetzner-cloud-provider/infrastructure"
	"github.com/yolo-sh/yolo/entities"
	"github.com/yolo-sh/yolo/queues"
	"github.com/yolo-sh/yolo/stepper"
)

type EnvInfrastructure struct {
	Firewall *infrastructure.Firewall `json:"firewall"`
	SSHKey   *infrastructure.SSHKey   `json:"ssh_key"`
	Server   *infrastructure.Server   `json:"server"`
}

func (h *Hetzner) CreateEnv(
	stepper stepper.Stepper,
	config *entities.Config,
	cluster *entities.Cluster,
	env *entities.Env,
) error {

	var clusterInfra *ClusterInfrastructure
	err := json.Unmarshal([]byte(cluster.InfrastructureJSON), &clusterInfra)

	if err != nil {
		return err
	}

	envInfra := &EnvInfrastructure{}
	if len(env.InfrastructureJSON) > 0 {
		err := json.Unmarshal([]byte(env.InfrastructureJSON), envInfra)

		if err != nil {
			return err
		}
	}

	prefixResource := prefixEnvResource(cluster.GetNameSlug(), env.GetNameSlug())
	hetznerClient := hcloud.NewClient(hcloud.WithToken(h.config.Credentials.APIToken))

	envInfraQueue := queues.InfrastructureQueue[*EnvInfrastructure]{}

	createFirewall := func(infra *EnvInfrastructure) error {
		if infra.Firewall != nil {
			return nil
		}

		yoloSSHServerListenPort := constants.SSHServerListenPort
		_, sourceIP, err := net.ParseCIDR("0.0.0.0/0")

		if err != nil {
			return err
		}

		firewall, err := infrastructure.CreateFirewall(
			hetznerClient,
			prefixResource("firewall"),
			[]hcloud.FirewallRule{
				{
					Direction: hcloud.FirewallRuleDirectionIn,
					Protocol:  hcloud.FirewallRuleProtocolTCP,
					Port:      &yoloSSHServerListenPort,
					SourceIPs: []net.IPNet{
						{
							IP:   sourceIP.IP,
							Mask: sourceIP.Mask,
						},
					},
				},
			},
		)

		if err != nil {
			return err
		}

		infra.Firewall = firewall
		return nil
	}

	createSSHKey := func(infra *EnvInfrastructure) error {
		if infra.SSHKey != nil {
			return nil
		}

		sshKey, err := infrastructure.CreateSSHKey(
			hetznerClient,
			prefixResource("ssh-key"),
		)

		if err != nil {
			return err
		}

		infra.SSHKey = sshKey
		return nil
	}

	envInfraQueue = append(
		envInfraQueue,
		queues.InfrastructureQueueSteps[*EnvInfrastructure]{
			func(*EnvInfrastructure) error {
				stepper.StartTemporaryStep("Creating a firewall and an SSH key")
				return nil
			},
			createFirewall,
			createSSHKey,
		},
	)

	createServer := func(infra *EnvInfrastructure) error {
		if infra.Server != nil {
			return nil
		}

		server, err := infrastructure.CreateServer(
			hetznerClient,
			clusterInfra.Location.ID,
			clusterInfra.Network.ID,
			env.InstanceType,
			infra.Firewall.ID,
			prefixResource("server"),
			infra.SSHKey.ID,
		)

		if err != nil {
			return err
		}

		infra.Server = server
		infra.Firewall.AttachedToServer = true
		return nil
	}

	envInfraQueue = append(
		envInfraQueue,
		queues.InfrastructureQueueSteps[*EnvInfrastructure]{
			func(*EnvInfrastructure) error {
				stepper.StartTemporaryStep("Creating a server")
				return nil
			},
			createServer,
		},
	)

	lookupServerInitScriptResults := func(infra *EnvInfrastructure) error {
		if infra.Server.InitScriptResults != nil {
			return nil
		}

		initScriptResults, err := infrastructure.LookupServerInitScriptResults(
			envInfra.Server.PublicIPAddress,
			constants.SSHServerListenPort,
			entities.EnvRootUser,
			envInfra.SSHKey.PEMContent,
		)

		if err != nil {
			return err
		}

		infra.Server.InitScriptResults = initScriptResults
		return nil
	}

	envInfraQueue = append(
		envInfraQueue,
		queues.InfrastructureQueueSteps[*EnvInfrastructure]{
			func(*EnvInfrastructure) error {
				stepper.StartTemporaryStep("Configuring the server")
				return nil
			},
			lookupServerInitScriptResults,
		},
	)

	err = envInfraQueue.Run(envInfra)

	// Env infra could be updated in the queue even
	// in case of error (partial infrastructure)
	env.SetInfrastructureJSON(envInfra)

	if err != nil {
		return err
	}

	env.InstancePublicIPAddress = envInfra.Server.PublicIPAddress

	env.SSHHostKeys = envInfra.Server.InitScriptResults.SSHHostKeys
	env.SSHKeyPairPEMContent = envInfra.SSHKey.PEMContent

	return nil
}
