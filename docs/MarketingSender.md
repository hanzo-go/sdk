# MarketingSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** | Company is the sender&#39;s legal name in the footer. | [optional] 
**FromAddress** | Pointer to **string** | FromAddress is the sender&#39;s address. | [optional] 
**FromName** | Pointer to **string** | FromName is the sender&#39;s name. | [optional] 
**Postal** | Pointer to **string** | Postal is the sender&#39;s physical postal address (CAN-SPAM). | [optional] 
**Provider** | Pointer to **string** | Provider is ses or sendgrid; empty takes the platform&#39;s default. | [optional] 
**ReplyTo** | Pointer to **string** | ReplyTo is where replies go. | [optional] 
**Site** | Pointer to **string** | Site is the org&#39;s public site, as a host name (\&quot;acme.com\&quot;), which serves /newsletter/confirm and /newsletter/preferences: the pages a confirmation link and a preference link open. Empty takes the brand&#39;s site when the org is a brand; an org with neither cannot run a double opt-in. | [optional] 

## Methods

### NewMarketingSender

`func NewMarketingSender() *MarketingSender`

NewMarketingSender instantiates a new MarketingSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSenderWithDefaults

`func NewMarketingSenderWithDefaults() *MarketingSender`

NewMarketingSenderWithDefaults instantiates a new MarketingSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *MarketingSender) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *MarketingSender) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *MarketingSender) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *MarketingSender) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetFromAddress

`func (o *MarketingSender) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MarketingSender) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MarketingSender) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MarketingSender) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetFromName

`func (o *MarketingSender) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *MarketingSender) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *MarketingSender) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *MarketingSender) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetPostal

`func (o *MarketingSender) GetPostal() string`

GetPostal returns the Postal field if non-nil, zero value otherwise.

### GetPostalOk

`func (o *MarketingSender) GetPostalOk() (*string, bool)`

GetPostalOk returns a tuple with the Postal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostal

`func (o *MarketingSender) SetPostal(v string)`

SetPostal sets Postal field to given value.

### HasPostal

`func (o *MarketingSender) HasPostal() bool`

HasPostal returns a boolean if a field has been set.

### GetProvider

`func (o *MarketingSender) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MarketingSender) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MarketingSender) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *MarketingSender) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReplyTo

`func (o *MarketingSender) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MarketingSender) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *MarketingSender) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *MarketingSender) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSite

`func (o *MarketingSender) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MarketingSender) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MarketingSender) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *MarketingSender) HasSite() bool`

HasSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


