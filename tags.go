package rkstrings

// Tag stores a tag with an optional value
type Tag struct {
	Key   string
	Value string
}

// Tags wraps a map of tags
type Tags struct {
	Map map[string]string
}

// NewEmptyTags creates a new empty tags struct
func NewEmptyTags() Tags {
	return Tags{Map: make(map[string]string)}
}

// NewTags creates a new tags struct with the given values
func NewTags(tags []Tag) Tags {
	var result Tags
	result.Map = make(map[string]string)
	for _, tag := range tags {
		result.Map[tag.Key] = tag.Value
	}
	return result
}

// Count returns the count in tags
func (tags *Tags) Count() int {
	return len(tags.Map)
}

// Add the given tag to tags
func (tags *Tags) Add(tag Tag) {
	tags.Map[tag.Key] = tag.Value
}

// Remove the given tag from tags
func (tags *Tags) Remove(key string) {
	delete(tags.Map, key)
}

// Has returns true if the tag is present
func (tags *Tags) Has(key string) bool {
	_, ok := tags.Map[key]
	return ok
}

// MustTag returns the value for the given tag, or panics if the
// key is not present.
func (tags *Tags) MustTag(key string) string {
	v, ok := tags.Map[key]
	if !ok {
		panic("could not find tag: " + key)
	}
	return v
}

// GetWithDefault gets the value for the given tag, returning dflt if
// the key is not present.
func (tags *Tags) GetWithDefault(key, dflt string) string {
	v, ok := tags.Map[key]
	if !ok {
		return dflt
	}
	return v
}
