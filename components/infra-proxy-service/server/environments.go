package server

import (
	"context"
	"encoding/json"
	"sort"

	chef "github.com/go-chef/chef"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chef/automate/api/interservice/infra_proxy/request"
	"github.com/chef/automate/api/interservice/infra_proxy/response"
)

// GetEnvironments get environments list
func (s *Server) GetEnvironments(ctx context.Context, req *request.Environments) (*response.Environments, error) {

	c, err := s.createClient(ctx, req.OrgId, req.ServerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid org ID: %s", err.Error())
	}

	environments, err := c.client.Environments.List()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &response.Environments{
		Environments: fromAPIToListEnvironments(*environments),
	}, nil
}

// GetEnvironment gets the environment details
func (s *Server) GetEnvironment(ctx context.Context, req *request.Environment) (*response.Environment, error) {
	c, err := s.createClient(ctx, req.OrgId, req.ServerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid org ID: %s", err.Error())
	}

	en, err := c.client.Environments.Get(req.Name)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	defaultAttributes, err := json.Marshal(en.DefaultAttributes)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	overrideAttributes, err := json.Marshal(en.OverrideAttributes)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &response.Environment{
		Name:               en.Name,
		ChefType:           en.ChefType,
		Description:        en.Description,
		CookbookVersions:   fromAPIToListEnvCookbookVersions(en.CookbookVersions),
		JsonClass:          en.JsonClass,
		DefaultAttributes:  string(defaultAttributes),
		OverrideAttributes: string(overrideAttributes),
	}, nil

}

// DeleteEnvironment deletes the environment
func (s *Server) DeleteEnvironment(ctx context.Context, req *request.Environment) (*response.Environment, error) {
	c, err := s.createClient(ctx, req.OrgId, req.ServerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid org ID: %s", err.Error())
	}

	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "must supply environment name")
	}

	environment, err := c.client.Environments.Delete(req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &response.Environment{
		Name: environment.Name,
	}, nil

}

// fromAPIToListEnvironments a response.Environments from a struct of Environments
func fromAPIToListEnvironments(al chef.EnvironmentResult) []*response.EnvironmentListItem {
	cl := make([]*response.EnvironmentListItem, len(al))

	index := 0
	for c := range al {
		cl[index] = &response.EnvironmentListItem{
			Name: c,
		}
		index++
	}

	sort.Slice(cl, func(i, j int) bool {
		return cl[i].Name < cl[j].Name
	})

	return cl
}

func fromAPIToListEnvCookbookVersions(cookbooks map[string]string) []string {
	cl := make([]string, len(cookbooks))

	for i, c := range cl {
		cl[i] = c
	}
	return cl
}
