# ApiModelSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]ApiFieldSchema**](ApiFieldSchema.md) | Fields contains information about each field in the model | [optional] 
**Relationships** | Pointer to [**[]ApiRelationshipSchema**](ApiRelationshipSchema.md) | Relationships contains information about model relationships | [optional] 

## Methods

### NewApiModelSchema

`func NewApiModelSchema() *ApiModelSchema`

NewApiModelSchema instantiates a new ApiModelSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiModelSchemaWithDefaults

`func NewApiModelSchemaWithDefaults() *ApiModelSchema`

NewApiModelSchemaWithDefaults instantiates a new ApiModelSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *ApiModelSchema) GetFields() []ApiFieldSchema`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ApiModelSchema) GetFieldsOk() (*[]ApiFieldSchema, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ApiModelSchema) SetFields(v []ApiFieldSchema)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ApiModelSchema) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ApiModelSchema) GetRelationships() []ApiRelationshipSchema`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ApiModelSchema) GetRelationshipsOk() (*[]ApiRelationshipSchema, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ApiModelSchema) SetRelationships(v []ApiRelationshipSchema)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ApiModelSchema) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


