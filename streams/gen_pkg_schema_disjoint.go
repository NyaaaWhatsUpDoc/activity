// Code generated by astool. DO NOT EDIT.

package streams

import (
	typepropertyvalue "github.com/superseriousbusiness/activity/streams/impl/schema/type_propertyvalue"
	vocab "github.com/superseriousbusiness/activity/streams/vocab"
)

// SchemaPropertyValueIsDisjointWith returns true if PropertyValue is disjoint
// with the other's type.
func SchemaPropertyValueIsDisjointWith(other vocab.Type) bool {
	return typepropertyvalue.PropertyValueIsDisjointWith(other)
}
