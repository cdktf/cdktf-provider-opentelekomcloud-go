// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ctseventnotificationv3


type CtsEventNotificationV3Operations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/cts_event_notification_v3#resource_type CtsEventNotificationV3#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/cts_event_notification_v3#service_type CtsEventNotificationV3#service_type}.
	ServiceType *string `field:"required" json:"serviceType" yaml:"serviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/cts_event_notification_v3#trace_names CtsEventNotificationV3#trace_names}.
	TraceNames *[]*string `field:"required" json:"traceNames" yaml:"traceNames"`
}

