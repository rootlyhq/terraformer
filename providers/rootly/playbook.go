
// Generated by generators/rootly/rootly.js
 
package rootly

import (
	
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/rootlyhq/terraform-provider-rootly/v2/client"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/v2/schema"
)

type PlaybookGenerator struct {
	RootlyService
}

func (g* PlaybookGenerator) InitResources() error {
	page_size := 50
	page_num := 1

	client, err := g.RootlyClient()
	if err != nil {
		return err
	}

	for {
		resources, err := func(page_size, page_num int) ([]interface{}, error) {
			params := new(rootlygo.ListPlaybooksParams)
			params.PageSize = &page_size
			params.PageNumber = &page_num
			return client.ListPlaybooks(params)
		}(page_size, page_num)

		if err != nil {
			return err
		}

		if len(resources) == 0 {
			break
		}

  	for _, resource := range resources {
      tf_resource := g.createPlaybookResource(resource)
      g.Resources = append(g.Resources, tf_resource)
      child_playbook_task, err := g.createPlaybookTaskResources(tf_resource.InstanceState.ID)
      if err != nil {
        return err
      }
      g.Resources = append(g.Resources, child_playbook_task...)
  	}

		page_num += 1
	}

	return nil
}

func (g *PlaybookGenerator) createPlaybookResource(provider_resource interface{}) terraformutils.Resource {
	x, _ := provider_resource.(*client.Playbook)
	return terraformutils.NewSimpleResource(
		x.ID,
		x.ID,
		"rootly_playbook",
		g.ProviderName,
		[]string{},
	)
}


func (g *PlaybookGenerator) PostConvertHook() error {
  for _, resource := range g.Resources {
		
    if resource.InstanceInfo.Type != "rootly_playbook" {
      continue
    }
		
    
        for i, playbook_task := range g.Resources {
          if playbook_task.InstanceInfo.Type != "rootly_playbook_task" {
            continue
          }
          if playbook_task.InstanceState.Attributes["playbook_id"] == resource.InstanceState.ID {
            g.Resources[i].Item["playbook_id"] = "${" + resource.InstanceInfo.Type + "." + resource.ResourceName + ".id}"
          }
        }
      
  }

  return nil
}


func (g *PlaybookGenerator) createPlaybookTaskResources(parent_id string) ([]terraformutils.Resource, error) {
	page_size := 50
	page_num := 1

	client, err := g.RootlyClient()
	if err != nil {
		return nil, err
	}

  var tf_resources []terraformutils.Resource

	for {
		resources, err := func(page_size, page_num int) ([]interface{}, error) {
			params := new(rootlygo.ListPlaybookTasksParams)
			params.PageSize = &page_size
			params.PageNumber = &page_num
			return client.ListPlaybookTasks(parent_id, params)
		}(page_size, page_num)

		if err != nil {
			return nil, err
		}

		if len(resources) == 0 {
			break
		}

  	for _, resource := range resources {
      tf_resources = append(tf_resources, g.createPlaybookTaskResource(resource))
  	}

		page_num += 1
	}

	return tf_resources, nil
}

func (g *PlaybookGenerator) createPlaybookTaskResource(provider_resource interface{}) terraformutils.Resource {
	x, _ := provider_resource.(*client.PlaybookTask)
	return terraformutils.NewSimpleResource(
		x.ID,
		x.ID,
		"rootly_playbook_task",
		g.ProviderName,
		[]string{},
	)
}

