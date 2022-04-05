// Code generated by astool. DO NOT EDIT.

package propertypartof

import vocab "github.com/superseriousbusiness/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeCollectionActivityStreams returns the deserialization method
	// for the "ActivityStreamsCollection" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCollection, error)
	// DeserializeCollectionPageActivityStreams returns the deserialization
	// method for the "ActivityStreamsCollectionPage" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCollectionPage, error)
	// DeserializeLinkActivityStreams returns the deserialization method for
	// the "ActivityStreamsLink" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeLinkActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLink, error)
	// DeserializeMentionActivityStreams returns the deserialization method
	// for the "ActivityStreamsMention" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeMentionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsMention, error)
	// DeserializeOrderedCollectionActivityStreams returns the deserialization
	// method for the "ActivityStreamsOrderedCollection" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeOrderedCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrderedCollection, error)
	// DeserializeOrderedCollectionPageActivityStreams returns the
	// deserialization method for the
	// "ActivityStreamsOrderedCollectionPage" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeOrderedCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrderedCollectionPage, error)
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}
