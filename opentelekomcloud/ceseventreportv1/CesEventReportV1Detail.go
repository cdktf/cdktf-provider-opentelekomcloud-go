// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ceseventreportv1


type CesEventReportV1Detail struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/ces_event_report_v1#content CesEventReportV1#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/ces_event_report_v1#event_level CesEventReportV1#event_level}.
	EventLevel *string `field:"optional" json:"eventLevel" yaml:"eventLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/ces_event_report_v1#event_state CesEventReportV1#event_state}.
	EventState *string `field:"optional" json:"eventState" yaml:"eventState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/ces_event_report_v1#event_type CesEventReportV1#event_type}.
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/ces_event_report_v1#event_user CesEventReportV1#event_user}.
	EventUser *string `field:"optional" json:"eventUser" yaml:"eventUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/ces_event_report_v1#group_id CesEventReportV1#group_id}.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/ces_event_report_v1#resource_id CesEventReportV1#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/ces_event_report_v1#resource_name CesEventReportV1#resource_name}.
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
}

