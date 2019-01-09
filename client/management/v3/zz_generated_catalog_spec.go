package client

const (
	CatalogSpecType             = "catalogSpec"
	CatalogSpecFieldBranch      = "branch"
	CatalogSpecFieldCatalogKind = "catalogKind"
	CatalogSpecFieldDescription = "description"
	CatalogSpecFieldKeyRing     = "keyring"
	CatalogSpecFieldPassword    = "password"
	CatalogSpecFieldURL         = "url"
	CatalogSpecFieldUsername    = "username"
	CatalogSpecFieldVerify      = "verify"
)

type CatalogSpec struct {
	Branch      string `json:"branch,omitempty" yaml:"branch,omitempty"`
	CatalogKind string `json:"catalogKind,omitempty" yaml:"catalogKind,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	KeyRing     string `json:"keyring,omitempty" yaml:"keyring,omitempty"`
	Password    string `json:"password,omitempty" yaml:"password,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
	Username    string `json:"username,omitempty" yaml:"username,omitempty"`
	Verify      bool   `json:"verify,omitempty" yaml:"verify,omitempty"`
}
