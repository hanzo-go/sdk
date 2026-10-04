# PromptVersionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when this revision was appended, RFC 3339 UTC. | [optional] 
**Type** | Pointer to **string** | Type is the kind this revision was written with, which may differ from the current one. | [optional] 
**Version** | Pointer to **int64** | Version is this revision&#39;s number, 1 for the first. Numbers are dense and never reused: deleting the prompt drops the whole history with it. | [optional] 

## Methods

### NewPromptVersionView

`func NewPromptVersionView() *PromptVersionView`

NewPromptVersionView instantiates a new PromptVersionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptVersionViewWithDefaults

`func NewPromptVersionViewWithDefaults() *PromptVersionView`

NewPromptVersionViewWithDefaults instantiates a new PromptVersionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PromptVersionView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PromptVersionView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PromptVersionView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PromptVersionView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetType

`func (o *PromptVersionView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptVersionView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptVersionView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptVersionView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *PromptVersionView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PromptVersionView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PromptVersionView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PromptVersionView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


