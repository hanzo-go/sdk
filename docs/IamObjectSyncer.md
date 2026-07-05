# IamObjectSyncer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffiliationTable** | Pointer to **string** |  | [optional] 
**AvatarBaseUrl** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Database** | Pointer to **string** |  | [optional] 
**DatabaseType** | Pointer to **string** |  | [optional] 
**ErrorText** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**SshHost** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**SshPort** | Pointer to **int64** |  | [optional] 
**SshType** | Pointer to **string** |  | [optional] 
**SshUser** | Pointer to **string** |  | [optional] 
**SslMode** | Pointer to **string** |  | [optional] 
**SyncInterval** | Pointer to **int64** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**TableColumns** | Pointer to [**[]IamObjectTableColumn**](IamObjectTableColumn.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectSyncer

`func NewIamObjectSyncer() *IamObjectSyncer`

NewIamObjectSyncer instantiates a new IamObjectSyncer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectSyncerWithDefaults

`func NewIamObjectSyncerWithDefaults() *IamObjectSyncer`

NewIamObjectSyncerWithDefaults instantiates a new IamObjectSyncer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffiliationTable

`func (o *IamObjectSyncer) GetAffiliationTable() string`

GetAffiliationTable returns the AffiliationTable field if non-nil, zero value otherwise.

### GetAffiliationTableOk

`func (o *IamObjectSyncer) GetAffiliationTableOk() (*string, bool)`

GetAffiliationTableOk returns a tuple with the AffiliationTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliationTable

`func (o *IamObjectSyncer) SetAffiliationTable(v string)`

SetAffiliationTable sets AffiliationTable field to given value.

### HasAffiliationTable

`func (o *IamObjectSyncer) HasAffiliationTable() bool`

HasAffiliationTable returns a boolean if a field has been set.

### GetAvatarBaseUrl

`func (o *IamObjectSyncer) GetAvatarBaseUrl() string`

GetAvatarBaseUrl returns the AvatarBaseUrl field if non-nil, zero value otherwise.

### GetAvatarBaseUrlOk

`func (o *IamObjectSyncer) GetAvatarBaseUrlOk() (*string, bool)`

GetAvatarBaseUrlOk returns a tuple with the AvatarBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarBaseUrl

`func (o *IamObjectSyncer) SetAvatarBaseUrl(v string)`

SetAvatarBaseUrl sets AvatarBaseUrl field to given value.

### HasAvatarBaseUrl

`func (o *IamObjectSyncer) HasAvatarBaseUrl() bool`

HasAvatarBaseUrl returns a boolean if a field has been set.

### GetCert

`func (o *IamObjectSyncer) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *IamObjectSyncer) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *IamObjectSyncer) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *IamObjectSyncer) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectSyncer) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectSyncer) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectSyncer) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectSyncer) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDatabase

`func (o *IamObjectSyncer) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *IamObjectSyncer) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *IamObjectSyncer) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *IamObjectSyncer) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetDatabaseType

`func (o *IamObjectSyncer) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *IamObjectSyncer) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *IamObjectSyncer) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *IamObjectSyncer) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### GetErrorText

`func (o *IamObjectSyncer) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *IamObjectSyncer) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *IamObjectSyncer) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *IamObjectSyncer) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetHost

`func (o *IamObjectSyncer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IamObjectSyncer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IamObjectSyncer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IamObjectSyncer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIsEnabled

`func (o *IamObjectSyncer) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IamObjectSyncer) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IamObjectSyncer) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IamObjectSyncer) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *IamObjectSyncer) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *IamObjectSyncer) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *IamObjectSyncer) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *IamObjectSyncer) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetName

`func (o *IamObjectSyncer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectSyncer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectSyncer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectSyncer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *IamObjectSyncer) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamObjectSyncer) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamObjectSyncer) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamObjectSyncer) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectSyncer) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectSyncer) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectSyncer) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectSyncer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPassword

`func (o *IamObjectSyncer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamObjectSyncer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamObjectSyncer) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamObjectSyncer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *IamObjectSyncer) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IamObjectSyncer) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IamObjectSyncer) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *IamObjectSyncer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSshHost

`func (o *IamObjectSyncer) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *IamObjectSyncer) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *IamObjectSyncer) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *IamObjectSyncer) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPassword

`func (o *IamObjectSyncer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *IamObjectSyncer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *IamObjectSyncer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *IamObjectSyncer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshPort

`func (o *IamObjectSyncer) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *IamObjectSyncer) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *IamObjectSyncer) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *IamObjectSyncer) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshType

`func (o *IamObjectSyncer) GetSshType() string`

GetSshType returns the SshType field if non-nil, zero value otherwise.

### GetSshTypeOk

`func (o *IamObjectSyncer) GetSshTypeOk() (*string, bool)`

GetSshTypeOk returns a tuple with the SshType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshType

`func (o *IamObjectSyncer) SetSshType(v string)`

SetSshType sets SshType field to given value.

### HasSshType

`func (o *IamObjectSyncer) HasSshType() bool`

HasSshType returns a boolean if a field has been set.

### GetSshUser

`func (o *IamObjectSyncer) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *IamObjectSyncer) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *IamObjectSyncer) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *IamObjectSyncer) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSslMode

`func (o *IamObjectSyncer) GetSslMode() string`

GetSslMode returns the SslMode field if non-nil, zero value otherwise.

### GetSslModeOk

`func (o *IamObjectSyncer) GetSslModeOk() (*string, bool)`

GetSslModeOk returns a tuple with the SslMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMode

`func (o *IamObjectSyncer) SetSslMode(v string)`

SetSslMode sets SslMode field to given value.

### HasSslMode

`func (o *IamObjectSyncer) HasSslMode() bool`

HasSslMode returns a boolean if a field has been set.

### GetSyncInterval

`func (o *IamObjectSyncer) GetSyncInterval() int64`

GetSyncInterval returns the SyncInterval field if non-nil, zero value otherwise.

### GetSyncIntervalOk

`func (o *IamObjectSyncer) GetSyncIntervalOk() (*int64, bool)`

GetSyncIntervalOk returns a tuple with the SyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInterval

`func (o *IamObjectSyncer) SetSyncInterval(v int64)`

SetSyncInterval sets SyncInterval field to given value.

### HasSyncInterval

`func (o *IamObjectSyncer) HasSyncInterval() bool`

HasSyncInterval returns a boolean if a field has been set.

### GetTable

`func (o *IamObjectSyncer) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *IamObjectSyncer) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *IamObjectSyncer) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *IamObjectSyncer) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetTableColumns

`func (o *IamObjectSyncer) GetTableColumns() []IamObjectTableColumn`

GetTableColumns returns the TableColumns field if non-nil, zero value otherwise.

### GetTableColumnsOk

`func (o *IamObjectSyncer) GetTableColumnsOk() (*[]IamObjectTableColumn, bool)`

GetTableColumnsOk returns a tuple with the TableColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableColumns

`func (o *IamObjectSyncer) SetTableColumns(v []IamObjectTableColumn)`

SetTableColumns sets TableColumns field to given value.

### HasTableColumns

`func (o *IamObjectSyncer) HasTableColumns() bool`

HasTableColumns returns a boolean if a field has been set.

### GetType

`func (o *IamObjectSyncer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamObjectSyncer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamObjectSyncer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamObjectSyncer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *IamObjectSyncer) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamObjectSyncer) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamObjectSyncer) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamObjectSyncer) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


