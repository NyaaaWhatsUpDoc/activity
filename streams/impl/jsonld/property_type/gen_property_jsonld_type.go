// Code generated by astool. DO NOT EDIT.

package propertytype

import (
	"fmt"
	anyuri "github.com/superseriousbusiness/activity/streams/values/anyURI"
	string1 "github.com/superseriousbusiness/activity/streams/values/string"
	vocab "github.com/superseriousbusiness/activity/streams/vocab"
	"net/url"
)

// JSONLDTypePropertyIterator is an iterator for a property. It is permitted to be
// one of multiple value types. At most, one type of value can be present, or
// none at all. Setting a value will clear the other types of values so that
// only one of the 'Is' methods will return true. It is possible to clear all
// values, so that this property is empty.
type JSONLDTypePropertyIterator struct {
	xmlschemaAnyURIMember *url.URL
	xmlschemaStringMember string
	hasStringMember       bool
	unknown               interface{}
	alias                 string
	myIdx                 int
	parent                vocab.JSONLDTypeProperty
}

// NewJSONLDTypePropertyIterator creates a new JSONLDType property.
func NewJSONLDTypePropertyIterator() *JSONLDTypePropertyIterator {
	return &JSONLDTypePropertyIterator{alias: ""}
}

// deserializeJSONLDTypePropertyIterator creates an iterator from an element that
// has been unmarshalled from a text or binary format.
func deserializeJSONLDTypePropertyIterator(i interface{}, aliasMap map[string]string) (*JSONLDTypePropertyIterator, error) {
	alias := ""

	if v, err := anyuri.DeserializeAnyURI(i); err == nil {
		this := &JSONLDTypePropertyIterator{
			alias:                 alias,
			xmlschemaAnyURIMember: v,
		}
		return this, nil
	} else if v, err := string1.DeserializeString(i); err == nil {
		this := &JSONLDTypePropertyIterator{
			alias:                 alias,
			hasStringMember:       true,
			xmlschemaStringMember: v,
		}
		return this, nil
	}
	this := &JSONLDTypePropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this JSONLDTypePropertyIterator) GetIRI() *url.URL {
	return this.xmlschemaAnyURIMember
}

// GetXMLSchemaAnyURI returns the value of this property. When IsXMLSchemaAnyURI
// returns false, GetXMLSchemaAnyURI will return an arbitrary value.
func (this JSONLDTypePropertyIterator) GetXMLSchemaAnyURI() *url.URL {
	return this.xmlschemaAnyURIMember
}

// GetXMLSchemaString returns the value of this property. When IsXMLSchemaString
// returns false, GetXMLSchemaString will return an arbitrary value.
func (this JSONLDTypePropertyIterator) GetXMLSchemaString() string {
	return this.xmlschemaStringMember
}

// HasAny returns true if any of the different values is set.
func (this JSONLDTypePropertyIterator) HasAny() bool {
	return this.IsXMLSchemaAnyURI() ||
		this.IsXMLSchemaString()
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this JSONLDTypePropertyIterator) IsIRI() bool {
	return this.xmlschemaAnyURIMember != nil
}

// IsXMLSchemaAnyURI returns true if this property has a type of "anyURI". When
// true, use the GetXMLSchemaAnyURI and SetXMLSchemaAnyURI methods to access
// and set this property.
func (this JSONLDTypePropertyIterator) IsXMLSchemaAnyURI() bool {
	return this.xmlschemaAnyURIMember != nil
}

