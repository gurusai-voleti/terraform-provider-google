// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package securitycenterv2

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceSecurityCenterV2OrganizationNotificationConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityCenterV2OrganizationNotificationConfigCreate,
		Read:   resourceSecurityCenterV2OrganizationNotificationConfigRead,
		Update: resourceSecurityCenterV2OrganizationNotificationConfigUpdate,
		Delete: resourceSecurityCenterV2OrganizationNotificationConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecurityCenterV2OrganizationNotificationConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"config_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `This must be unique within the organization.`,
			},
			"organization": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The organization whose Cloud Security Command Center the Notification
Config lives in.`,
			},
			"pubsub_topic": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The Pub/Sub topic to send notifications to. Its format is
"projects/[project_id]/topics/[topic]".`,
			},
			"streaming_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The config for triggering streaming-based notifications.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Expression that defines the filter to apply across create/update
events of assets or findings as specified by the event type. The
expression is a list of zero or more restrictions combined via
logical operators AND and OR. Parentheses are supported, and OR
has higher precedence than AND.

Restrictions have the form <field> <operator> <value> and may have
a - character in front of them to indicate negation. The fields
map to those defined in the corresponding resource.

The supported operators are:

* = for all value types.
* >, <, >=, <= for integer values.
* :, meaning substring matching, for strings.

The supported value types are:

* string literals in quotes.
* integer literals without quotes.
* boolean literals true and false without quotes.

See
[Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications)
for information on how to write a filter.`,
						},
					},
				},
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
				Description:  `The description of the notification config (max of 1024 characters).`,
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `location Id is provided by organization. If not provided, Use global as default.`,
				Default:     "global",
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of this notification config, in the format
'organizations/{{organization}}/notificationConfigs/{{config_id}}'.`,
			},
			"service_account": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The service account that needs "pubsub.topics.publish" permission to
publish to the Pub/Sub topic.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecurityCenterV2OrganizationNotificationConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterV2OrganizationNotificationConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	pubsubTopicProp, err := expandSecurityCenterV2OrganizationNotificationConfigPubsubTopic(d.Get("pubsub_topic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("pubsub_topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(pubsubTopicProp)) && (ok || !reflect.DeepEqual(v, pubsubTopicProp)) {
		obj["pubsubTopic"] = pubsubTopicProp
	}
	streamingConfigProp, err := expandSecurityCenterV2OrganizationNotificationConfigStreamingConfig(d.Get("streaming_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("streaming_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(streamingConfigProp)) && (ok || !reflect.DeepEqual(v, streamingConfigProp)) {
		obj["streamingConfig"] = streamingConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}organizations/{{organization}}/locations/{{location}}/notificationConfigs?configId={{config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new OrganizationNotificationConfig: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating OrganizationNotificationConfig: %s", err)
	}
	if err := d.Set("name", flattenSecurityCenterV2OrganizationNotificationConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	log.Printf("[DEBUG] Finished creating OrganizationNotificationConfig %q: %#v", d.Id(), res)

	return resourceSecurityCenterV2OrganizationNotificationConfigRead(d, meta)
}

func resourceSecurityCenterV2OrganizationNotificationConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecurityCenterV2OrganizationNotificationConfig %q", d.Id()))
	}

	if err := d.Set("name", flattenSecurityCenterV2OrganizationNotificationConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationNotificationConfig: %s", err)
	}
	if err := d.Set("description", flattenSecurityCenterV2OrganizationNotificationConfigDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationNotificationConfig: %s", err)
	}
	if err := d.Set("pubsub_topic", flattenSecurityCenterV2OrganizationNotificationConfigPubsubTopic(res["pubsubTopic"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationNotificationConfig: %s", err)
	}
	if err := d.Set("service_account", flattenSecurityCenterV2OrganizationNotificationConfigServiceAccount(res["serviceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationNotificationConfig: %s", err)
	}
	if err := d.Set("streaming_config", flattenSecurityCenterV2OrganizationNotificationConfigStreamingConfig(res["streamingConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationNotificationConfig: %s", err)
	}

	return nil
}

func resourceSecurityCenterV2OrganizationNotificationConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterV2OrganizationNotificationConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	pubsubTopicProp, err := expandSecurityCenterV2OrganizationNotificationConfigPubsubTopic(d.Get("pubsub_topic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("pubsub_topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pubsubTopicProp)) {
		obj["pubsubTopic"] = pubsubTopicProp
	}
	streamingConfigProp, err := expandSecurityCenterV2OrganizationNotificationConfigStreamingConfig(d.Get("streaming_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("streaming_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, streamingConfigProp)) {
		obj["streamingConfig"] = streamingConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating OrganizationNotificationConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("pubsub_topic") {
		updateMask = append(updateMask, "pubsubTopic")
	}

	if d.HasChange("streaming_config") {
		updateMask = append(updateMask, "streamingConfig.filter")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating OrganizationNotificationConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating OrganizationNotificationConfig %q: %#v", d.Id(), res)
		}

	}

	return resourceSecurityCenterV2OrganizationNotificationConfigRead(d, meta)
}

func resourceSecurityCenterV2OrganizationNotificationConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting OrganizationNotificationConfig %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "OrganizationNotificationConfig")
	}

	log.Printf("[DEBUG] Finished deleting OrganizationNotificationConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceSecurityCenterV2OrganizationNotificationConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	stringParts := strings.Split(d.Get("name").(string), "/")
	if len(stringParts) != 6 {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s",
			d.Get("name"),
			"organizations/{{organization}}/locations/{{location}}/notificationConfigs/{{config_id}}",
		)
	}

	if err := d.Set("organization", stringParts[1]); err != nil {
		return nil, fmt.Errorf("Error setting organization: %s", err)
	}
	if err := d.Set("location", stringParts[3]); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("config_id", stringParts[5]); err != nil {
		return nil, fmt.Errorf("Error setting config_id: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenSecurityCenterV2OrganizationNotificationConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2OrganizationNotificationConfigDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2OrganizationNotificationConfigPubsubTopic(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2OrganizationNotificationConfigServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2OrganizationNotificationConfigStreamingConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["filter"] =
		flattenSecurityCenterV2OrganizationNotificationConfigStreamingConfigFilter(original["filter"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityCenterV2OrganizationNotificationConfigStreamingConfigFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecurityCenterV2OrganizationNotificationConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2OrganizationNotificationConfigPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2OrganizationNotificationConfigStreamingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFilter, err := expandSecurityCenterV2OrganizationNotificationConfigStreamingConfigFilter(original["filter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["filter"] = transformedFilter
	}

	return transformed, nil
}

func expandSecurityCenterV2OrganizationNotificationConfigStreamingConfigFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
