package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testItems = `
<?xml version="1.0"?>
<output>
    <items>
        <item uid="dash-advanced://go/Method/file:///Users/moli/Library/Application%20Support/Dash/DocSets/JavaScript/JavaScript.docset/Contents/Resources/Documents/developer.mozilla.org/en-US/docs/Web/API/History/go.html" arg="0" autocomplete="go">
            <title>go</title>
            <text type="copy">go</text>
            <text type="largetype">go</text>
            <subtitle mod="cmd">Open &quot;go&quot; in browser</subtitle>
            <subtitle mod="alt">Copy &quot;go&quot; to clipboard</subtitle>
            <subtitle>JavaScript - History.go</subtitle>
            <icon>/Users/moli/Library/Application Support/Dash/Data/Alfred/e72b678a9d0dc354d848f87f7d536242.png</icon>
            <quicklookurl>http://127.0.0.1:51737/Dash/skavsvsh/developer.mozilla.org/en-US/docs/Web/API/History/go.html</quicklookurl>
        </item>
        <item uid="dash-advanced://Go/Type/file:///Users/moli/Library/Application%20Support/Dash/DocSets/Go/Go.docset/Contents/Resources/Documents/godoc.org/golang.org/x/tools/go/ssa.html%23//dash_ref_Go/Type/Go/0" arg="1" autocomplete="Go">
            <title>Go</title>
            <text type="copy">Go</text>
            <text type="largetype">Go</text>
            <subtitle mod="cmd">Open &quot;Go&quot; in browser</subtitle>
            <subtitle mod="alt">Copy &quot;Go&quot; to clipboard</subtitle>
            <subtitle>Go - golang.org/x/tools/go/ssa.Go</subtitle>
            <icon>/Users/moli/Library/Application Support/Dash/Data/Alfred/4b3592d25f1bd44118feee39ed162d2f.png</icon>
            <quicklookurl>http://127.0.0.1:51737/Dash/taoqungt/godoc.org/golang.org/x/tools/go/ssa.html#//dash_ref_Go/Type/Go/0</quicklookurl>
        </item>
	</items>
</output>
`

func Test_parseItems(t *testing.T) {
	got, err := parseItems(testItems)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(got), 2)
}
