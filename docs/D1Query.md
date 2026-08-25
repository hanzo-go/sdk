# D1Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Params** | Pointer to **[]interface{}** | Params are the statement&#39;s bound values, in the order its &#x60;?&#x60; placeholders appear — a string, a number, a boolean or null, whatever the column takes. Absent means the statement carries no placeholders; bind values here rather than interpolating them into the statement. | [optional] 
**Sql** | Pointer to **string** | SQL is the statement to run. Blank (or absent) is refused before anything reaches D1. | [optional] 

## Methods

### NewD1Query

`func NewD1Query() *D1Query`

NewD1Query instantiates a new D1Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewD1QueryWithDefaults

`func NewD1QueryWithDefaults() *D1Query`

NewD1QueryWithDefaults instantiates a new D1Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParams

`func (o *D1Query) GetParams() []interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *D1Query) GetParamsOk() (*[]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *D1Query) SetParams(v []interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *D1Query) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSql

`func (o *D1Query) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *D1Query) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *D1Query) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *D1Query) HasSql() bool`

HasSql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


