// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.feed.getPostThread

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
)

// FeedGetPostThread_Output is the output of a app.bsky.feed.getPostThread call.
type FeedGetPostThread_Output struct {
	Thread *FeedGetPostThread_Output_Thread `json:"thread" cborgen:"thread"`
}

type FeedGetPostThread_Output_Thread struct {
	FeedDefs_ThreadViewPost *FeedDefs_ThreadViewPost
	FeedDefs_NotFoundPost   *FeedDefs_NotFoundPost
	FeedDefs_BlockedPost    *FeedDefs_BlockedPost
}

func (t *FeedGetPostThread_Output_Thread) MarshalJSON() ([]byte, error) {
	if t.FeedDefs_ThreadViewPost != nil {
		t.FeedDefs_ThreadViewPost.LexiconTypeID = "app.bsky.feed.defs#threadViewPost"
		return json.Marshal(t.FeedDefs_ThreadViewPost)
	}
	if t.FeedDefs_NotFoundPost != nil {
		t.FeedDefs_NotFoundPost.LexiconTypeID = "app.bsky.feed.defs#notFoundPost"
		return json.Marshal(t.FeedDefs_NotFoundPost)
	}
	if t.FeedDefs_BlockedPost != nil {
		t.FeedDefs_BlockedPost.LexiconTypeID = "app.bsky.feed.defs#blockedPost"
		return json.Marshal(t.FeedDefs_BlockedPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedGetPostThread_Output_Thread) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.defs#threadViewPost":
		t.FeedDefs_ThreadViewPost = new(FeedDefs_ThreadViewPost)
		return json.Unmarshal(b, t.FeedDefs_ThreadViewPost)
	case "app.bsky.feed.defs#notFoundPost":
		t.FeedDefs_NotFoundPost = new(FeedDefs_NotFoundPost)
		return json.Unmarshal(b, t.FeedDefs_NotFoundPost)
	case "app.bsky.feed.defs#blockedPost":
		t.FeedDefs_BlockedPost = new(FeedDefs_BlockedPost)
		return json.Unmarshal(b, t.FeedDefs_BlockedPost)

	default:
		return nil
	}
}

// FeedGetPostThread calls the XRPC method "app.bsky.feed.getPostThread".
//
// depth: How many levels of reply depth should be included in response.
// parentHeight: How many levels of parent (and grandparent, etc) post to include.
// uri: Reference (AT-URI) to post record.
func FeedGetPostThread(ctx context.Context, c *xrpc.Client, depth int64, parentHeight int64, uri string) (*FeedGetPostThread_Output, error) {
	var out FeedGetPostThread_Output

	params := map[string]interface{}{
		"depth":        depth,
		"parentHeight": parentHeight,
		"uri":          uri,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.feed.getPostThread", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}