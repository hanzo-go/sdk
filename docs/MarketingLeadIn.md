# MarketingLeadIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** | Company is the person&#39;s company. Optional; clipped at 200 characters. | [optional] 
**Country** | Pointer to **string** | Country is where the person is. Optional; clipped at 200 characters. | [optional] 
**Email** | Pointer to **string** | Email is where sales replies: the person&#39;s work email. Required, and it must look like an address. It is not verified. | [optional] 
**Name** | Pointer to **string** | Name is the person&#39;s name. Optional; clipped at 200 characters. | [optional] 
**Need** | Pointer to **string** | Need is what the person wants from sales, in their own words. Optional; clipped at 4 KiB. | [optional] 
**Role** | Pointer to **string** | Role is the person&#39;s job title. Optional; clipped at 200 characters. | [optional] 
**Size** | Pointer to **string** | Size is the company&#39;s headcount band: 1-49, 50-249, 250-999, 1000-4999 or 5000+. Any other value is recorded as unstated rather than refused. | [optional] 
**Source** | Pointer to **string** | Source is the page the form was filed from, such as hanzo.ai/contact-sales. Optional; clipped at 200 characters. | [optional] 
**Visit** | Pointer to [**MarketingVisit**](MarketingVisit.md) | Visit is what the visitor&#39;s page knows about the ad click that brought them and the consent they gave, so the lead reaches the ad platforms matched to the click and under the visitor&#39;s own choice. Optional. | [optional] 

## Methods

### NewMarketingLeadIn

`func NewMarketingLeadIn() *MarketingLeadIn`

NewMarketingLeadIn instantiates a new MarketingLeadIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingLeadInWithDefaults

`func NewMarketingLeadInWithDefaults() *MarketingLeadIn`

NewMarketingLeadInWithDefaults instantiates a new MarketingLeadIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *MarketingLeadIn) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *MarketingLeadIn) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *MarketingLeadIn) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *MarketingLeadIn) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountry

`func (o *MarketingLeadIn) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MarketingLeadIn) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MarketingLeadIn) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MarketingLeadIn) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *MarketingLeadIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MarketingLeadIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MarketingLeadIn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MarketingLeadIn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *MarketingLeadIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingLeadIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingLeadIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingLeadIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeed

`func (o *MarketingLeadIn) GetNeed() string`

GetNeed returns the Need field if non-nil, zero value otherwise.

### GetNeedOk

`func (o *MarketingLeadIn) GetNeedOk() (*string, bool)`

GetNeedOk returns a tuple with the Need field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeed

`func (o *MarketingLeadIn) SetNeed(v string)`

SetNeed sets Need field to given value.

### HasNeed

`func (o *MarketingLeadIn) HasNeed() bool`

HasNeed returns a boolean if a field has been set.

### GetRole

`func (o *MarketingLeadIn) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MarketingLeadIn) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MarketingLeadIn) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MarketingLeadIn) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSize

`func (o *MarketingLeadIn) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MarketingLeadIn) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MarketingLeadIn) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *MarketingLeadIn) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *MarketingLeadIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketingLeadIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketingLeadIn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketingLeadIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVisit

`func (o *MarketingLeadIn) GetVisit() MarketingVisit`

GetVisit returns the Visit field if non-nil, zero value otherwise.

### GetVisitOk

`func (o *MarketingLeadIn) GetVisitOk() (*MarketingVisit, bool)`

GetVisitOk returns a tuple with the Visit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisit

`func (o *MarketingLeadIn) SetVisit(v MarketingVisit)`

SetVisit sets Visit field to given value.

### HasVisit

`func (o *MarketingLeadIn) HasVisit() bool`

HasVisit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


