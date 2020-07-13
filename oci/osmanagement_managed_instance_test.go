// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	managedInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
	}

	managedInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `osms-instance`},
		"os_family":      Representation{repType: Optional, create: `ALL`},
	}

	ManagedInstanceResourceConfig = ManagedInstanceManagementResourceDependencies
)

func TestOsmanagementManagedInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_osmanagement_managed_instances.test_managed_instances"
	singularDatasourceName := "data.oci_osmanagement_managed_instance.test_managed_instance"

	resourceName := "oci_osmanagement_managed_instance_management.test_managed_instance_management"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// create dependencies
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceResourceConfig,
				Check: func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] OS Management Resource should be created after 5 minutes as OS Agent times to activate")
					time.Sleep(5 * time.Minute)
					return nil
				},
			},
			// verify create
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceResourceConfig +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Required, Create, ManagedInstanceManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_osmanagement_managed_instances", "test_managed_instances", Required, Create, managedInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedInstanceResourceConfig +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Required, Create, ManagedInstanceManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", Required, Create, managedInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedInstanceResourceConfig +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_management", "test_managed_instance_management", Required, Create, ManagedInstanceManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_reboot_required"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "last_boot"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "last_checkin"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_family"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_kernel_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "updates_available"),
				),
			},
		},
	})
}
