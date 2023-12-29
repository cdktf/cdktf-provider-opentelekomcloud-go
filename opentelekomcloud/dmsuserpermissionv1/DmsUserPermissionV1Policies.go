// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsuserpermissionv1


type DmsUserPermissionV1Policies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/dms_user_permission_v1#access_policy DmsUserPermissionV1#access_policy}.
	AccessPolicy *string `field:"required" json:"accessPolicy" yaml:"accessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/dms_user_permission_v1#username DmsUserPermissionV1#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

