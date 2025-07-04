---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/MirroringDeployment.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Network Security"
description: |-
  A deployment represents a zonal mirroring backend ready to accept
  GENEVE-encapsulated replica traffic, e.
---

# google_network_security_mirroring_deployment

A deployment represents a zonal mirroring backend ready to accept
GENEVE-encapsulated replica traffic, e.g. a zonal instance group fronted by
an internal passthrough load balancer. Deployments are always part of a
global deployment group which represents a global mirroring service.


To get more information about MirroringDeployment, see:

* [API documentation](https://cloud.google.com/network-security-integration/docs/reference/rest/v1/projects.locations.mirroringDeployments)
* How-to Guides
    * [Mirroring deployment overview](https://cloud.google.com/network-security-integration/docs/out-of-band/deployments-overview)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=network_security_mirroring_deployment_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Network Security Mirroring Deployment Basic


```hcl
resource "google_compute_network" "network" {
  name                    = "example-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnetwork" {
  name          = "example-subnet"
  region        = "us-central1"
  ip_cidr_range = "10.1.0.0/16"
  network       = google_compute_network.network.name
}

resource "google_compute_region_health_check" "health_check" {
  name     = "example-hc"
  region   = "us-central1"
  http_health_check {
    port = 80
  }
}

resource "google_compute_region_backend_service" "backend_service" {
  name                  = "example-bs"
  region                = "us-central1"
  health_checks         = [google_compute_region_health_check.health_check.id]
  protocol              = "UDP"
  load_balancing_scheme = "INTERNAL"
}

resource "google_compute_forwarding_rule" "forwarding_rule" {
  name                   = "example-fwr"
  region                 = "us-central1"
  network                = google_compute_network.network.name
  subnetwork             = google_compute_subnetwork.subnetwork.name
  backend_service        = google_compute_region_backend_service.backend_service.id
  load_balancing_scheme  = "INTERNAL"
  ports                  = [6081]
  ip_protocol            = "UDP"
  is_mirroring_collector = true
}

resource "google_network_security_mirroring_deployment_group" "deployment_group" {
  mirroring_deployment_group_id = "example-dg"
  location                      = "global"
  network                       = google_compute_network.network.id
}

resource "google_network_security_mirroring_deployment" "default" {
  mirroring_deployment_id    = "example-deployment"
  location                   = "us-central1-a"
  forwarding_rule            = google_compute_forwarding_rule.forwarding_rule.id
  mirroring_deployment_group = google_network_security_mirroring_deployment_group.deployment_group.id
  description                = "some description"
  labels = {
    foo = "bar"
  }
}
```

## Argument Reference

The following arguments are supported:


* `forwarding_rule` -
  (Required)
  The regional forwarding rule that fronts the mirroring collectors, for
  example: `projects/123456789/regions/us-central1/forwardingRules/my-rule`.
  See https://google.aip.dev/124.

* `mirroring_deployment_group` -
  (Required)
  The deployment group that this deployment is a part of, for example:
  `projects/123456789/locations/global/mirroringDeploymentGroups/my-dg`.
  See https://google.aip.dev/124.

* `location` -
  (Required)
  The cloud location of the deployment, e.g. `us-central1-a` or `asia-south1-b`.

* `mirroring_deployment_id` -
  (Required)
  The ID to use for the new deployment, which will become the final
  component of the deployment's resource name.


* `labels` -
  (Optional)
  Labels are key/value pairs that help to organize and filter resources.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `description` -
  (Optional)
  User-provided description of the deployment.
  Used as additional context for the deployment.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/mirroringDeployments/{{mirroring_deployment_id}}`

* `name` -
  The resource name of this deployment, for example:
  `projects/123456789/locations/us-central1-a/mirroringDeployments/my-dep`.
  See https://google.aip.dev/122 for more details.

* `create_time` -
  The timestamp when the resource was created.
  See https://google.aip.dev/148#timestamps.

* `update_time` -
  The timestamp when the resource was most recently updated.
  See https://google.aip.dev/148#timestamps.

* `state` -
  The current state of the deployment.
  See https://google.aip.dev/216.
  Possible values:
  STATE_UNSPECIFIED
  ACTIVE
  CREATING
  DELETING
  OUT_OF_SYNC
  DELETE_FAILED

* `reconciling` -
  The current state of the resource does not match the user's intended state,
  and the system is working to reconcile them. This part of the normal
  operation (e.g. linking a new association to the parent group).
  See https://google.aip.dev/128.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


MirroringDeployment can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/mirroringDeployments/{{mirroring_deployment_id}}`
* `{{project}}/{{location}}/{{mirroring_deployment_id}}`
* `{{location}}/{{mirroring_deployment_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import MirroringDeployment using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/mirroringDeployments/{{mirroring_deployment_id}}"
  to = google_network_security_mirroring_deployment.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), MirroringDeployment can be imported using one of the formats above. For example:

```
$ terraform import google_network_security_mirroring_deployment.default projects/{{project}}/locations/{{location}}/mirroringDeployments/{{mirroring_deployment_id}}
$ terraform import google_network_security_mirroring_deployment.default {{project}}/{{location}}/{{mirroring_deployment_id}}
$ terraform import google_network_security_mirroring_deployment.default {{location}}/{{mirroring_deployment_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
