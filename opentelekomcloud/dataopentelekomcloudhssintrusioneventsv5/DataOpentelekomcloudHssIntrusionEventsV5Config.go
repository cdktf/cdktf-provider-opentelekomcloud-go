// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhssintrusioneventsv5

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudHssIntrusionEventsV5Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Event category. Its value can be: host (host security event) or container (container security event).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#category DataOpentelekomcloudHssIntrusionEventsV5#category}
	Category *string `field:"required" json:"category" yaml:"category"`
	// Customized start time of a segment.
	//
	// The timestamp is accurate to seconds. The begin_time should be no more than two days earlier than the end_time. This parameter is mutually exclusive with the queried duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#begin_time DataOpentelekomcloudHssIntrusionEventsV5#begin_time}
	BeginTime *string `field:"optional" json:"beginTime" yaml:"beginTime"`
	// Container instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#container_name DataOpentelekomcloudHssIntrusionEventsV5#container_name}
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Number of days to be queried. This parameter is mutually exclusive with begin_time and end_time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#days DataOpentelekomcloudHssIntrusionEventsV5#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Customized end time of a segment.
	//
	// The timestamp is accurate to seconds. The begin_time should be no more than two days earlier than the end_time. This parameter is mutually exclusive with the queried duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#end_time DataOpentelekomcloudHssIntrusionEventsV5#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Enterprise project ID.
	//
	// The value 0 indicates the default enterprise project. To query all enterprise projects, set this parameter to all_granted_eps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#enterprise_project_id DataOpentelekomcloudHssIntrusionEventsV5#enterprise_project_id}
	EnterpriseProjectId *string `field:"optional" json:"enterpriseProjectId" yaml:"enterpriseProjectId"`
	// Intrusion types.
	//
	// Possible values include:
	// 1001: Malware
	// 1010: Rootkit
	// 1011: Ransomware
	// 1015: Web shell
	// 1017: Reverse shell
	// 2001: Common vulnerability exploit
	// 3002: File privilege escalation
	// 3003: Process privilege escalation
	// 3004: Important file change
	// 3005: File/Directory change
	// 3007: Abnormal process behavior
	// 3015: High-risk command execution
	// 3018: Abnormal shell
	// 3027: Suspicious crontab tasks
	// 4002: Brute-force attack
	// 4004: Abnormal login
	// 4006: Invalid system account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#event_types DataOpentelekomcloudHssIntrusionEventsV5#event_types}
	EventTypes *[]*string `field:"optional" json:"eventTypes" yaml:"eventTypes"`
	// Status. Possible values: unhandled, handled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#handle_status DataOpentelekomcloudHssIntrusionEventsV5#handle_status}
	HandleStatus *string `field:"optional" json:"handleStatus" yaml:"handleStatus"`
	// Host ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#host_id DataOpentelekomcloudHssIntrusionEventsV5#host_id}
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// Server name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#host_name DataOpentelekomcloudHssIntrusionEventsV5#host_name}
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#id DataOpentelekomcloudHssIntrusionEventsV5#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Server IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#private_ip DataOpentelekomcloudHssIntrusionEventsV5#private_ip}
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Threat level. Possible values: Security, Low, Medium, High, Critical.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/hss_intrusion_events_v5#severity DataOpentelekomcloudHssIntrusionEventsV5#severity}
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
}

