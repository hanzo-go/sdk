# BaseBaseView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int64** | Bytes is the store&#39;s size on disk, present only once it exists. It is what this Base occupies, not a quota. | [optional] 
**Exists** | Pointer to **bool** | Exists reports whether this Base&#39;s store has been provisioned. False is an org nobody has stored anything for yet, which is a state to name rather than an error: the store is created the first time anything writes. | [optional] 
**Org** | Pointer to **string** | Org is the org this Base belongs to. It is the address every other Base call is scoped by, and a Base has no name of its own. | [optional] 

## Methods

### NewBaseBaseView

`func NewBaseBaseView() *BaseBaseView`

NewBaseBaseView instantiates a new BaseBaseView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseBaseViewWithDefaults

`func NewBaseBaseViewWithDefaults() *BaseBaseView`

NewBaseBaseViewWithDefaults instantiates a new BaseBaseView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *BaseBaseView) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *BaseBaseView) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *BaseBaseView) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *BaseBaseView) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetExists

`func (o *BaseBaseView) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *BaseBaseView) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *BaseBaseView) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *BaseBaseView) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetOrg

`func (o *BaseBaseView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BaseBaseView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BaseBaseView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BaseBaseView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


