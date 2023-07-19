package ctseventnotificationv3


type CtsEventNotificationV3NotifyUserListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.3/docs/resources/cts_event_notification_v3#user_group CtsEventNotificationV3#user_group}.
	UserGroup *string `field:"required" json:"userGroup" yaml:"userGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.3/docs/resources/cts_event_notification_v3#user_list CtsEventNotificationV3#user_list}.
	UserList *[]*string `field:"required" json:"userList" yaml:"userList"`
}

