# EsignSigner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address the link was issued to. | [optional] 
**Id** | Pointer to **string** | ID is the recipient id. | [optional] 
**Name** | Pointer to **string** | Name is the display name recorded for them, empty when none was given. | [optional] 
**Role** | Pointer to **string** | Role is the role they were added with. | [optional] 
**SigningStatus** | Pointer to **string** | SigningStatus is NOT_SIGNED until they finish or decline. | [optional] 

## Methods

### NewEsignSigner

`func NewEsignSigner() *EsignSigner`

NewEsignSigner instantiates a new EsignSigner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignSignerWithDefaults

`func NewEsignSignerWithDefaults() *EsignSigner`

NewEsignSignerWithDefaults instantiates a new EsignSigner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EsignSigner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EsignSigner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EsignSigner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EsignSigner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *EsignSigner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignSigner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignSigner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignSigner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EsignSigner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsignSigner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsignSigner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EsignSigner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *EsignSigner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EsignSigner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EsignSigner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EsignSigner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSigningStatus

`func (o *EsignSigner) GetSigningStatus() string`

GetSigningStatus returns the SigningStatus field if non-nil, zero value otherwise.

### GetSigningStatusOk

`func (o *EsignSigner) GetSigningStatusOk() (*string, bool)`

GetSigningStatusOk returns a tuple with the SigningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningStatus

`func (o *EsignSigner) SetSigningStatus(v string)`

SetSigningStatus sets SigningStatus field to given value.

### HasSigningStatus

`func (o *EsignSigner) HasSigningStatus() bool`

HasSigningStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


