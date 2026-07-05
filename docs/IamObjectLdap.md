# IamObjectLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSelfSignedCert** | Pointer to **bool** |  | [optional] 
**AutoSync** | Pointer to **int64** |  | [optional] 
**BaseDn** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**CustomAttributes** | Pointer to  |  | [optional] 
**DefaultGroup** | Pointer to **string** |  | [optional] 
**EnableSsl** | Pointer to **bool** |  | [optional] 
**Filter** | Pointer to **string** |  | [optional] 
**FilterFields** | Pointer to **[]string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordType** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectLdap

`func NewIamObjectLdap() *IamObjectLdap`

NewIamObjectLdap instantiates a new IamObjectLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectLdapWithDefaults

`func NewIamObjectLdapWithDefaults() *IamObjectLdap`

NewIamObjectLdapWithDefaults instantiates a new IamObjectLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSelfSignedCert

`func (o *IamObjectLdap) GetAllowSelfSignedCert() bool`

GetAllowSelfSignedCert returns the AllowSelfSignedCert field if non-nil, zero value otherwise.

### GetAllowSelfSignedCertOk

`func (o *IamObjectLdap) GetAllowSelfSignedCertOk() (*bool, bool)`

GetAllowSelfSignedCertOk returns a tuple with the AllowSelfSignedCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSelfSignedCert

`func (o *IamObjectLdap) SetAllowSelfSignedCert(v bool)`

SetAllowSelfSignedCert sets AllowSelfSignedCert field to given value.

### HasAllowSelfSignedCert

`func (o *IamObjectLdap) HasAllowSelfSignedCert() bool`

HasAllowSelfSignedCert returns a boolean if a field has been set.

### GetAutoSync

`func (o *IamObjectLdap) GetAutoSync() int64`

GetAutoSync returns the AutoSync field if non-nil, zero value otherwise.

### GetAutoSyncOk

`func (o *IamObjectLdap) GetAutoSyncOk() (*int64, bool)`

GetAutoSyncOk returns a tuple with the AutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSync

`func (o *IamObjectLdap) SetAutoSync(v int64)`

SetAutoSync sets AutoSync field to given value.

### HasAutoSync

`func (o *IamObjectLdap) HasAutoSync() bool`

HasAutoSync returns a boolean if a field has been set.

### GetBaseDn

`func (o *IamObjectLdap) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *IamObjectLdap) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *IamObjectLdap) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *IamObjectLdap) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectLdap) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectLdap) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectLdap) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectLdap) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *IamObjectLdap) GetCustomAttributes() map[string]string`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *IamObjectLdap) GetCustomAttributesOk() (*map[string]string, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *IamObjectLdap) SetCustomAttributes(v map[string]string)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *IamObjectLdap) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *IamObjectLdap) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *IamObjectLdap) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetDefaultGroup

`func (o *IamObjectLdap) GetDefaultGroup() string`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *IamObjectLdap) GetDefaultGroupOk() (*string, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *IamObjectLdap) SetDefaultGroup(v string)`

SetDefaultGroup sets DefaultGroup field to given value.

### HasDefaultGroup

`func (o *IamObjectLdap) HasDefaultGroup() bool`

HasDefaultGroup returns a boolean if a field has been set.

### GetEnableSsl

`func (o *IamObjectLdap) GetEnableSsl() bool`

GetEnableSsl returns the EnableSsl field if non-nil, zero value otherwise.

### GetEnableSslOk

`func (o *IamObjectLdap) GetEnableSslOk() (*bool, bool)`

GetEnableSslOk returns a tuple with the EnableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSsl

`func (o *IamObjectLdap) SetEnableSsl(v bool)`

SetEnableSsl sets EnableSsl field to given value.

### HasEnableSsl

`func (o *IamObjectLdap) HasEnableSsl() bool`

HasEnableSsl returns a boolean if a field has been set.

### GetFilter

`func (o *IamObjectLdap) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IamObjectLdap) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IamObjectLdap) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IamObjectLdap) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFilterFields

`func (o *IamObjectLdap) GetFilterFields() []string`

GetFilterFields returns the FilterFields field if non-nil, zero value otherwise.

### GetFilterFieldsOk

`func (o *IamObjectLdap) GetFilterFieldsOk() (*[]string, bool)`

GetFilterFieldsOk returns a tuple with the FilterFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterFields

`func (o *IamObjectLdap) SetFilterFields(v []string)`

SetFilterFields sets FilterFields field to given value.

### HasFilterFields

`func (o *IamObjectLdap) HasFilterFields() bool`

HasFilterFields returns a boolean if a field has been set.

### GetHost

`func (o *IamObjectLdap) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IamObjectLdap) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IamObjectLdap) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IamObjectLdap) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *IamObjectLdap) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamObjectLdap) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamObjectLdap) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamObjectLdap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSync

`func (o *IamObjectLdap) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *IamObjectLdap) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *IamObjectLdap) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *IamObjectLdap) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectLdap) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectLdap) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectLdap) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectLdap) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPassword

`func (o *IamObjectLdap) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamObjectLdap) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamObjectLdap) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamObjectLdap) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordType

`func (o *IamObjectLdap) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *IamObjectLdap) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *IamObjectLdap) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *IamObjectLdap) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetPort

`func (o *IamObjectLdap) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IamObjectLdap) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IamObjectLdap) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *IamObjectLdap) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServerName

`func (o *IamObjectLdap) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *IamObjectLdap) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *IamObjectLdap) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *IamObjectLdap) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetUsername

`func (o *IamObjectLdap) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IamObjectLdap) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IamObjectLdap) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IamObjectLdap) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


