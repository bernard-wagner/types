package client

const (
	RancherKubernetesEngineConfigType                     = "rancherKubernetesEngineConfig"
	RancherKubernetesEngineConfigFieldAddons              = "addons"
	RancherKubernetesEngineConfigFieldAuthentication      = "authentication"
	RancherKubernetesEngineConfigFieldAuthorization       = "authorization"
	RancherKubernetesEngineConfigFieldNetwork             = "network"
	RancherKubernetesEngineConfigFieldNodes               = "nodes"
	RancherKubernetesEngineConfigFieldSSHKeyPath          = "sshKeyPath"
	RancherKubernetesEngineConfigFieldServices            = "services"
	RancherKubernetesEngineConfigFieldStrictDockerVersion = "strictDockerVersion"
	RancherKubernetesEngineConfigFieldSystemImages        = "systemImages"
)

type RancherKubernetesEngineConfig struct {
	Addons              string             `json:"addons,omitempty"`
	Authentication      *AuthnConfig       `json:"authentication,omitempty"`
	Authorization       *AuthzConfig       `json:"authorization,omitempty"`
	Network             *NetworkConfig     `json:"network,omitempty"`
	Nodes               []RKEConfigNode    `json:"nodes,omitempty"`
	SSHKeyPath          string             `json:"sshKeyPath,omitempty"`
	Services            *RKEConfigServices `json:"services,omitempty"`
	StrictDockerVersion *bool              `json:"strictDockerVersion,omitempty"`
	SystemImages        map[string]string  `json:"systemImages,omitempty"`
}
