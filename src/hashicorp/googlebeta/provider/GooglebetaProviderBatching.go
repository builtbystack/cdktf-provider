package provider


type GooglebetaProviderBatching struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs#enable_batching GooglebetaProvider#enable_batching}.
	EnableBatching interface{} `field:"optional" json:"enableBatching" yaml:"enableBatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs#send_after GooglebetaProvider#send_after}.
	SendAfter *string `field:"optional" json:"sendAfter" yaml:"sendAfter"`
}

