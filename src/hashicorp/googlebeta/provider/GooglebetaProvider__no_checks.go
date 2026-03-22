//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GooglebetaProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GooglebetaProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateGooglebetaProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGooglebetaProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGooglebetaProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGooglebetaProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GooglebetaProvider) validateSetAddTerraformAttributionLabelParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GooglebetaProvider) validateSetBatchingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GooglebetaProvider) validateSetExternalCredentialsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GooglebetaProvider) validateSetUserProjectOverrideParameters(val interface{}) error {
	return nil
}

func validateNewGooglebetaProviderParameters(scope constructs.Construct, id *string, config *GooglebetaProviderConfig) error {
	return nil
}

