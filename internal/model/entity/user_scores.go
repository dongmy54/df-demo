// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserScores is the golang structure for table user_scores.
type UserScores struct {
	Id     uint   `json:"id"     orm:"id"      description:""`
	UserId uint   `json:"userId" orm:"user_id" description:""`
	Score  uint   `json:"score"  orm:"score"   description:""`
	Course string `json:"course" orm:"course"  description:""`
}
