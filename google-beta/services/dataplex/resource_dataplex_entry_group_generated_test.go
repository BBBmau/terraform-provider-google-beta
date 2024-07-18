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

package dataplex_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccDataplexEntryGroup_dataplexEntryGroupBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexEntryGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexEntryGroup_dataplexEntryGroupBasicExample(context),
			},
			{
				ResourceName:            "google_dataplex_entry_group.test_entry_group_basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"entry_group_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDataplexEntryGroup_dataplexEntryGroupBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_entry_group" "test_entry_group_basic" {
  entry_group_id = "tf-test-entry-group-basic%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"
}
`, context)
}

func TestAccDataplexEntryGroup_dataplexEntryGroupFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexEntryGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexEntryGroup_dataplexEntryGroupFullExample(context),
			},
			{
				ResourceName:            "google_dataplex_entry_group.test_entry_group_full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"entry_group_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDataplexEntryGroup_dataplexEntryGroupFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_entry_group" "test_entry_group_full" {
  entry_group_id = "tf-test-entry-group-full%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"

  labels = { "tag": "test-tf" }
  display_name = "terraform entry group"
  description = "entry group created by Terraform"
}
`, context)
}

func testAccCheckDataplexEntryGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dataplex_entry_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataplexBasePath}}projects/{{project}}/locations/{{location}}/entryGroups/{{entry_group_id}}")
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
				return fmt.Errorf("DataplexEntryGroup still exists at %s", url)
			}
		}

		return nil
	}
}
