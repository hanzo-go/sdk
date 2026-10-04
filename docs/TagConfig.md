# TagConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** |  | [optional] 
**Consent** | Pointer to [**Decision**](Decision.md) |  | [optional] 
**Tags** | Pointer to [**[]BrowserTagOut**](BrowserTagOut.md) |  | [optional] 

## Methods

### NewTagConfig

`func NewTagConfig() *TagConfig`

NewTagConfig instantiates a new TagConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagConfigWithDefaults

`func NewTagConfigWithDefaults() *TagConfig`

NewTagConfigWithDefaults instantiates a new TagConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *TagConfig) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *TagConfig) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *TagConfig) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *TagConfig) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetConsent

`func (o *TagConfig) GetConsent() Decision`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *TagConfig) GetConsentOk() (*Decision, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *TagConfig) SetConsent(v Decision)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *TagConfig) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetTags

`func (o *TagConfig) GetTags() []BrowserTagOut`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagConfig) GetTagsOk() (*[]BrowserTagOut, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagConfig) SetTags(v []BrowserTagOut)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TagConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


