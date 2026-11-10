# MarketingMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentAt** | Pointer to **string** | ConsentAt is a column holding when the person subscribed. | [optional] 
**Email** | Pointer to **string** | Email is the column holding the address. | [optional] 
**Fields** | Pointer to **map[string]string** | Fields maps a custom field key to the column holding it. | [optional] 
**First** | Pointer to **string** | First and Last are columns holding the name in two parts. | [optional] 
**Language** | Pointer to **string** | Language is a column holding a language or locale. | [optional] 
**Last** | Pointer to **string** | Last is the family name&#39;s column. | [optional] 
**Name** | Pointer to **string** | Name is a column holding the whole name. | [optional] 
**Status** | Pointer to **string** | Status is a column holding a status from another provider (subscribed, unsubscribed, cleaned, bounced, complained and their usual spellings). | [optional] 
**Tags** | Pointer to **string** | Tags is a column of tags separated by commas or semicolons. | [optional] 

## Methods

### NewMarketingMapping

`func NewMarketingMapping() *MarketingMapping`

NewMarketingMapping instantiates a new MarketingMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingMappingWithDefaults

`func NewMarketingMappingWithDefaults() *MarketingMapping`

NewMarketingMappingWithDefaults instantiates a new MarketingMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentAt

`func (o *MarketingMapping) GetConsentAt() string`

GetConsentAt returns the ConsentAt field if non-nil, zero value otherwise.

### GetConsentAtOk

`func (o *MarketingMapping) GetConsentAtOk() (*string, bool)`

GetConsentAtOk returns a tuple with the ConsentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentAt

`func (o *MarketingMapping) SetConsentAt(v string)`

SetConsentAt sets ConsentAt field to given value.

### HasConsentAt

`func (o *MarketingMapping) HasConsentAt() bool`

HasConsentAt returns a boolean if a field has been set.

### GetEmail

`func (o *MarketingMapping) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MarketingMapping) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MarketingMapping) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MarketingMapping) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFields

`func (o *MarketingMapping) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MarketingMapping) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MarketingMapping) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MarketingMapping) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFirst

`func (o *MarketingMapping) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *MarketingMapping) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *MarketingMapping) SetFirst(v string)`

SetFirst sets First field to given value.

### HasFirst

`func (o *MarketingMapping) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLanguage

`func (o *MarketingMapping) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MarketingMapping) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MarketingMapping) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MarketingMapping) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLast

`func (o *MarketingMapping) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *MarketingMapping) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *MarketingMapping) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *MarketingMapping) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetName

`func (o *MarketingMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingMapping) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingMapping) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *MarketingMapping) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketingMapping) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketingMapping) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketingMapping) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *MarketingMapping) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MarketingMapping) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MarketingMapping) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MarketingMapping) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


