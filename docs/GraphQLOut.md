# GraphQLOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Errors** | Pointer to [**[]GraphQLError**](GraphQLError.md) |  | [optional] 

## Methods

### NewGraphQLOut

`func NewGraphQLOut() *GraphQLOut`

NewGraphQLOut instantiates a new GraphQLOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLOutWithDefaults

`func NewGraphQLOutWithDefaults() *GraphQLOut`

NewGraphQLOutWithDefaults instantiates a new GraphQLOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GraphQLOut) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GraphQLOut) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GraphQLOut) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GraphQLOut) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GraphQLOut) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GraphQLOut) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetErrors

`func (o *GraphQLOut) GetErrors() []GraphQLError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GraphQLOut) GetErrorsOk() (*[]GraphQLError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GraphQLOut) SetErrors(v []GraphQLError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GraphQLOut) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


