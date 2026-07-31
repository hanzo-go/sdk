# CommerceUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Username** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**BillingAddress** | Pointer to [**CommerceAddress**](CommerceAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**CommerceAddress**](CommerceAddress.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**IsAffiliate** | Pointer to **bool** |  | [optional] 
**AffiliateId** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Test** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCommerceUser

`func NewCommerceUser() *CommerceUser`

NewCommerceUser instantiates a new CommerceUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceUserWithDefaults

`func NewCommerceUserWithDefaults() *CommerceUser`

NewCommerceUserWithDefaults instantiates a new CommerceUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommerceUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommerceUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommerceUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommerceUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *CommerceUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CommerceUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CommerceUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CommerceUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *CommerceUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CommerceUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CommerceUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CommerceUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CommerceUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CommerceUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CommerceUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CommerceUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *CommerceUser) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CommerceUser) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CommerceUser) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CommerceUser) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetPhone

`func (o *CommerceUser) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CommerceUser) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CommerceUser) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CommerceUser) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *CommerceUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CommerceUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CommerceUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CommerceUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetBillingAddress

`func (o *CommerceUser) GetBillingAddress() CommerceAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CommerceUser) GetBillingAddressOk() (*CommerceAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CommerceUser) SetBillingAddress(v CommerceAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CommerceUser) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *CommerceUser) GetShippingAddress() CommerceAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *CommerceUser) GetShippingAddressOk() (*CommerceAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *CommerceUser) SetShippingAddress(v CommerceAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *CommerceUser) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetEnabled

`func (o *CommerceUser) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CommerceUser) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CommerceUser) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CommerceUser) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsAffiliate

`func (o *CommerceUser) GetIsAffiliate() bool`

GetIsAffiliate returns the IsAffiliate field if non-nil, zero value otherwise.

### GetIsAffiliateOk

`func (o *CommerceUser) GetIsAffiliateOk() (*bool, bool)`

GetIsAffiliateOk returns a tuple with the IsAffiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAffiliate

`func (o *CommerceUser) SetIsAffiliate(v bool)`

SetIsAffiliate sets IsAffiliate field to given value.

### HasIsAffiliate

`func (o *CommerceUser) HasIsAffiliate() bool`

HasIsAffiliate returns a boolean if a field has been set.

### GetAffiliateId

`func (o *CommerceUser) GetAffiliateId() string`

GetAffiliateId returns the AffiliateId field if non-nil, zero value otherwise.

### GetAffiliateIdOk

`func (o *CommerceUser) GetAffiliateIdOk() (*string, bool)`

GetAffiliateIdOk returns a tuple with the AffiliateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateId

`func (o *CommerceUser) SetAffiliateId(v string)`

SetAffiliateId sets AffiliateId field to given value.

### HasAffiliateId

`func (o *CommerceUser) HasAffiliateId() bool`

HasAffiliateId returns a boolean if a field has been set.

### GetMetadata

`func (o *CommerceUser) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommerceUser) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommerceUser) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CommerceUser) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTest

`func (o *CommerceUser) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CommerceUser) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CommerceUser) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CommerceUser) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommerceUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommerceUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommerceUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommerceUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommerceUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommerceUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommerceUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommerceUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


