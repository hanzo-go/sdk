# IdentityView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enrollment** | Pointer to [**EnrollmentView**](EnrollmentView.md) | Enrollment is present only while the identity holds an un-used one-time token — on create, and on a listed identity that has not yet enrolled, so a mislaid JWT can be read again until it is spent or lapses. | [optional] 
**Id** | Pointer to **string** | ID is the identity&#39;s fabric id — the key DELETE addresses. | [optional] 
**Name** | Pointer to **string** | Name is the identity&#39;s name within the org. | [optional] 
**Roles** | Pointer to **[]string** | Roles are the identity&#39;s role attributes as the fabric holds them: the org&#39;s own \&quot;org-&lt;org&gt;\&quot; plus any org-scoped roles it was minted with. | [optional] 

## Methods

### NewIdentityView

`func NewIdentityView() *IdentityView`

NewIdentityView instantiates a new IdentityView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityViewWithDefaults

`func NewIdentityViewWithDefaults() *IdentityView`

NewIdentityViewWithDefaults instantiates a new IdentityView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnrollment

`func (o *IdentityView) GetEnrollment() EnrollmentView`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *IdentityView) GetEnrollmentOk() (*EnrollmentView, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *IdentityView) SetEnrollment(v EnrollmentView)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *IdentityView) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetId

`func (o *IdentityView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *IdentityView) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IdentityView) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IdentityView) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IdentityView) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


