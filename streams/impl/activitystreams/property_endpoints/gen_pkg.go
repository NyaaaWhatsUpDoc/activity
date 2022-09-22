// Code generated by astool. DO NOT EDIT.

package propertyendpoints

import vocab "github.com/superseriousbusiness/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeEndpointCollectionActivityStreams returns the
	// deserialization method for the "ActivityStreamsEndpointCollection"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeEndpointCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsEndpointCollection, error)
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}
