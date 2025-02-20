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

package secretmanager_test

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

func TestAccSecretManagerSecretVersion_secretVersionBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecretVersion_secretVersionBasicExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret_version.secret-version-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret", "secret_data_wo_version"},
			},
		},
	})
}

func testAccSecretManagerSecretVersion_secretVersionBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-version%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    auto {}
  }
}


resource "google_secret_manager_secret_version" "secret-version-basic" {
  secret = google_secret_manager_secret.secret-basic.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}
`, context)
}

func TestAccSecretManagerSecretVersion_secretVersionBasicWriteOnlyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecretVersion_secretVersionBasicWriteOnlyExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret_version.secret-version-basic-write-only",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret", "secret_data_wo_version"},
			},
		},
	})
}

func testAccSecretManagerSecretVersion_secretVersionBasicWriteOnlyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-basic-write-only" {
  secret_id = "tf-test-secret-version-write-only%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    auto {}
  }
}


resource "google_secret_manager_secret_version" "secret-version-basic-write-only" {
  secret = google_secret_manager_secret.secret-basic-write-only.id
  secret_data_wo_version = 1
  secret_data_wo = "tf-test-secret-data-write-only%{random_suffix}"
}
`, context)
}

func TestAccSecretManagerSecretVersion_secretVersionDeletionPolicyAbandonExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecretVersion_secretVersionDeletionPolicyAbandonExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret_version.secret-version-deletion-policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_policy", "secret", "secret_data_wo_version"},
			},
		},
	})
}

func testAccSecretManagerSecretVersion_secretVersionDeletionPolicyAbandonExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-version%{random_suffix}"

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
    }
  }
}

resource "google_secret_manager_secret_version" "secret-version-deletion-policy" {
  secret = google_secret_manager_secret.secret-basic.id

  secret_data = "tf-test-secret-data%{random_suffix}"
  deletion_policy = "ABANDON"
}
`, context)
}

func TestAccSecretManagerSecretVersion_secretVersionDeletionPolicyDisableExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecretVersion_secretVersionDeletionPolicyDisableExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret_version.secret-version-deletion-policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_policy", "secret", "secret_data_wo_version"},
			},
		},
	})
}

func testAccSecretManagerSecretVersion_secretVersionDeletionPolicyDisableExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-version%{random_suffix}"

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
    }
  }
}

resource "google_secret_manager_secret_version" "secret-version-deletion-policy" {
  secret = google_secret_manager_secret.secret-basic.id

  secret_data = "tf-test-secret-data%{random_suffix}"
  deletion_policy = "DISABLE"
}
`, context)
}

func TestAccSecretManagerSecretVersion_secretVersionWithBase64StringSecretDataExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"data":          "./test-fixtures/binary-file.pfx",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecretVersion_secretVersionWithBase64StringSecretDataExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret_version.secret-version-base64",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_secret_data_base64", "secret", "secret_data_wo_version"},
			},
		},
	})
}

func testAccSecretManagerSecretVersion_secretVersionWithBase64StringSecretDataExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-version%{random_suffix}"

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
    }
  }
}

resource "google_secret_manager_secret_version" "secret-version-base64" {
  secret = google_secret_manager_secret.secret-basic.id

  is_secret_data_base64 = true
  secret_data = filebase64("%{data}")
}
`, context)
}

func TestAccSecretManagerSecretVersion_secretVersionWithBase64StringSecretDataWriteOnlyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"data":          "./test-fixtures/binary-file.pfx",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecretVersion_secretVersionWithBase64StringSecretDataWriteOnlyExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret_version.secret-version-base64-write-only",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_secret_data_base64", "secret", "secret_data_wo_version"},
			},
		},
	})
}

func testAccSecretManagerSecretVersion_secretVersionWithBase64StringSecretDataWriteOnlyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-version-base64-write-only%{random_suffix}"

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
    }
  }
}

resource "google_secret_manager_secret_version" "secret-version-base64-write-only" {
  secret = google_secret_manager_secret.secret-basic.id

  is_secret_data_base64 = true
  secret_data_wo_version = 1
  secret_data_wo = filebase64("%{data}")
}
`, context)
}

func testAccCheckSecretManagerSecretVersionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_secret_manager_secret_version" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SecretManagerBasePath}}{{name}}")
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
				return fmt.Errorf("SecretManagerSecretVersion still exists at %s", url)
			}
		}

		return nil
	}
}
