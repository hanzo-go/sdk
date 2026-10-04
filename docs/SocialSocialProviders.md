# SocialSocialProviders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SocialSocialProvider**](SocialSocialProvider.md) | Data is one row per supported network, in the product&#39;s fixed order: x, facebook, instagram, linkedin, tiktok, youtube, threads. | [optional] 

## Methods

### NewSocialSocialProviders

`func NewSocialSocialProviders() *SocialSocialProviders`

NewSocialSocialProviders instantiates a new SocialSocialProviders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialSocialProvidersWithDefaults

`func NewSocialSocialProvidersWithDefaults() *SocialSocialProviders`

NewSocialSocialProvidersWithDefaults instantiates a new SocialSocialProviders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SocialSocialProviders) GetData() []SocialSocialProvider`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SocialSocialProviders) GetDataOk() (*[]SocialSocialProvider, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SocialSocialProviders) SetData(v []SocialSocialProvider)`

SetData sets Data field to given value.

### HasData

`func (o *SocialSocialProviders) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


