package lustreinstance


type LustreInstanceAccessRulesOptionsAccessRules struct {
	// An array of IP address strings or CIDR ranges that this rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/lustre_instance#ip_address_ranges LustreInstance#ip_address_ranges}
	IpAddressRanges *[]*string `field:"required" json:"ipAddressRanges" yaml:"ipAddressRanges"`
	// A unique identifier for the access rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/lustre_instance#name LustreInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The squash mode for this specific rule. Currently, only "NO_SQUASH" is supported for exceptions. Possible values: ["NO_SQUASH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/lustre_instance#squash_mode LustreInstance#squash_mode}
	SquashMode *string `field:"required" json:"squashMode" yaml:"squashMode"`
}

