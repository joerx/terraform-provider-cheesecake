package cheesecake

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Cheesecake data source
func dataSourceCheesecakes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCheesecakesRead,
		Schema: map[string]*schema.Schema{
			"cheesecakes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id":          {Type: schema.TypeInt, Computed: true},
						"name":        {Type: schema.TypeString, Computed: true},
						"teaser":      {Type: schema.TypeString, Computed: true},
						"description": {Type: schema.TypeString, Computed: true},
						"price":       {Type: schema.TypeInt, Computed: true},
						"image":       {Type: schema.TypeString, Computed: true},
						"ingredients": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ingredient_id": {Type: schema.TypeInt, Computed: true},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceCheesecakesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// just mock the response so we don't need any external services running
	data := `
	[
		{
			"id": 6,
			"name": "Strawberry Cheesecake",
			"teaser": "Spike your cholesterol with a plate of creamy, fatty, tangy goodness",
			"description": "",
			"price": 250,
			"image": "strawberry-cheesecake.png",
			"ingredients": [
				{
					"ingredient_id": 1
				},
				{
					"ingredient_id": 5
				}
			]
		}
	]
	`

	cheesecakes := make([]map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(data), &cheesecakes); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("cheesecakes", cheesecakes); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
