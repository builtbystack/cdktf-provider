package developerconnectaccountconnector


type DeveloperConnectAccountConnectorProviderOauthConfig struct {
	// Required.
	//
	// User selected scopes to apply to the Oauth config
	// In the event of changing scopes, user records under AccountConnector will
	// be deleted and users will re-auth again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/developer_connect_account_connector#scopes DeveloperConnectAccountConnector#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// List of providers that are owned by Developer Connect.
	//
	// Creation of
	// new non-SCM providers Account Connectors is not possible at this
	// time.
	//
	// Possible values:
	// GITHUB
	// GITLAB
	// GOOGLE
	// SENTRY
	// ROVO
	// NEW_RELIC
	// DATASTAX
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/developer_connect_account_connector#system_provider_id DeveloperConnectAccountConnector#system_provider_id}
	SystemProviderId *string `field:"optional" json:"systemProviderId" yaml:"systemProviderId"`
}

