const inflect = require('inflect')
const path = require('path')
const fs = require('fs')
const providerTpl = require(path.join(__dirname, 'templates', 'provider.go.js'))
const resourceTpl = require(path.join(__dirname, 'templates', 'resource.go.js'))
const swaggerUrl = process.env.SWAGGER_URL || "https://rootly-heroku.s3.amazonaws.com/swagger/v1/swagger.tf.json"
const cleanSwagger = require('./clean-swagger.js')

const excluded = [
    "alert",
    "alert_urgency",
    "alert_group",
    "audits",
    "catalog",
    "catalog_field",
    "catalog_entity",
    "catalog_entity_property",
    "custom_field_option",
    "custom_field",
    "errors",
    "escalation_policy",
    "on_call_shadows",
    "live_call_router",
    "incident_status_page_event",
    "incident_action_item",
    "incident_custom_field_selection",
    "incident_event_functionality",
    "incident_event_service",
    "incident_event",
    "incident_feedback",
    "incident_form_field_selection",
    "incident_post_mortem",
    "incident",
    "ip_ranges",
    "post_mortem_template",
    "pulse",
    "retrospective_configuration",
    "retrospective_process",
    "retrospective_step",
    "secret",
    "shift",
    "user",
    "user_notification_rule",
    "webhooks_delivery",
    "workflow_custom_field_selection",
    "workflow_runs",
]

async function main() {
  const swagger = await getSwagger()
  const resources = getResources(swagger)
  const children = getChildren(swagger, resources)
  const connections = getConnections(swagger, resources)
  const rootResources = resources.filter((name) => !Object.values(children).flat().includes(name))
  
  writeProvider(swagger, rootResources, connections)

  rootResources.forEach((name) => writeResource(swagger, name, children[name]))
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
  return fetch(swaggerUrl).then((res) => res.json()).then(cleanSwagger)
}

function writeResource(swagger, name, childName) {
  const resourcePath = path.join(__dirname, '..', '..', 'providers', 'rootly', `${name}.go`)
  fs.writeFileSync(resourcePath, resourceTpl(swagger, name, childName))
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
      if (get?.operationId === "listEscalationLevelsPolicies" && `list${inflect.pluralize(inflect.camelize(name))}` === "listEscalationLevels") {
        return true
      }
      return (
        get &&
        get.operationId.replace(/ /g, "") ===
          `list${inflect.pluralize(inflect.camelize(name))}`
      );
    })
    .map((url) => swagger.paths[url])[0];
}

main()
