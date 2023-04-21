package cesalarmrule


type CesAlarmruleMetric struct {
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ces_alarmrule#dimensions CesAlarmrule#dimensions}
	Dimensions interface{} `field:"required" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ces_alarmrule#metric_name CesAlarmrule#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ces_alarmrule#namespace CesAlarmrule#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

