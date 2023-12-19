package rootly

import (
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/rootlyhq/terraform-provider-rootly/client"
)

type RootlyService struct { //nolint
	terraformutils.Service
}

func (s *RootlyService) RootlyClient() (*client.Client, error) {
	host  := s.GetArgs()["api_host"].(string)
	token := s.GetArgs()["api_token"].(string)
	return client.NewClient(host, token, RootlyUserAgent("2023-12-15"))
}

func RootlyUserAgent(version string) string {
	return fmt.Sprintf("Terraformer/%s (+https://github.com/rootlyhq/terraformer)", version)
}
