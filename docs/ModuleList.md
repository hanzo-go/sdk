# ModuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ModuleInfo**](ModuleInfo.md) | Data is every module compiled into this binary, with the DocTypes it installs. | [optional] 

## Methods

### NewModuleList

`func NewModuleList() *ModuleList`

NewModuleList instantiates a new ModuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleListWithDefaults

`func NewModuleListWithDefaults() *ModuleList`

NewModuleListWithDefaults instantiates a new ModuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModuleList) GetData() []ModuleInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModuleList) GetDataOk() (*[]ModuleInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModuleList) SetData(v []ModuleInfo)`

SetData sets Data field to given value.

### HasData

`func (o *ModuleList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


