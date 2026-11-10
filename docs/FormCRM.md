# FormCRM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** | Company is the ref of the field whose answer is the person&#39;s company. | [optional] 
**Email** | Pointer to **string** | Email is the ref of the field whose answer is the person&#39;s address. Empty, the first email field is. A signed-in respondent is always the address their account holds. | [optional] 
**List** | Pointer to **string** | List is the CRM list the person joins, by name (a subscribe form&#39;s mailing list). Empty joins none. | [optional] 
**Name** | Pointer to **string** | Name is the ref of the field whose answer is the person&#39;s name. | [optional] 

## Methods

### NewFormCRM

`func NewFormCRM() *FormCRM`

NewFormCRM instantiates a new FormCRM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormCRMWithDefaults

`func NewFormCRMWithDefaults() *FormCRM`

NewFormCRMWithDefaults instantiates a new FormCRM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *FormCRM) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *FormCRM) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *FormCRM) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *FormCRM) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEmail

`func (o *FormCRM) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FormCRM) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FormCRM) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FormCRM) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetList

`func (o *FormCRM) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *FormCRM) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *FormCRM) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *FormCRM) HasList() bool`

HasList returns a boolean if a field has been set.

### GetName

`func (o *FormCRM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormCRM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormCRM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormCRM) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


