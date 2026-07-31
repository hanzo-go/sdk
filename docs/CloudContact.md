# CloudContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | City is where the person is based. | [optional] 
**CompanyId** | Pointer to **string** | CompanyID links the contact to one of the org&#39;s companies; empty when the contact stands alone, and cleared when its company is deleted. A write naming a company the org does not own is refused with 422. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is the unix second the contact was created. Server-owned. | [optional] 
**Email** | Pointer to **string** | Email is the person&#39;s email address. | [optional] 
**FirstName** | Pointer to **string** | FirstName is the person&#39;s given name. | [optional] 
**Id** | Pointer to **string** | ID is the server-minted contact id (\&quot;cont_\&quot; + 128 random bits). | [optional] 
**JobTitle** | Pointer to **string** | JobTitle is the person&#39;s role at their company. | [optional] 
**LastName** | Pointer to **string** | LastName is the person&#39;s family name. | [optional] 
**LinkedinLink** | Pointer to **string** | Linkedin is the person&#39;s LinkedIn URL. | [optional] 
**Phone** | Pointer to **string** | Phone is the person&#39;s phone number. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second of the last write. Server-owned. | [optional] 
**XLink** | Pointer to **string** | XLink is the person&#39;s X (Twitter) URL. | [optional] 

## Methods

### NewCloudContact

`func NewCloudContact() *CloudContact`

NewCloudContact instantiates a new CloudContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudContactWithDefaults

`func NewCloudContactWithDefaults() *CloudContact`

NewCloudContactWithDefaults instantiates a new CloudContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *CloudContact) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CloudContact) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CloudContact) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CloudContact) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyId

`func (o *CloudContact) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CloudContact) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CloudContact) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CloudContact) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudContact) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudContact) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudContact) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudContact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *CloudContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *CloudContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CloudContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CloudContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CloudContact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *CloudContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobTitle

`func (o *CloudContact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *CloudContact) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *CloudContact) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *CloudContact) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastName

`func (o *CloudContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CloudContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CloudContact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CloudContact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLinkedinLink

`func (o *CloudContact) GetLinkedinLink() string`

GetLinkedinLink returns the LinkedinLink field if non-nil, zero value otherwise.

### GetLinkedinLinkOk

`func (o *CloudContact) GetLinkedinLinkOk() (*string, bool)`

GetLinkedinLinkOk returns a tuple with the LinkedinLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinLink

`func (o *CloudContact) SetLinkedinLink(v string)`

SetLinkedinLink sets LinkedinLink field to given value.

### HasLinkedinLink

`func (o *CloudContact) HasLinkedinLink() bool`

HasLinkedinLink returns a boolean if a field has been set.

### GetPhone

`func (o *CloudContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CloudContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CloudContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CloudContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudContact) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudContact) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudContact) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudContact) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetXLink

`func (o *CloudContact) GetXLink() string`

GetXLink returns the XLink field if non-nil, zero value otherwise.

### GetXLinkOk

`func (o *CloudContact) GetXLinkOk() (*string, bool)`

GetXLinkOk returns a tuple with the XLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLink

`func (o *CloudContact) SetXLink(v string)`

SetXLink sets XLink field to given value.

### HasXLink

`func (o *CloudContact) HasXLink() bool`

HasXLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


