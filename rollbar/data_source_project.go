package rollbar

import (
	"fmt"

	"github.com/babbel/rollbar-go/rollbar"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProjectRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed values
			"account_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"date_created": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"mock": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
		},
	}
}

func dataSourceProjectRead(d *schema.ResourceData, meta interface{}) error {
	if d.Get("mock").(bool) {
		return nil
	}

	name := d.Get("name").(string)

	client := meta.(*rollbar.Client)
	project, err := client.GetProjectByName(name)
	if err != nil {
		return err
	}
	if project == nil {
		d.SetId("")
		return fmt.Errorf("No project with the name %s found", name)
	}

	id := fmt.Sprintf("%d", project.ID)
	d.SetId(id)
	d.Set("account_id", project.AccountID)
	d.Set("date_created", project.DateCreated)

	return nil
}
