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

package networksecurity

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceNetworkSecurityClientTlsPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityClientTlsPolicyCreate,
		Read:   resourceNetworkSecurityClientTlsPolicyRead,
		Update: resourceNetworkSecurityClientTlsPolicyUpdate,
		Delete: resourceNetworkSecurityClientTlsPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityClientTlsPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the ClientTlsPolicy resource.`,
			},
			"client_certificate": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_provider_instance": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"plugin_instance": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.`,
									},
								},
							},
							ExactlyOneOf: []string{},
						},
						"grpc_endpoint": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `gRPC specific configuration to access the gRPC server to obtain the cert and private key.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_uri": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".`,
									},
								},
							},
							ExactlyOneOf: []string{},
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Set of label tags associated with the ClientTlsPolicy resource.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The location of the client tls policy.
The default value is 'global'.`,
				Default: "global",
			},
			"server_validation_ca": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_provider_instance": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"plugin_instance": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.`,
									},
								},
							},
							ExactlyOneOf: []string{},
						},
						"grpc_endpoint": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `gRPC specific configuration to access the gRPC server to obtain the cert and private key.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_uri": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".`,
									},
								},
							},
							ExactlyOneOf: []string{},
						},
					},
				},
			},
			"sni": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the ClientTlsPolicy was created in UTC.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the ClientTlsPolicy was updated in UTC.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceNetworkSecurityClientTlsPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityClientTlsPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sniProp, err := expandNetworkSecurityClientTlsPolicySni(d.Get("sni"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sni"); !tpgresource.IsEmptyValue(reflect.ValueOf(sniProp)) && (ok || !reflect.DeepEqual(v, sniProp)) {
		obj["sni"] = sniProp
	}
	clientCertificateProp, err := expandNetworkSecurityClientTlsPolicyClientCertificate(d.Get("client_certificate"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_certificate"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientCertificateProp)) && (ok || !reflect.DeepEqual(v, clientCertificateProp)) {
		obj["clientCertificate"] = clientCertificateProp
	}
	serverValidationCaProp, err := expandNetworkSecurityClientTlsPolicyServerValidationCa(d.Get("server_validation_ca"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("server_validation_ca"); !tpgresource.IsEmptyValue(reflect.ValueOf(serverValidationCaProp)) && (ok || !reflect.DeepEqual(v, serverValidationCaProp)) {
		obj["serverValidationCa"] = serverValidationCaProp
	}
	labelsProp, err := expandNetworkSecurityClientTlsPolicyEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/clientTlsPolicies?clientTlsPolicyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ClientTlsPolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ClientTlsPolicy: %s", err)
	}
	billingProject = project

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
		return fmt.Errorf("Error creating ClientTlsPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating ClientTlsPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ClientTlsPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ClientTlsPolicy %q: %#v", d.Id(), res)

	return resourceNetworkSecurityClientTlsPolicyRead(d, meta)
}

func resourceNetworkSecurityClientTlsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ClientTlsPolicy: %s", err)
	}
	billingProject = project

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityClientTlsPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ClientTlsPolicy: %s", err)
	}

	if err := d.Set("create_time", flattenNetworkSecurityClientTlsPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ClientTlsPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityClientTlsPolicyUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ClientTlsPolicy: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecurityClientTlsPolicyLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ClientTlsPolicy: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecurityClientTlsPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ClientTlsPolicy: %s", err)
	}
	if err := d.Set("sni", flattenNetworkSecurityClientTlsPolicySni(res["sni"], d, config)); err != nil {
		return fmt.Errorf("Error reading ClientTlsPolicy: %s", err)
	}
	if err := d.Set("client_certificate", flattenNetworkSecurityClientTlsPolicyClientCertificate(res["clientCertificate"], d, config)); err != nil {
		return fmt.Errorf("Error reading ClientTlsPolicy: %s", err)
	}
	if err := d.Set("server_validation_ca", flattenNetworkSecurityClientTlsPolicyServerValidationCa(res["serverValidationCa"], d, config)); err != nil {
		return fmt.Errorf("Error reading ClientTlsPolicy: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkSecurityClientTlsPolicyTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ClientTlsPolicy: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkSecurityClientTlsPolicyEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ClientTlsPolicy: %s", err)
	}

	return nil
}

func resourceNetworkSecurityClientTlsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ClientTlsPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityClientTlsPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sniProp, err := expandNetworkSecurityClientTlsPolicySni(d.Get("sni"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sni"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sniProp)) {
		obj["sni"] = sniProp
	}
	clientCertificateProp, err := expandNetworkSecurityClientTlsPolicyClientCertificate(d.Get("client_certificate"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_certificate"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientCertificateProp)) {
		obj["clientCertificate"] = clientCertificateProp
	}
	serverValidationCaProp, err := expandNetworkSecurityClientTlsPolicyServerValidationCa(d.Get("server_validation_ca"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("server_validation_ca"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, serverValidationCaProp)) {
		obj["serverValidationCa"] = serverValidationCaProp
	}
	labelsProp, err := expandNetworkSecurityClientTlsPolicyEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ClientTlsPolicy %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("sni") {
		updateMask = append(updateMask, "sni")
	}

	if d.HasChange("client_certificate") {
		updateMask = append(updateMask, "clientCertificate")
	}

	if d.HasChange("server_validation_ca") {
		updateMask = append(updateMask, "serverValidationCa")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
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
			return fmt.Errorf("Error updating ClientTlsPolicy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ClientTlsPolicy %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating ClientTlsPolicy", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecurityClientTlsPolicyRead(d, meta)
}

func resourceNetworkSecurityClientTlsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ClientTlsPolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ClientTlsPolicy %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ClientTlsPolicy")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting ClientTlsPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ClientTlsPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityClientTlsPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/clientTlsPolicies/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityClientTlsPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityClientTlsPolicyUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityClientTlsPolicyLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkSecurityClientTlsPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityClientTlsPolicySni(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityClientTlsPolicyClientCertificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["grpc_endpoint"] =
		flattenNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpoint(original["grpcEndpoint"], d, config)
	transformed["certificate_provider_instance"] =
		flattenNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance(original["certificateProviderInstance"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpoint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target_uri"] =
		flattenNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpointTargetUri(original["targetUri"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpointTargetUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["plugin_instance"] =
		flattenNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstancePluginInstance(original["pluginInstance"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstancePluginInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityClientTlsPolicyServerValidationCa(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"grpc_endpoint":                 flattenNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint(original["grpcEndpoint"], d, config),
			"certificate_provider_instance": flattenNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance(original["certificateProviderInstance"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target_uri"] =
		flattenNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpointTargetUri(original["targetUri"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpointTargetUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["plugin_instance"] =
		flattenNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstancePluginInstance(original["pluginInstance"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstancePluginInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityClientTlsPolicyTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkSecurityClientTlsPolicyEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityClientTlsPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityClientTlsPolicySni(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityClientTlsPolicyClientCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGrpcEndpoint, err := expandNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpoint(original["grpc_endpoint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrpcEndpoint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["grpcEndpoint"] = transformedGrpcEndpoint
	}

	transformedCertificateProviderInstance, err := expandNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance(original["certificate_provider_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateProviderInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["certificateProviderInstance"] = transformedCertificateProviderInstance
	}

	return transformed, nil
}

func expandNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetUri, err := expandNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpointTargetUri(original["target_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetUri"] = transformedTargetUri
	}

	return transformed, nil
}

func expandNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpointTargetUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPluginInstance, err := expandNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstancePluginInstance(original["plugin_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPluginInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pluginInstance"] = transformedPluginInstance
	}

	return transformed, nil
}

func expandNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstancePluginInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityClientTlsPolicyServerValidationCa(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGrpcEndpoint, err := expandNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint(original["grpc_endpoint"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGrpcEndpoint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["grpcEndpoint"] = transformedGrpcEndpoint
		}

		transformedCertificateProviderInstance, err := expandNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance(original["certificate_provider_instance"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCertificateProviderInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["certificateProviderInstance"] = transformedCertificateProviderInstance
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetUri, err := expandNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpointTargetUri(original["target_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetUri"] = transformedTargetUri
	}

	return transformed, nil
}

func expandNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpointTargetUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPluginInstance, err := expandNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstancePluginInstance(original["plugin_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPluginInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pluginInstance"] = transformedPluginInstance
	}

	return transformed, nil
}

func expandNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstancePluginInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityClientTlsPolicyEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
