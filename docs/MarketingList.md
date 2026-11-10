# MarketingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** | Company is the sender&#39;s legal name in the footer; empty takes the org&#39;s. | [optional] 
**Counts** | Pointer to **map[string]int64** | Counts are the list&#39;s subscribers by status. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the list was created, unix seconds. | [optional] 
**Description** | Pointer to **string** | Description says what subscribers receive. | [optional] 
**DoubleOptIn** | Pointer to **bool** | DoubleOptIn is whether a public subscription waits for the address to confirm before it is subscribed. | [optional] 
**Footer** | Pointer to **string** | Footer is text closing every mail of the list, above the unsubscribe line. | [optional] 
**FromAddress** | Pointer to **string** | FromAddress is the sender&#39;s address; empty takes the org&#39;s. | [optional] 
**FromName** | Pointer to **string** | FromName is the sender&#39;s name on the list&#39;s mail; empty takes the org&#39;s. | [optional] 
**Id** | Pointer to **string** | ID is the list&#39;s id (\&quot;lst_\&quot; and 96 random bits in hex). | [optional] 
**Language** | Pointer to **string** | Language is the list&#39;s language, a BCP 47 tag. | [optional] 
**Name** | Pointer to **string** | Name is what the list is called. | [optional] 
**Postal** | Pointer to **string** | Postal is the sender&#39;s physical postal address, which every commercial mail carries (CAN-SPAM); empty takes the org&#39;s. | [optional] 
**Provider** | Pointer to **string** | Provider is ses or sendgrid; empty takes the org&#39;s. | [optional] 
**ReplyTo** | Pointer to **string** | ReplyTo is where replies go; empty takes the org&#39;s. | [optional] 
**Slug** | Pointer to **string** | Slug is the list&#39;s name in a URL, unique in the org: a public form may name the list by it (\&quot;newsletter\&quot;). | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the last change to the list itself, unix seconds. | [optional] 
**Visibility** | Pointer to **string** | Visibility is public, which the public subscribe and the preference page offer, or private, which only the org&#39;s own calls fill. | [optional] 

## Methods

### NewMarketingList

`func NewMarketingList() *MarketingList`

NewMarketingList instantiates a new MarketingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingListWithDefaults

`func NewMarketingListWithDefaults() *MarketingList`

NewMarketingListWithDefaults instantiates a new MarketingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *MarketingList) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *MarketingList) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *MarketingList) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *MarketingList) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCounts

`func (o *MarketingList) GetCounts() map[string]int64`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *MarketingList) GetCountsOk() (*map[string]int64, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *MarketingList) SetCounts(v map[string]int64)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *MarketingList) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketingList) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingList) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingList) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketingList) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *MarketingList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketingList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketingList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketingList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDoubleOptIn

`func (o *MarketingList) GetDoubleOptIn() bool`

GetDoubleOptIn returns the DoubleOptIn field if non-nil, zero value otherwise.

### GetDoubleOptInOk

`func (o *MarketingList) GetDoubleOptInOk() (*bool, bool)`

GetDoubleOptInOk returns a tuple with the DoubleOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleOptIn

`func (o *MarketingList) SetDoubleOptIn(v bool)`

SetDoubleOptIn sets DoubleOptIn field to given value.

### HasDoubleOptIn

`func (o *MarketingList) HasDoubleOptIn() bool`

HasDoubleOptIn returns a boolean if a field has been set.

### GetFooter

`func (o *MarketingList) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *MarketingList) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *MarketingList) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *MarketingList) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetFromAddress

`func (o *MarketingList) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MarketingList) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MarketingList) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MarketingList) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetFromName

`func (o *MarketingList) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *MarketingList) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *MarketingList) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *MarketingList) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetId

`func (o *MarketingList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *MarketingList) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MarketingList) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MarketingList) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MarketingList) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetName

`func (o *MarketingList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostal

`func (o *MarketingList) GetPostal() string`

GetPostal returns the Postal field if non-nil, zero value otherwise.

### GetPostalOk

`func (o *MarketingList) GetPostalOk() (*string, bool)`

GetPostalOk returns a tuple with the Postal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostal

`func (o *MarketingList) SetPostal(v string)`

SetPostal sets Postal field to given value.

### HasPostal

`func (o *MarketingList) HasPostal() bool`

HasPostal returns a boolean if a field has been set.

### GetProvider

`func (o *MarketingList) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MarketingList) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MarketingList) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *MarketingList) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReplyTo

`func (o *MarketingList) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MarketingList) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *MarketingList) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *MarketingList) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSlug

`func (o *MarketingList) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *MarketingList) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *MarketingList) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *MarketingList) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarketingList) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketingList) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketingList) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketingList) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisibility

`func (o *MarketingList) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MarketingList) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MarketingList) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *MarketingList) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


