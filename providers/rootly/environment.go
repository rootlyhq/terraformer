package rootly

import (
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/schema"
)

type EnvironmentGenerator struct {
	RootlyService
}

func (g* EnvironmentGenerator) InitResources() error {
	page_size := 50
	page_num := 1

	client, err := g.RootlyClient()
	if err != nil {
		return err
	}

	var environments []interface{}

	for {
		results, err := func(page_size, page_num int) ([]interface{}, error) {
			params := new(rootlygo.ListEnvironmentsParams)
			params.PageSize = &page_size
			params.PageNumber = &page_num
			return client.ListEnvironments(params)
		}(page_size, page_num)

		if err != nil {
			return err
		}

		if len(results) == 0 {
			break
		}

		environments = append(environments, results...)
		page_num += 1
	}

	g.Resources = g.createResources(environments)

	return nil
}

func (g *EnvironmentGenerator) createResources(environments []interface{}) []terraformutils.Resource {
	var resources []terraformutils.Resource

	for _, environment := range environments {
		x, _ := environment.(*client.Environment)
		resources = append(resources, terraformutils.NewResource(
			x.ID,
			fmt.Sprintf("%s-%s", x.ID, x.Slug),
			"rootly_environment",
			g.ProviderName,
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}
