# UserOptinReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | Handle is the display name shown on a listed row: 1-40 characters of letters, digits, space, dot, underscore, apostrophe or hyphen. Left empty on a listing opt-in it defaults to the caller&#39;s username. | [optional] 
**Listed** | Pointer to **bool** | Listed publishes the caller&#39;s row to other viewers of the board when true, and anonymizes it when false. | [optional] 

## Methods

### NewUserOptinReq

`func NewUserOptinReq() *UserOptinReq`

NewUserOptinReq instantiates a new UserOptinReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOptinReqWithDefaults

`func NewUserOptinReqWithDefaults() *UserOptinReq`

NewUserOptinReqWithDefaults instantiates a new UserOptinReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *UserOptinReq) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UserOptinReq) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UserOptinReq) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UserOptinReq) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetListed

`func (o *UserOptinReq) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *UserOptinReq) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *UserOptinReq) SetListed(v bool)`

SetListed sets Listed field to given value.

### HasListed

`func (o *UserOptinReq) HasListed() bool`

HasListed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


