
// Generated by generators/rootly/rootly.js
 
package rootly

import (
	
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/rootlyhq/terraform-provider-rootly/v2/client"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/v2/schema"
)

type RetrospectiveProcessGenerator struct {
	RootlyService
}

func (g* RetrospectiveProcessGenerator) InitResources() error {
	page_size := 50
	page_num := 1

	client, err := g.RootlyClient()
	if err != nil {
		return err
	}

	for {
		resources, err := func(page_size, page_num int) ([]interface{}, error) {
			params := new(rootlygo.ListRetrospectiveProcessesParams)
			params.PageSize = &page_size
			params.PageNumber = &page_num
			return client.ListRetrospectiveProcesses(params)
		}(page_size, page_num)

		if err != nil {
			return err
		}

		if len(resources) == 0 {
			break
		}

  	for _, resource := range resources {
      tf_resource := g.createRetrospectiveProcessResource(resource)
      g.Resources = append(g.Resources, tf_resource)
      child_retrospective_step, err := g.createRetrospectiveStepResources(tf_resource.InstanceState.ID)
      if err != nil {
        return err
      }
      g.Resources = append(g.Resources, child_retrospective_step...)
  	}

		page_num += 1
	}

	return nil
}

func (g *RetrospectiveProcessGenerator) createRetrospectiveProcessResource(provider_resource interface{}) terraformutils.Resource {
	x, _ := provider_resource.(*client.RetrospectiveProcess)
	return terraformutils.NewSimpleResource(
		x.ID,
		x.ID,
		"rootly_retrospective_process",
		g.ProviderName,
		[]string{},
	)
}


func (g *RetrospectiveProcessGenerator) PostConvertHook() error {
  for _, resource := range g.Resources {
		
    if resource.InstanceInfo.Type != "rootly_retrospective_process" {
      continue
    }
		
    
        for i, retrospective_step := range g.Resources {
          if retrospective_step.InstanceInfo.Type != "rootly_retrospective_step" {
            continue
          }
          if retrospective_step.InstanceState.Attributes["retrospective_process_id"] == resource.InstanceState.ID {
            g.Resources[i].Item["retrospective_process_id"] = "${" + resource.InstanceInfo.Type + "." + resource.ResourceName + ".id}"
          }
        }
      
  }

  return nil
}


func (g *RetrospectiveProcessGenerator) createRetrospectiveStepResources(parent_id string) ([]terraformutils.Resource, error) {
	page_size := 50
	page_num := 1

	client, err := g.RootlyClient()
	if err != nil {
		return nil, err
	}

  var tf_resources []terraformutils.Resource

	for {
		resources, err := func(page_size, page_num int) ([]interface{}, error) {
			params := new(rootlygo.ListRetrospectiveStepsParams)
			params.PageSize = &page_size
			params.PageNumber = &page_num
			return client.ListRetrospectiveSteps(parent_id, params)
		}(page_size, page_num)

		if err != nil {
			return nil, err
		}

		if len(resources) == 0 {
			break
		}

  	for _, resource := range resources {
      tf_resources = append(tf_resources, g.createRetrospectiveStepResource(resource))
  	}

		page_num += 1
	}

	return tf_resources, nil
}

func (g *RetrospectiveProcessGenerator) createRetrospectiveStepResource(provider_resource interface{}) terraformutils.Resource {
	x, _ := provider_resource.(*client.RetrospectiveStep)
	return terraformutils.NewSimpleResource(
		x.ID,
		x.Slug,
		"rootly_retrospective_step",
		g.ProviderName,
		[]string{},
	)
}

