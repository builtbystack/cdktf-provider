package provider


type GooglebetaProviderExternalCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs#audience GooglebetaProvider#audience}.
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs#identity_token GooglebetaProvider#identity_token}.
	IdentityToken *string `field:"required" json:"identityToken" yaml:"identityToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs#service_account_email GooglebetaProvider#service_account_email}.
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

