const inflect = require('inflect')
const path = require('path')
const fs = require('fs')
const providerTpl = require(path.join(__dirname, 'templates', 'provider.go.js'))
const resourceTpl = require(path.join(__dirname, 'templates', 'resource.go.js'))
const swaggerUrl = process.env.SWAGGER_URL || "https://rootly-heroku.s3.amazonaws.com/swagger/v1/swagger.tf.json"

const excluded = [
  "alert",
  "audits",
  "custom_field",
  "custom_field_option",
  "dashboard_panel",
  "errors",
  "incident",
  "incident_action_item",
  "incident_custom_field_selection",
  "incident_event",
  "incident_event_functionality",
  "incident_event_service",
  "incident_feedback",
  "incident_form_field_selection",
  "incident_post_mortem",
  "incident_status_page_event",
  "ip_ranges",
  "pulse",
  "user",
  "webhooks_delivery",
  "workflow_runs",
]

async function main() {
  const swagger = await getSwagger()
  const resources = getResources(swagger)
  const children = getChildren(swagger, resources)
  const connections = getConnections(swagger, resources)
  const rootResources = resources.filter((name) => !Object.values(children).flat().includes(name))
  
  writeProvider(swagger, rootResources, connections)

  rootResources.forEach((name) => writeResource(name, children[name]))
}

function getResources(swagger) {
  return Object.keys(swagger.components.schemas)
    .filter((name) => name.match(/_list$/))
    .map((name) => name.replace(/_list$/, ''))
    .filter((name) => !excluded.includes(name))
}

function getConnections(swagger, resources) {
  const connections = {}
  resources.forEach((name) => {
    const field = parentIdField(swagger, name)
    if (field) connections[name] = field
  })
  return connections
}

function getChildren(swagger, resources) {
  const children = {}
  resources.forEach((name) => {
    const field = parentIdField(swagger, name)
    if (field) {
      children[field.replace(/_id$/, '')] ||= []
      children[field.replace(/_id$/, '')].push(name)
    }
  })
  return children
}

function getSwagger() {
  return fetch(swaggerUrl).then((res) => res.json())
}

function writeResource(name, childName) {
  const resourcePath = path.join(__dirname, '..', '..', 'providers', 'rootly', `${name}.go`)
  fs.writeFileSync(resourcePath, resourceTpl(name, childName))
}

function writeProvider(swagger, resources, connections) {
  const providerPath = path.join(__dirname, '..', '..', 'providers', 'rootly', 'rootly_provider.go')
  fs.writeFileSync(providerPath, providerTpl(swagger, resources, connections))
}

function parentIdField(swagger, name) {
  const collectionSchema = collectionPathSchema(swagger, name);
  const parentIdField =
    collectionSchema &&
    collectionSchema.parameters &&
    collectionSchema.parameters[0] &&
    collectionSchema.parameters[0].name;
  return parentIdField
}

function collectionPathSchema(swagger, name) {
  return Object.keys(swagger.paths)
    .filter((url) => {
      const get = swagger.paths[url].get;
      return (
        get &&
        get.operationId.replace(/ /g, "") ===
          `list${inflect.pluralize(inflect.camelize(name))}`
      );
    })
    .map((url) => swagger.paths[url])[0];
}

main()
