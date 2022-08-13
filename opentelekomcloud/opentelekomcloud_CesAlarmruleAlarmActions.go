// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type CesAlarmruleAlarmActions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/ces_alarmrule#notification_list CesAlarmrule#notification_list}.
	NotificationList *[]*string `field:"required" json:"notificationList" yaml:"notificationList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/ces_alarmrule#type CesAlarmrule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

