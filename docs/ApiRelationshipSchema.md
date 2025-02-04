# ApiRelationshipSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the relationship | [optional] 
**RelatedModel** | Pointer to **string** | Name of the related model | [optional] 
**Type** | Pointer to **string** | Type of relationship (hasOne, hasMany) | [optional] 

## Methods

### NewApiRelationshipSchema

`func NewApiRelationshipSchema() *ApiRelationshipSchema`

NewApiRelationshipSchema instantiates a new ApiRelationshipSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRelationshipSchemaWithDefaults

`func NewApiRelationshipSchemaWithDefaults() *ApiRelationshipSchema`

NewApiRelationshipSchemaWithDefaults instantiates a new ApiRelationshipSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiRelationshipSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiRelationshipSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiRelationshipSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiRelationshipSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelatedModel

`func (o *ApiRelationshipSchema) GetRelatedModel() string`

GetRelatedModel returns the RelatedModel field if non-nil, zero value otherwise.

### GetRelatedModelOk

`func (o *ApiRelationshipSchema) GetRelatedModelOk() (*string, bool)`

GetRelatedModelOk returns a tuple with the RelatedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedModel

`func (o *ApiRelationshipSchema) SetRelatedModel(v string)`

SetRelatedModel sets RelatedModel field to given value.

### HasRelatedModel

`func (o *ApiRelationshipSchema) HasRelatedModel() bool`

HasRelatedModel returns a boolean if a field has been set.

### GetType

`func (o *ApiRelationshipSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiRelationshipSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiRelationshipSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiRelationshipSchema) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


