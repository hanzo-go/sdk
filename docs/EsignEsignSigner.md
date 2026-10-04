# EsignEsignSigner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address the link was issued to. | [optional] 
**Id** | Pointer to **string** | ID is the recipient id. | [optional] 
**Name** | Pointer to **string** | Name is the display name recorded for them, empty when none was given. | [optional] 
**Role** | Pointer to **string** | Role is the role they were added with. | [optional] 
**SigningStatus** | Pointer to **string** | SigningStatus is NOT_SIGNED until they finish or decline. | [optional] 

## Methods

### NewEsignEsignSigner

`func NewEsignEsignSigner() *EsignEsignSigner`

NewEsignEsignSigner instantiates a new EsignEsignSigner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEsignSignerWithDefaults

`func NewEsignEsignSignerWithDefaults() *EsignEsignSigner`

NewEsignEsignSignerWithDefaults instantiates a new EsignEsignSigner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EsignEsignSigner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EsignEsignSigner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EsignEsignSigner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EsignEsignSigner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *EsignEsignSigner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignEsignSigner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignEsignSigner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignEsignSigner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EsignEsignSigner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsignEsignSigner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsignEsignSigner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EsignEsignSigner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *EsignEsignSigner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EsignEsignSigner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EsignEsignSigner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EsignEsignSigner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSigningStatus

`func (o *EsignEsignSigner) GetSigningStatus() string`

GetSigningStatus returns the SigningStatus field if non-nil, zero value otherwise.

### GetSigningStatusOk

`func (o *EsignEsignSigner) GetSigningStatusOk() (*string, bool)`

GetSigningStatusOk returns a tuple with the SigningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningStatus

`func (o *EsignEsignSigner) SetSigningStatus(v string)`

SetSigningStatus sets SigningStatus field to given value.

### HasSigningStatus

`func (o *EsignEsignSigner) HasSigningStatus() bool`

HasSigningStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


