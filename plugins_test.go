// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"testing"
	"context"
)

func TestClientPlugins(t *testing.T) {
	client, err := NewClient(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Plugins()
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientHasPlugin(t *testing.T) {
	client, err := NewClient(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	found, err := client.HasPlugin("no-such-plugin")
	if err != nil {
		t.Fatal(err)
	}
	if found {
		t.Fatalf("expected to not find plugin %q", "no-such-plugin")
	}
}
