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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccNetworkConnectivityServiceConnectionPolicy_networkConnectivityPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"service_class_name": "gcp-memorystore-redis",
		"random_suffix":      acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityServiceConnectionPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityServiceConnectionPolicy_networkConnectivityPolicyBasicExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_service_connection_policy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNetworkConnectivityServiceConnectionPolicy_networkConnectivityPolicyBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "producer_net" {
  name                    = "tf-test-producer-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "producer_subnet" {
  name          = "tf-test-producer-subnet%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.producer_net.id
}

resource "google_network_connectivity_service_connection_policy" "default" {
  name = "tf-test-my-network-connectivity-policy%{random_suffix}"
  location = "us-central1"
  service_class = "%{service_class_name}"
  description   = "my basic service connection policy"
  network = google_compute_network.producer_net.id
  psc_config {
    subnetworks = [google_compute_subnetwork.producer_subnet.id]
    limit = 2
  }
}
`, context)
}

func testAccCheckNetworkConnectivityServiceConnectionPolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_connectivity_service_connection_policy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/{{location}}/serviceConnectionPolicies/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("NetworkConnectivityServiceConnectionPolicy still exists at %s", url)
			}
		}

		return nil
	}
}