package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/builtbystack/cdktf-provider/src/hashicorp/googlebeta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/builtbystack/cdktf-provider/src/hashicorp/googlebeta/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs googlebeta}.
type GooglebetaProvider interface {
	cdktf.TerraformProvider
	AccessApprovalCustomEndpoint() *string
	SetAccessApprovalCustomEndpoint(val *string)
	AccessApprovalCustomEndpointInput() *string
	AccessContextManagerCustomEndpoint() *string
	SetAccessContextManagerCustomEndpoint(val *string)
	AccessContextManagerCustomEndpointInput() *string
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	ActiveDirectoryCustomEndpoint() *string
	SetActiveDirectoryCustomEndpoint(val *string)
	ActiveDirectoryCustomEndpointInput() *string
	AddTerraformAttributionLabel() interface{}
	SetAddTerraformAttributionLabel(val interface{})
	AddTerraformAttributionLabelInput() interface{}
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AlloydbCustomEndpoint() *string
	SetAlloydbCustomEndpoint(val *string)
	AlloydbCustomEndpointInput() *string
	ApiGatewayCustomEndpoint() *string
	SetApiGatewayCustomEndpoint(val *string)
	ApiGatewayCustomEndpointInput() *string
	ApigeeCustomEndpoint() *string
	SetApigeeCustomEndpoint(val *string)
	ApigeeCustomEndpointInput() *string
	ApihubCustomEndpoint() *string
	SetApihubCustomEndpoint(val *string)
	ApihubCustomEndpointInput() *string
	ApikeysCustomEndpoint() *string
	SetApikeysCustomEndpoint(val *string)
	ApikeysCustomEndpointInput() *string
	AppEngineCustomEndpoint() *string
	SetAppEngineCustomEndpoint(val *string)
	AppEngineCustomEndpointInput() *string
	ApphubCustomEndpoint() *string
	SetApphubCustomEndpoint(val *string)
	ApphubCustomEndpointInput() *string
	ArtifactRegistryCustomEndpoint() *string
	SetArtifactRegistryCustomEndpoint(val *string)
	ArtifactRegistryCustomEndpointInput() *string
	AssuredWorkloadsCustomEndpoint() *string
	SetAssuredWorkloadsCustomEndpoint(val *string)
	AssuredWorkloadsCustomEndpointInput() *string
	BackupDrCustomEndpoint() *string
	SetBackupDrCustomEndpoint(val *string)
	BackupDrCustomEndpointInput() *string
	Batching() interface{}
	SetBatching(val interface{})
	BatchingInput() interface{}
	BeyondcorpCustomEndpoint() *string
	SetBeyondcorpCustomEndpoint(val *string)
	BeyondcorpCustomEndpointInput() *string
	BiglakeCustomEndpoint() *string
	SetBiglakeCustomEndpoint(val *string)
	BiglakeCustomEndpointInput() *string
	BiglakeIcebergCustomEndpoint() *string
	SetBiglakeIcebergCustomEndpoint(val *string)
	BiglakeIcebergCustomEndpointInput() *string
	BigqueryAnalyticsHubCustomEndpoint() *string
	SetBigqueryAnalyticsHubCustomEndpoint(val *string)
	BigqueryAnalyticsHubCustomEndpointInput() *string
	BigqueryConnectionCustomEndpoint() *string
	SetBigqueryConnectionCustomEndpoint(val *string)
	BigqueryConnectionCustomEndpointInput() *string
	BigQueryCustomEndpoint() *string
	SetBigQueryCustomEndpoint(val *string)
	BigQueryCustomEndpointInput() *string
	BigqueryDatapolicyCustomEndpoint() *string
	SetBigqueryDatapolicyCustomEndpoint(val *string)
	BigqueryDatapolicyCustomEndpointInput() *string
	BigqueryDatapolicyv2CustomEndpoint() *string
	SetBigqueryDatapolicyv2CustomEndpoint(val *string)
	BigqueryDatapolicyv2CustomEndpointInput() *string
	BigqueryDataTransferCustomEndpoint() *string
	SetBigqueryDataTransferCustomEndpoint(val *string)
	BigqueryDataTransferCustomEndpointInput() *string
	BigqueryReservationCustomEndpoint() *string
	SetBigqueryReservationCustomEndpoint(val *string)
	BigqueryReservationCustomEndpointInput() *string
	BigtableCustomEndpoint() *string
	SetBigtableCustomEndpoint(val *string)
	BigtableCustomEndpointInput() *string
	BillingCustomEndpoint() *string
	SetBillingCustomEndpoint(val *string)
	BillingCustomEndpointInput() *string
	BillingProject() *string
	SetBillingProject(val *string)
	BillingProjectInput() *string
	BinaryAuthorizationCustomEndpoint() *string
	SetBinaryAuthorizationCustomEndpoint(val *string)
	BinaryAuthorizationCustomEndpointInput() *string
	BlockchainNodeEngineCustomEndpoint() *string
	SetBlockchainNodeEngineCustomEndpoint(val *string)
	BlockchainNodeEngineCustomEndpointInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateManagerCustomEndpoint() *string
	SetCertificateManagerCustomEndpoint(val *string)
	CertificateManagerCustomEndpointInput() *string
	CesCustomEndpoint() *string
	SetCesCustomEndpoint(val *string)
	CesCustomEndpointInput() *string
	ChronicleCustomEndpoint() *string
	SetChronicleCustomEndpoint(val *string)
	ChronicleCustomEndpointInput() *string
	CloudAssetCustomEndpoint() *string
	SetCloudAssetCustomEndpoint(val *string)
	CloudAssetCustomEndpointInput() *string
	CloudBillingCustomEndpoint() *string
	SetCloudBillingCustomEndpoint(val *string)
	CloudBillingCustomEndpointInput() *string
	CloudBuildCustomEndpoint() *string
	SetCloudBuildCustomEndpoint(val *string)
	CloudBuildCustomEndpointInput() *string
	Cloudbuildv2CustomEndpoint() *string
	SetCloudbuildv2CustomEndpoint(val *string)
	Cloudbuildv2CustomEndpointInput() *string
	ClouddeployCustomEndpoint() *string
	SetClouddeployCustomEndpoint(val *string)
	ClouddeployCustomEndpointInput() *string
	ClouddomainsCustomEndpoint() *string
	SetClouddomainsCustomEndpoint(val *string)
	ClouddomainsCustomEndpointInput() *string
	Cloudfunctions2CustomEndpoint() *string
	SetCloudfunctions2CustomEndpoint(val *string)
	Cloudfunctions2CustomEndpointInput() *string
	CloudFunctionsCustomEndpoint() *string
	SetCloudFunctionsCustomEndpoint(val *string)
	CloudFunctionsCustomEndpointInput() *string
	CloudIdentityCustomEndpoint() *string
	SetCloudIdentityCustomEndpoint(val *string)
	CloudIdentityCustomEndpointInput() *string
	CloudIdsCustomEndpoint() *string
	SetCloudIdsCustomEndpoint(val *string)
	CloudIdsCustomEndpointInput() *string
	CloudQuotasCustomEndpoint() *string
	SetCloudQuotasCustomEndpoint(val *string)
	CloudQuotasCustomEndpointInput() *string
	CloudResourceManagerCustomEndpoint() *string
	SetCloudResourceManagerCustomEndpoint(val *string)
	CloudResourceManagerCustomEndpointInput() *string
	CloudRunCustomEndpoint() *string
	SetCloudRunCustomEndpoint(val *string)
	CloudRunCustomEndpointInput() *string
	CloudRunV2CustomEndpoint() *string
	SetCloudRunV2CustomEndpoint(val *string)
	CloudRunV2CustomEndpointInput() *string
	CloudSchedulerCustomEndpoint() *string
	SetCloudSchedulerCustomEndpoint(val *string)
	CloudSchedulerCustomEndpointInput() *string
	CloudSecurityComplianceCustomEndpoint() *string
	SetCloudSecurityComplianceCustomEndpoint(val *string)
	CloudSecurityComplianceCustomEndpointInput() *string
	CloudTasksCustomEndpoint() *string
	SetCloudTasksCustomEndpoint(val *string)
	CloudTasksCustomEndpointInput() *string
	ColabCustomEndpoint() *string
	SetColabCustomEndpoint(val *string)
	ColabCustomEndpointInput() *string
	ComposerCustomEndpoint() *string
	SetComposerCustomEndpoint(val *string)
	ComposerCustomEndpointInput() *string
	ComputeCustomEndpoint() *string
	SetComputeCustomEndpoint(val *string)
	ComputeCustomEndpointInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContactCenterInsightsCustomEndpoint() *string
	SetContactCenterInsightsCustomEndpoint(val *string)
	ContactCenterInsightsCustomEndpointInput() *string
	ContainerAnalysisCustomEndpoint() *string
	SetContainerAnalysisCustomEndpoint(val *string)
	ContainerAnalysisCustomEndpointInput() *string
	ContainerAttachedCustomEndpoint() *string
	SetContainerAttachedCustomEndpoint(val *string)
	ContainerAttachedCustomEndpointInput() *string
	ContainerAwsCustomEndpoint() *string
	SetContainerAwsCustomEndpoint(val *string)
	ContainerAwsCustomEndpointInput() *string
	ContainerAzureCustomEndpoint() *string
	SetContainerAzureCustomEndpoint(val *string)
	ContainerAzureCustomEndpointInput() *string
	ContainerCustomEndpoint() *string
	SetContainerCustomEndpoint(val *string)
	ContainerCustomEndpointInput() *string
	CoreBillingCustomEndpoint() *string
	SetCoreBillingCustomEndpoint(val *string)
	CoreBillingCustomEndpointInput() *string
	Credentials() *string
	SetCredentials(val *string)
	CredentialsInput() *string
	DatabaseMigrationServiceCustomEndpoint() *string
	SetDatabaseMigrationServiceCustomEndpoint(val *string)
	DatabaseMigrationServiceCustomEndpointInput() *string
	DataCatalogCustomEndpoint() *string
	SetDataCatalogCustomEndpoint(val *string)
	DataCatalogCustomEndpointInput() *string
	DataflowCustomEndpoint() *string
	SetDataflowCustomEndpoint(val *string)
	DataflowCustomEndpointInput() *string
	DataformCustomEndpoint() *string
	SetDataformCustomEndpoint(val *string)
	DataformCustomEndpointInput() *string
	DataFusionCustomEndpoint() *string
	SetDataFusionCustomEndpoint(val *string)
	DataFusionCustomEndpointInput() *string
	DataLossPreventionCustomEndpoint() *string
	SetDataLossPreventionCustomEndpoint(val *string)
	DataLossPreventionCustomEndpointInput() *string
	DataPipelineCustomEndpoint() *string
	SetDataPipelineCustomEndpoint(val *string)
	DataPipelineCustomEndpointInput() *string
	DataplexCustomEndpoint() *string
	SetDataplexCustomEndpoint(val *string)
	DataplexCustomEndpointInput() *string
	DataprocCustomEndpoint() *string
	SetDataprocCustomEndpoint(val *string)
	DataprocCustomEndpointInput() *string
	DataprocGdcCustomEndpoint() *string
	SetDataprocGdcCustomEndpoint(val *string)
	DataprocGdcCustomEndpointInput() *string
	DataprocMetastoreCustomEndpoint() *string
	SetDataprocMetastoreCustomEndpoint(val *string)
	DataprocMetastoreCustomEndpointInput() *string
	DatastreamCustomEndpoint() *string
	SetDatastreamCustomEndpoint(val *string)
	DatastreamCustomEndpointInput() *string
	DefaultLabels() *map[string]*string
	SetDefaultLabels(val *map[string]*string)
	DefaultLabelsInput() *map[string]*string
	DeploymentManagerCustomEndpoint() *string
	SetDeploymentManagerCustomEndpoint(val *string)
	DeploymentManagerCustomEndpointInput() *string
	DeveloperConnectCustomEndpoint() *string
	SetDeveloperConnectCustomEndpoint(val *string)
	DeveloperConnectCustomEndpointInput() *string
	DialogflowCustomEndpoint() *string
	SetDialogflowCustomEndpoint(val *string)
	DialogflowCustomEndpointInput() *string
	DialogflowCxCustomEndpoint() *string
	SetDialogflowCxCustomEndpoint(val *string)
	DialogflowCxCustomEndpointInput() *string
	DiscoveryEngineCustomEndpoint() *string
	SetDiscoveryEngineCustomEndpoint(val *string)
	DiscoveryEngineCustomEndpointInput() *string
	DnsCustomEndpoint() *string
	SetDnsCustomEndpoint(val *string)
	DnsCustomEndpointInput() *string
	DocumentAiCustomEndpoint() *string
	SetDocumentAiCustomEndpoint(val *string)
	DocumentAiCustomEndpointInput() *string
	DocumentAiWarehouseCustomEndpoint() *string
	SetDocumentAiWarehouseCustomEndpoint(val *string)
	DocumentAiWarehouseCustomEndpointInput() *string
	EdgecontainerCustomEndpoint() *string
	SetEdgecontainerCustomEndpoint(val *string)
	EdgecontainerCustomEndpointInput() *string
	EdgenetworkCustomEndpoint() *string
	SetEdgenetworkCustomEndpoint(val *string)
	EdgenetworkCustomEndpointInput() *string
	EssentialContactsCustomEndpoint() *string
	SetEssentialContactsCustomEndpoint(val *string)
	EssentialContactsCustomEndpointInput() *string
	EventarcCustomEndpoint() *string
	SetEventarcCustomEndpoint(val *string)
	EventarcCustomEndpointInput() *string
	ExternalCredentials() interface{}
	SetExternalCredentials(val interface{})
	ExternalCredentialsInput() interface{}
	FilestoreCustomEndpoint() *string
	SetFilestoreCustomEndpoint(val *string)
	FilestoreCustomEndpointInput() *string
	FirebaseAiLogicCustomEndpoint() *string
	SetFirebaseAiLogicCustomEndpoint(val *string)
	FirebaseAiLogicCustomEndpointInput() *string
	FirebaseAppCheckCustomEndpoint() *string
	SetFirebaseAppCheckCustomEndpoint(val *string)
	FirebaseAppCheckCustomEndpointInput() *string
	FirebaseAppHostingCustomEndpoint() *string
	SetFirebaseAppHostingCustomEndpoint(val *string)
	FirebaseAppHostingCustomEndpointInput() *string
	FirebaseCustomEndpoint() *string
	SetFirebaseCustomEndpoint(val *string)
	FirebaseCustomEndpointInput() *string
	FirebaseDatabaseCustomEndpoint() *string
	SetFirebaseDatabaseCustomEndpoint(val *string)
	FirebaseDatabaseCustomEndpointInput() *string
	FirebaseDataConnectCustomEndpoint() *string
	SetFirebaseDataConnectCustomEndpoint(val *string)
	FirebaseDataConnectCustomEndpointInput() *string
	FirebaseExtensionsCustomEndpoint() *string
	SetFirebaseExtensionsCustomEndpoint(val *string)
	FirebaseExtensionsCustomEndpointInput() *string
	FirebaseHostingCustomEndpoint() *string
	SetFirebaseHostingCustomEndpoint(val *string)
	FirebaseHostingCustomEndpointInput() *string
	FirebaserulesCustomEndpoint() *string
	SetFirebaserulesCustomEndpoint(val *string)
	FirebaserulesCustomEndpointInput() *string
	FirebaseStorageCustomEndpoint() *string
	SetFirebaseStorageCustomEndpoint(val *string)
	FirebaseStorageCustomEndpointInput() *string
	FirestoreCustomEndpoint() *string
	SetFirestoreCustomEndpoint(val *string)
	FirestoreCustomEndpointInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeminiCustomEndpoint() *string
	SetGeminiCustomEndpoint(val *string)
	GeminiCustomEndpointInput() *string
	GkeBackupCustomEndpoint() *string
	SetGkeBackupCustomEndpoint(val *string)
	GkeBackupCustomEndpointInput() *string
	GkeHub2CustomEndpoint() *string
	SetGkeHub2CustomEndpoint(val *string)
	GkeHub2CustomEndpointInput() *string
	GkeHubCustomEndpoint() *string
	SetGkeHubCustomEndpoint(val *string)
	GkeHubCustomEndpointInput() *string
	GkeonpremCustomEndpoint() *string
	SetGkeonpremCustomEndpoint(val *string)
	GkeonpremCustomEndpointInput() *string
	HealthcareCustomEndpoint() *string
	SetHealthcareCustomEndpoint(val *string)
	HealthcareCustomEndpointInput() *string
	HypercomputeclusterCustomEndpoint() *string
	SetHypercomputeclusterCustomEndpoint(val *string)
	HypercomputeclusterCustomEndpointInput() *string
	Iam2CustomEndpoint() *string
	SetIam2CustomEndpoint(val *string)
	Iam2CustomEndpointInput() *string
	Iam3CustomEndpoint() *string
	SetIam3CustomEndpoint(val *string)
	Iam3CustomEndpointInput() *string
	IamBetaCustomEndpoint() *string
	SetIamBetaCustomEndpoint(val *string)
	IamBetaCustomEndpointInput() *string
	IamCredentialsCustomEndpoint() *string
	SetIamCredentialsCustomEndpoint(val *string)
	IamCredentialsCustomEndpointInput() *string
	IamCustomEndpoint() *string
	SetIamCustomEndpoint(val *string)
	IamCustomEndpointInput() *string
	IamWorkforcePoolCustomEndpoint() *string
	SetIamWorkforcePoolCustomEndpoint(val *string)
	IamWorkforcePoolCustomEndpointInput() *string
	IapCustomEndpoint() *string
	SetIapCustomEndpoint(val *string)
	IapCustomEndpointInput() *string
	IdentityPlatformCustomEndpoint() *string
	SetIdentityPlatformCustomEndpoint(val *string)
	IdentityPlatformCustomEndpointInput() *string
	ImpersonateServiceAccount() *string
	SetImpersonateServiceAccount(val *string)
	ImpersonateServiceAccountDelegates() *[]*string
	SetImpersonateServiceAccountDelegates(val *[]*string)
	ImpersonateServiceAccountDelegatesInput() *[]*string
	ImpersonateServiceAccountInput() *string
	IntegrationConnectorsCustomEndpoint() *string
	SetIntegrationConnectorsCustomEndpoint(val *string)
	IntegrationConnectorsCustomEndpointInput() *string
	IntegrationsCustomEndpoint() *string
	SetIntegrationsCustomEndpoint(val *string)
	IntegrationsCustomEndpointInput() *string
	KmsCustomEndpoint() *string
	SetKmsCustomEndpoint(val *string)
	KmsCustomEndpointInput() *string
	LoggingCustomEndpoint() *string
	SetLoggingCustomEndpoint(val *string)
	LoggingCustomEndpointInput() *string
	LookerCustomEndpoint() *string
	SetLookerCustomEndpoint(val *string)
	LookerCustomEndpointInput() *string
	LustreCustomEndpoint() *string
	SetLustreCustomEndpoint(val *string)
	LustreCustomEndpointInput() *string
	ManagedKafkaCustomEndpoint() *string
	SetManagedKafkaCustomEndpoint(val *string)
	ManagedKafkaCustomEndpointInput() *string
	MemcacheCustomEndpoint() *string
	SetMemcacheCustomEndpoint(val *string)
	MemcacheCustomEndpointInput() *string
	MemorystoreCustomEndpoint() *string
	SetMemorystoreCustomEndpoint(val *string)
	MemorystoreCustomEndpointInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	MigrationCenterCustomEndpoint() *string
	SetMigrationCenterCustomEndpoint(val *string)
	MigrationCenterCustomEndpointInput() *string
	MlEngineCustomEndpoint() *string
	SetMlEngineCustomEndpoint(val *string)
	MlEngineCustomEndpointInput() *string
	ModelArmorCustomEndpoint() *string
	SetModelArmorCustomEndpoint(val *string)
	ModelArmorCustomEndpointInput() *string
	ModelArmorGlobalCustomEndpoint() *string
	SetModelArmorGlobalCustomEndpoint(val *string)
	ModelArmorGlobalCustomEndpointInput() *string
	MonitoringCustomEndpoint() *string
	SetMonitoringCustomEndpoint(val *string)
	MonitoringCustomEndpointInput() *string
	NetappCustomEndpoint() *string
	SetNetappCustomEndpoint(val *string)
	NetappCustomEndpointInput() *string
	NetworkConnectivityCustomEndpoint() *string
	SetNetworkConnectivityCustomEndpoint(val *string)
	NetworkConnectivityCustomEndpointInput() *string
	NetworkConnectivityv1CustomEndpoint() *string
	SetNetworkConnectivityv1CustomEndpoint(val *string)
	NetworkConnectivityv1CustomEndpointInput() *string
	NetworkManagementCustomEndpoint() *string
	SetNetworkManagementCustomEndpoint(val *string)
	NetworkManagementCustomEndpointInput() *string
	NetworkSecurityCustomEndpoint() *string
	SetNetworkSecurityCustomEndpoint(val *string)
	NetworkSecurityCustomEndpointInput() *string
	NetworkServicesCustomEndpoint() *string
	SetNetworkServicesCustomEndpoint(val *string)
	NetworkServicesCustomEndpointInput() *string
	// The tree node.
	Node() constructs.Node
	NotebooksCustomEndpoint() *string
	SetNotebooksCustomEndpoint(val *string)
	NotebooksCustomEndpointInput() *string
	ObservabilityCustomEndpoint() *string
	SetObservabilityCustomEndpoint(val *string)
	ObservabilityCustomEndpointInput() *string
	OracleDatabaseCustomEndpoint() *string
	SetOracleDatabaseCustomEndpoint(val *string)
	OracleDatabaseCustomEndpointInput() *string
	OrgPolicyCustomEndpoint() *string
	SetOrgPolicyCustomEndpoint(val *string)
	OrgPolicyCustomEndpointInput() *string
	OsConfigCustomEndpoint() *string
	SetOsConfigCustomEndpoint(val *string)
	OsConfigCustomEndpointInput() *string
	OsConfigV2CustomEndpoint() *string
	SetOsConfigV2CustomEndpoint(val *string)
	OsConfigV2CustomEndpointInput() *string
	OsLoginCustomEndpoint() *string
	SetOsLoginCustomEndpoint(val *string)
	OsLoginCustomEndpointInput() *string
	ParallelstoreCustomEndpoint() *string
	SetParallelstoreCustomEndpoint(val *string)
	ParallelstoreCustomEndpointInput() *string
	ParameterManagerCustomEndpoint() *string
	SetParameterManagerCustomEndpoint(val *string)
	ParameterManagerCustomEndpointInput() *string
	ParameterManagerRegionalCustomEndpoint() *string
	SetParameterManagerRegionalCustomEndpoint(val *string)
	ParameterManagerRegionalCustomEndpointInput() *string
	PollInterval() *string
	SetPollInterval(val *string)
	PollIntervalInput() *string
	PrivatecaCustomEndpoint() *string
	SetPrivatecaCustomEndpoint(val *string)
	PrivatecaCustomEndpointInput() *string
	PrivilegedAccessManagerCustomEndpoint() *string
	SetPrivilegedAccessManagerCustomEndpoint(val *string)
	PrivilegedAccessManagerCustomEndpointInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	PublicCaCustomEndpoint() *string
	SetPublicCaCustomEndpoint(val *string)
	PublicCaCustomEndpointInput() *string
	PubsubCustomEndpoint() *string
	SetPubsubCustomEndpoint(val *string)
	PubsubCustomEndpointInput() *string
	PubsubLiteCustomEndpoint() *string
	SetPubsubLiteCustomEndpoint(val *string)
	PubsubLiteCustomEndpointInput() *string
	// Experimental.
	RawOverrides() interface{}
	RecaptchaEnterpriseCustomEndpoint() *string
	SetRecaptchaEnterpriseCustomEndpoint(val *string)
	RecaptchaEnterpriseCustomEndpointInput() *string
	RedisCustomEndpoint() *string
	SetRedisCustomEndpoint(val *string)
	RedisCustomEndpointInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RequestReason() *string
	SetRequestReason(val *string)
	RequestReasonInput() *string
	RequestTimeout() *string
	SetRequestTimeout(val *string)
	RequestTimeoutInput() *string
	ResourceManager3CustomEndpoint() *string
	SetResourceManager3CustomEndpoint(val *string)
	ResourceManager3CustomEndpointInput() *string
	ResourceManagerCustomEndpoint() *string
	SetResourceManagerCustomEndpoint(val *string)
	ResourceManagerCustomEndpointInput() *string
	ResourceManagerV3CustomEndpoint() *string
	SetResourceManagerV3CustomEndpoint(val *string)
	ResourceManagerV3CustomEndpointInput() *string
	RuntimeconfigCustomEndpoint() *string
	SetRuntimeconfigCustomEndpoint(val *string)
	RuntimeConfigCustomEndpoint() *string
	SetRuntimeConfigCustomEndpoint(val *string)
	RuntimeconfigCustomEndpointInput() *string
	RuntimeConfigCustomEndpointInput() *string
	SaasRuntimeCustomEndpoint() *string
	SetSaasRuntimeCustomEndpoint(val *string)
	SaasRuntimeCustomEndpointInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	SecretManagerCustomEndpoint() *string
	SetSecretManagerCustomEndpoint(val *string)
	SecretManagerCustomEndpointInput() *string
	SecretManagerRegionalCustomEndpoint() *string
	SetSecretManagerRegionalCustomEndpoint(val *string)
	SecretManagerRegionalCustomEndpointInput() *string
	SecureSourceManagerCustomEndpoint() *string
	SetSecureSourceManagerCustomEndpoint(val *string)
	SecureSourceManagerCustomEndpointInput() *string
	SecurityCenterCustomEndpoint() *string
	SetSecurityCenterCustomEndpoint(val *string)
	SecurityCenterCustomEndpointInput() *string
	SecurityCenterManagementCustomEndpoint() *string
	SetSecurityCenterManagementCustomEndpoint(val *string)
	SecurityCenterManagementCustomEndpointInput() *string
	SecurityCenterV2CustomEndpoint() *string
	SetSecurityCenterV2CustomEndpoint(val *string)
	SecurityCenterV2CustomEndpointInput() *string
	SecuritypostureCustomEndpoint() *string
	SetSecuritypostureCustomEndpoint(val *string)
	SecuritypostureCustomEndpointInput() *string
	SecurityScannerCustomEndpoint() *string
	SetSecurityScannerCustomEndpoint(val *string)
	SecurityScannerCustomEndpointInput() *string
	ServiceDirectoryCustomEndpoint() *string
	SetServiceDirectoryCustomEndpoint(val *string)
	ServiceDirectoryCustomEndpointInput() *string
	ServiceManagementCustomEndpoint() *string
	SetServiceManagementCustomEndpoint(val *string)
	ServiceManagementCustomEndpointInput() *string
	ServiceNetworkingCustomEndpoint() *string
	SetServiceNetworkingCustomEndpoint(val *string)
	ServiceNetworkingCustomEndpointInput() *string
	ServiceUsageCustomEndpoint() *string
	SetServiceUsageCustomEndpoint(val *string)
	ServiceUsageCustomEndpointInput() *string
	SiteVerificationCustomEndpoint() *string
	SetSiteVerificationCustomEndpoint(val *string)
	SiteVerificationCustomEndpointInput() *string
	SourceRepoCustomEndpoint() *string
	SetSourceRepoCustomEndpoint(val *string)
	SourceRepoCustomEndpointInput() *string
	SpannerCustomEndpoint() *string
	SetSpannerCustomEndpoint(val *string)
	SpannerCustomEndpointInput() *string
	SqlCustomEndpoint() *string
	SetSqlCustomEndpoint(val *string)
	SqlCustomEndpointInput() *string
	StorageBatchOperationsCustomEndpoint() *string
	SetStorageBatchOperationsCustomEndpoint(val *string)
	StorageBatchOperationsCustomEndpointInput() *string
	StorageControlCustomEndpoint() *string
	SetStorageControlCustomEndpoint(val *string)
	StorageControlCustomEndpointInput() *string
	StorageCustomEndpoint() *string
	SetStorageCustomEndpoint(val *string)
	StorageCustomEndpointInput() *string
	StorageInsightsCustomEndpoint() *string
	SetStorageInsightsCustomEndpoint(val *string)
	StorageInsightsCustomEndpointInput() *string
	StorageTransferCustomEndpoint() *string
	SetStorageTransferCustomEndpoint(val *string)
	StorageTransferCustomEndpointInput() *string
	TagsCustomEndpoint() *string
	SetTagsCustomEndpoint(val *string)
	TagsCustomEndpointInput() *string
	TagsLocationCustomEndpoint() *string
	SetTagsLocationCustomEndpoint(val *string)
	TagsLocationCustomEndpointInput() *string
	TerraformAttributionLabelAdditionStrategy() *string
	SetTerraformAttributionLabelAdditionStrategy(val *string)
	TerraformAttributionLabelAdditionStrategyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	TpuV2CustomEndpoint() *string
	SetTpuV2CustomEndpoint(val *string)
	TpuV2CustomEndpointInput() *string
	TranscoderCustomEndpoint() *string
	SetTranscoderCustomEndpoint(val *string)
	TranscoderCustomEndpointInput() *string
	UniverseDomain() *string
	SetUniverseDomain(val *string)
	UniverseDomainInput() *string
	UserProjectOverride() interface{}
	SetUserProjectOverride(val interface{})
	UserProjectOverrideInput() interface{}
	VectorSearchCustomEndpoint() *string
	SetVectorSearchCustomEndpoint(val *string)
	VectorSearchCustomEndpointInput() *string
	VertexAiCustomEndpoint() *string
	SetVertexAiCustomEndpoint(val *string)
	VertexAiCustomEndpointInput() *string
	VmwareengineCustomEndpoint() *string
	SetVmwareengineCustomEndpoint(val *string)
	VmwareengineCustomEndpointInput() *string
	VpcAccessCustomEndpoint() *string
	SetVpcAccessCustomEndpoint(val *string)
	VpcAccessCustomEndpointInput() *string
	WorkbenchCustomEndpoint() *string
	SetWorkbenchCustomEndpoint(val *string)
	WorkbenchCustomEndpointInput() *string
	WorkflowsCustomEndpoint() *string
	SetWorkflowsCustomEndpoint(val *string)
	WorkflowsCustomEndpointInput() *string
	WorkstationsCustomEndpoint() *string
	SetWorkstationsCustomEndpoint(val *string)
	WorkstationsCustomEndpointInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccessApprovalCustomEndpoint()
	ResetAccessContextManagerCustomEndpoint()
	ResetAccessToken()
	ResetActiveDirectoryCustomEndpoint()
	ResetAddTerraformAttributionLabel()
	ResetAlias()
	ResetAlloydbCustomEndpoint()
	ResetApiGatewayCustomEndpoint()
	ResetApigeeCustomEndpoint()
	ResetApihubCustomEndpoint()
	ResetApikeysCustomEndpoint()
	ResetAppEngineCustomEndpoint()
	ResetApphubCustomEndpoint()
	ResetArtifactRegistryCustomEndpoint()
	ResetAssuredWorkloadsCustomEndpoint()
	ResetBackupDrCustomEndpoint()
	ResetBatching()
	ResetBeyondcorpCustomEndpoint()
	ResetBiglakeCustomEndpoint()
	ResetBiglakeIcebergCustomEndpoint()
	ResetBigqueryAnalyticsHubCustomEndpoint()
	ResetBigqueryConnectionCustomEndpoint()
	ResetBigQueryCustomEndpoint()
	ResetBigqueryDatapolicyCustomEndpoint()
	ResetBigqueryDatapolicyv2CustomEndpoint()
	ResetBigqueryDataTransferCustomEndpoint()
	ResetBigqueryReservationCustomEndpoint()
	ResetBigtableCustomEndpoint()
	ResetBillingCustomEndpoint()
	ResetBillingProject()
	ResetBinaryAuthorizationCustomEndpoint()
	ResetBlockchainNodeEngineCustomEndpoint()
	ResetCertificateManagerCustomEndpoint()
	ResetCesCustomEndpoint()
	ResetChronicleCustomEndpoint()
	ResetCloudAssetCustomEndpoint()
	ResetCloudBillingCustomEndpoint()
	ResetCloudBuildCustomEndpoint()
	ResetCloudbuildv2CustomEndpoint()
	ResetClouddeployCustomEndpoint()
	ResetClouddomainsCustomEndpoint()
	ResetCloudfunctions2CustomEndpoint()
	ResetCloudFunctionsCustomEndpoint()
	ResetCloudIdentityCustomEndpoint()
	ResetCloudIdsCustomEndpoint()
	ResetCloudQuotasCustomEndpoint()
	ResetCloudResourceManagerCustomEndpoint()
	ResetCloudRunCustomEndpoint()
	ResetCloudRunV2CustomEndpoint()
	ResetCloudSchedulerCustomEndpoint()
	ResetCloudSecurityComplianceCustomEndpoint()
	ResetCloudTasksCustomEndpoint()
	ResetColabCustomEndpoint()
	ResetComposerCustomEndpoint()
	ResetComputeCustomEndpoint()
	ResetContactCenterInsightsCustomEndpoint()
	ResetContainerAnalysisCustomEndpoint()
	ResetContainerAttachedCustomEndpoint()
	ResetContainerAwsCustomEndpoint()
	ResetContainerAzureCustomEndpoint()
	ResetContainerCustomEndpoint()
	ResetCoreBillingCustomEndpoint()
	ResetCredentials()
	ResetDatabaseMigrationServiceCustomEndpoint()
	ResetDataCatalogCustomEndpoint()
	ResetDataflowCustomEndpoint()
	ResetDataformCustomEndpoint()
	ResetDataFusionCustomEndpoint()
	ResetDataLossPreventionCustomEndpoint()
	ResetDataPipelineCustomEndpoint()
	ResetDataplexCustomEndpoint()
	ResetDataprocCustomEndpoint()
	ResetDataprocGdcCustomEndpoint()
	ResetDataprocMetastoreCustomEndpoint()
	ResetDatastreamCustomEndpoint()
	ResetDefaultLabels()
	ResetDeploymentManagerCustomEndpoint()
	ResetDeveloperConnectCustomEndpoint()
	ResetDialogflowCustomEndpoint()
	ResetDialogflowCxCustomEndpoint()
	ResetDiscoveryEngineCustomEndpoint()
	ResetDnsCustomEndpoint()
	ResetDocumentAiCustomEndpoint()
	ResetDocumentAiWarehouseCustomEndpoint()
	ResetEdgecontainerCustomEndpoint()
	ResetEdgenetworkCustomEndpoint()
	ResetEssentialContactsCustomEndpoint()
	ResetEventarcCustomEndpoint()
	ResetExternalCredentials()
	ResetFilestoreCustomEndpoint()
	ResetFirebaseAiLogicCustomEndpoint()
	ResetFirebaseAppCheckCustomEndpoint()
	ResetFirebaseAppHostingCustomEndpoint()
	ResetFirebaseCustomEndpoint()
	ResetFirebaseDatabaseCustomEndpoint()
	ResetFirebaseDataConnectCustomEndpoint()
	ResetFirebaseExtensionsCustomEndpoint()
	ResetFirebaseHostingCustomEndpoint()
	ResetFirebaserulesCustomEndpoint()
	ResetFirebaseStorageCustomEndpoint()
	ResetFirestoreCustomEndpoint()
	ResetGeminiCustomEndpoint()
	ResetGkeBackupCustomEndpoint()
	ResetGkeHub2CustomEndpoint()
	ResetGkeHubCustomEndpoint()
	ResetGkeonpremCustomEndpoint()
	ResetHealthcareCustomEndpoint()
	ResetHypercomputeclusterCustomEndpoint()
	ResetIam2CustomEndpoint()
	ResetIam3CustomEndpoint()
	ResetIamBetaCustomEndpoint()
	ResetIamCredentialsCustomEndpoint()
	ResetIamCustomEndpoint()
	ResetIamWorkforcePoolCustomEndpoint()
	ResetIapCustomEndpoint()
	ResetIdentityPlatformCustomEndpoint()
	ResetImpersonateServiceAccount()
	ResetImpersonateServiceAccountDelegates()
	ResetIntegrationConnectorsCustomEndpoint()
	ResetIntegrationsCustomEndpoint()
	ResetKmsCustomEndpoint()
	ResetLoggingCustomEndpoint()
	ResetLookerCustomEndpoint()
	ResetLustreCustomEndpoint()
	ResetManagedKafkaCustomEndpoint()
	ResetMemcacheCustomEndpoint()
	ResetMemorystoreCustomEndpoint()
	ResetMigrationCenterCustomEndpoint()
	ResetMlEngineCustomEndpoint()
	ResetModelArmorCustomEndpoint()
	ResetModelArmorGlobalCustomEndpoint()
	ResetMonitoringCustomEndpoint()
	ResetNetappCustomEndpoint()
	ResetNetworkConnectivityCustomEndpoint()
	ResetNetworkConnectivityv1CustomEndpoint()
	ResetNetworkManagementCustomEndpoint()
	ResetNetworkSecurityCustomEndpoint()
	ResetNetworkServicesCustomEndpoint()
	ResetNotebooksCustomEndpoint()
	ResetObservabilityCustomEndpoint()
	ResetOracleDatabaseCustomEndpoint()
	ResetOrgPolicyCustomEndpoint()
	ResetOsConfigCustomEndpoint()
	ResetOsConfigV2CustomEndpoint()
	ResetOsLoginCustomEndpoint()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParallelstoreCustomEndpoint()
	ResetParameterManagerCustomEndpoint()
	ResetParameterManagerRegionalCustomEndpoint()
	ResetPollInterval()
	ResetPrivatecaCustomEndpoint()
	ResetPrivilegedAccessManagerCustomEndpoint()
	ResetProject()
	ResetPublicCaCustomEndpoint()
	ResetPubsubCustomEndpoint()
	ResetPubsubLiteCustomEndpoint()
	ResetRecaptchaEnterpriseCustomEndpoint()
	ResetRedisCustomEndpoint()
	ResetRegion()
	ResetRequestReason()
	ResetRequestTimeout()
	ResetResourceManager3CustomEndpoint()
	ResetResourceManagerCustomEndpoint()
	ResetResourceManagerV3CustomEndpoint()
	ResetRuntimeconfigCustomEndpoint()
	ResetRuntimeConfigCustomEndpoint()
	ResetSaasRuntimeCustomEndpoint()
	ResetScopes()
	ResetSecretManagerCustomEndpoint()
	ResetSecretManagerRegionalCustomEndpoint()
	ResetSecureSourceManagerCustomEndpoint()
	ResetSecurityCenterCustomEndpoint()
	ResetSecurityCenterManagementCustomEndpoint()
	ResetSecurityCenterV2CustomEndpoint()
	ResetSecuritypostureCustomEndpoint()
	ResetSecurityScannerCustomEndpoint()
	ResetServiceDirectoryCustomEndpoint()
	ResetServiceManagementCustomEndpoint()
	ResetServiceNetworkingCustomEndpoint()
	ResetServiceUsageCustomEndpoint()
	ResetSiteVerificationCustomEndpoint()
	ResetSourceRepoCustomEndpoint()
	ResetSpannerCustomEndpoint()
	ResetSqlCustomEndpoint()
	ResetStorageBatchOperationsCustomEndpoint()
	ResetStorageControlCustomEndpoint()
	ResetStorageCustomEndpoint()
	ResetStorageInsightsCustomEndpoint()
	ResetStorageTransferCustomEndpoint()
	ResetTagsCustomEndpoint()
	ResetTagsLocationCustomEndpoint()
	ResetTerraformAttributionLabelAdditionStrategy()
	ResetTpuV2CustomEndpoint()
	ResetTranscoderCustomEndpoint()
	ResetUniverseDomain()
	ResetUserProjectOverride()
	ResetVectorSearchCustomEndpoint()
	ResetVertexAiCustomEndpoint()
	ResetVmwareengineCustomEndpoint()
	ResetVpcAccessCustomEndpoint()
	ResetWorkbenchCustomEndpoint()
	ResetWorkflowsCustomEndpoint()
	ResetWorkstationsCustomEndpoint()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GooglebetaProvider
type jsiiProxy_GooglebetaProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_GooglebetaProvider) AccessApprovalCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessApprovalCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AccessApprovalCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessApprovalCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AccessContextManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessContextManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AccessContextManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessContextManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ActiveDirectoryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ActiveDirectoryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AddTerraformAttributionLabel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addTerraformAttributionLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AddTerraformAttributionLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addTerraformAttributionLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AlloydbCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alloydbCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AlloydbCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alloydbCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ApiGatewayCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGatewayCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ApiGatewayCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGatewayCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ApigeeCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apigeeCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ApigeeCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apigeeCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ApihubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apihubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ApihubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apihubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ApikeysCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apikeysCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ApikeysCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apikeysCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AppEngineCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEngineCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AppEngineCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEngineCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ApphubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apphubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ApphubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apphubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ArtifactRegistryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactRegistryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ArtifactRegistryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactRegistryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AssuredWorkloadsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assuredWorkloadsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) AssuredWorkloadsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assuredWorkloadsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BackupDrCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDrCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BackupDrCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDrCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Batching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BatchingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BeyondcorpCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beyondcorpCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BeyondcorpCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beyondcorpCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BiglakeCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"biglakeCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BiglakeCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"biglakeCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BiglakeIcebergCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"biglakeIcebergCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BiglakeIcebergCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"biglakeIcebergCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryAnalyticsHubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryAnalyticsHubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryAnalyticsHubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryAnalyticsHubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryConnectionCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryConnectionCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryConnectionCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryConnectionCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigQueryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigQueryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigQueryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigQueryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryDatapolicyCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDatapolicyCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryDatapolicyCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDatapolicyCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryDatapolicyv2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDatapolicyv2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryDatapolicyv2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDatapolicyv2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryDataTransferCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDataTransferCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryDataTransferCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDataTransferCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryReservationCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryReservationCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigqueryReservationCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryReservationCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigtableCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigtableCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BigtableCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigtableCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BillingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BillingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BillingProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BillingProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BinaryAuthorizationCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAuthorizationCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BinaryAuthorizationCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAuthorizationCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BlockchainNodeEngineCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockchainNodeEngineCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) BlockchainNodeEngineCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockchainNodeEngineCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CertificateManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CertificateManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CesCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cesCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CesCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cesCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ChronicleCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chronicleCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ChronicleCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chronicleCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudAssetCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAssetCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudAssetCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAssetCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudBillingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBillingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudBillingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBillingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudBuildCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuildCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudBuildCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuildCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Cloudbuildv2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudbuildv2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Cloudbuildv2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudbuildv2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ClouddeployCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clouddeployCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ClouddeployCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clouddeployCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ClouddomainsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clouddomainsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ClouddomainsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clouddomainsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Cloudfunctions2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfunctions2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Cloudfunctions2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfunctions2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudFunctionsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFunctionsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudFunctionsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFunctionsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudIdentityCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdentityCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudIdentityCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdentityCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudIdsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudIdsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudQuotasCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudQuotasCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudQuotasCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudQuotasCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudResourceManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudResourceManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudResourceManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudResourceManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudRunCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRunCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudRunCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRunCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudRunV2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRunV2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudRunV2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRunV2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudSchedulerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSchedulerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudSchedulerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSchedulerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudSecurityComplianceCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSecurityComplianceCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudSecurityComplianceCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSecurityComplianceCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudTasksCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudTasksCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CloudTasksCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudTasksCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ColabCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colabCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ColabCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colabCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ComposerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ComposerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ComputeCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ComputeCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContactCenterInsightsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactCenterInsightsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContactCenterInsightsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactCenterInsightsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContainerAnalysisCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAnalysisCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContainerAnalysisCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAnalysisCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContainerAttachedCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAttachedCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContainerAttachedCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAttachedCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContainerAwsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAwsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContainerAwsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAwsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContainerAzureCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAzureCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContainerAzureCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAzureCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContainerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ContainerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CoreBillingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreBillingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CoreBillingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreBillingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Credentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) CredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DatabaseMigrationServiceCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMigrationServiceCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DatabaseMigrationServiceCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMigrationServiceCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataCatalogCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCatalogCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataCatalogCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCatalogCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataflowCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataflowCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataflowCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataflowCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataformCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataformCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataformCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataformCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataFusionCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFusionCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataFusionCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFusionCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataLossPreventionCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLossPreventionCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataLossPreventionCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLossPreventionCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataPipelineCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPipelineCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataPipelineCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPipelineCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataplexCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataplexCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataplexCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataplexCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataprocCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataprocCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataprocGdcCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocGdcCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataprocGdcCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocGdcCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataprocMetastoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocMetastoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DataprocMetastoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocMetastoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DatastreamCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastreamCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DatastreamCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastreamCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DefaultLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DefaultLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DeploymentManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DeploymentManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DeveloperConnectCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerConnectCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DeveloperConnectCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerConnectCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DialogflowCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DialogflowCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DialogflowCxCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCxCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DialogflowCxCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCxCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DiscoveryEngineCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryEngineCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DiscoveryEngineCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryEngineCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DnsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DnsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DocumentAiCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentAiCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DocumentAiCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentAiCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DocumentAiWarehouseCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentAiWarehouseCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) DocumentAiWarehouseCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentAiWarehouseCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) EdgecontainerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgecontainerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) EdgecontainerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgecontainerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) EdgenetworkCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgenetworkCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) EdgenetworkCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgenetworkCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) EssentialContactsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"essentialContactsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) EssentialContactsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"essentialContactsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) EventarcCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventarcCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) EventarcCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventarcCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ExternalCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ExternalCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FilestoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filestoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FilestoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filestoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseAiLogicCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseAiLogicCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseAiLogicCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseAiLogicCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseAppCheckCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseAppCheckCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseAppCheckCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseAppCheckCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseAppHostingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseAppHostingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseAppHostingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseAppHostingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseDatabaseCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseDatabaseCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseDatabaseCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseDatabaseCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseDataConnectCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseDataConnectCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseDataConnectCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseDataConnectCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseExtensionsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseExtensionsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseExtensionsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseExtensionsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseHostingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseHostingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseHostingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseHostingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaserulesCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaserulesCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaserulesCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaserulesCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseStorageCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseStorageCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirebaseStorageCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseStorageCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirestoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firestoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FirestoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firestoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) GeminiCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geminiCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) GeminiCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geminiCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) GkeBackupCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeBackupCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) GkeBackupCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeBackupCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) GkeHub2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeHub2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) GkeHub2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeHub2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) GkeHubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeHubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) GkeHubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeHubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) GkeonpremCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeonpremCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) GkeonpremCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeonpremCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) HealthcareCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthcareCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) HealthcareCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthcareCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) HypercomputeclusterCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypercomputeclusterCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) HypercomputeclusterCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypercomputeclusterCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Iam2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iam2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Iam2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iam2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Iam3CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iam3CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Iam3CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iam3CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IamBetaCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamBetaCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IamBetaCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamBetaCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IamCredentialsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCredentialsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IamCredentialsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCredentialsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IamCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IamCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IamWorkforcePoolCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamWorkforcePoolCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IamWorkforcePoolCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamWorkforcePoolCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IapCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iapCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IapCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iapCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IdentityPlatformCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPlatformCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IdentityPlatformCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPlatformCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ImpersonateServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impersonateServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ImpersonateServiceAccountDelegates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"impersonateServiceAccountDelegates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ImpersonateServiceAccountDelegatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"impersonateServiceAccountDelegatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ImpersonateServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impersonateServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IntegrationConnectorsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationConnectorsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IntegrationConnectorsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationConnectorsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IntegrationsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) IntegrationsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) KmsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) KmsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) LoggingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) LoggingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) LookerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) LookerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) LustreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lustreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) LustreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lustreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ManagedKafkaCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedKafkaCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ManagedKafkaCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedKafkaCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MemcacheCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memcacheCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MemcacheCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memcacheCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MemorystoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memorystoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MemorystoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memorystoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MigrationCenterCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationCenterCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MigrationCenterCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationCenterCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MlEngineCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlEngineCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MlEngineCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlEngineCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ModelArmorCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelArmorCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ModelArmorCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelArmorCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ModelArmorGlobalCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelArmorGlobalCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ModelArmorGlobalCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelArmorGlobalCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MonitoringCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) MonitoringCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetappCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netappCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetappCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netappCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetworkConnectivityCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetworkConnectivityCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetworkConnectivityv1CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityv1CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetworkConnectivityv1CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityv1CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetworkManagementCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkManagementCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetworkManagementCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkManagementCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetworkSecurityCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetworkSecurityCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetworkServicesCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkServicesCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NetworkServicesCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkServicesCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NotebooksCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebooksCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) NotebooksCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebooksCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ObservabilityCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"observabilityCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ObservabilityCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"observabilityCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) OracleDatabaseCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oracleDatabaseCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) OracleDatabaseCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oracleDatabaseCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) OrgPolicyCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgPolicyCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) OrgPolicyCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgPolicyCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) OsConfigCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osConfigCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) OsConfigCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osConfigCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) OsConfigV2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osConfigV2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) OsConfigV2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osConfigV2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) OsLoginCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osLoginCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) OsLoginCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osLoginCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ParallelstoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parallelstoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ParallelstoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parallelstoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ParameterManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ParameterManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ParameterManagerRegionalCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterManagerRegionalCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ParameterManagerRegionalCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterManagerRegionalCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PollInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PollIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PrivatecaCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privatecaCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PrivatecaCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privatecaCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PrivilegedAccessManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedAccessManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PrivilegedAccessManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedAccessManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PublicCaCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicCaCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PublicCaCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicCaCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PubsubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PubsubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PubsubLiteCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubLiteCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) PubsubLiteCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubLiteCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RecaptchaEnterpriseCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recaptchaEnterpriseCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RecaptchaEnterpriseCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recaptchaEnterpriseCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RedisCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RedisCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RequestReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RequestReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RequestTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RequestTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ResourceManager3CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManager3CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ResourceManager3CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManager3CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ResourceManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ResourceManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ResourceManagerV3CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerV3CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ResourceManagerV3CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerV3CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RuntimeconfigCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeconfigCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RuntimeConfigCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeConfigCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RuntimeconfigCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeconfigCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) RuntimeConfigCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeConfigCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SaasRuntimeCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saasRuntimeCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SaasRuntimeCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saasRuntimeCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecretManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecretManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecretManagerRegionalCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerRegionalCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecretManagerRegionalCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerRegionalCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecureSourceManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureSourceManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecureSourceManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureSourceManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecurityCenterCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityCenterCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecurityCenterCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityCenterCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecurityCenterManagementCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityCenterManagementCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecurityCenterManagementCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityCenterManagementCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecurityCenterV2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityCenterV2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecurityCenterV2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityCenterV2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecuritypostureCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitypostureCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecuritypostureCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitypostureCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecurityScannerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityScannerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SecurityScannerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityScannerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ServiceDirectoryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectoryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ServiceDirectoryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectoryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ServiceManagementCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceManagementCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ServiceManagementCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceManagementCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ServiceNetworkingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ServiceNetworkingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ServiceUsageCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUsageCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ServiceUsageCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUsageCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SiteVerificationCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteVerificationCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SiteVerificationCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteVerificationCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SourceRepoCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRepoCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SourceRepoCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRepoCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SpannerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spannerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SpannerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spannerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SqlCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) SqlCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) StorageBatchOperationsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBatchOperationsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) StorageBatchOperationsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBatchOperationsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) StorageControlCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageControlCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) StorageControlCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageControlCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) StorageCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) StorageCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) StorageInsightsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageInsightsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) StorageInsightsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageInsightsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) StorageTransferCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTransferCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) StorageTransferCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTransferCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TagsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TagsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TagsLocationCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsLocationCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TagsLocationCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsLocationCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TerraformAttributionLabelAdditionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttributionLabelAdditionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TerraformAttributionLabelAdditionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttributionLabelAdditionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TpuV2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuV2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TpuV2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuV2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TranscoderCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transcoderCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) TranscoderCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transcoderCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) UniverseDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"universeDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) UniverseDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"universeDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) UserProjectOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userProjectOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) UserProjectOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userProjectOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) VectorSearchCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorSearchCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) VectorSearchCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorSearchCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) VertexAiCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vertexAiCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) VertexAiCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vertexAiCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) VmwareengineCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmwareengineCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) VmwareengineCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmwareengineCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) VpcAccessCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcAccessCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) VpcAccessCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcAccessCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) WorkbenchCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workbenchCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) WorkbenchCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workbenchCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) WorkflowsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) WorkflowsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) WorkstationsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workstationsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) WorkstationsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workstationsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglebetaProvider) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs googlebeta} Resource.
func NewGooglebetaProvider(scope constructs.Construct, id *string, config *GooglebetaProviderConfig) GooglebetaProvider {
	_init_.Initialize()

	if err := validateNewGooglebetaProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglebetaProvider{}

	_jsii_.Create(
		"googlebeta.provider.GooglebetaProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs googlebeta} Resource.
func NewGooglebetaProvider_Override(g GooglebetaProvider, scope constructs.Construct, id *string, config *GooglebetaProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"googlebeta.provider.GooglebetaProvider",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetAccessApprovalCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"accessApprovalCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetAccessContextManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"accessContextManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetAccessToken(val *string) {
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetActiveDirectoryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"activeDirectoryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetAddTerraformAttributionLabel(val interface{}) {
	if err := j.validateSetAddTerraformAttributionLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addTerraformAttributionLabel",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetAlloydbCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"alloydbCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetApiGatewayCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apiGatewayCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetApigeeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apigeeCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetApihubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apihubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetApikeysCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apikeysCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetAppEngineCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"appEngineCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetApphubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apphubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetArtifactRegistryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"artifactRegistryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetAssuredWorkloadsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"assuredWorkloadsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBackupDrCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"backupDrCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBatching(val interface{}) {
	if err := j.validateSetBatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batching",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBeyondcorpCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"beyondcorpCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBiglakeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"biglakeCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBiglakeIcebergCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"biglakeIcebergCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBigqueryAnalyticsHubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryAnalyticsHubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBigqueryConnectionCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryConnectionCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBigQueryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigQueryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBigqueryDatapolicyCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryDatapolicyCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBigqueryDatapolicyv2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryDatapolicyv2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBigqueryDataTransferCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryDataTransferCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBigqueryReservationCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryReservationCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBigtableCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigtableCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBillingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"billingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBillingProject(val *string) {
	_jsii_.Set(
		j,
		"billingProject",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBinaryAuthorizationCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"binaryAuthorizationCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetBlockchainNodeEngineCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"blockchainNodeEngineCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCertificateManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"certificateManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCesCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cesCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetChronicleCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"chronicleCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudAssetCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudAssetCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudBillingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudBillingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudBuildCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudBuildCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudbuildv2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudbuildv2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetClouddeployCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"clouddeployCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetClouddomainsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"clouddomainsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudfunctions2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudfunctions2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudFunctionsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudFunctionsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudIdentityCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudIdentityCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudIdsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudIdsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudQuotasCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudQuotasCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudResourceManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudResourceManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudRunCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudRunCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudRunV2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudRunV2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudSchedulerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudSchedulerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudSecurityComplianceCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudSecurityComplianceCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCloudTasksCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudTasksCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetColabCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"colabCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetComposerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"composerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetComputeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"computeCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetContactCenterInsightsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"contactCenterInsightsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetContainerAnalysisCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAnalysisCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetContainerAttachedCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAttachedCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetContainerAwsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAwsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetContainerAzureCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAzureCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetContainerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCoreBillingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"coreBillingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetCredentials(val *string) {
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDatabaseMigrationServiceCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"databaseMigrationServiceCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDataCatalogCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataCatalogCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDataflowCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataflowCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDataformCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataformCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDataFusionCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataFusionCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDataLossPreventionCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataLossPreventionCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDataPipelineCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataPipelineCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDataplexCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataplexCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDataprocCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataprocCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDataprocGdcCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataprocGdcCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDataprocMetastoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataprocMetastoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDatastreamCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"datastreamCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDefaultLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"defaultLabels",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDeploymentManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"deploymentManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDeveloperConnectCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"developerConnectCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDialogflowCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dialogflowCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDialogflowCxCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dialogflowCxCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDiscoveryEngineCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"discoveryEngineCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDnsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dnsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDocumentAiCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"documentAiCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetDocumentAiWarehouseCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"documentAiWarehouseCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetEdgecontainerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"edgecontainerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetEdgenetworkCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"edgenetworkCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetEssentialContactsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"essentialContactsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetEventarcCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"eventarcCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetExternalCredentials(val interface{}) {
	if err := j.validateSetExternalCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalCredentials",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFilestoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"filestoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirebaseAiLogicCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaseAiLogicCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirebaseAppCheckCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaseAppCheckCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirebaseAppHostingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaseAppHostingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirebaseCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaseCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirebaseDatabaseCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaseDatabaseCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirebaseDataConnectCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaseDataConnectCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirebaseExtensionsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaseExtensionsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirebaseHostingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaseHostingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirebaserulesCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaserulesCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirebaseStorageCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaseStorageCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetFirestoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firestoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetGeminiCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"geminiCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetGkeBackupCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gkeBackupCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetGkeHub2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gkeHub2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetGkeHubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gkeHubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetGkeonpremCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gkeonpremCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetHealthcareCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"healthcareCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetHypercomputeclusterCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"hypercomputeclusterCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetIam2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iam2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetIam3CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iam3CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetIamBetaCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamBetaCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetIamCredentialsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamCredentialsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetIamCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetIamWorkforcePoolCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamWorkforcePoolCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetIapCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iapCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetIdentityPlatformCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"identityPlatformCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetImpersonateServiceAccount(val *string) {
	_jsii_.Set(
		j,
		"impersonateServiceAccount",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetImpersonateServiceAccountDelegates(val *[]*string) {
	_jsii_.Set(
		j,
		"impersonateServiceAccountDelegates",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetIntegrationConnectorsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"integrationConnectorsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetIntegrationsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"integrationsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetKmsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"kmsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetLoggingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"loggingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetLookerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"lookerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetLustreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"lustreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetManagedKafkaCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"managedKafkaCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetMemcacheCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"memcacheCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetMemorystoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"memorystoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetMigrationCenterCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"migrationCenterCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetMlEngineCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"mlEngineCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetModelArmorCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"modelArmorCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetModelArmorGlobalCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"modelArmorGlobalCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetMonitoringCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"monitoringCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetNetappCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"netappCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetNetworkConnectivityCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkConnectivityCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetNetworkConnectivityv1CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkConnectivityv1CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetNetworkManagementCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkManagementCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetNetworkSecurityCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkSecurityCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetNetworkServicesCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkServicesCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetNotebooksCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"notebooksCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetObservabilityCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"observabilityCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetOracleDatabaseCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"oracleDatabaseCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetOrgPolicyCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"orgPolicyCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetOsConfigCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"osConfigCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetOsConfigV2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"osConfigV2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetOsLoginCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"osLoginCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetParallelstoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"parallelstoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetParameterManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"parameterManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetParameterManagerRegionalCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"parameterManagerRegionalCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetPollInterval(val *string) {
	_jsii_.Set(
		j,
		"pollInterval",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetPrivatecaCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"privatecaCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetPrivilegedAccessManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"privilegedAccessManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetProject(val *string) {
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetPublicCaCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"publicCaCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetPubsubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"pubsubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetPubsubLiteCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"pubsubLiteCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetRecaptchaEnterpriseCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"recaptchaEnterpriseCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetRedisCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"redisCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetRequestReason(val *string) {
	_jsii_.Set(
		j,
		"requestReason",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetRequestTimeout(val *string) {
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetResourceManager3CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"resourceManager3CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetResourceManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"resourceManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetResourceManagerV3CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"resourceManagerV3CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetRuntimeconfigCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"runtimeconfigCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetRuntimeConfigCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"runtimeConfigCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSaasRuntimeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"saasRuntimeCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSecretManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"secretManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSecretManagerRegionalCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"secretManagerRegionalCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSecureSourceManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"secureSourceManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSecurityCenterCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"securityCenterCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSecurityCenterManagementCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"securityCenterManagementCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSecurityCenterV2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"securityCenterV2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSecuritypostureCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"securitypostureCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSecurityScannerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"securityScannerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetServiceDirectoryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceDirectoryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetServiceManagementCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceManagementCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetServiceNetworkingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceNetworkingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetServiceUsageCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceUsageCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSiteVerificationCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"siteVerificationCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSourceRepoCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"sourceRepoCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSpannerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"spannerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetSqlCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"sqlCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetStorageBatchOperationsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageBatchOperationsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetStorageControlCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageControlCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetStorageCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetStorageInsightsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageInsightsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetStorageTransferCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageTransferCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetTagsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"tagsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetTagsLocationCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"tagsLocationCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetTerraformAttributionLabelAdditionStrategy(val *string) {
	_jsii_.Set(
		j,
		"terraformAttributionLabelAdditionStrategy",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetTpuV2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"tpuV2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetTranscoderCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"transcoderCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetUniverseDomain(val *string) {
	_jsii_.Set(
		j,
		"universeDomain",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetUserProjectOverride(val interface{}) {
	if err := j.validateSetUserProjectOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userProjectOverride",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetVectorSearchCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"vectorSearchCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetVertexAiCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"vertexAiCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetVmwareengineCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"vmwareengineCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetVpcAccessCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"vpcAccessCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetWorkbenchCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"workbenchCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetWorkflowsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"workflowsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetWorkstationsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"workstationsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GooglebetaProvider)SetZone(val *string) {
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a GooglebetaProvider resource upon running "cdktf plan <stack-name>".
func GooglebetaProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGooglebetaProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"googlebeta.provider.GooglebetaProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GooglebetaProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglebetaProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"googlebeta.provider.GooglebetaProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GooglebetaProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglebetaProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"googlebeta.provider.GooglebetaProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GooglebetaProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglebetaProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"googlebeta.provider.GooglebetaProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GooglebetaProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"googlebeta.provider.GooglebetaProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GooglebetaProvider) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GooglebetaProvider) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetAccessApprovalCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessApprovalCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetAccessContextManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessContextManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetAccessToken() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetActiveDirectoryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetActiveDirectoryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetAddTerraformAttributionLabel() {
	_jsii_.InvokeVoid(
		g,
		"resetAddTerraformAttributionLabel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		g,
		"resetAlias",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetAlloydbCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAlloydbCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetApiGatewayCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetApiGatewayCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetApigeeCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetApigeeCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetApihubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetApihubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetApikeysCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetApikeysCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetAppEngineCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAppEngineCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetApphubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetApphubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetArtifactRegistryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetArtifactRegistryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetAssuredWorkloadsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAssuredWorkloadsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBackupDrCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBackupDrCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBatching() {
	_jsii_.InvokeVoid(
		g,
		"resetBatching",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBeyondcorpCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBeyondcorpCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBiglakeCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBiglakeCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBiglakeIcebergCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBiglakeIcebergCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBigqueryAnalyticsHubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryAnalyticsHubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBigqueryConnectionCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryConnectionCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBigQueryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigQueryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBigqueryDatapolicyCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryDatapolicyCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBigqueryDatapolicyv2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryDatapolicyv2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBigqueryDataTransferCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryDataTransferCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBigqueryReservationCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryReservationCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBigtableCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigtableCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBillingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBillingProject() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBinaryAuthorizationCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBinaryAuthorizationCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetBlockchainNodeEngineCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBlockchainNodeEngineCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCertificateManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCesCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCesCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetChronicleCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetChronicleCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudAssetCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudAssetCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudBillingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudBillingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudBuildCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudBuildCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudbuildv2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudbuildv2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetClouddeployCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetClouddeployCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetClouddomainsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetClouddomainsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudfunctions2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudfunctions2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudFunctionsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudFunctionsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudIdentityCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudIdentityCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudIdsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudIdsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudQuotasCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudQuotasCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudResourceManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudResourceManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudRunCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudRunCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudRunV2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudRunV2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudSchedulerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudSchedulerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudSecurityComplianceCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudSecurityComplianceCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCloudTasksCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudTasksCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetColabCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetColabCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetComposerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetComposerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetComputeCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetContactCenterInsightsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContactCenterInsightsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetContainerAnalysisCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAnalysisCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetContainerAttachedCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAttachedCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetContainerAwsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAwsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetContainerAzureCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAzureCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetContainerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCoreBillingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCoreBillingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDatabaseMigrationServiceCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseMigrationServiceCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDataCatalogCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataCatalogCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDataflowCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataflowCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDataformCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataformCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDataFusionCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataFusionCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDataLossPreventionCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataLossPreventionCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDataPipelineCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataPipelineCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDataplexCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataplexCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDataprocCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataprocCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDataprocGdcCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataprocGdcCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDataprocMetastoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataprocMetastoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDatastreamCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDatastreamCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDefaultLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDeploymentManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDeveloperConnectCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDeveloperConnectCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDialogflowCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDialogflowCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDialogflowCxCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDialogflowCxCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDiscoveryEngineCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDiscoveryEngineCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDnsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDocumentAiCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDocumentAiCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetDocumentAiWarehouseCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDocumentAiWarehouseCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetEdgecontainerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEdgecontainerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetEdgenetworkCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEdgenetworkCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetEssentialContactsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEssentialContactsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetEventarcCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEventarcCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetExternalCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFilestoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFilestoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirebaseAiLogicCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaseAiLogicCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirebaseAppCheckCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaseAppCheckCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirebaseAppHostingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaseAppHostingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirebaseCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaseCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirebaseDatabaseCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaseDatabaseCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirebaseDataConnectCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaseDataConnectCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirebaseExtensionsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaseExtensionsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirebaseHostingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaseHostingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirebaserulesCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaserulesCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirebaseStorageCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaseStorageCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetFirestoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirestoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetGeminiCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGeminiCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetGkeBackupCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeBackupCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetGkeHub2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeHub2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetGkeHubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeHubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetGkeonpremCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeonpremCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetHealthcareCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthcareCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetHypercomputeclusterCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetHypercomputeclusterCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetIam2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIam2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetIam3CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIam3CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetIamBetaCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamBetaCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetIamCredentialsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamCredentialsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetIamCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetIamWorkforcePoolCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamWorkforcePoolCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetIapCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIapCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetIdentityPlatformCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIdentityPlatformCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetImpersonateServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetImpersonateServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetImpersonateServiceAccountDelegates() {
	_jsii_.InvokeVoid(
		g,
		"resetImpersonateServiceAccountDelegates",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetIntegrationConnectorsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIntegrationConnectorsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetIntegrationsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIntegrationsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetKmsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetLoggingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetLookerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetLookerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetLustreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetLustreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetManagedKafkaCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetManagedKafkaCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetMemcacheCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMemcacheCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetMemorystoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMemorystoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetMigrationCenterCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMigrationCenterCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetMlEngineCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMlEngineCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetModelArmorCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetModelArmorCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetModelArmorGlobalCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetModelArmorGlobalCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetMonitoringCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoringCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetNetappCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetappCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetNetworkConnectivityCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConnectivityCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetNetworkConnectivityv1CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConnectivityv1CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetNetworkManagementCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkManagementCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetNetworkSecurityCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkSecurityCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetNetworkServicesCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkServicesCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetNotebooksCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNotebooksCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetObservabilityCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetObservabilityCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetOracleDatabaseCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOracleDatabaseCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetOrgPolicyCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOrgPolicyCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetOsConfigCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOsConfigCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetOsConfigV2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOsConfigV2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetOsLoginCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOsLoginCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetParallelstoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetParallelstoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetParameterManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetParameterManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetParameterManagerRegionalCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetParameterManagerRegionalCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetPollInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetPollInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetPrivatecaCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivatecaCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetPrivilegedAccessManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivilegedAccessManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetPublicCaCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicCaCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetPubsubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetPubsubLiteCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubLiteCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetRecaptchaEnterpriseCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetRecaptchaEnterpriseCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetRedisCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetRedisCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetRequestReason() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestReason",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetResourceManager3CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManager3CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetResourceManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetResourceManagerV3CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerV3CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetRuntimeconfigCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeconfigCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetRuntimeConfigCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeConfigCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSaasRuntimeCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSaasRuntimeCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSecretManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSecretManagerRegionalCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretManagerRegionalCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSecureSourceManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecureSourceManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSecurityCenterCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityCenterCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSecurityCenterManagementCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityCenterManagementCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSecurityCenterV2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityCenterV2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSecuritypostureCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecuritypostureCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSecurityScannerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityScannerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetServiceDirectoryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetServiceManagementCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceManagementCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetServiceNetworkingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceNetworkingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetServiceUsageCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceUsageCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSiteVerificationCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSiteVerificationCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSourceRepoCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceRepoCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSpannerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSpannerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetSqlCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSqlCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetStorageBatchOperationsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageBatchOperationsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetStorageControlCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageControlCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetStorageCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetStorageInsightsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageInsightsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetStorageTransferCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageTransferCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetTagsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetTagsLocationCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsLocationCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetTerraformAttributionLabelAdditionStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetTerraformAttributionLabelAdditionStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetTpuV2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetTpuV2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetTranscoderCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetTranscoderCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetUniverseDomain() {
	_jsii_.InvokeVoid(
		g,
		"resetUniverseDomain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetUserProjectOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetUserProjectOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetVectorSearchCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetVectorSearchCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetVertexAiCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetVertexAiCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetVmwareengineCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetVmwareengineCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetVpcAccessCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcAccessCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetWorkbenchCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkbenchCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetWorkflowsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkflowsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetWorkstationsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkstationsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglebetaProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglebetaProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglebetaProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglebetaProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglebetaProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglebetaProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

