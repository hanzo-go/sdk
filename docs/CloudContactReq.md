# CloudContactReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | City is where the person is based. | [optional] 
**CompanyId** | Pointer to **string** | CompanyID links the contact to one of the org&#39;s companies. | [optional] 
**Email** | Pointer to **string** | Email is the person&#39;s email address. | [optional] 
**FirstName** | Pointer to **string** | FirstName is the person&#39;s given name. | [optional] 
**Id** | Pointer to **string** | ID names the contact to update and comes from the path. A create ignores it: the server mints the id. | [optional] 
**JobTitle** | Pointer to **string** | JobTitle is the person&#39;s role at their company. | [optional] 
**LastName** | Pointer to **string** | LastName is the person&#39;s family name. | [optional] 
**LinkedinLink** | Pointer to **string** | Linkedin is the person&#39;s LinkedIn URL. | [optional] 
**Phone** | Pointer to **string** | Phone is the person&#39;s phone number. | [optional] 
**XLink** | Pointer to **string** | XLink is the person&#39;s X (Twitter) URL. | [optional] 

## Methods

### NewCloudContactReq

`func NewCloudContactReq() *CloudContactReq`

NewCloudContactReq instantiates a new CloudContactReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudContactReqWithDefaults

`func NewCloudContactReqWithDefaults() *CloudContactReq`

NewCloudContactReqWithDefaults instantiates a new CloudContactReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *CloudContactReq) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CloudContactReq) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CloudContactReq) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CloudContactReq) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyId

`func (o *CloudContactReq) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CloudContactReq) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CloudContactReq) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CloudContactReq) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetEmail

`func (o *CloudContactReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudContactReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudContactReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudContactReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *CloudContactReq) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CloudContactReq) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CloudContactReq) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CloudContactReq) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *CloudContactReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudContactReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudContactReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudContactReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobTitle

`func (o *CloudContactReq) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *CloudContactReq) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *CloudContactReq) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *CloudContactReq) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastName

`func (o *CloudContactReq) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CloudContactReq) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CloudContactReq) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CloudContactReq) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLinkedinLink

`func (o *CloudContactReq) GetLinkedinLink() string`

GetLinkedinLink returns the LinkedinLink field if non-nil, zero value otherwise.

### GetLinkedinLinkOk

`func (o *CloudContactReq) GetLinkedinLinkOk() (*string, bool)`

GetLinkedinLinkOk returns a tuple with the LinkedinLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinLink

`func (o *CloudContactReq) SetLinkedinLink(v string)`

SetLinkedinLink sets LinkedinLink field to given value.

### HasLinkedinLink

`func (o *CloudContactReq) HasLinkedinLink() bool`

HasLinkedinLink returns a boolean if a field has been set.

### GetPhone

`func (o *CloudContactReq) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CloudContactReq) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CloudContactReq) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CloudContactReq) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetXLink

`func (o *CloudContactReq) GetXLink() string`

GetXLink returns the XLink field if non-nil, zero value otherwise.

### GetXLinkOk

`func (o *CloudContactReq) GetXLinkOk() (*string, bool)`

GetXLinkOk returns a tuple with the XLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLink

`func (o *CloudContactReq) SetXLink(v string)`

SetXLink sets XLink field to given value.

### HasXLink

`func (o *CloudContactReq) HasXLink() bool`

HasXLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


