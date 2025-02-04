# QueryEntityWithRelations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MainEntity** | Pointer to **map[string]interface{}** |  | [optional] 
**Relations** | Pointer to **map[string][]map[string]interface{}** |  | [optional] 

## Methods

### NewQueryEntityWithRelations

`func NewQueryEntityWithRelations() *QueryEntityWithRelations`

NewQueryEntityWithRelations instantiates a new QueryEntityWithRelations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryEntityWithRelationsWithDefaults

`func NewQueryEntityWithRelationsWithDefaults() *QueryEntityWithRelations`

NewQueryEntityWithRelationsWithDefaults instantiates a new QueryEntityWithRelations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMainEntity

`func (o *QueryEntityWithRelations) GetMainEntity() map[string]interface{}`

GetMainEntity returns the MainEntity field if non-nil, zero value otherwise.

### GetMainEntityOk

`func (o *QueryEntityWithRelations) GetMainEntityOk() (*map[string]interface{}, bool)`

GetMainEntityOk returns a tuple with the MainEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainEntity

`func (o *QueryEntityWithRelations) SetMainEntity(v map[string]interface{})`

SetMainEntity sets MainEntity field to given value.

### HasMainEntity

`func (o *QueryEntityWithRelations) HasMainEntity() bool`

HasMainEntity returns a boolean if a field has been set.

### GetRelations

`func (o *QueryEntityWithRelations) GetRelations() map[string][]map[string]interface{}`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *QueryEntityWithRelations) GetRelationsOk() (*map[string][]map[string]interface{}, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *QueryEntityWithRelations) SetRelations(v map[string][]map[string]interface{})`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *QueryEntityWithRelations) HasRelations() bool`

HasRelations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


