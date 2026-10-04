# EnvironmentEnvList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EnvironmentEnvironment**](EnvironmentEnvironment.md) | Data is one entry per codebase that has an environment, by name. | [optional] 

## Methods

### NewEnvironmentEnvList

`func NewEnvironmentEnvList() *EnvironmentEnvList`

NewEnvironmentEnvList instantiates a new EnvironmentEnvList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEnvListWithDefaults

`func NewEnvironmentEnvListWithDefaults() *EnvironmentEnvList`

NewEnvironmentEnvListWithDefaults instantiates a new EnvironmentEnvList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnvironmentEnvList) GetData() []EnvironmentEnvironment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnvironmentEnvList) GetDataOk() (*[]EnvironmentEnvironment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnvironmentEnvList) SetData(v []EnvironmentEnvironment)`

SetData sets Data field to given value.

### HasData

`func (o *EnvironmentEnvList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


