// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.graph.listblock

import (
	"github.com/bluesky-social/indigo/lex/util"
)

func init() {
	util.RegisterType("app.bsky.graph.listblock", &GraphListblock{})
} //
// RECORDTYPE: GraphListblock
type GraphListblock struct {
	LexiconTypeID string `json:"$type,const=app.bsky.graph.listblock" cborgen:"$type,const=app.bsky.graph.listblock"`
	CreatedAt     string `json:"createdAt" cborgen:"createdAt"`
	// subject: Reference (AT-URI) to the mod list record.
	Subject string `json:"subject" cborgen:"subject"`
}