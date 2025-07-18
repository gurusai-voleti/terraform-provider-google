---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/logging/OrganizationSettings.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud (Stackdriver) Logging"
description: |-
  Default resource settings control whether CMEK is required for new log buckets.
---

# google_logging_organization_settings

Default resource settings control whether CMEK is required for new log buckets. These settings also determine the storage location for the _Default and _Required log buckets, and whether the _Default sink is enabled or disabled.


To get more information about OrganizationSettings, see:

* [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/TopLevel/getSettings)
* How-to Guides
    * [Configure default settings for organizations and folders](https://cloud.google.com/logging/docs/default-settings)

## Example Usage - Logging Organization Settings All


```hcl
resource "google_logging_organization_settings" "example" {
  disable_default_sink = true
  kms_key_name         = "kms-key"
  organization         = "123456789"
  storage_location     = "us-central1"
  depends_on           = [ google_kms_crypto_key_iam_member.iam ]
}

data "google_logging_organization_settings" "settings" {
  organization = "123456789"
}

resource "google_kms_crypto_key_iam_member" "iam" {
  crypto_key_id = "kms-key"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:${data.google_logging_organization_settings.settings.kms_service_account_id}"
}
```

## Argument Reference

The following arguments are supported:


* `organization` -
  (Required)
  The organization for which to retrieve or configure settings.


* `kms_key_name` -
  (Optional)
  The resource name for the configured Cloud KMS key.

* `storage_location` -
  (Optional)
  The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly provided.

* `disable_default_sink` -
  (Optional)
  If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The _Default sink can be re-enabled manually if needed.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `organizations/{{organization}}/settings`

* `name` -
  The resource name of the settings.

* `kms_service_account_id` -
  The service account that will be used by the Log Router to access your Cloud KMS key.

* `logging_service_account_id` -
  The service account for the given container. Sinks use this service account as their writerIdentity if no custom service account is provided.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


OrganizationSettings can be imported using any of these accepted formats:

* `organizations/{{organization}}/settings`
* `{{organization}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import OrganizationSettings using one of the formats above. For example:

```tf
import {
  id = "organizations/{{organization}}/settings"
  to = google_logging_organization_settings.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), OrganizationSettings can be imported using one of the formats above. For example:

```
$ terraform import google_logging_organization_settings.default organizations/{{organization}}/settings
$ terraform import google_logging_organization_settings.default {{organization}}
```
