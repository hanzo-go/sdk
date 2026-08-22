# GraphQLIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationName** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGraphQLIn

`func NewGraphQLIn() *GraphQLIn`

NewGraphQLIn instantiates a new GraphQLIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLInWithDefaults

`func NewGraphQLInWithDefaults() *GraphQLIn`

NewGraphQLInWithDefaults instantiates a new GraphQLIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationName

`func (o *GraphQLIn) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *GraphQLIn) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *GraphQLIn) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *GraphQLIn) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### GetQuery

`func (o *GraphQLIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GraphQLIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GraphQLIn) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GraphQLIn) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetVariables

`func (o *GraphQLIn) GetVariables() interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *GraphQLIn) GetVariablesOk() (*interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *GraphQLIn) SetVariables(v interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *GraphQLIn) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *GraphQLIn) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *GraphQLIn) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


