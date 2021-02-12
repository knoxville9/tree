// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"lqc.com/tree/app/dao/internal"
)

// postDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type postDao struct {
	*internal.PostDao
}

var (
	// Post is globally public accessible object for table post operations.
	Post = &postDao{
		internal.Post,
	}
)

// Fill with you ideas below.
