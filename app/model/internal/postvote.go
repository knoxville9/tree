// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Postvote is the golang structure for table postvote.
type Postvote struct {
	Id       uint        `orm:"id,primary" json:"id"`       //
	Pid      uint        `orm:"pid"        json:"pid"`      //
	Userid   uint        `orm:"userid"     json:"userid"`   //
	Vote     int         `orm:"vote"       json:"vote"`     // 点赞1,踩0
	CreateAt *gtime.Time `orm:"CreateAt"   json:"createAt"` //
	Deleted  int         `orm:"deleted"    json:"-"`        // 删除为1
}
