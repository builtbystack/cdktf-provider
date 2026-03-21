package computeregionbackendbucket


type ComputeRegionBackendBucketTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/compute_region_backend_bucket#create ComputeRegionBackendBucket#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/compute_region_backend_bucket#delete ComputeRegionBackendBucket#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/compute_region_backend_bucket#update ComputeRegionBackendBucket#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