// IsXMLSchemaString returns true if this property has a type of "string". When
// true, use the GetXMLSchemaString and SetXMLSchemaString methods to access
// and set this property.
func (this JSONLDTypePropertyIterator) IsXMLSchemaString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this JSONLDTypePropertyIterator) JSONLDContext() map[string]string {
	var m map[string]string
	var child map[string]string

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
func (this JSONLDTypePropertyIterator) KindIndex() int {
	if this.IsXMLSchemaAnyURI() {
		return 0
	}
	if this.IsXMLSchemaString() {
		return 1
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
func (this JSONLDTypePropertyIterator) LessThan(o vocab.JSONLDTypePropertyIterator) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsXMLSchemaAnyURI() {
		return anyuri.LessAnyURI(this.GetXMLSchemaAnyURI(), o.GetXMLSchemaAnyURI())
	} else if this.IsXMLSchemaString() {
		return string1.LessString(this.GetXMLSchemaString(), o.GetXMLSchemaString())
	}
	return false
}

// Name returns the name of this property: "JSONLDType".
func (this JSONLDTypePropertyIterator) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "JSONLDType"
	} else {
		return "JSONLDType"
	}
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this JSONLDTypePropertyIterator) Next() vocab.JSONLDTypePropertyIterator {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this JSONLDTypePropertyIterator) Prev() vocab.JSONLDTypePropertyIterator {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *JSONLDTypePropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.SetXMLSchemaAnyURI(v)
}

// SetXMLSchemaAnyURI sets the value of this property. Calling IsXMLSchemaAnyURI
// afterwards returns true.
func (this *JSONLDTypePropertyIterator) SetXMLSchemaAnyURI(v *url.URL) {
	this.clear()
	this.xmlschemaAnyURIMember = v
}

// SetXMLSchemaString sets the value of this property. Calling IsXMLSchemaString
// afterwards returns true.
func (this *JSONLDTypePropertyIterator) SetXMLSchemaString(v string) {
	this.clear()
	this.xmlschemaStringMember = v
	this.hasStringMember = true
}

// clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *JSONLDTypePropertyIterator) clear() {
	this.xmlschemaAnyURIMember = nil
	this.hasStringMember = false
	this.unknown = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this JSONLDTypePropertyIterator) serialize() (interface{}, error) {
	if this.IsXMLSchemaAnyURI() {
		return anyuri.SerializeAnyURI(this.GetXMLSchemaAnyURI())
	} else if this.IsXMLSchemaString() {
		return string1.SerializeString(this.GetXMLSchemaString())
	}
	return this.unknown, nil
}

// JSONLDTypeProperty is the non-functional property "type". It is permitted to
// have one or more values, and of different value types.
type JSONLDTypeProperty struct {
	properties []*JSONLDTypePropertyIterator
	alias      string
}

// DeserializeTypeProperty creates a "type" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeTypeProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.JSONLDTypeProperty, error) {
	alias := ""

	propName := "type"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "type")
	}
	i, ok := m[propName]

	if ok {
		this := &JSONLDTypeProperty{
			alias:      alias,
			properties: []*JSONLDTypePropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeJSONLDTypePropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeJSONLDTypePropertyIterator(i, aliasMap); err != nil {
				return this, err
			} else if p != nil {
				this.properties = append(this.properties, p)
			}
		}
		// Set up the properties for iteration.
		for idx, ele := range this.properties {
			ele.parent = this
			ele.myIdx = idx
		}
		return this, nil
	}
	return nil, nil
}

// NewJSONLDTypeProperty creates a new type property.
func NewJSONLDTypeProperty() *JSONLDTypeProperty {
	return &JSONLDTypeProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property "type"
func (this *JSONLDTypeProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &JSONLDTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaAnyURIMember: v,
	})
}

// AppendXMLSchemaAnyURI appends a anyURI value to the back of a list of the
// property "type". Invalidates iterators that are traversing using Prev.
func (this *JSONLDTypeProperty) AppendXMLSchemaAnyURI(v *url.URL) {
	this.properties = append(this.properties, &JSONLDTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaAnyURIMember: v,
	})
}

