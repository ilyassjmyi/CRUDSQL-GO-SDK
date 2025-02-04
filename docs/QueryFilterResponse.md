# QueryFilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | @Description Array of entities matching the filter conditions @Description For relationship queries, includes related entities based on the filter | [optional] 
**Page** | Pointer to **int32** | @Description Current page number (1-based indexing) @Description Example: page&#x3D;2 returns the second page of results | [optional] 
**PageSize** | Pointer to **int32** | @Description Number of items per page (default may vary) @Description Example: page_size&#x3D;20 returns 20 items per page | [optional] 
**Total** | Pointer to **int32** | @Description Total number of records matching the filter conditions @Description Used for calculating pagination metadata | [optional] 
**TotalPages** | Pointer to **int32** | @Description Total number of pages based on total records and page size @Description Calculated as ceil(total/page_size) | [optional] 

## Methods

### NewQueryFilterResponse

`func NewQueryFilterResponse() *QueryFilterResponse`

NewQueryFilterResponse instantiates a new QueryFilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryFilterResponseWithDefaults

`func NewQueryFilterResponseWithDefaults() *QueryFilterResponse`

NewQueryFilterResponseWithDefaults instantiates a new QueryFilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *QueryFilterResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueryFilterResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueryFilterResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *QueryFilterResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPage

`func (o *QueryFilterResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *QueryFilterResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *QueryFilterResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *QueryFilterResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *QueryFilterResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *QueryFilterResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *QueryFilterResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *QueryFilterResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotal

`func (o *QueryFilterResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryFilterResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryFilterResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryFilterResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalPages

`func (o *QueryFilterResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *QueryFilterResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *QueryFilterResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *QueryFilterResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


