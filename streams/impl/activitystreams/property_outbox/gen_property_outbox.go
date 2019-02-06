package propertyoutbox

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// OutboxProperty is the functional property "outbox". It is permitted to be a
// single nilable value type.
type OutboxProperty struct {
	OrderedCollectionMember vocab.OrderedCollectionInterface
	unknown                 interface{}
	iri                     *url.URL
	alias                   string
}

// DeserializeOutboxProperty creates a "outbox" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeOutboxProperty(m map[string]interface{}, aliasMap map[string]string) (*OutboxProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "outbox"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "outbox")
	}
	if i, ok := m[propName]; ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &OutboxProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if m, ok := i.(map[string]interface{}); ok {
			if v, err := mgr.DeserializeOrderedCollectionActivityStreams()(m, aliasMap); err == nil {
				this := &OutboxProperty{
					OrderedCollectionMember: v,
					alias:                   alias,
				}
				return this, nil
			}
		}
		this := &OutboxProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewOutboxProperty creates a new outbox property.
func NewOutboxProperty() *OutboxProperty {
	return &OutboxProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsOrderedCollection
// afterwards will return false.
func (this *OutboxProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.OrderedCollectionMember = nil
}

// Get returns the value of this property. When IsOrderedCollection returns false,
// Get will return any arbitrary value.
func (this OutboxProperty) Get() vocab.OrderedCollectionInterface {
	return this.OrderedCollectionMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this OutboxProperty) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this OutboxProperty) HasAny() bool {
	return this.IsOrderedCollection() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this OutboxProperty) IsIRI() bool {
	return this.iri != nil
}

// IsOrderedCollection returns true if this property is set and not an IRI.
func (this OutboxProperty) IsOrderedCollection() bool {
	return this.OrderedCollectionMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this OutboxProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsOrderedCollection() {
		child = this.Get().JSONLDContext()
	}
	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this OutboxProperty) KindIndex() int {
	if this.IsOrderedCollection() {
		return 0
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this OutboxProperty) LessThan(o vocab.OutboxPropertyInterface) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsOrderedCollection() && !o.IsOrderedCollection() {
		// Both are unknowns.
		return false
	} else if this.IsOrderedCollection() && !o.IsOrderedCollection() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsOrderedCollection() && o.IsOrderedCollection() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return this.Get().LessThan(o.Get())
	}
}

// Name returns the name of this property: "outbox".
func (this OutboxProperty) Name() string {
	return "outbox"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this OutboxProperty) Serialize() (interface{}, error) {
	if this.IsOrderedCollection() {
		return this.Get().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsOrderedCollection afterwards
// will return true.
func (this *OutboxProperty) Set(v vocab.OrderedCollectionInterface) {
	this.Clear()
	this.OrderedCollectionMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *OutboxProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}