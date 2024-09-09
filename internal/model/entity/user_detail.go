// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserDetail is the golang structure for table user_detail.
type UserDetail struct {
	Id      uint   `json:"id"      orm:"id"      description:""`
	UserId  uint   `json:"userId"  orm:"user_id" description:""`
	Address string `json:"address" orm:"address" description:""`
}
