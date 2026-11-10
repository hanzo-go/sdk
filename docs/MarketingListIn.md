# MarketingListIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** | Company is the sender&#39;s legal name in the footer; at most 200 characters. | [optional] 
**Description** | Pointer to **string** | Description says what subscribers receive; at most 2000 characters. | [optional] 
**DoubleOptIn** | Pointer to **bool** | DoubleOptIn is whether a public subscription must be confirmed; a new list asks for confirmation. | [optional] 
**Footer** | Pointer to **string** | Footer is text closing every mail of the list; at most 2000 characters. | [optional] 
**FromAddress** | Pointer to **string** | FromAddress is the sender&#39;s address. | [optional] 
**FromName** | Pointer to **string** | FromName is the sender&#39;s name; at most 200 characters. | [optional] 
**Id** | Pointer to **string** | ID is the list to change, from the path; ignored on create. | [optional] 
**Language** | Pointer to **string** | Language is a BCP 47 tag; a new list takes the platform&#39;s default. | [optional] 
**Name** | Pointer to **string** | Name is what the list is called. Required on create; at most 200 characters. | [optional] 
**Postal** | Pointer to **string** | Postal is the sender&#39;s physical postal address; at most 500 characters. | [optional] 
**Provider** | Pointer to **string** | Provider is ses, sendgrid, or empty for the org&#39;s own. | [optional] 
**ReplyTo** | Pointer to **string** | ReplyTo is where replies go. | [optional] 
**Slug** | Pointer to **string** | Slug is the list&#39;s name in a URL: lower-case letters, digits and hyphens, at most 63. On create it defaults to the name, made one. | [optional] 
**Visibility** | Pointer to **string** | Visibility is public or private; a new list is private. | [optional] 

## Methods

### NewMarketingListIn

`func NewMarketingListIn() *MarketingListIn`

NewMarketingListIn instantiates a new MarketingListIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingListInWithDefaults

`func NewMarketingListInWithDefaults() *MarketingListIn`

NewMarketingListInWithDefaults instantiates a new MarketingListIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *MarketingListIn) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *MarketingListIn) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *MarketingListIn) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *MarketingListIn) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDescription

`func (o *MarketingListIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketingListIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketingListIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketingListIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDoubleOptIn

`func (o *MarketingListIn) GetDoubleOptIn() bool`

GetDoubleOptIn returns the DoubleOptIn field if non-nil, zero value otherwise.

### GetDoubleOptInOk

`func (o *MarketingListIn) GetDoubleOptInOk() (*bool, bool)`

GetDoubleOptInOk returns a tuple with the DoubleOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleOptIn

`func (o *MarketingListIn) SetDoubleOptIn(v bool)`

SetDoubleOptIn sets DoubleOptIn field to given value.

### HasDoubleOptIn

`func (o *MarketingListIn) HasDoubleOptIn() bool`

HasDoubleOptIn returns a boolean if a field has been set.

### GetFooter

`func (o *MarketingListIn) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *MarketingListIn) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *MarketingListIn) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *MarketingListIn) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetFromAddress

`func (o *MarketingListIn) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MarketingListIn) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MarketingListIn) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MarketingListIn) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetFromName

`func (o *MarketingListIn) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *MarketingListIn) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *MarketingListIn) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *MarketingListIn) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetId

`func (o *MarketingListIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingListIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingListIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingListIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *MarketingListIn) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MarketingListIn) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MarketingListIn) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MarketingListIn) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetName

`func (o *MarketingListIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingListIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingListIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingListIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostal

`func (o *MarketingListIn) GetPostal() string`

GetPostal returns the Postal field if non-nil, zero value otherwise.

### GetPostalOk

`func (o *MarketingListIn) GetPostalOk() (*string, bool)`

GetPostalOk returns a tuple with the Postal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostal

`func (o *MarketingListIn) SetPostal(v string)`

SetPostal sets Postal field to given value.

### HasPostal

`func (o *MarketingListIn) HasPostal() bool`

HasPostal returns a boolean if a field has been set.

### GetProvider

`func (o *MarketingListIn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MarketingListIn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MarketingListIn) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *MarketingListIn) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReplyTo

`func (o *MarketingListIn) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MarketingListIn) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *MarketingListIn) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *MarketingListIn) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSlug

`func (o *MarketingListIn) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *MarketingListIn) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *MarketingListIn) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *MarketingListIn) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetVisibility

`func (o *MarketingListIn) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MarketingListIn) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MarketingListIn) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *MarketingListIn) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


