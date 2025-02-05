package openapi

import (
	"encoding/json"
)

// checks if the QueryFilterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryFilterResponse{}

// QueryFilterResponse Paginated response containing filtered entities and metadata
type QueryFilterResponse struct {
	// Data can be either an array of entities or a map
	Data json.RawMessage `json:"data,omitempty"`
	// Current page number (1-based indexing)
	Page *int32 `json:"page,omitempty"`
	// Number of items per page
	PageSize *int32 `json:"page_size,omitempty"`
	// Total number of records matching the filter conditions
	Total *int32 `json:"total,omitempty"`
	// Total number of pages based on total records and page size
	TotalPages *int32 `json:"total_pages,omitempty"`
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

// GetData returns the Data field value if set, zero value otherwise
func (o *QueryFilterResponse) GetData() json.RawMessage {
	if o == nil || IsNil(o.Data) {
		return nil
	}
	return o.Data
}

// GetDataAsArray attempts to parse the Data field as an array of interface{}
func (o *QueryFilterResponse) GetDataAsArray() ([]interface{}, error) {
	if o == nil || IsNil(o.Data) {
		return nil, nil
	}
	var result []interface{}
	err := json.Unmarshal(o.Data, &result)
	return result, err
}

// GetDataAsMap attempts to parse the Data field as a map[string]interface{}
func (o *QueryFilterResponse) GetDataAsMap() (map[string]interface{}, error) {
	if o == nil || IsNil(o.Data) {
		return nil, nil
	}
	var result map[string]interface{}
	err := json.Unmarshal(o.Data, &result)
	return result, err
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryFilterResponse) GetDataOk() (json.RawMessage, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *QueryFilterResponse) HasData() bool {
	return o != nil && !IsNil(o.Data)
}

// SetData sets the Data field
func (o *QueryFilterResponse) SetData(v interface{}) error {
	if v == nil {
		o.Data = nil
		return nil
	}

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	o.Data = data
	return nil
}

// GetPage returns the Page field value if set, zero value otherwise
func (o *QueryFilterResponse) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryFilterResponse) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *QueryFilterResponse) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}
	return false
}

// SetPage sets the Page field
func (o *QueryFilterResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise
func (o *QueryFilterResponse) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryFilterResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *QueryFilterResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}
	return false
}

// SetPageSize sets the PageSize field
func (o *QueryFilterResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotal returns the Total field value if set, zero value otherwise
func (o *QueryFilterResponse) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryFilterResponse) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryFilterResponse) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}
	return false
}

// SetTotal sets the Total field
func (o *QueryFilterResponse) SetTotal(v int32) {
	o.Total = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise
func (o *QueryFilterResponse) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryFilterResponse) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *QueryFilterResponse) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}
	return false
}

// SetTotalPages sets the TotalPages field
func (o *QueryFilterResponse) SetTotalPages(v int32) {
	o.TotalPages = &v
}

func (o QueryFilterResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryFilterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		var data interface{}
		err := json.Unmarshal(o.Data, &data)
		if err != nil {
			return nil, err
		}
		toSerialize["data"] = data
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
