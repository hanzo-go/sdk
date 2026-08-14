# Contact

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

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *Contact) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Contact) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Contact) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Contact) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyId

`func (o *Contact) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Contact) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Contact) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *Contact) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Contact) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Contact) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Contact) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Contact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Contact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *Contact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Contact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Contact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Contact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *Contact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Contact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobTitle

`func (o *Contact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Contact) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Contact) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Contact) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastName

`func (o *Contact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Contact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Contact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Contact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLinkedinLink

`func (o *Contact) GetLinkedinLink() string`

GetLinkedinLink returns the LinkedinLink field if non-nil, zero value otherwise.

### GetLinkedinLinkOk

`func (o *Contact) GetLinkedinLinkOk() (*string, bool)`

GetLinkedinLinkOk returns a tuple with the LinkedinLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinLink

`func (o *Contact) SetLinkedinLink(v string)`

SetLinkedinLink sets LinkedinLink field to given value.

### HasLinkedinLink

`func (o *Contact) HasLinkedinLink() bool`

HasLinkedinLink returns a boolean if a field has been set.

### GetPhone

`func (o *Contact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Contact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Contact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Contact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Contact) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Contact) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Contact) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Contact) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetXLink

`func (o *Contact) GetXLink() string`

GetXLink returns the XLink field if non-nil, zero value otherwise.

### GetXLinkOk

`func (o *Contact) GetXLinkOk() (*string, bool)`

GetXLinkOk returns a tuple with the XLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLink

`func (o *Contact) SetXLink(v string)`

SetXLink sets XLink field to given value.

### HasXLink

`func (o *Contact) HasXLink() bool`

HasXLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


