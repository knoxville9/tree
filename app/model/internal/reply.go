// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Reply is the golang structure for table reply.
type Reply struct {
	Id       uint        `orm:"id,primary" json:"id"`       //
	Pid      uint        `orm:"pid"        json:"pid"`      //
	Userid   uint        `orm:"userid"     json:"userid"`   //
	Content  string      `orm:"content"    json:"content"`  //
	CreateAt *gtime.Time `orm:"CreateAt"   json:"createAt"` //
	Deleted  int         `orm:"deleted"    json:"deleted"`  // 删除为1
}
