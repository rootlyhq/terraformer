
// Generated by generators/rootly/rootly.js
 
package rootly

import (
	
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/rootlyhq/terraform-provider-rootly/v2/client"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/v2/schema"
)

type ScheduleGenerator struct {
	RootlyService
}

func (g* ScheduleGenerator) InitResources() error {
	page_size := 50
	page_num := 1

	client, err := g.RootlyClient()
	if err != nil {
		return err
	}

	for {
		resources, err := func(page_size, page_num int) ([]interface{}, error) {
			params := new(rootlygo.ListSchedulesParams)
			params.PageSize = &page_size
			params.PageNumber = &page_num
			return client.ListSchedules(params)
		}(page_size, page_num)

		if err != nil {
			return err
		}

		if len(resources) == 0 {
			break
		}

  	for _, resource := range resources {
      tf_resource := g.createScheduleResource(resource)
      g.Resources = append(g.Resources, tf_resource)
      child_override_shift, err := g.createOverrideShiftResources(tf_resource.InstanceState.ID)
      if err != nil {
        return err
      }
      g.Resources = append(g.Resources, child_override_shift...)
child_schedule_rotation, err := g.createScheduleRotationResources(tf_resource.InstanceState.ID)
      if err != nil {
        return err
      }
      g.Resources = append(g.Resources, child_schedule_rotation...)
  	}

		page_num += 1
	}

	return nil
}

func (g *ScheduleGenerator) createScheduleResource(provider_resource interface{}) terraformutils.Resource {
	x, _ := provider_resource.(*client.Schedule)
	return terraformutils.NewSimpleResource(
		x.ID,
		x.ID,
		"rootly_schedule",
		g.ProviderName,
		[]string{},
	)
}


func (g *ScheduleGenerator) PostConvertHook() error {
  for _, resource := range g.Resources {
		
    if resource.InstanceInfo.Type != "rootly_schedule" {
      continue
    }
		
    
        for i, override_shift := range g.Resources {
          if override_shift.InstanceInfo.Type != "rootly_override_shift" {
            continue
          }
          if override_shift.InstanceState.Attributes["schedule_id"] == resource.InstanceState.ID {
            g.Resources[i].Item["schedule_id"] = "${" + resource.InstanceInfo.Type + "." + resource.ResourceName + ".id}"
          }
        }
      

        for i, schedule_rotation := range g.Resources {
          if schedule_rotation.InstanceInfo.Type != "rootly_schedule_rotation" {
            continue
          }
          if schedule_rotation.InstanceState.Attributes["schedule_id"] == resource.InstanceState.ID {
            g.Resources[i].Item["schedule_id"] = "${" + resource.InstanceInfo.Type + "." + resource.ResourceName + ".id}"
          }
        }
      
  }

  return nil
}


func (g *ScheduleGenerator) createOverrideShiftResources(parent_id string) ([]terraformutils.Resource, error) {
	page_size := 50
	page_num := 1

	client, err := g.RootlyClient()
	if err != nil {
		return nil, err
	}

  var tf_resources []terraformutils.Resource

	for {
		resources, err := func(page_size, page_num int) ([]interface{}, error) {
			params := new(rootlygo.ListOverrideShiftsParams)
			params.PageSize = &page_size
			params.PageNumber = &page_num
			return client.ListOverrideShifts(parent_id, params)
		}(page_size, page_num)

		if err != nil {
			return nil, err
		}

		if len(resources) == 0 {
			break
		}

  	for _, resource := range resources {
      tf_resources = append(tf_resources, g.createOverrideShiftResource(resource))
  	}

		page_num += 1
	}

	return tf_resources, nil
}

func (g *ScheduleGenerator) createOverrideShiftResource(provider_resource interface{}) terraformutils.Resource {
	x, _ := provider_resource.(*client.OverrideShift)
	return terraformutils.NewSimpleResource(
		x.ID,
		x.ID,
		"rootly_override_shift",
		g.ProviderName,
		[]string{},
	)
}


func (g *ScheduleGenerator) createScheduleRotationResources(parent_id string) ([]terraformutils.Resource, error) {
	page_size := 50
	page_num := 1

	client, err := g.RootlyClient()
	if err != nil {
		return nil, err
	}

  var tf_resources []terraformutils.Resource

	for {
		resources, err := func(page_size, page_num int) ([]interface{}, error) {
			params := new(rootlygo.ListScheduleRotationsParams)
			params.PageSize = &page_size
			params.PageNumber = &page_num
			return client.ListScheduleRotations(parent_id, params)
		}(page_size, page_num)

		if err != nil {
			return nil, err
		}

		if len(resources) == 0 {
			break
		}

  	for _, resource := range resources {
      tf_resources = append(tf_resources, g.createScheduleRotationResource(resource))
  	}

		page_num += 1
	}

	return tf_resources, nil
}

func (g *ScheduleGenerator) createScheduleRotationResource(provider_resource interface{}) terraformutils.Resource {
	x, _ := provider_resource.(*client.ScheduleRotation)
	return terraformutils.NewSimpleResource(
		x.ID,
		x.ID,
		"rootly_schedule_rotation",
		g.ProviderName,
		[]string{},
	)
}
