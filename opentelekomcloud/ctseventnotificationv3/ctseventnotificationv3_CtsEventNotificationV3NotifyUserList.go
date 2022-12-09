package ctseventnotificationv3


type CtsEventNotificationV3NotifyUserList struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_event_notification_v3#user_group CtsEventNotificationV3#user_group}.
	UserGroup *string `field:"required" json:"userGroup" yaml:"userGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_event_notification_v3#user_list CtsEventNotificationV3#user_list}.
	UserList *[]*string `field:"required" json:"userList" yaml:"userList"`
}

