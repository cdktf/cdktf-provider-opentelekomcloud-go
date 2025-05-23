// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cssconfigurationv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CssConfigurationV1Config struct {
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
	// The CSS cluster ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#cluster_id CssConfigurationV1#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Whether to auto-create index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#auto_create_index CssConfigurationV1#auto_create_index}
	AutoCreateIndex *string `field:"optional" json:"autoCreateIndex" yaml:"autoCreateIndex"`
	// Whether to return the Access-Control-Allow-Credentials of the header during cross-domain access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#http_cors_allow_credentials CssConfigurationV1#http_cors_allow_credentials}
	HttpCorsAllowCredentials *string `field:"optional" json:"httpCorsAllowCredentials" yaml:"httpCorsAllowCredentials"`
	// Headers allowed for cross-domain access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#http_cors_allow_headers CssConfigurationV1#http_cors_allow_headers}
	HttpCorsAllowHeaders *string `field:"optional" json:"httpCorsAllowHeaders" yaml:"httpCorsAllowHeaders"`
	// Methods allowed for cross-domain access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#http_cors_allow_methods CssConfigurationV1#http_cors_allow_methods}
	HttpCorsAllowMethods *string `field:"optional" json:"httpCorsAllowMethods" yaml:"httpCorsAllowMethods"`
	// Origin IP address allowed for cross-domain access, for example, **122.122.122.122:9200**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#http_cors_allow_origin CssConfigurationV1#http_cors_allow_origin}
	HttpCorsAllowOrigin *string `field:"optional" json:"httpCorsAllowOrigin" yaml:"httpCorsAllowOrigin"`
	// Whether to allow cross-domain access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#http_cors_enabled CssConfigurationV1#http_cors_enabled}
	HttpCorsEnabled *string `field:"optional" json:"httpCorsEnabled" yaml:"httpCorsEnabled"`
	// Cache duration of the browser. The cache is automatically cleared after the time range you specify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#http_cors_max_age CssConfigurationV1#http_cors_max_age}
	HttpCorsMaxAge *string `field:"optional" json:"httpCorsMaxAge" yaml:"httpCorsMaxAge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#id CssConfigurationV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Cache size in the query phase. Value range: **1** to **100**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#indices_queries_cache_size CssConfigurationV1#indices_queries_cache_size}
	IndicesQueriesCacheSize *string `field:"optional" json:"indicesQueriesCacheSize" yaml:"indicesQueriesCacheSize"`
	// Configured for migrating data from the current cluster to the target cluster through the reindex API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#reindex_remote_whitelist CssConfigurationV1#reindex_remote_whitelist}
	ReindexRemoteWhitelist *string `field:"optional" json:"reindexRemoteWhitelist" yaml:"reindexRemoteWhitelist"`
	// Queue size in the force merge thread pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#thread_pool_force_merge_size CssConfigurationV1#thread_pool_force_merge_size}
	ThreadPoolForceMergeSize *string `field:"optional" json:"threadPoolForceMergeSize" yaml:"threadPoolForceMergeSize"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/css_configuration_v1#timeouts CssConfigurationV1#timeouts}
	Timeouts *CssConfigurationV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

