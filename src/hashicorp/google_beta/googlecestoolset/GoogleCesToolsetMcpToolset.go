package googlecestoolset


type GoogleCesToolsetMcpToolset struct {
	// The address of the MCP server, for example, "https://example.com/mcp/". If the server is built with the MCP SDK, the url should be suffixed with "/mcp/". Only Streamable HTTP transport based servers are supported. See https://modelcontextprotocol.io/specification/2025-03-26/basic/transports#streamable-http for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_ces_toolset#server_address GoogleCesToolset#server_address}
	ServerAddress *string `field:"required" json:"serverAddress" yaml:"serverAddress"`
	// api_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_ces_toolset#api_authentication GoogleCesToolset#api_authentication}
	ApiAuthentication *GoogleCesToolsetMcpToolsetApiAuthentication `field:"optional" json:"apiAuthentication" yaml:"apiAuthentication"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_ces_toolset#service_directory_config GoogleCesToolset#service_directory_config}
	ServiceDirectoryConfig *GoogleCesToolsetMcpToolsetServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.24.0/docs/resources/google_ces_toolset#tls_config GoogleCesToolset#tls_config}
	TlsConfig *GoogleCesToolsetMcpToolsetTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

