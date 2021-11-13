// Code generated by astool. DO NOT EDIT.

package typeorderedcollectionpage

import vocab "github.com/superseriousbusiness/activity/streams/vocab"

var mgr privateManager

var typePropertyConstructor func() vocab.JSONLDTypeProperty

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeAltitudePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsAltitudeProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeAltitudePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAltitudeProperty, error)
	// DeserializeAttachmentPropertyActivityStreams returns the
	// deserialization method for the "ActivityStreamsAttachmentProperty"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeAttachmentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAttachmentProperty, error)
	// DeserializeAttributedToPropertyActivityStreams returns the
	// deserialization method for the
	// "ActivityStreamsAttributedToProperty" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeAttributedToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAttributedToProperty, error)
	// DeserializeAudiencePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsAudienceProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeAudiencePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAudienceProperty, error)
	// DeserializeBccPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsBccProperty" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeBccPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsBccProperty, error)
	// DeserializeBtoPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsBtoProperty" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeBtoPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsBtoProperty, error)
	// DeserializeCcPropertyActivityStreams returns the deserialization method
	// for the "ActivityStreamsCcProperty" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeCcPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCcProperty, error)
	// DeserializeContentPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsContentProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeContentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsContentProperty, error)
	// DeserializeContextPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsContextProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeContextPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsContextProperty, error)
	// DeserializeCurrentPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsCurrentProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeCurrentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCurrentProperty, error)
	// DeserializeDurationPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsDurationProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeDurationPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDurationProperty, error)
	// DeserializeEarlyItemsPropertyForgeFed returns the deserialization
	// method for the "ForgeFedEarlyItemsProperty" non-functional property
	// in the vocabulary "ForgeFed"
	DeserializeEarlyItemsPropertyForgeFed() func(map[string]interface{}, map[string]string) (vocab.ForgeFedEarlyItemsProperty, error)
	// DeserializeEndTimePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsEndTimeProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeEndTimePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsEndTimeProperty, error)
	// DeserializeFirstPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsFirstProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeFirstPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsFirstProperty, error)
	// DeserializeGeneratorPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsGeneratorProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeGeneratorPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsGeneratorProperty, error)
	// DeserializeIconPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsIconProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeIconPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsIconProperty, error)
	// DeserializeIdPropertyJSONLD returns the deserialization method for the
	// "JSONLDIdProperty" non-functional property in the vocabulary
	// "JSONLD"
	DeserializeIdPropertyJSONLD() func(map[string]interface{}, map[string]string) (vocab.JSONLDIdProperty, error)
	// DeserializeImagePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsImageProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeImagePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsImageProperty, error)
	// DeserializeInReplyToPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsInReplyToProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeInReplyToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsInReplyToProperty, error)
	// DeserializeLastPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsLastProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeLastPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLastProperty, error)
	// DeserializeLikesPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsLikesProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeLikesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLikesProperty, error)
	// DeserializeLocationPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsLocationProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeLocationPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLocationProperty, error)
	// DeserializeMediaTypePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsMediaTypeProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeMediaTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsMediaTypeProperty, error)
	// DeserializeNamePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsNameProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeNamePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsNameProperty, error)
	// DeserializeNextPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsNextProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeNextPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsNextProperty, error)
	// DeserializeObjectPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsObjectProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeObjectPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsObjectProperty, error)
	// DeserializeOrderedItemsPropertyActivityStreams returns the
	// deserialization method for the
	// "ActivityStreamsOrderedItemsProperty" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeOrderedItemsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrderedItemsProperty, error)
	// DeserializePartOfPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsPartOfProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializePartOfPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPartOfProperty, error)
	// DeserializePrevPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsPrevProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializePrevPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPrevProperty, error)
	// DeserializePreviewPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsPreviewProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializePreviewPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPreviewProperty, error)
	// DeserializePublishedPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsPublishedProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializePublishedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPublishedProperty, error)
	// DeserializeRepliesPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsRepliesProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeRepliesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRepliesProperty, error)
	// DeserializeSensitivePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsSensitiveProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeSensitivePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsSensitiveProperty, error)
	// DeserializeSharesPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsSharesProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeSharesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsSharesProperty, error)
	// DeserializeSourcePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsSourceProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeSourcePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsSourceProperty, error)
	// DeserializeStartIndexPropertyActivityStreams returns the
	// deserialization method for the "ActivityStreamsStartIndexProperty"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeStartIndexPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsStartIndexProperty, error)
	// DeserializeStartTimePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsStartTimeProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeStartTimePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsStartTimeProperty, error)
	// DeserializeSummaryPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsSummaryProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeSummaryPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsSummaryProperty, error)
	// DeserializeTagPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsTagProperty" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeTagPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTagProperty, error)
	// DeserializeTeamPropertyForgeFed returns the deserialization method for
	// the "ForgeFedTeamProperty" non-functional property in the
	// vocabulary "ForgeFed"
	DeserializeTeamPropertyForgeFed() func(map[string]interface{}, map[string]string) (vocab.ForgeFedTeamProperty, error)
	// DeserializeTicketsTrackedByPropertyForgeFed returns the deserialization
	// method for the "ForgeFedTicketsTrackedByProperty" non-functional
	// property in the vocabulary "ForgeFed"
	DeserializeTicketsTrackedByPropertyForgeFed() func(map[string]interface{}, map[string]string) (vocab.ForgeFedTicketsTrackedByProperty, error)
	// DeserializeToPropertyActivityStreams returns the deserialization method
	// for the "ActivityStreamsToProperty" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsToProperty, error)
	// DeserializeTotalItemsPropertyActivityStreams returns the
	// deserialization method for the "ActivityStreamsTotalItemsProperty"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeTotalItemsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTotalItemsProperty, error)
	// DeserializeTracksTicketsForPropertyForgeFed returns the deserialization
	// method for the "ForgeFedTracksTicketsForProperty" non-functional
	// property in the vocabulary "ForgeFed"
	DeserializeTracksTicketsForPropertyForgeFed() func(map[string]interface{}, map[string]string) (vocab.ForgeFedTracksTicketsForProperty, error)
	// DeserializeTypePropertyJSONLD returns the deserialization method for
	// the "JSONLDTypeProperty" non-functional property in the vocabulary
	// "JSONLD"
	DeserializeTypePropertyJSONLD() func(map[string]interface{}, map[string]string) (vocab.JSONLDTypeProperty, error)
	// DeserializeUpdatedPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsUpdatedProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeUpdatedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUpdatedProperty, error)
	// DeserializeUrlPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsUrlProperty" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeUrlPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUrlProperty, error)
}

// jsonldContexter is a private interface to determine the JSON-LD contexts and
// aliases needed for functional and non-functional properties. It is a helper
// interface for this implementation.
type jsonldContexter interface {
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}

// SetTypePropertyConstructor sets the "type" property's constructor in the
// package-global variable. For internal use only, do not use as part of
// Application behavior. Must be called at golang init time. Permits
// ActivityStreams types to correctly set their "type" property at
// construction time, so users don't have to remember to do so each time. It
// is dependency injected so other go-fed compatible implementations could
// inject their own type.
func SetTypePropertyConstructor(f func() vocab.JSONLDTypeProperty) {
	typePropertyConstructor = f
}
