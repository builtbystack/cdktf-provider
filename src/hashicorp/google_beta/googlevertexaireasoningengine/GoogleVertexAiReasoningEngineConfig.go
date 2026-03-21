package googlevertexaireasoningengine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVertexAiReasoningEngineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The display name of the ReasoningEngine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_vertex_ai_reasoning_engine#display_name GoogleVertexAiReasoningEngine#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The description of the ReasoningEngine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_vertex_ai_reasoning_engine#description GoogleVertexAiReasoningEngine#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_vertex_ai_reasoning_engine#encryption_spec GoogleVertexAiReasoningEngine#encryption_spec}
	EncryptionSpec *GoogleVertexAiReasoningEngineEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_vertex_ai_reasoning_engine#id GoogleVertexAiReasoningEngine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_vertex_ai_reasoning_engine#project GoogleVertexAiReasoningEngine#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the reasoning engine. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_vertex_ai_reasoning_engine#region GoogleVertexAiReasoningEngine#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_vertex_ai_reasoning_engine#spec GoogleVertexAiReasoningEngine#spec}
	Spec *GoogleVertexAiReasoningEngineSpec `field:"optional" json:"spec" yaml:"spec"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_vertex_ai_reasoning_engine#timeouts GoogleVertexAiReasoningEngine#timeouts}
	Timeouts *GoogleVertexAiReasoningEngineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

