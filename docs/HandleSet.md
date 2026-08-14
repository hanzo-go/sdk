# HandleSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | Handle is the display name as STORED, echoed back after trimming. Empty means the caller opted out: it keeps its rank and still sees its own row, it is just no longer listed to anyone else. | [optional] 

## Methods

### NewHandleSet

`func NewHandleSet() *HandleSet`

NewHandleSet instantiates a new HandleSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandleSetWithDefaults

`func NewHandleSetWithDefaults() *HandleSet`

NewHandleSetWithDefaults instantiates a new HandleSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *HandleSet) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *HandleSet) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *HandleSet) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *HandleSet) HasHandle() bool`

HasHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


