// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucketreplication


type ObsBucketReplicationRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/obs_bucket_replication#delete_data ObsBucketReplication#delete_data}.
	DeleteData interface{} `field:"optional" json:"deleteData" yaml:"deleteData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/obs_bucket_replication#enabled ObsBucketReplication#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/obs_bucket_replication#history_enabled ObsBucketReplication#history_enabled}.
	HistoryEnabled interface{} `field:"optional" json:"historyEnabled" yaml:"historyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/obs_bucket_replication#prefix ObsBucketReplication#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/obs_bucket_replication#storage_class ObsBucketReplication#storage_class}.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

