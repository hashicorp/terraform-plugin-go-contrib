package v5tov6

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func Convert(server tfprotov5.ProviderServer) tfprotov6.ProviderServer {
	return v5tov6Server{
		v5Server: server,
	}
}

func v6tov5GetProviderSchemaRequest(in *tfprotov6.GetProviderSchemaRequest) *tfprotov5.GetProviderSchemaRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.GetProviderSchemaRequest{}
}

func v5tov6StringKind(in tfprotov5.StringKind) tfprotov6.StringKind {
	return tfprotov6.StringKind(in)
}

func v5tov6SchemaAttribute(in *tfprotov5.SchemaAttribute) *tfprotov6.SchemaAttribute {
	if in == nil {
		return nil
	}
	return &tfprotov6.SchemaAttribute{
		Name:            in.Name,
		Type:            in.Type,
		Description:     in.Description,
		Required:        in.Required,
		Optional:        in.Optional,
		Computed:        in.Computed,
		Sensitive:       in.Sensitive,
		DescriptionKind: v5tov6StringKind(in.DescriptionKind),
		Deprecated:      in.Deprecated,
	}
}

func v5tov6SchemaNestedBlock(in *tfprotov5.SchemaNestedBlock) *tfprotov6.SchemaNestedBlock {
	if in == nil {
		return nil
	}
	return &tfprotov6.SchemaNestedBlock{
		TypeName: in.TypeName,
		Block:    v5tov6SchemaBlock(in.Block),
		Nesting:  tfprotov6.SchemaNestedBlockNestingMode(in.Nesting),
		MinItems: in.MinItems,
		MaxItems: in.MaxItems,
	}
}

func v5tov6SchemaBlock(in *tfprotov5.SchemaBlock) *tfprotov6.SchemaBlock {
	if in == nil {
		return nil
	}
	var attrs []*tfprotov6.SchemaAttribute
	if in.Attributes != nil {
		attrs = make([]*tfprotov6.SchemaAttribute, 0, len(in.Attributes))
		for _, attr := range in.Attributes {
			attrs = append(attrs, v5tov6SchemaAttribute(attr))
		}
	}
	var nestedBlocks []*tfprotov6.SchemaNestedBlock
	if in.BlockTypes != nil {
		nestedBlocks = make([]*tfprotov6.SchemaNestedBlock, 0, len(in.BlockTypes))
		for _, block := range in.BlockTypes {
			nestedBlocks = append(nestedBlocks, v5tov6SchemaNestedBlock(block))
		}
	}
	return &tfprotov6.SchemaBlock{
		Version:         in.Version,
		Attributes:      attrs,
		BlockTypes:      nestedBlocks,
		Description:     in.Description,
		DescriptionKind: v5tov6StringKind(in.DescriptionKind),
		Deprecated:      in.Deprecated,
	}
}

func v5tov6Schema(in *tfprotov5.Schema) *tfprotov6.Schema {
	if in == nil {
		return nil
	}
	return &tfprotov6.Schema{
		Version: in.Version,
		Block:   v5tov6SchemaBlock(in.Block),
	}
}

func v5tov6Diagnostics(in []*tfprotov5.Diagnostic) []*tfprotov6.Diagnostic {
	if in == nil {
		return nil
	}
	diags := make([]*tfprotov6.Diagnostic, 0, len(in))
	for _, diag := range in {
		if diag == nil {
			diags = append(diags, nil)
			continue
		}
		diags = append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverity(diag.Severity),
			Summary:   diag.Summary,
			Detail:    diag.Detail,
			Attribute: diag.Attribute,
		})
	}
	return diags
}

func v5tov6GetProviderSchemaResponse(in *tfprotov5.GetProviderSchemaResponse) *tfprotov6.GetProviderSchemaResponse {
	if in == nil {
		return nil
	}
	resourceSchemas := make(map[string]*tfprotov6.Schema, len(in.ResourceSchemas))
	for k, v := range in.ResourceSchemas {
		resourceSchemas[k] = v5tov6Schema(v)
	}
	dataSourceSchemas := make(map[string]*tfprotov6.Schema, len(in.DataSourceSchemas))
	for k, v := range in.DataSourceSchemas {
		dataSourceSchemas[k] = v5tov6Schema(v)
	}
	return &tfprotov6.GetProviderSchemaResponse{
		Provider:          v5tov6Schema(in.Provider),
		ProviderMeta:      v5tov6Schema(in.ProviderMeta),
		ResourceSchemas:   resourceSchemas,
		DataSourceSchemas: dataSourceSchemas,
		Diagnostics:       v5tov6Diagnostics(in.Diagnostics),
	}
}

