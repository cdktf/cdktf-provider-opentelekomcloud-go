package ddsinstancev3


type DdsInstanceV3Datastore struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dds_instance_v3#type DdsInstanceV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dds_instance_v3#version DdsInstanceV3#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dds_instance_v3#storage_engine DdsInstanceV3#storage_engine}.
	StorageEngine *string `field:"optional" json:"storageEngine" yaml:"storageEngine"`
}

