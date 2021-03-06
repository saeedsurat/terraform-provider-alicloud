package alicloud

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"

	"github.com/hashicorp/terraform/helper/acctest"
)

func TestAccAlicloudAlikafkaInstancesDataSource(t *testing.T) {

	rand := acctest.RandInt()
	resourceId := "data.alicloud_alikafka_instances.default"
	name := fmt.Sprintf("tf-testacc-alikafkainstance%v", rand)

	testAccConfig := dataSourceTestAccConfigFunc(resourceId, name, dataSourceAlikafkaInstancesConfigDependence)

	nameRegexConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"name_regex": "${alicloud_alikafka_instance.default.name}",
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"name_regex": "fake_tf-testacc*",
		}),
	}

	idsConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"ids": []string{"${alicloud_alikafka_instance.default.id}"},
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"ids": []string{"${alicloud_alikafka_instance.default.id}_fake"},
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"ids":        []string{"${alicloud_alikafka_instance.default.id}"},
			"name_regex": "${alicloud_alikafka_instance.default.name}",
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"ids":        []string{"${alicloud_alikafka_instance.default.id}_fake"},
			"name_regex": "${alicloud_alikafka_instance.default.name}",
		}),
	}

	var existAlikafkaInstancesMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":                   "1",
			"instances.#":             "1",
			"instances.0.name":        fmt.Sprintf("tf-testacc-alikafkainstance%v", rand),
			"instances.0.topic_quota": "50",
			"instances.0.disk_type":   "1",
			"instances.0.disk_size":   "500",
			"instances.0.deploy_type": "5",
			"instances.0.io_max":      "20",
			"instances.0.paid_type":   "PostPaid",
			"instances.0.spec_type":   "normal",
		}
	}

	var fakeAlikafkaInstancesMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":       "0",
			"instances.#": "0",
		}
	}

	var alikafkaInstancesCheckInfo = dataSourceAttr{
		resourceId:   resourceId,
		existMapFunc: existAlikafkaInstancesMapFunc,
		fakeMapFunc:  fakeAlikafkaInstancesMapFunc,
	}
	preCheck := func() {
		testAccPreCheckWithRegions(t, true, connectivity.AlikafkaSupportedRegions)
		testAccPreCheckWithNoDefaultVswitch(t)
	}
	alikafkaInstancesCheckInfo.dataSourceTestCheckWithPreCheck(t, rand, preCheck, nameRegexConf, idsConf, allConf)
}

func dataSourceAlikafkaInstancesConfigDependence(name string) string {
	return fmt.Sprintf(`
		variable "name" {
		 default = "%v"
		}

        data "alicloud_vswitches" "default" {
		  is_default = "true"
		}

		resource "alicloud_alikafka_instance" "default" {
          name = "${var.name}"
		  topic_quota = "50"
		  disk_type = "1"
		  disk_size = "500"
		  deploy_type = "5"
		  io_max = "20"
          vswitch_id = "${data.alicloud_vswitches.default.ids.0}"
		}
		`, name)
}
