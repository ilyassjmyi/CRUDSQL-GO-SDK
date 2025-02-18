/*
CrudSQL API

A powerful dynamic CRUD API engine that automatically generates RESTful endpoints for your data models CrudSQL provides automatic CRUD operations, filtering, pagination, and sorting capabilities for any data model. Features: - Automatic REST API generation - Dynamic model support - Complex filtering and querying - Pagination and sorting - Swagger documentation - Multiple database support (SQL & NoSQL)

API version: 1.0.0
Contact: taqi@mobix.biz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the QueryQueryFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryQueryFilter{}

// QueryQueryFilter Filter conditions for querying entities using complex expressions Supports logical operations (AND, OR, NOT), field comparisons, and relationship filtering Example: { \"expressions\": [ {\"field\": \"age\", \"operator\": \"gte\", \"value\": 18}, { \"operator\": \"AND\", \"expressions\": [ {\"field\": \"status\", \"operator\": \"eq\", \"value\": \"active\"} ] } ] }
type QueryQueryFilter struct {
	// @Description Array of expressions to filter entities @Description Each expression can be a FieldExpression, LogicalExpression, or RelationshipExpression
	Expressions []map[string]interface{} `json:"expressions,omitempty"`
}

// NewQueryQueryFilter instantiates a new QueryQueryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryQueryFilter() *QueryQueryFilter {
	this := QueryQueryFilter{}
	return &this
}

// NewQueryQueryFilterWithDefaults instantiates a new QueryQueryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryQueryFilterWithDefaults() *QueryQueryFilter {
	this := QueryQueryFilter{}
	return &this
}

// GetExpressions returns the Expressions field value if set, zero value otherwise.
func (o *QueryQueryFilter) GetExpressions() []map[string]interface{} {
	if o == nil || IsNil(o.Expressions) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Expressions
}

// GetExpressionsOk returns a tuple with the Expressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryQueryFilter) GetExpressionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Expressions) {
		return nil, false
	}
	return o.Expressions, true
}

// HasExpressions returns a boolean if a field has been set.
func (o *QueryQueryFilter) HasExpressions() bool {
	if o != nil && !IsNil(o.Expressions) {
		return true
	}

	return false
}

// SetExpressions gets a reference to the given []map[string]interface{} and assigns it to the Expressions field.
func (o *QueryQueryFilter) SetExpressions(v []map[string]interface{}) {
	o.Expressions = v
}

func (o QueryQueryFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryQueryFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expressions) {
		toSerialize["expressions"] = o.Expressions
	}
	return toSerialize, nil
}

type NullableQueryQueryFilter struct {
	value *QueryQueryFilter
	isSet bool
}

func (v NullableQueryQueryFilter) Get() *QueryQueryFilter {
	return v.value
}

func (v *NullableQueryQueryFilter) Set(val *QueryQueryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryQueryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryQueryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryQueryFilter(val *QueryQueryFilter) *NullableQueryQueryFilter {
	return &NullableQueryQueryFilter{value: val, isSet: true}
}

func (v NullableQueryQueryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryQueryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


