package log

import (
	"testing"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/stretchr/testify/require"
)

func TestLabelsBuilder_Get(t *testing.T) {
	b := NewLabelsBuilder()
	b.Reset(labels.Labels{labels.Label{Name: "already", Value: "in"}})
	b.Set("foo", "bar")
	b.Set("bar", "buzz")
	b.Del("foo")
	_, ok := b.Get("foo")
	require.False(t, ok)
	v, ok := b.Get("bar")
	require.True(t, ok)
	require.Equal(t, "buzz", v)
	v, ok = b.Get("already")
	require.True(t, ok)
	require.Equal(t, "in", v)
	b.Del("bar")
	_, ok = b.Get("bar")
	require.False(t, ok)
	b.Del("already")
	_, ok = b.Get("already")
	require.False(t, ok)
}