// AppendXMLSchemaString appends a string value to the back of a list of the
// property "type". Invalidates iterators that are traversing using Prev.
func (this *JSONLDTypeProperty) AppendXMLSchemaString(v string) {
	this.properties = append(this.properties, &JSONLDTypePropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaStringMember: v,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this JSONLDTypeProperty) At(index int) vocab.JSONLDTypePropertyIterator {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this JSONLDTypeProperty) Begin() vocab.JSONLDTypePropertyIterator {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this JSONLDTypeProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this JSONLDTypeProperty) End() vocab.JSONLDTypePropertyIterator {
	return nil
}

// Insert inserts an IRI value at the specified index for a property "type".
// Existing elements at that index and higher are shifted back once.
// Invalidates all iterators.
func (this *JSONLDTypeProperty) InsertIRI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &JSONLDTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// InsertXMLSchemaAnyURI inserts a anyURI value at the specified index for a
// property "type". Existing elements at that index and higher are shifted
// back once. Invalidates all iterators.
func (this *JSONLDTypeProperty) InsertXMLSchemaAnyURI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &JSONLDTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// InsertXMLSchemaString inserts a string value at the specified index for a
// property "type". Existing elements at that index and higher are shifted
// back once. Invalidates all iterators.
func (this *JSONLDTypeProperty) InsertXMLSchemaString(idx int, v string) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &JSONLDTypePropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 idx,
		parent:                this,
		xmlschemaStringMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this JSONLDTypeProperty) JSONLDContext() map[string]string {
	var m map[string]string
	for _, elem := range this.properties {
		child := elem.JSONLDContext()
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API method specifically needed only for alternate implementations
// for go-fed. Applications should not use this method. Panics if the index is
// out of bounds.
func (this JSONLDTypeProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "type" property.
func (this JSONLDTypeProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this JSONLDTypeProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].GetXMLSchemaAnyURI()
			rhs := this.properties[j].GetXMLSchemaAnyURI()
			return anyuri.LessAnyURI(lhs, rhs)
		} else if idx1 == 1 {
			lhs := this.properties[i].GetXMLSchemaString()
			rhs := this.properties[j].GetXMLSchemaString()
			return string1.LessString(lhs, rhs)
		} else if idx1 == -2 {
			lhs := this.properties[i].GetIRI()
			rhs := this.properties[j].GetIRI()
			return lhs.String() < rhs.String()
		}
	}
	return false
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this JSONLDTypeProperty) LessThan(o vocab.JSONLDTypeProperty) bool {
	l1 := this.Len()
	l2 := o.Len()
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if this.properties[i].LessThan(o.At(i)) {
			return true
		} else if o.At(i).LessThan(this.properties[i]) {
			return false
		}
	}
	return l1 < l2
}

// Name returns the name of this property ("type") with any alias.
func (this JSONLDTypeProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "type"
	} else {
		return "type"
	}
}

// PrependIRI prepends an IRI value to the front of a list of the property "type".
func (this *JSONLDTypeProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*JSONLDTypePropertyIterator{{
		alias:                 this.alias,
		myIdx:                 0,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependXMLSchemaAnyURI prepends a anyURI value to the front of a list of the
// property "type". Invalidates all iterators.
func (this *JSONLDTypeProperty) PrependXMLSchemaAnyURI(v *url.URL) {
	this.properties = append([]*JSONLDTypePropertyIterator{{
		alias:                 this.alias,
		myIdx:                 0,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependXMLSchemaString prepends a string value to the front of a list of the
// property "type". Invalidates all iterators.
func (this *JSONLDTypeProperty) PrependXMLSchemaString(v string) {
	this.properties = append([]*JSONLDTypePropertyIterator{{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 0,
		parent:                this,
		xmlschemaStringMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "type", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *JSONLDTypeProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &JSONLDTypePropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this JSONLDTypeProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	// Shortcut: if serializing one value, don't return an array -- pretty sure other Fediverse software would choke on a "type" value with array, for example.
	if len(s) == 1 {
		return s[0], nil
	}
	return s, nil
}

// SetIRI sets an IRI value to be at the specified index for the property "type".
// Panics if the index is out of bounds.
func (this *JSONLDTypeProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &JSONLDTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
}

// SetXMLSchemaAnyURI sets a anyURI value to be at the specified index for the
// property "type". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *JSONLDTypeProperty) SetXMLSchemaAnyURI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &JSONLDTypePropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
}

// SetXMLSchemaString sets a string value to be at the specified index for the
// property "type". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *JSONLDTypeProperty) SetXMLSchemaString(idx int, v string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &JSONLDTypePropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 idx,
		parent:                this,
		xmlschemaStringMember: v,
	}
}

// Swap swaps the location of values at two indices for the "type" property.
func (this JSONLDTypeProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
