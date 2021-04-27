package v5tov6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type v5tov6Server struct {
	v5Server tfprotov5.ProviderServer
}

func (v v5tov6Server) GetProviderSchema(ctx context.Context, req *tfprotov6.GetProviderSchemaRequest) (*tfprotov6.GetProviderSchemaResponse, error) {
	v5Req := v6tov5GetProviderSchemaRequest(req)
	resp, err := v.v5Server.GetProviderSchema(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6GetProviderSchemaResponse(resp), nil
}

func (v v5tov6Server) ValidateProviderConfig(ctx context.Context, req *tfprotov6.ValidateProviderConfigRequest) (*tfprotov6.ValidateProviderConfigResponse, error) {
	v5Req := v6tov5ValidateProviderConfigRequest(req)
	resp, err := v.v5Server.PrepareProviderConfig(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6ValidateProviderConfigResponse(resp), nil
}

func (v v5tov6Server) ConfigureProvider(ctx context.Context, req *tfprotov6.ConfigureProviderRequest) (*tfprotov6.ConfigureProviderResponse, error) {
	v5Req := v6tov5ConfigureProviderRequest(req)
	resp, err := v.v5Server.ConfigureProvider(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6ConfigureProviderResponse(resp), nil
}

func (v v5tov6Server) StopProvider(ctx context.Context, req *tfprotov6.StopProviderRequest) (*tfprotov6.StopProviderResponse, error) {
	v5Req := v6tov5StopProviderRequest(req)
	resp, err := v.v5Server.StopProvider(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6StopProviderResponse(resp), nil
}

func (v v5tov6Server) ValidateResourceConfig(ctx context.Context, req *tfprotov6.ValidateResourceConfigRequest) (*tfprotov6.ValidateResourceConfigResponse, error) {
	v5Req := v6tov5ValidateResourceConfigRequest(req)
	resp, err := v.v5Server.ValidateResourceTypeConfig(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6ValidateResourceConfigResponse(resp), nil
}

func (v v5tov6Server) UpgradeResourceState(ctx context.Context, req *tfprotov6.UpgradeResourceStateRequest) (*tfprotov6.UpgradeResourceStateResponse, error) {
	v5Req := v6tov5UpgradeResourceStateRequest(req)
	resp, err := v.v5Server.UpgradeResourceState(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6UpgradeResourceStateResponse(resp), nil
}

func (v v5tov6Server) ReadResource(ctx context.Context, req *tfprotov6.ReadResourceRequest) (*tfprotov6.ReadResourceResponse, error) {
	v5Req := v6tov5ReadResourceRequest(req)
	resp, err := v.v5Server.ReadResource(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6ReadResourceResponse(resp), nil
}

func (v v5tov6Server) PlanResourceChange(ctx context.Context, req *tfprotov6.PlanResourceChangeRequest) (*tfprotov6.PlanResourceChangeResponse, error) {
	v5Req := v6tov5PlanResourceChangeRequest(req)
	resp, err := v.v5Server.PlanResourceChange(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6PlanResourceChangeResponse(resp), nil
}

func (v v5tov6Server) ApplyResourceChange(ctx context.Context, req *tfprotov6.ApplyResourceChangeRequest) (*tfprotov6.ApplyResourceChangeResponse, error) {
	v5Req := v6tov5ApplyResourceChangeRequest(req)
	resp, err := v.v5Server.ApplyResourceChange(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6ApplyResourceChangeResponse(resp), nil
}

func (v v5tov6Server) ImportResourceState(ctx context.Context, req *tfprotov6.ImportResourceStateRequest) (*tfprotov6.ImportResourceStateResponse, error) {
	v5Req := v6tov5ImportResourceStateRequest(req)
	resp, err := v.v5Server.ImportResourceState(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6ImportResourceStateResponse(resp), nil
}

func (v v5tov6Server) ValidateDataResourceConfig(ctx context.Context, req *tfprotov6.ValidateDataResourceConfigRequest) (*tfprotov6.ValidateDataResourceConfigResponse, error) {
	v5Req := v6tov5ValidateDataResourceConfigRequest(req)
	resp, err := v.v5Server.ValidateDataSourceConfig(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6ValidateDataResourceConfigResponse(resp), nil
}

func (v v5tov6Server) ReadDataSource(ctx context.Context, req *tfprotov6.ReadDataSourceRequest) (*tfprotov6.ReadDataSourceResponse, error) {
	v5Req := v6tov5ReadDataSourceRequest(req)
	resp, err := v.v5Server.ReadDataSource(ctx, v5Req)
	if err != nil {
		return nil, err
	}
	return v5tov6ReadDataSourceResponse(resp), nil
}
