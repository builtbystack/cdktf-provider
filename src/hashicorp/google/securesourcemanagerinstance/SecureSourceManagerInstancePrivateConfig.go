package securesourcemanagerinstance


type SecureSourceManagerInstancePrivateConfig struct {
	// 'Indicate if it's private instance.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/secure_source_manager_instance#is_private SecureSourceManagerInstance#is_private}
	IsPrivate interface{} `field:"required" json:"isPrivate" yaml:"isPrivate"`
	// CA pool resource, resource must in the format of 'projects/{project}/locations/{location}/caPools/{ca_pool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/secure_source_manager_instance#ca_pool SecureSourceManagerInstance#ca_pool}
	CaPool *string `field:"optional" json:"caPool" yaml:"caPool"`
}

