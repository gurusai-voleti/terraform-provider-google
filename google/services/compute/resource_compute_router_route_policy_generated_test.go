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

package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccComputeRouterRoutePolicy_routerRoutePolicyExportExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterRoutePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterRoutePolicy_routerRoutePolicyExportExample(context),
			},
			{
				ResourceName:            "google_compute_router_route_policy.rp-export",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "router"},
			},
		},
	})
}

func testAccComputeRouterRoutePolicy_routerRoutePolicyExportExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "net" {
  name                    = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet" {
  name          = "tf-test-my-subnetwork%{random_suffix}"
  network       = google_compute_network.net.id
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_router" "router" {
  name    = "tf-test-my-router%{random_suffix}"
  region  = google_compute_subnetwork.subnet.region
  network = google_compute_network.net.id
}

resource "google_compute_router_route_policy" "rp-export" {
  router = google_compute_router.router.name
  region = google_compute_router.router.region
	name = "tf-test-my-rp1%{random_suffix}"
	type = "ROUTE_POLICY_TYPE_EXPORT"
	terms {
    priority = 1
    match {
      expression = "destination == '10.0.0.0/12'"
	  }
    actions {
      expression = "accept()"
    }
  }
}
`, context)
}

func TestAccComputeRouterRoutePolicy_routerRoutePolicyImportExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRouterRoutePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouterRoutePolicy_routerRoutePolicyImportExample(context),
			},
			{
				ResourceName:            "google_compute_router_route_policy.rp-import",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "router"},
			},
		},
	})
}

func testAccComputeRouterRoutePolicy_routerRoutePolicyImportExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "net" {
  name                    = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet" {
  name          = "tf-test-my-subnetwork%{random_suffix}"
  network       = google_compute_network.net.id
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_router" "router" {
  name    = "tf-test-my-router%{random_suffix}"
  region  = google_compute_subnetwork.subnet.region
  network = google_compute_network.net.id
}

resource "google_compute_router_route_policy" "rp-import" {
  name = "tf-test-my-rp2%{random_suffix}"
  router = google_compute_router.router.name
  region = google_compute_router.router.region
	type = "ROUTE_POLICY_TYPE_IMPORT"
	terms {
    priority = 2
    match {
      expression = "destination == '10.0.0.0/12'"
	  }
    actions {
      expression = "accept()"
    }
  }
}
`, context)
}

func testAccCheckComputeRouterRoutePolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_router_route_policy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/routers/{{router}}/getRoutePolicy?policy={{name}}")
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
				return fmt.Errorf("ComputeRouterRoutePolicy still exists at %s", url)
			}
		}

		return nil
	}
}
