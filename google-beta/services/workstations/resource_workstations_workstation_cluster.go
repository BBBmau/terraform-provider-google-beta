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

package workstations

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

func ResourceWorkstationsWorkstationCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkstationsWorkstationClusterCreate,
		Read:   resourceWorkstationsWorkstationClusterRead,
		Update: resourceWorkstationsWorkstationClusterUpdate,
		Delete: resourceWorkstationsWorkstationClusterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceWorkstationsWorkstationClusterImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.SetAnnotationsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The relative resource name of the VPC network on which the instance can be accessed.
It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".`,
			},
			"subnetwork": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
Must be part of the subnetwork specified for this cluster.`,
			},
			"workstation_cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID to use for the workstation cluster.`,
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Client-specified annotations. This is distinct from labels.

**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
Please refer to the field 'effective_annotations' for all of the annotations present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Human-readable name for this resource.`,
			},
			"domain_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration options for a custom domain.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `Domain used by Workstations for HTTP ingress.`,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The location where the workstation cluster should reside.`,
			},
			"private_cluster_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration for private cluster.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_private_endpoint": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: `Whether Workstations endpoint is private.`,
						},
						"allowed_projects": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Description: `Additional project IDs that are allowed to attach to the workstation cluster's service attachment.
By default, the workstation cluster's project and the VPC host project (if different) are allowed.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"cluster_hostname": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `Hostname for the workstation cluster.
This field will be populated only when private endpoint is enabled.
To access workstations in the cluster, create a new DNS zone mapping this domain name to an internal IP address and a forwarding rule mapping that address to the service attachment.`,
						},
						"service_attachment_uri": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `Service attachment URI for the workstation cluster.
The service attachment is created when private endpoint is enabled.
To access workstations in the cluster, configure access to the managed service using (Private Service Connect)[https://cloud.google.com/vpc/docs/configure-private-service-connect-services].`,
						},
					},
				},
			},
			"conditions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Status conditions describing the current resource state.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The status code, which should be an enum value of google.rpc.Code.`,
						},
						"details": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `A list of messages that carry the error details.`,
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Human readable message indicating details about the current status.`,
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time when this resource was created.`,
			},
			"degraded": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `Whether this resource is in degraded mode, in which case it may require user action to restore full functionality.
Details can be found in the conditions field.`,
			},
			"effective_annotations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Checksum computed by the server.
