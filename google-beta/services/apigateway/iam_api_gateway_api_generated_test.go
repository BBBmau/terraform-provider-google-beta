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

package apigateway_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccApiGatewayApiIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_api_gateway_api_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/apis/%s roles/apigateway.viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-api%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccApiGatewayApiIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_api_gateway_api_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/apis/%s roles/apigateway.viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-api%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccApiGatewayApiIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccApiGatewayApiIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_api_gateway_api_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/apis/%s roles/apigateway.viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-api%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccApiGatewayApiIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_api_gateway_api_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_api_gateway_api_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/apis/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-api%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccApiGatewayApiIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_api_gateway_api_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/global/apis/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-api%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccApiGatewayApiIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_iam_member" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccApiGatewayApiIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_api_gateway_api_iam_policy" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_api_gateway_api_iam_policy" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  depends_on = [
    google_api_gateway_api_iam_policy.foo
  ]
}
`, context)
}

func testAccApiGatewayApiIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_api_gateway_api_iam_policy" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccApiGatewayApiIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_iam_binding" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccApiGatewayApiIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_iam_binding" "foo" {
  provider = google-beta
  project = google_api_gateway_api.api.project
  api = google_api_gateway_api.api.api_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
