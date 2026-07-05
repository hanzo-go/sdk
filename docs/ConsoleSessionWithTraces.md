# ConsoleSessionWithTraces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Traces** | Pointer to [**[]ConsoleTrace**](ConsoleTrace.md) |  | [optional] 

## Methods

### NewConsoleSessionWithTraces

`func NewConsoleSessionWithTraces() *ConsoleSessionWithTraces`

NewConsoleSessionWithTraces instantiates a new ConsoleSessionWithTraces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleSessionWithTracesWithDefaults

`func NewConsoleSessionWithTracesWithDefaults() *ConsoleSessionWithTraces`

NewConsoleSessionWithTracesWithDefaults instantiates a new ConsoleSessionWithTraces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleSessionWithTraces) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleSessionWithTraces) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleSessionWithTraces) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleSessionWithTraces) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConsoleSessionWithTraces) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConsoleSessionWithTraces) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConsoleSessionWithTraces) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConsoleSessionWithTraces) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProjectId

`func (o *ConsoleSessionWithTraces) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ConsoleSessionWithTraces) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ConsoleSessionWithTraces) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ConsoleSessionWithTraces) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetEnvironment

`func (o *ConsoleSessionWithTraces) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConsoleSessionWithTraces) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConsoleSessionWithTraces) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConsoleSessionWithTraces) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetTraces

`func (o *ConsoleSessionWithTraces) GetTraces() []ConsoleTrace`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *ConsoleSessionWithTraces) GetTracesOk() (*[]ConsoleTrace, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *ConsoleSessionWithTraces) SetTraces(v []ConsoleTrace)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *ConsoleSessionWithTraces) HasTraces() bool`

HasTraces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


