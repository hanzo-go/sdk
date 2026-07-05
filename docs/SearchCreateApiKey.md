# SearchCreateApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Actions** | **[]string** | Actions like search, documents.add, indexes.create, etc. | 
**Indexes** | **[]string** | Index UIDs this key can access (* for all) | 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSearchCreateApiKey

`func NewSearchCreateApiKey(actions []string, indexes []string, ) *SearchCreateApiKey`

NewSearchCreateApiKey instantiates a new SearchCreateApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCreateApiKeyWithDefaults

`func NewSearchCreateApiKeyWithDefaults() *SearchCreateApiKey`

NewSearchCreateApiKeyWithDefaults instantiates a new SearchCreateApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SearchCreateApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchCreateApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchCreateApiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchCreateApiKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SearchCreateApiKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchCreateApiKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchCreateApiKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchCreateApiKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUid

`func (o *SearchCreateApiKey) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SearchCreateApiKey) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SearchCreateApiKey) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SearchCreateApiKey) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetActions

`func (o *SearchCreateApiKey) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *SearchCreateApiKey) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *SearchCreateApiKey) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetIndexes

`func (o *SearchCreateApiKey) GetIndexes() []string`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *SearchCreateApiKey) GetIndexesOk() (*[]string, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *SearchCreateApiKey) SetIndexes(v []string)`

SetIndexes sets Indexes field to given value.


### GetExpiresAt

`func (o *SearchCreateApiKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SearchCreateApiKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SearchCreateApiKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SearchCreateApiKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


