// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"lqc.com/tree/app/model/internal"
)

// Replyvote is the golang structure for table replyvote.
type Replyvote internal.Replyvote

// Fill with you ideas below.

type ReplyvoteDoVote struct {
	Userid  int
	Replyid *int `json:"Replyid" v:"required"`
	Vote    *int `json:"vote" v:"required|min:0|max:1"`
}
