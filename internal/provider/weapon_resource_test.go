package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccExampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + testAccWeaponResourceConfig("one"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("gqldenring_weapon.test", "name", "one"),
					resource.TestCheckResourceAttr("gqldenring_weapon.test", "custom", "true"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "gqldenring_weapon.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: providerConfig + testAccWeaponResourceConfig("two"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("gqldenring_weapon.test", "name", "two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccWeaponResourceConfig(name string) string {
	return fmt.Sprintf(`
resource "gqldenring_weapon" "test" {
 name = %[1]q
}
`, name)
}
