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

package storage_test

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

func TestAccStorageAnywhereCache_storageAnywhereCacheBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccCheckStorageAnywhereCacheDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageAnywhereCache_storageAnywhereCacheBasicExample(context),
			},
			{
				ResourceName:            "google_storage_anywhere_cache.cache",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bucket"},
			},
		},
	})
}

func testAccStorageAnywhereCache_storageAnywhereCacheBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "bucket" {
  name                        = "tf-test-bucket-name%{random_suffix}"
  location                    = "US"
}

resource "time_sleep" "destroy_wait_5000_seconds" {
  depends_on = [google_storage_bucket.bucket]
  destroy_duration = "5000s"
}

resource "google_storage_anywhere_cache" "cache" {
  bucket = google_storage_bucket.bucket.name
  zone = "us-central1-f"
  ttl = "3601s"
  depends_on = [time_sleep.destroy_wait_5000_seconds]
}
`, context)
}

func testAccCheckStorageAnywhereCacheDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_storage_anywhere_cache" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{StorageBasePath}}b/{{bucket}}/anywhereCaches/{{anywhere_cache_id}}")
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
				return fmt.Errorf("StorageAnywhereCache still exists at %s", url)
			}
		}

		return nil
	}
}
