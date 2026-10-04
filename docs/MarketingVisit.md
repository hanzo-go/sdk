# MarketingVisit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | Pointer to **string** | Cid is the Google Analytics client id, the last two fields of _ga. | [optional] 
**Clicked** | Pointer to **int64** | Clicked is when the page captured the click, in unix seconds. | [optional] 
**Consent** | Pointer to **string** | Consent is the Cookie Preferences the visitor granted, comma separated: any of analytics, marketing and ads. Empty is a visitor who refused them all, which is a choice; absent is one who stated none. | [optional] 
**Fbc** | Pointer to **string** | Fbc is Meta&#39;s click cookie, fb.1.&lt;ms&gt;.&lt;fbclid&gt;. | [optional] 
**Fbclid** | Pointer to **string** | Fbclid is the Meta click id the visitor arrived with. | [optional] 
**Fbp** | Pointer to **string** | Fbp is Meta&#39;s browser cookie, fb.1.&lt;ms&gt;.&lt;random&gt;. | [optional] 
**Gbraid** | Pointer to **string** | Gbraid is the Google Ads click id for an iOS app click. | [optional] 
**Gclid** | Pointer to **string** | Gclid is the Google Ads click id the visitor arrived with. | [optional] 
**Tags** | Pointer to **string** | Tags is which of the page&#39;s ad tags already saw this lead: ga, meta. | [optional] 
**Url** | Pointer to **string** | URL is the page the form was filed from, in full. | [optional] 
**UserAgent** | Pointer to **string** | UserAgent is the visitor&#39;s browser, as navigator.userAgent reads. | [optional] 
**Wbraid** | Pointer to **string** | Wbraid is the Google Ads click id for an iOS web click. | [optional] 

## Methods

### NewMarketingVisit

`func NewMarketingVisit() *MarketingVisit`

NewMarketingVisit instantiates a new MarketingVisit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingVisitWithDefaults

`func NewMarketingVisitWithDefaults() *MarketingVisit`

NewMarketingVisitWithDefaults instantiates a new MarketingVisit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *MarketingVisit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *MarketingVisit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *MarketingVisit) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *MarketingVisit) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetClicked

`func (o *MarketingVisit) GetClicked() int64`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *MarketingVisit) GetClickedOk() (*int64, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *MarketingVisit) SetClicked(v int64)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *MarketingVisit) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetConsent

`func (o *MarketingVisit) GetConsent() string`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *MarketingVisit) GetConsentOk() (*string, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *MarketingVisit) SetConsent(v string)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *MarketingVisit) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetFbc

`func (o *MarketingVisit) GetFbc() string`

GetFbc returns the Fbc field if non-nil, zero value otherwise.

### GetFbcOk

`func (o *MarketingVisit) GetFbcOk() (*string, bool)`

GetFbcOk returns a tuple with the Fbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFbc

`func (o *MarketingVisit) SetFbc(v string)`

SetFbc sets Fbc field to given value.

### HasFbc

`func (o *MarketingVisit) HasFbc() bool`

HasFbc returns a boolean if a field has been set.

### GetFbclid

`func (o *MarketingVisit) GetFbclid() string`

GetFbclid returns the Fbclid field if non-nil, zero value otherwise.

### GetFbclidOk

`func (o *MarketingVisit) GetFbclidOk() (*string, bool)`

GetFbclidOk returns a tuple with the Fbclid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFbclid

`func (o *MarketingVisit) SetFbclid(v string)`

SetFbclid sets Fbclid field to given value.

### HasFbclid

`func (o *MarketingVisit) HasFbclid() bool`

HasFbclid returns a boolean if a field has been set.

### GetFbp

`func (o *MarketingVisit) GetFbp() string`

GetFbp returns the Fbp field if non-nil, zero value otherwise.

### GetFbpOk

`func (o *MarketingVisit) GetFbpOk() (*string, bool)`

GetFbpOk returns a tuple with the Fbp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFbp

`func (o *MarketingVisit) SetFbp(v string)`

SetFbp sets Fbp field to given value.

### HasFbp

`func (o *MarketingVisit) HasFbp() bool`

HasFbp returns a boolean if a field has been set.

### GetGbraid

`func (o *MarketingVisit) GetGbraid() string`

GetGbraid returns the Gbraid field if non-nil, zero value otherwise.

### GetGbraidOk

`func (o *MarketingVisit) GetGbraidOk() (*string, bool)`

GetGbraidOk returns a tuple with the Gbraid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbraid

`func (o *MarketingVisit) SetGbraid(v string)`

SetGbraid sets Gbraid field to given value.

### HasGbraid

`func (o *MarketingVisit) HasGbraid() bool`

HasGbraid returns a boolean if a field has been set.

### GetGclid

`func (o *MarketingVisit) GetGclid() string`

GetGclid returns the Gclid field if non-nil, zero value otherwise.

### GetGclidOk

`func (o *MarketingVisit) GetGclidOk() (*string, bool)`

GetGclidOk returns a tuple with the Gclid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGclid

`func (o *MarketingVisit) SetGclid(v string)`

SetGclid sets Gclid field to given value.

### HasGclid

`func (o *MarketingVisit) HasGclid() bool`

HasGclid returns a boolean if a field has been set.

### GetTags

`func (o *MarketingVisit) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MarketingVisit) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MarketingVisit) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MarketingVisit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUrl

`func (o *MarketingVisit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MarketingVisit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MarketingVisit) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MarketingVisit) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserAgent

`func (o *MarketingVisit) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *MarketingVisit) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *MarketingVisit) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *MarketingVisit) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetWbraid

`func (o *MarketingVisit) GetWbraid() string`

GetWbraid returns the Wbraid field if non-nil, zero value otherwise.

### GetWbraidOk

`func (o *MarketingVisit) GetWbraidOk() (*string, bool)`

GetWbraidOk returns a tuple with the Wbraid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbraid

`func (o *MarketingVisit) SetWbraid(v string)`

SetWbraid sets Wbraid field to given value.

### HasWbraid

`func (o *MarketingVisit) HasWbraid() bool`

HasWbraid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


