# MarketingSubscribeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address to subscribe. Required, at most 320 characters, and it must look like one address. It is mailed a link to confirm, and nothing is recorded until the link is pressed. | [optional] 
**Locale** | Pointer to **string** | Locale is the reader&#39;s language as a BCP 47 tag, such as en or pt-BR. Optional; any other value is recorded as none. The mail is in English. | [optional] 
**Source** | Pointer to **string** | Source is where the form was, such as hanzo.ai or hanzo.ai/blog. Optional; clipped at 200 characters and recorded with the subscription once it is confirmed. | [optional] 

## Methods

### NewMarketingSubscribeIn

`func NewMarketingSubscribeIn() *MarketingSubscribeIn`

NewMarketingSubscribeIn instantiates a new MarketingSubscribeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSubscribeInWithDefaults

`func NewMarketingSubscribeInWithDefaults() *MarketingSubscribeIn`

NewMarketingSubscribeInWithDefaults instantiates a new MarketingSubscribeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *MarketingSubscribeIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MarketingSubscribeIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MarketingSubscribeIn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MarketingSubscribeIn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLocale

`func (o *MarketingSubscribeIn) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MarketingSubscribeIn) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MarketingSubscribeIn) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MarketingSubscribeIn) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetSource

`func (o *MarketingSubscribeIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketingSubscribeIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketingSubscribeIn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketingSubscribeIn) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


