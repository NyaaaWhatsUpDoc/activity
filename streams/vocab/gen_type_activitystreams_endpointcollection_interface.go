// Code generated by astool. DO NOT EDIT.

package vocab

// A json object which maps additional (typically server/domain-wide) endpoints
// which may be useful either for this actor or someone referencing this
// actor. This mapping may be nested inside the actor document as the value or
// may be a link to a JSON-LD document with these properties.
type ActivityStreamsEndpointCollection interface {
	// GetActivityStreamsSharedInbox returns the "sharedInbox" property if it
	// exists, and nil otherwise.
	GetActivityStreamsSharedInbox() ActivityStreamsSharedInboxProperty
	// GetJSONLDId returns the "id" property if it exists, and nil otherwise.
	GetJSONLDId() JSONLDIdProperty
	// GetJSONLDType returns the "type" property if it exists, and nil
	// otherwise.
	GetJSONLDType() JSONLDTypeProperty
	// GetTypeName returns the name of this type.
	GetTypeName() string
	// GetUnknownProperties returns the unknown properties for the
	// EndpointCollection type. Note that this should not be used by app
	// developers. It is only used to help determine which implementation
	// is LessThan the other. Developers who are creating a different
	// implementation of this type's interface can use this method in
	// their LessThan implementation, but routine ActivityPub applications
	// should not use this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// IsExtending returns true if the EndpointCollection type extends from
	// the other type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this EndpointCollection is lesser, with an
	// arbitrary but stable determination.
	LessThan(o ActivityStreamsEndpointCollection) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetActivityStreamsSharedInbox sets the "sharedInbox" property.
	SetActivityStreamsSharedInbox(i ActivityStreamsSharedInboxProperty)
	// SetJSONLDId sets the "id" property.
	SetJSONLDId(i JSONLDIdProperty)
	// SetJSONLDType sets the "type" property.
	SetJSONLDType(i JSONLDTypeProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}