func v6tov5DynamicValue(in *tfprotov6.DynamicValue) *tfprotov5.DynamicValue {
	if in == nil {
		return nil
	}
	return &tfprotov5.DynamicValue{
		MsgPack: in.MsgPack,
		JSON:    in.JSON,
	}
}

func v6tov5ValidateProviderConfigRequest(in *tfprotov6.ValidateProviderConfigRequest) *tfprotov5.PrepareProviderConfigRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.PrepareProviderConfigRequest{
		Config: v6tov5DynamicValue(in.Config),
	}
}

func v5tov6DynamicValue(in *tfprotov5.DynamicValue) *tfprotov6.DynamicValue {
	if in == nil {
		return nil
	}
	return &tfprotov6.DynamicValue{
		MsgPack: in.MsgPack,
		JSON:    in.JSON,
	}
}

func v5tov6ValidateProviderConfigResponse(in *tfprotov5.PrepareProviderConfigResponse) *tfprotov6.ValidateProviderConfigResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.ValidateProviderConfigResponse{
		PreparedConfig: v5tov6DynamicValue(in.PreparedConfig),
		Diagnostics:    v5tov6Diagnostics(in.Diagnostics),
	}
}

func v6tov5ConfigureProviderRequest(in *tfprotov6.ConfigureProviderRequest) *tfprotov5.ConfigureProviderRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.ConfigureProviderRequest{
		TerraformVersion: in.TerraformVersion,
		Config:           v6tov5DynamicValue(in.Config),
	}
}

func v5tov6ConfigureProviderResponse(in *tfprotov5.ConfigureProviderResponse) *tfprotov6.ConfigureProviderResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.ConfigureProviderResponse{
		Diagnostics: v5tov6Diagnostics(in.Diagnostics),
	}
}

func v6tov5StopProviderRequest(in *tfprotov6.StopProviderRequest) *tfprotov5.StopProviderRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.StopProviderRequest{}
}

func v5tov6StopProviderResponse(in *tfprotov5.StopProviderResponse) *tfprotov6.StopProviderResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.StopProviderResponse{
		Error: in.Error,
	}
}

func v6tov5ValidateResourceConfigRequest(in *tfprotov6.ValidateResourceConfigRequest) *tfprotov5.ValidateResourceTypeConfigRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.ValidateResourceTypeConfigRequest{
		TypeName: in.TypeName,
		Config:   v6tov5DynamicValue(in.Config),
	}
}

func v5tov6ValidateResourceConfigResponse(in *tfprotov5.ValidateResourceTypeConfigResponse) *tfprotov6.ValidateResourceConfigResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.ValidateResourceConfigResponse{
		Diagnostics: v5tov6Diagnostics(in.Diagnostics),
	}
}

func v6tov5RawState(in *tfprotov6.RawState) *tfprotov5.RawState {
	if in == nil {
		return nil
	}
	return &tfprotov5.RawState{
		JSON:    in.JSON,
		Flatmap: in.Flatmap,
	}
}

func v6tov5UpgradeResourceStateRequest(in *tfprotov6.UpgradeResourceStateRequest) *tfprotov5.UpgradeResourceStateRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.UpgradeResourceStateRequest{
		TypeName: in.TypeName,
		Version:  in.Version,
		RawState: v6tov5RawState(in.RawState),
	}
}

func v5tov6UpgradeResourceStateResponse(in *tfprotov5.UpgradeResourceStateResponse) *tfprotov6.UpgradeResourceStateResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.UpgradeResourceStateResponse{
		UpgradedState: v5tov6DynamicValue(in.UpgradedState),
		Diagnostics:   v5tov6Diagnostics(in.Diagnostics),
	}
}

func v6tov5ReadResourceRequest(in *tfprotov6.ReadResourceRequest) *tfprotov5.ReadResourceRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.ReadResourceRequest{
		TypeName:     in.TypeName,
		CurrentState: v6tov5DynamicValue(in.CurrentState),
		Private:      in.Private,
		ProviderMeta: v6tov5DynamicValue(in.ProviderMeta),
	}
}

func v5tov6ReadResourceResponse(in *tfprotov5.ReadResourceResponse) *tfprotov6.ReadResourceResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.ReadResourceResponse{
		NewState:    v5tov6DynamicValue(in.NewState),
		Diagnostics: v5tov6Diagnostics(in.Diagnostics),
		Private:     in.Private,
	}
}

