const inflect = require("inflect")

module.exports = (swagger, resources, connections) => `// Generated by generators/rootly/rootly.js
package rootly

import (
	"errors"
	"os"
	"github.com/zclconf/go-cty/cty"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type RootlyProvider struct { //nolint
	terraformutils.Provider
	apiKey        string
	apiUrl        string
}

func (p *RootlyProvider) Init(args []string) error {
	if apiUrl := os.Getenv("ROOTLY_API_URL"); apiUrl != "" {
		p.apiUrl = os.Getenv("ROOTLY_API_URL")
	}
	if args[0] != "" {
		p.apiUrl = args[0]
	}
	if p.apiUrl == "" {
		p.apiUrl = "https://api.rootly.com"
	}

	if apiKey := os.Getenv("ROOTLY_API_TOKEN"); apiKey != "" {
		p.apiKey = os.Getenv("ROOTLY_API_TOKEN")
	}
	if args[0] != "" {
		p.apiKey = args[0]
	}
	if p.apiKey == "" {
		return errors.New("required API key missing")
	}

	return nil
}

func (p *RootlyProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	// SetArgs are used for fetching details within other files in the terraformer code.
	p.Service.SetArgs(map[string]interface{}{
		"api_token": p.apiKey,
		"api_host":  p.apiUrl,
	})
	return nil
}

func (p *RootlyProvider) GetName() string {
	return "rootly"
}

func (p *RootlyProvider) GetSource() string {
	return "rootlyhq/rootly"
}

func (p *RootlyProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"api_token": cty.StringVal(p.apiKey),
		"api_host":  cty.StringVal(p.apiUrl),
	})
}

func (p RootlyProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (p *RootlyProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		${Object.keys(connections).map((parent) => {
			return resourceType(swagger, parent).map((childResource) => {
				const parentResources = resourceType(swagger, connections[parent].replace(/_id$/, ''))
				return`${childResource}: {
					${parentResources.map((parentResource) => {
						return `
							${parentResource}: {
								"${connections[parent]}", "id",
							},`
					}).join('\n')}
				},`
			}).join('\n')
		}).join('\n    ')}
	}
}

func (p *RootlyProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		${resources.map((name) => {
			return `"${name}": &${inflect.camelize(name)}Generator{},`  
		}).join('\n		')}
	}
}`

const resourceType = (swagger, name) => {
	switch (name) {
		case "workflow":
			return [`"rootly_workflow_incident"`, `"rootly_workflow_action_item"`, `"rootly_workflow_post_mortem"`, `"rootly_workflow_pulse"`, `"rootly_workflow_alert"`, `"rootly_workflow_simple"`]
		case "workflow_task":
			return Object.keys(swagger.components.schemas).filter((name) => name.match(/_task_params$/)).map((name) => `"rootly_workflow_task_${name.replace(/_task_params/, '')}"`)
		default:
			return [`"rootly_${name}"`]
	}
}