package lustreinstance


type LustreInstanceAccessRulesOptions struct {
	// Set to "ROOT_SQUASH" to enable root squashing by default. Other values include "NO_SQUASH". Possible values: ["ROOT_SQUASH", "NO_SQUASH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/lustre_instance#default_squash_mode LustreInstance#default_squash_mode}
	DefaultSquashMode *string `field:"required" json:"defaultSquashMode" yaml:"defaultSquashMode"`
	// access_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/lustre_instance#access_rules LustreInstance#access_rules}
	AccessRules interface{} `field:"optional" json:"accessRules" yaml:"accessRules"`
	// The GID to map the root user to when root squashing is enabled (e.g., 65534 for nobody).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/lustre_instance#default_squash_gid LustreInstance#default_squash_gid}
	DefaultSquashGid *float64 `field:"optional" json:"defaultSquashGid" yaml:"defaultSquashGid"`
	// The UID to map the root user to when root squashing is enabled (e.g., 65534 for nobody).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/lustre_instance#default_squash_uid LustreInstance#default_squash_uid}
	DefaultSquashUid *float64 `field:"optional" json:"defaultSquashUid" yaml:"defaultSquashUid"`
}