func v6tov5PlanResourceChangeRequest(in *tfprotov6.PlanResourceChangeRequest) *tfprotov5.PlanResourceChangeRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.PlanResourceChangeRequest{
		TypeName:         in.TypeName,
		PriorState:       v6tov5DynamicValue(in.PriorState),
		ProposedNewState: v6tov5DynamicValue(in.ProposedNewState),
		Config:           v6tov5DynamicValue(in.Config),
		PriorPrivate:     in.PriorPrivate,
		ProviderMeta:     v6tov5DynamicValue(in.ProviderMeta),
	}
}

func v5tov6PlanResourceChangeResponse(in *tfprotov5.PlanResourceChangeResponse) *tfprotov6.PlanResourceChangeResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.PlanResourceChangeResponse{
		PlannedState:                v5tov6DynamicValue(in.PlannedState),
		RequiresReplace:             in.RequiresReplace,
		PlannedPrivate:              in.PlannedPrivate,
		Diagnostics:                 v5tov6Diagnostics(in.Diagnostics),
		UnsafeToUseLegacyTypeSystem: in.UnsafeToUseLegacyTypeSystem,
	}
}

func v6tov5ApplyResourceChangeRequest(in *tfprotov6.ApplyResourceChangeRequest) *tfprotov5.ApplyResourceChangeRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.ApplyResourceChangeRequest{
		TypeName:       in.TypeName,
		PriorState:     v6tov5DynamicValue(in.PriorState),
		PlannedState:   v6tov5DynamicValue(in.PlannedState),
		Config:         v6tov5DynamicValue(in.Config),
		PlannedPrivate: in.PlannedPrivate,
		ProviderMeta:   v6tov5DynamicValue(in.ProviderMeta),
	}
}

func v5tov6ApplyResourceChangeResponse(in *tfprotov5.ApplyResourceChangeResponse) *tfprotov6.ApplyResourceChangeResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.ApplyResourceChangeResponse{
		NewState:                    v5tov6DynamicValue(in.NewState),
		Private:                     in.Private,
		Diagnostics:                 v5tov6Diagnostics(in.Diagnostics),
		UnsafeToUseLegacyTypeSystem: in.UnsafeToUseLegacyTypeSystem,
	}
}

func v6tov5ImportResourceStateRequest(in *tfprotov6.ImportResourceStateRequest) *tfprotov5.ImportResourceStateRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.ImportResourceStateRequest{
		TypeName: in.TypeName,
		ID:       in.ID,
	}
}

func v5tov6ImportedResources(in []*tfprotov5.ImportedResource) []*tfprotov6.ImportedResource {
	if in == nil {
		return nil
	}
	res := make([]*tfprotov6.ImportedResource, 0, len(in))
	for _, imp := range in {
		if imp == nil {
			res = append(res, nil)
			continue
		}
		res = append(res, &tfprotov6.ImportedResource{
			TypeName: imp.TypeName,
			State:    v5tov6DynamicValue(imp.State),
			Private:  imp.Private,
		})
	}
	return res
}

func v5tov6ImportResourceStateResponse(in *tfprotov5.ImportResourceStateResponse) *tfprotov6.ImportResourceStateResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.ImportResourceStateResponse{
		ImportedResources: v5tov6ImportedResources(in.ImportedResources),
		Diagnostics:       v5tov6Diagnostics(in.Diagnostics),
	}
}

func v6tov5ValidateDataResourceConfigRequest(in *tfprotov6.ValidateDataResourceConfigRequest) *tfprotov5.ValidateDataSourceConfigRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.ValidateDataSourceConfigRequest{
		TypeName: in.TypeName,
		Config:   v6tov5DynamicValue(in.Config),
	}
}

func v5tov6ValidateDataResourceConfigResponse(in *tfprotov5.ValidateDataSourceConfigResponse) *tfprotov6.ValidateDataResourceConfigResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.ValidateDataResourceConfigResponse{
		Diagnostics: v5tov6Diagnostics(in.Diagnostics),
	}
}

func v6tov5ReadDataSourceRequest(in *tfprotov6.ReadDataSourceRequest) *tfprotov5.ReadDataSourceRequest {
	if in == nil {
		return nil
	}
	return &tfprotov5.ReadDataSourceRequest{
		TypeName:     in.TypeName,
		Config:       v6tov5DynamicValue(in.Config),
		ProviderMeta: v6tov5DynamicValue(in.ProviderMeta),
	}
}

func v5tov6ReadDataSourceResponse(in *tfprotov5.ReadDataSourceResponse) *tfprotov6.ReadDataSourceResponse {
	if in == nil {
		return nil
	}
	return &tfprotov6.ReadDataSourceResponse{
		State:       v5tov6DynamicValue(in.State),
		Diagnostics: v5tov6Diagnostics(in.Diagnostics),
	}
}
