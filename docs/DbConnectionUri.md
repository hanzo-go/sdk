# DbConnectionUri

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionUri** | Pointer to **string** | PostgreSQL connection string | [optional] 
**ConnectionParameters** | Pointer to [**DbConnectionUriConnectionParameters**](DbConnectionUriConnectionParameters.md) |  | [optional] 

## Methods

### NewDbConnectionUri

`func NewDbConnectionUri() *DbConnectionUri`

NewDbConnectionUri instantiates a new DbConnectionUri object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbConnectionUriWithDefaults

`func NewDbConnectionUriWithDefaults() *DbConnectionUri`

NewDbConnectionUriWithDefaults instantiates a new DbConnectionUri object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionUri

`func (o *DbConnectionUri) GetConnectionUri() string`

GetConnectionUri returns the ConnectionUri field if non-nil, zero value otherwise.

### GetConnectionUriOk

`func (o *DbConnectionUri) GetConnectionUriOk() (*string, bool)`

GetConnectionUriOk returns a tuple with the ConnectionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUri

`func (o *DbConnectionUri) SetConnectionUri(v string)`

SetConnectionUri sets ConnectionUri field to given value.

### HasConnectionUri

`func (o *DbConnectionUri) HasConnectionUri() bool`

HasConnectionUri returns a boolean if a field has been set.

### GetConnectionParameters

`func (o *DbConnectionUri) GetConnectionParameters() DbConnectionUriConnectionParameters`

GetConnectionParameters returns the ConnectionParameters field if non-nil, zero value otherwise.

### GetConnectionParametersOk

`func (o *DbConnectionUri) GetConnectionParametersOk() (*DbConnectionUriConnectionParameters, bool)`

GetConnectionParametersOk returns a tuple with the ConnectionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionParameters

`func (o *DbConnectionUri) SetConnectionParameters(v DbConnectionUriConnectionParameters)`

SetConnectionParameters sets ConnectionParameters field to given value.

### HasConnectionParameters

`func (o *DbConnectionUri) HasConnectionParameters() bool`

HasConnectionParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


