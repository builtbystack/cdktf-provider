package googlelustreinstance


type GoogleLustreInstanceAccessRulesOptions struct {
	// Set to "ROOT_SQUASH" to enable root squashing by default. Other values include "NO_SQUASH". Possible values: ["ROOT_SQUASH", "NO_SQUASH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_lustre_instance#default_squash_mode GoogleLustreInstance#default_squash_mode}
	DefaultSquashMode *string `field:"required" json:"defaultSquashMode" yaml:"defaultSquashMode"`
	// access_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_lustre_instance#access_rules GoogleLustreInstance#access_rules}
	AccessRules interface{} `field:"optional" json:"accessRules" yaml:"accessRules"`
	// The GID to map the root user to when root squashing is enabled (e.g., 65534 for nobody).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_lustre_instance#default_squash_gid GoogleLustreInstance#default_squash_gid}
	DefaultSquashGid *float64 `field:"optional" json:"defaultSquashGid" yaml:"defaultSquashGid"`
	// The UID to map the root user to when root squashing is enabled (e.g., 65534 for nobody).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_lustre_instance#default_squash_uid GoogleLustreInstance#default_squash_uid}
	DefaultSquashUid *float64 `field:"optional" json:"defaultSquashUid" yaml:"defaultSquashUid"`
}

