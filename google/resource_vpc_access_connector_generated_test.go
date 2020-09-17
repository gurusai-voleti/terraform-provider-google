// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
)

func TestAccVPCAccessConnector_vpcAccessConnectorExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckVPCAccessConnectorDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCAccessConnector_vpcAccessConnectorExample(context),
			},
			{
				ResourceName:            "google_vpc_access_connector.connector",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_link", "region"},
			},
		},
	})
}

func testAccVPCAccessConnector_vpcAccessConnectorExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vpc_access_connector" "connector" {
  name          = "vpcconn%{random_suffix}"
  region        = "us-central1"
  ip_cidr_range = "10.8.0.0/28"
  network       = "default"
}
`, context)
}

func testAccCheckVPCAccessConnectorDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vpc_access_connector" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{VPCAccessBasePath}}projects/{{project}}/locations/{{region}}/connectors/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("VPCAccessConnector still exists at %s", url)
			}
		}

		return nil
	}
}