May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the cluster resource.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The system-generated UID of the resource.`,
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

func resourceWorkstationsWorkstationClusterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	networkProp, err := expandWorkstationsWorkstationClusterNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	subnetworkProp, err := expandWorkstationsWorkstationClusterSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subnetwork"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	displayNameProp, err := expandWorkstationsWorkstationClusterDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	etagProp, err := expandWorkstationsWorkstationClusterEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	privateClusterConfigProp, err := expandWorkstationsWorkstationClusterPrivateClusterConfig(d.Get("private_cluster_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("private_cluster_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(privateClusterConfigProp)) && (ok || !reflect.DeepEqual(v, privateClusterConfigProp)) {
		obj["privateClusterConfig"] = privateClusterConfigProp
	}
	domainConfigProp, err := expandWorkstationsWorkstationClusterDomainConfig(d.Get("domain_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("domain_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(domainConfigProp)) && (ok || !reflect.DeepEqual(v, domainConfigProp)) {
		obj["domainConfig"] = domainConfigProp
	}
	labelsProp, err := expandWorkstationsWorkstationClusterEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandWorkstationsWorkstationClusterEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{WorkstationsBasePath}}projects/{{project}}/locations/{{location}}/workstationClusters?workstationClusterId={{workstation_cluster_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new WorkstationCluster: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WorkstationCluster: %s", err)
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
		return fmt.Errorf("Error creating WorkstationCluster: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = WorkstationsOperationWaitTime(
		config, res, project, "Creating WorkstationCluster", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create WorkstationCluster: %s", err)
	}

	log.Printf("[DEBUG] Finished creating WorkstationCluster %q: %#v", d.Id(), res)

	return resourceWorkstationsWorkstationClusterRead(d, meta)
}

func resourceWorkstationsWorkstationClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{WorkstationsBasePath}}projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WorkstationCluster: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("WorkstationsWorkstationCluster %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}

	if err := d.Set("name", flattenWorkstationsWorkstationClusterName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("uid", flattenWorkstationsWorkstationClusterUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("labels", flattenWorkstationsWorkstationClusterLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("network", flattenWorkstationsWorkstationClusterNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("subnetwork", flattenWorkstationsWorkstationClusterSubnetwork(res["subnetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("display_name", flattenWorkstationsWorkstationClusterDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("degraded", flattenWorkstationsWorkstationClusterDegraded(res["degraded"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("annotations", flattenWorkstationsWorkstationClusterAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("etag", flattenWorkstationsWorkstationClusterEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("create_time", flattenWorkstationsWorkstationClusterCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("private_cluster_config", flattenWorkstationsWorkstationClusterPrivateClusterConfig(res["privateClusterConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("domain_config", flattenWorkstationsWorkstationClusterDomainConfig(res["domainConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("conditions", flattenWorkstationsWorkstationClusterConditions(res["conditions"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("terraform_labels", flattenWorkstationsWorkstationClusterTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("effective_labels", flattenWorkstationsWorkstationClusterEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}
	if err := d.Set("effective_annotations", flattenWorkstationsWorkstationClusterEffectiveAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkstationCluster: %s", err)
	}

	return nil
}

func resourceWorkstationsWorkstationClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WorkstationCluster: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandWorkstationsWorkstationClusterDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	etagProp, err := expandWorkstationsWorkstationClusterEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	privateClusterConfigProp, err := expandWorkstationsWorkstationClusterPrivateClusterConfig(d.Get("private_cluster_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("private_cluster_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, privateClusterConfigProp)) {
		obj["privateClusterConfig"] = privateClusterConfigProp
	}
	domainConfigProp, err := expandWorkstationsWorkstationClusterDomainConfig(d.Get("domain_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("domain_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, domainConfigProp)) {
		obj["domainConfig"] = domainConfigProp
	}
	labelsProp, err := expandWorkstationsWorkstationClusterEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandWorkstationsWorkstationClusterEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{WorkstationsBasePath}}projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating WorkstationCluster %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("etag") {
		updateMask = append(updateMask, "etag")
	}

	if d.HasChange("private_cluster_config") {
		updateMask = append(updateMask, "privateClusterConfig")
	}

	if d.HasChange("domain_config") {
		updateMask = append(updateMask, "domainConfig")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("effective_annotations") {
		updateMask = append(updateMask, "annotations")
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
			return fmt.Errorf("Error updating WorkstationCluster %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating WorkstationCluster %q: %#v", d.Id(), res)
		}

		err = WorkstationsOperationWaitTime(
			config, res, project, "Updating WorkstationCluster", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceWorkstationsWorkstationClusterRead(d, meta)
}

func resourceWorkstationsWorkstationClusterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WorkstationCluster: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{WorkstationsBasePath}}projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting WorkstationCluster %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "WorkstationCluster")
	}

	err = WorkstationsOperationWaitTime(
		config, res, project, "Deleting WorkstationCluster", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting WorkstationCluster %q: %#v", d.Id(), res)
	return nil
}

func resourceWorkstationsWorkstationClusterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/workstationClusters/(?P<workstation_cluster_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenWorkstationsWorkstationClusterName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenWorkstationsWorkstationClusterNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterSubnetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterDegraded(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("annotations"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenWorkstationsWorkstationClusterEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterPrivateClusterConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enable_private_endpoint"] =
		flattenWorkstationsWorkstationClusterPrivateClusterConfigEnablePrivateEndpoint(original["enablePrivateEndpoint"], d, config)
	transformed["cluster_hostname"] =
		flattenWorkstationsWorkstationClusterPrivateClusterConfigClusterHostname(original["clusterHostname"], d, config)
	transformed["service_attachment_uri"] =
		flattenWorkstationsWorkstationClusterPrivateClusterConfigServiceAttachmentUri(original["serviceAttachmentUri"], d, config)
	transformed["allowed_projects"] =
		flattenWorkstationsWorkstationClusterPrivateClusterConfigAllowedProjects(original["allowedProjects"], d, config)
	return []interface{}{transformed}
}
func flattenWorkstationsWorkstationClusterPrivateClusterConfigEnablePrivateEndpoint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterPrivateClusterConfigClusterHostname(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterPrivateClusterConfigServiceAttachmentUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterPrivateClusterConfigAllowedProjects(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterDomainConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["domain"] =
		flattenWorkstationsWorkstationClusterDomainConfigDomain(original["domain"], d, config)
	return []interface{}{transformed}
}
func flattenWorkstationsWorkstationClusterDomainConfigDomain(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterConditions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"code":    flattenWorkstationsWorkstationClusterConditionsCode(original["code"], d, config),
			"message": flattenWorkstationsWorkstationClusterConditionsMessage(original["message"], d, config),
			"details": flattenWorkstationsWorkstationClusterConditionsDetails(original["details"], d, config),
		})
	}
	return transformed
}
func flattenWorkstationsWorkstationClusterConditionsCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenWorkstationsWorkstationClusterConditionsMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterConditionsDetails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenWorkstationsWorkstationClusterEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationClusterEffectiveAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandWorkstationsWorkstationClusterNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterSubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterPrivateClusterConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnablePrivateEndpoint, err := expandWorkstationsWorkstationClusterPrivateClusterConfigEnablePrivateEndpoint(original["enable_private_endpoint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnablePrivateEndpoint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enablePrivateEndpoint"] = transformedEnablePrivateEndpoint
	}

	transformedClusterHostname, err := expandWorkstationsWorkstationClusterPrivateClusterConfigClusterHostname(original["cluster_hostname"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterHostname); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clusterHostname"] = transformedClusterHostname
	}

	transformedServiceAttachmentUri, err := expandWorkstationsWorkstationClusterPrivateClusterConfigServiceAttachmentUri(original["service_attachment_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAttachmentUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAttachmentUri"] = transformedServiceAttachmentUri
	}

	transformedAllowedProjects, err := expandWorkstationsWorkstationClusterPrivateClusterConfigAllowedProjects(original["allowed_projects"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedProjects); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedProjects"] = transformedAllowedProjects
	}

	return transformed, nil
}

func expandWorkstationsWorkstationClusterPrivateClusterConfigEnablePrivateEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterPrivateClusterConfigClusterHostname(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterPrivateClusterConfigServiceAttachmentUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterPrivateClusterConfigAllowedProjects(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterDomainConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDomain, err := expandWorkstationsWorkstationClusterDomainConfigDomain(original["domain"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["domain"] = transformedDomain
	}

	return transformed, nil
}

func expandWorkstationsWorkstationClusterDomainConfigDomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationClusterEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandWorkstationsWorkstationClusterEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
