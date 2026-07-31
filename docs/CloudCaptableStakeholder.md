# CloudCaptableStakeholder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | City is the stakeholder&#39;s city, if recorded. | [optional] 
**CompanyName** | Pointer to **string** | CompanyName is the name of the company whose cap table this is. | [optional] 
**Country** | Pointer to **string** | Country is the stakeholder&#39;s two-letter country code. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the stakeholder was added, in unix milliseconds. | [optional] 
**CurrentRelationship** | Pointer to **string** | CurrentRelationship is how the stakeholder relates to the company, e.g. FOUNDER, INVESTOR or EMPLOYEE. | [optional] 
**Email** | Pointer to **string** | Email is the stakeholder&#39;s email, unique within the company. | [optional] 
**Id** | Pointer to **string** | ID is the stakeholder id. | [optional] 
**InstitutionName** | Pointer to **string** | InstitutionName names the institution, when the stakeholder is one. | [optional] 
**Name** | Pointer to **string** | Name is the stakeholder&#39;s full name. | [optional] 
**StakeholderType** | Pointer to **string** | StakeholderType is INDIVIDUAL or INSTITUTION. | [optional] 
**State** | Pointer to **string** | State is the stakeholder&#39;s state or province, if recorded. | [optional] 
**StreetAddress** | Pointer to **string** | StreetAddress is the stakeholder&#39;s street address, if recorded. | [optional] 
**TaxId** | Pointer to **string** | TaxID is the stakeholder&#39;s tax identifier, if recorded. | [optional] 
**Zipcode** | Pointer to **string** | Zipcode is the stakeholder&#39;s postal code, if recorded. | [optional] 

## Methods

### NewCloudCaptableStakeholder

`func NewCloudCaptableStakeholder() *CloudCaptableStakeholder`

NewCloudCaptableStakeholder instantiates a new CloudCaptableStakeholder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableStakeholderWithDefaults

`func NewCloudCaptableStakeholderWithDefaults() *CloudCaptableStakeholder`

NewCloudCaptableStakeholderWithDefaults instantiates a new CloudCaptableStakeholder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *CloudCaptableStakeholder) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CloudCaptableStakeholder) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CloudCaptableStakeholder) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CloudCaptableStakeholder) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *CloudCaptableStakeholder) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CloudCaptableStakeholder) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CloudCaptableStakeholder) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CloudCaptableStakeholder) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *CloudCaptableStakeholder) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CloudCaptableStakeholder) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CloudCaptableStakeholder) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CloudCaptableStakeholder) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudCaptableStakeholder) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudCaptableStakeholder) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudCaptableStakeholder) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudCaptableStakeholder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentRelationship

`func (o *CloudCaptableStakeholder) GetCurrentRelationship() string`

GetCurrentRelationship returns the CurrentRelationship field if non-nil, zero value otherwise.

### GetCurrentRelationshipOk

`func (o *CloudCaptableStakeholder) GetCurrentRelationshipOk() (*string, bool)`

GetCurrentRelationshipOk returns a tuple with the CurrentRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRelationship

`func (o *CloudCaptableStakeholder) SetCurrentRelationship(v string)`

SetCurrentRelationship sets CurrentRelationship field to given value.

### HasCurrentRelationship

`func (o *CloudCaptableStakeholder) HasCurrentRelationship() bool`

HasCurrentRelationship returns a boolean if a field has been set.

### GetEmail

`func (o *CloudCaptableStakeholder) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudCaptableStakeholder) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudCaptableStakeholder) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudCaptableStakeholder) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *CloudCaptableStakeholder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableStakeholder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableStakeholder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableStakeholder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstitutionName

`func (o *CloudCaptableStakeholder) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *CloudCaptableStakeholder) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *CloudCaptableStakeholder) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.

### HasInstitutionName

`func (o *CloudCaptableStakeholder) HasInstitutionName() bool`

HasInstitutionName returns a boolean if a field has been set.

### GetName

`func (o *CloudCaptableStakeholder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCaptableStakeholder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCaptableStakeholder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCaptableStakeholder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStakeholderType

`func (o *CloudCaptableStakeholder) GetStakeholderType() string`

GetStakeholderType returns the StakeholderType field if non-nil, zero value otherwise.

### GetStakeholderTypeOk

`func (o *CloudCaptableStakeholder) GetStakeholderTypeOk() (*string, bool)`

GetStakeholderTypeOk returns a tuple with the StakeholderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderType

`func (o *CloudCaptableStakeholder) SetStakeholderType(v string)`

SetStakeholderType sets StakeholderType field to given value.

### HasStakeholderType

`func (o *CloudCaptableStakeholder) HasStakeholderType() bool`

HasStakeholderType returns a boolean if a field has been set.

### GetState

`func (o *CloudCaptableStakeholder) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudCaptableStakeholder) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudCaptableStakeholder) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudCaptableStakeholder) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreetAddress

`func (o *CloudCaptableStakeholder) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *CloudCaptableStakeholder) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *CloudCaptableStakeholder) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *CloudCaptableStakeholder) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetTaxId

`func (o *CloudCaptableStakeholder) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CloudCaptableStakeholder) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CloudCaptableStakeholder) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CloudCaptableStakeholder) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetZipcode

`func (o *CloudCaptableStakeholder) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *CloudCaptableStakeholder) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *CloudCaptableStakeholder) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *CloudCaptableStakeholder) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


