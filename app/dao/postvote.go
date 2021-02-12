// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"lqc.com/tree/app/dao/internal"
)

// postvoteDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type postvoteDao struct {
	*internal.PostvoteDao
}

var (
	// Postvote is globally public accessible object for table postvote operations.
	Postvote = &postvoteDao{
		internal.Postvote,
	}
)

// Fill with you ideas below.
