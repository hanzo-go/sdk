# BaseView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int32** | Bytes is the store&#39;s size on disk, present only once it exists. It is what this Base occupies, not a quota. | [optional] 
**Exists** | Pointer to **bool** | Exists reports whether this Base&#39;s store has been provisioned. False is an org nobody has stored anything for yet, which is a state to name rather than an error: the store is created the first time anything writes. | [optional] 
**Org** | Pointer to **string** | Org is the org this Base belongs to. It is the address every other Base call is scoped by, and a Base has no name of its own. | [optional] 

## Methods

### NewBaseView

`func NewBaseView() *BaseView`

NewBaseView instantiates a new BaseView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseViewWithDefaults

`func NewBaseViewWithDefaults() *BaseView`

NewBaseViewWithDefaults instantiates a new BaseView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *BaseView) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *BaseView) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *BaseView) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *BaseView) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetExists

`func (o *BaseView) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *BaseView) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *BaseView) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *BaseView) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetOrg

`func (o *BaseView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BaseView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BaseView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BaseView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


