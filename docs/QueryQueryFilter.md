# QueryQueryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expressions** | Pointer to **[]map[string]interface{}** | @Description Array of expressions to filter entities @Description Each expression can be a FieldExpression, LogicalExpression, or RelationshipExpression | [optional] 

## Methods

### NewQueryQueryFilter

`func NewQueryQueryFilter() *QueryQueryFilter`

NewQueryQueryFilter instantiates a new QueryQueryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryQueryFilterWithDefaults

`func NewQueryQueryFilterWithDefaults() *QueryQueryFilter`

NewQueryQueryFilterWithDefaults instantiates a new QueryQueryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpressions

`func (o *QueryQueryFilter) GetExpressions() []map[string]interface{}`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *QueryQueryFilter) GetExpressionsOk() (*[]map[string]interface{}, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *QueryQueryFilter) SetExpressions(v []map[string]interface{})`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *QueryQueryFilter) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


