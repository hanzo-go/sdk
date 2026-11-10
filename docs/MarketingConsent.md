# MarketingConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is when it was given, unix seconds; for an import, the date the file carried when it carried one. | [optional] 
**Basis** | Pointer to **string** | Basis is explicit, legitimate_interest or unverified. | [optional] 
**Dnt** | Pointer to **bool** | DNT records that the request carried Do Not Track (DNT: 1). | [optional] 
**Gpc** | Pointer to **bool** | GPC records that the request carried Global Privacy Control (Sec-GPC: 1). | [optional] 
**Ip** | Pointer to **string** | IP is the network address it was given from, when it was given here. | [optional] 
**Language** | Pointer to **string** | Language is the language it was asked in. | [optional] 
**Method** | Pointer to **string** | Method is how it was given: double_opt_in, single_opt_in, import, api, preferences or commerce. | [optional] 
**Page** | Pointer to **string** | Page is the form or page it was given on. | [optional] 

## Methods

### NewMarketingConsent

`func NewMarketingConsent() *MarketingConsent`

NewMarketingConsent instantiates a new MarketingConsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingConsentWithDefaults

`func NewMarketingConsentWithDefaults() *MarketingConsent`

NewMarketingConsentWithDefaults instantiates a new MarketingConsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketingConsent) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketingConsent) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketingConsent) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketingConsent) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBasis

`func (o *MarketingConsent) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *MarketingConsent) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *MarketingConsent) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *MarketingConsent) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetDnt

`func (o *MarketingConsent) GetDnt() bool`

GetDnt returns the Dnt field if non-nil, zero value otherwise.

### GetDntOk

`func (o *MarketingConsent) GetDntOk() (*bool, bool)`

GetDntOk returns a tuple with the Dnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnt

`func (o *MarketingConsent) SetDnt(v bool)`

SetDnt sets Dnt field to given value.

### HasDnt

`func (o *MarketingConsent) HasDnt() bool`

HasDnt returns a boolean if a field has been set.

### GetGpc

`func (o *MarketingConsent) GetGpc() bool`

GetGpc returns the Gpc field if non-nil, zero value otherwise.

### GetGpcOk

`func (o *MarketingConsent) GetGpcOk() (*bool, bool)`

GetGpcOk returns a tuple with the Gpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpc

`func (o *MarketingConsent) SetGpc(v bool)`

SetGpc sets Gpc field to given value.

### HasGpc

`func (o *MarketingConsent) HasGpc() bool`

HasGpc returns a boolean if a field has been set.

### GetIp

`func (o *MarketingConsent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *MarketingConsent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *MarketingConsent) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *MarketingConsent) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLanguage

`func (o *MarketingConsent) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MarketingConsent) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MarketingConsent) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MarketingConsent) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMethod

`func (o *MarketingConsent) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MarketingConsent) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MarketingConsent) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *MarketingConsent) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPage

`func (o *MarketingConsent) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *MarketingConsent) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *MarketingConsent) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *MarketingConsent) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


