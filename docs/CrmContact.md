# CrmContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**CompanyId** | Pointer to **string** | Optional in-org relation to a Company | [optional] 
**LinkedinLink** | Pointer to **string** |  | [optional] 
**XLink** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** | Unix seconds | [optional] 
**UpdatedAt** | Pointer to **int64** | Unix seconds | [optional] 

## Methods

### NewCrmContact

`func NewCrmContact() *CrmContact`

NewCrmContact instantiates a new CrmContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmContactWithDefaults

`func NewCrmContactWithDefaults() *CrmContact`

NewCrmContactWithDefaults instantiates a new CrmContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *CrmContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CrmContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CrmContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CrmContact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CrmContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CrmContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CrmContact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CrmContact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *CrmContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CrmContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CrmContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CrmContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *CrmContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CrmContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CrmContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CrmContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetJobTitle

`func (o *CrmContact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *CrmContact) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *CrmContact) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *CrmContact) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetCity

`func (o *CrmContact) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CrmContact) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CrmContact) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CrmContact) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyId

`func (o *CrmContact) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CrmContact) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CrmContact) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CrmContact) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetLinkedinLink

`func (o *CrmContact) GetLinkedinLink() string`

GetLinkedinLink returns the LinkedinLink field if non-nil, zero value otherwise.

### GetLinkedinLinkOk

`func (o *CrmContact) GetLinkedinLinkOk() (*string, bool)`

GetLinkedinLinkOk returns a tuple with the LinkedinLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinLink

`func (o *CrmContact) SetLinkedinLink(v string)`

SetLinkedinLink sets LinkedinLink field to given value.

### HasLinkedinLink

`func (o *CrmContact) HasLinkedinLink() bool`

HasLinkedinLink returns a boolean if a field has been set.

### GetXLink

`func (o *CrmContact) GetXLink() string`

GetXLink returns the XLink field if non-nil, zero value otherwise.

### GetXLinkOk

`func (o *CrmContact) GetXLinkOk() (*string, bool)`

GetXLinkOk returns a tuple with the XLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLink

`func (o *CrmContact) SetXLink(v string)`

SetXLink sets XLink field to given value.

### HasXLink

`func (o *CrmContact) HasXLink() bool`

HasXLink returns a boolean if a field has been set.

### GetId

`func (o *CrmContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrmContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrmContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CrmContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CrmContact) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CrmContact) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CrmContact) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CrmContact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CrmContact) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CrmContact) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CrmContact) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CrmContact) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


