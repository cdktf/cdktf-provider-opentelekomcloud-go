// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltsnotificationtemplatev2


type LtsNotificationTemplateV2Templates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_notification_template_v2#content LtsNotificationTemplateV2#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_notification_template_v2#sub_type LtsNotificationTemplateV2#sub_type}.
	SubType *string `field:"required" json:"subType" yaml:"subType"`
}

