package managedkafkaconnectcluster


type ManagedKafkaConnectClusterGcpConfigAccessConfig struct {
	// network_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.24.0/docs/resources/managed_kafka_connect_cluster#network_configs ManagedKafkaConnectCluster#network_configs}
	NetworkConfigs interface{} `field:"required" json:"networkConfigs" yaml:"networkConfigs"`
}

