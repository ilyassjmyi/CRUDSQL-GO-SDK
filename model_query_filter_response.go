package openapi

import (
	"encoding/json"
)

// QueryFilterResponse Paginated response containing filtered entities and metadata
type QueryFilterResponse struct {
	Data       DataWrapper `json:"data,omitempty"`
	Page       *int32      `json:"page,omitempty"`
	PageSize   *int32      `json:"page_size,omitempty"`
	Total      *int32      `json:"total,omitempty"`
	TotalPages *int32      `json:"total_pages,omitempty"`
}

// DataWrapper is a custom type to handle both single map and array of maps
type DataWrapper struct {
	Value interface{}
}

// UnmarshalJSON custom unmarshaler for DataWrapper
func (dw *DataWrapper) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a single map
	var singleMap map[string]interface{}
	if err := json.Unmarshal(data, &singleMap); err == nil {
		dw.Value = singleMap
		return nil
	}

	// Try to unmarshal as an array of maps
	var arrayOfMaps []map[string]interface{}
	if err := json.Unmarshal(data, &arrayOfMaps); err == nil {
		dw.Value = arrayOfMaps
		return nil
	}

	return json.Unmarshal(data, &dw.Value)
}

// MarshalJSON custom marshaler for DataWrapper
func (dw DataWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(dw.Value)
}

// GetData returns the Data field value
func (o *QueryFilterResponse) GetData() interface{} {
	if o == nil {
		return nil
	}
	return o.Data.Value
}

// SetData sets the Data field value
func (o *QueryFilterResponse) SetData(v interface{}) {
	o.Data.Value = v
}

// NewQueryFilterResponse instantiates a new QueryFilterResponse object
func NewQueryFilterResponse() *QueryFilterResponse {
	this := QueryFilterResponse{}
	return &this
}

// NewQueryFilterResponseWithDefaults instantiates a new QueryFilterResponse object with default values
func NewQueryFilterResponseWithDefaults() *QueryFilterResponse {
	this := QueryFilterResponse{}
	return &this
}

// ToMap converts the QueryFilterResponse struct to a map
func (o QueryFilterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.Value != nil {
		toSerialize["data"] = o.Data.Value
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	return toSerialize, nil
}

// MarshalJSON converts the QueryFilterResponse struct to JSON
func (o QueryFilterResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

type NullableQueryFilterResponse struct {
	value *QueryFilterResponse
	isSet bool
}

func (v NullableQueryFilterResponse) Get() *QueryFilterResponse {
	return v.value
}

func (v *NullableQueryFilterResponse) Set(val *QueryFilterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryFilterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryFilterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryFilterResponse(val *QueryFilterResponse) *NullableQueryFilterResponse {
	return &NullableQueryFilterResponse{value: val, isSet: true}
}

func (v NullableQueryFilterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryFilterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
