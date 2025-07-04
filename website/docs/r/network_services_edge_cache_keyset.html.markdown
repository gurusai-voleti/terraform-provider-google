---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networkservices/EdgeCacheKeyset.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Network Services"
description: |-
  EdgeCacheKeyset represents a collection of public keys used for validating signed requests.
---

# google_network_services_edge_cache_keyset

EdgeCacheKeyset represents a collection of public keys used for validating signed requests.


To get more information about EdgeCacheKeyset, see:

* [API documentation](https://cloud.google.com/media-cdn/docs/reference/rest/v1/projects.locations.edgeCacheKeysets)
* How-to Guides
    * [Create keysets](https://cloud.google.com/media-cdn/docs/create-keyset)

~> **Warning:** All arguments including the following potentially sensitive
values will be stored in the raw state as plain text: `public_key.public_key.value`.
[Read more about sensitive data in state](https://www.terraform.io/language/state/sensitive-data).

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=network_services_edge_cache_keyset_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Network Services Edge Cache Keyset Basic


```hcl

resource "google_network_services_edge_cache_keyset" "default" {
  name                 = "my-keyset"
  description          = "The default keyset"
  public_key {
    id = "my-public-key"
    value = "FHsTyFHNmvNpw4o7-rp-M1yqMyBF8vXSBRkZtkQ0RKY"
  }
  public_key {
    id = "my-public-key-2"
    value = "hzd03llxB1u5FOLKFkZ6_wCJqC7jtN0bg7xlBqS6WVM"
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=network_services_edge_cache_keyset_dual_token&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Network Services Edge Cache Keyset Dual Token


```hcl
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "secret-name"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret-version-basic" {
  secret = google_secret_manager_secret.secret-basic.id

  secret_data = "secret-data"
}

resource "google_network_services_edge_cache_keyset" "default" {
  name        = "my-keyset"
  description = "The default keyset"
  public_key {
    id      = "my-public-key"
    managed = true
  }
  validation_shared_keys {
    secret_version = google_secret_manager_secret_version.secret-version-basic.id
  }
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  Name of the resource; provided by the client when the resource is created.
  The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
  and all following characters must be a dash, underscore, letter or digit.


* `description` -
  (Optional)
  A human-readable description of the resource.

* `labels` -
  (Optional)
  Set of label tags associated with the EdgeCache resource.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `public_key` -
  (Optional)
  An ordered list of Ed25519 public keys to use for validating signed requests.
  You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
  You may specify no more than one Google-managed public key.
  If you specify `public_keys`, you must specify at least one (1) key and may specify up to three (3) keys.
  Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
  Ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
  Structure is [documented below](#nested_public_key).

* `validation_shared_keys` -
  (Optional)
  An ordered list of shared keys to use for validating signed requests.
  Shared keys are secret.  Ensure that only authorized users can add `validation_shared_keys` to a keyset.
  You can rotate keys by appending (pushing) a new key to the list of `validation_shared_keys` and removing any superseded keys.
  You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
  Structure is [documented below](#nested_validation_shared_keys).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_public_key"></a>The `public_key` block supports:

* `id` -
  (Required)
  The ID of the public key. The ID must be 1-63 characters long, and comply with RFC1035.
  The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]*
  which means the first character must be a letter, and all following characters must be a dash, underscore, letter or digit.

* `value` -
  (Optional)
  The base64-encoded value of the Ed25519 public key. The base64 encoding can be padded (44 bytes) or unpadded (43 bytes).
  Representations or encodings of the public key other than this will be rejected with an error.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `managed` -
  (Optional)
  Set to true to have the CDN automatically manage this public key value.

<a name="nested_validation_shared_keys"></a>The `validation_shared_keys` block supports:

* `secret_version` -
  (Required)
  The name of the secret version in Secret Manager.
  The resource name of the secret version must be in the format `projects/*/secrets/*/versions/*` where the `*` values are replaced by the secrets themselves.
  The secrets must be at least 16 bytes large.  The recommended secret size depends on the signature algorithm you are using.
  * If you are using HMAC-SHA1, we suggest 20-byte secrets.
  * If you are using HMAC-SHA256, we suggest 32-byte secrets.
  See RFC 2104, Section 3 for more details on these recommendations.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}`

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 90 minutes.
- `update` - Default is 90 minutes.
- `delete` - Default is 90 minutes.

## Import


EdgeCacheKeyset can be imported using any of these accepted formats:

* `projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}`
* `{{project}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import EdgeCacheKeyset using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}"
  to = google_network_services_edge_cache_keyset.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), EdgeCacheKeyset can be imported using one of the formats above. For example:

```
$ terraform import google_network_services_edge_cache_keyset.default projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}
$ terraform import google_network_services_edge_cache_keyset.default {{project}}/{{name}}
$ terraform import google_network_services_edge_cache_keyset.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
