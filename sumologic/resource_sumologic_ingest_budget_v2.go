// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//	This file is automatically generated by Sumo Logic and manual
//	changes will be clobbered when the file is regenerated. Do not submit
//	changes to this file.
//
// ----------------------------------------------------------------------------
package sumologic

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSumologicIngestBudgetV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicIngestBudgetV2Create,
		Read:   resourceSumologicIngestBudgetV2Read,
		Update: resourceSumologicIngestBudgetV2Update,
		Delete: resourceSumologicIngestBudgetV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},

			"timezone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"action": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"capacity_bytes": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},

			"reset_time": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"audit_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
			},

			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

func resourceSumologicIngestBudgetV2Create(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	if d.Id() == "" {
		ingestBudgetV2 := resourceToIngestBudgetV2(d)
		id, err := c.CreateIngestBudgetV2(ingestBudgetV2)
		if err != nil {
			return err
		}

		d.SetId(id)
	}

	return resourceSumologicIngestBudgetV2Read(d, meta)
}

func resourceSumologicIngestBudgetV2Read(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	id := d.Id()
	ingestBudgetV2, err := c.GetIngestBudgetV2(id)
	if err != nil {
		return err
	}

	if ingestBudgetV2 == nil {
		log.Printf("[WARN] IngestBudgetV2 not found, removing from state: %v - %v", id, err)
		d.SetId("")
		return nil
	}

	d.Set("name", ingestBudgetV2.Name)
	d.Set("scope", ingestBudgetV2.Scope)
	d.Set("timezone", ingestBudgetV2.Timezone)
	d.Set("reset_time", ingestBudgetV2.ResetTime)
	d.Set("audit_threshold", ingestBudgetV2.AuditThreshold)
	d.Set("description", ingestBudgetV2.Description)
	d.Set("action", ingestBudgetV2.Action)
	d.Set("capacity_bytes", ingestBudgetV2.CapacityBytes)

	return nil
}

func resourceSumologicIngestBudgetV2Delete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	return c.DeleteIngestBudgetV2(d.Id())
}

func resourceSumologicIngestBudgetV2Update(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	ingestBudgetV2 := resourceToIngestBudgetV2(d)
	err := c.UpdateIngestBudgetV2(ingestBudgetV2)
	if err != nil {
		return err
	}

	return resourceSumologicIngestBudgetV2Read(d, meta)
}

func resourceToIngestBudgetV2(d *schema.ResourceData) IngestBudgetV2 {

	return IngestBudgetV2{
		Scope:          d.Get("scope").(string),
		Name:           d.Get("name").(string),
		ResetTime:      d.Get("reset_time").(string),
		Timezone:       d.Get("timezone").(string),
		ID:             d.Id(),
		Action:         d.Get("action").(string),
		Description:    d.Get("description").(string),
		AuditThreshold: d.Get("audit_threshold").(int),
		CapacityBytes:  d.Get("capacity_bytes").(int),
	}
}
