// Copyright 2015 The Pull Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package feed

import (
	"io/ioutil"
	"testing"
)

var (
	rss2, _ = ioutil.ReadFile("testdata/rss2.xml")
	atom1, _ = ioutil.ReadFile("testdata/atom1.xml")
	opml, _ = ioutil.ReadFile("testdata/opml.xml")
	feedburner, _ = ioutil.ReadFile("testdata/feedburner.xml")
)

func TestParseFeed(t *testing.T) {
	rss, err := ParseFeedContent(rss2)
	if err != nil {
		t.Error(err)
	}
	t.Log(rss.Title, rss.Link, rss.ItemList[0].PubDate)
	rss, err = ParseFeedContent(atom1)
	if err != nil {
		t.Error(err)
	}
	t.Log(rss.Title, rss.Link)
}

func TestFeedBurner(t *testing.T) {
	rss, err := ParseFeedContent(feedburner)
	if err != nil {
		t.Error(err)
	}
	t.Log(rss.Title, rss.Link, rss.ItemList[0].PubDate)
}
func TestParseOPML(t *testing.T) {
	opml, err := ParseOPMLContent(opml)
	if err != nil {
		t.Error(err)
	}
	t.Log(opml.Head.Title)
}