# RegistryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the OCI registry host clients push to and pull from. | [optional] 
**Oci** | Pointer to **bool** | Oci is true when the OCI registry answered its /v2/ probe. | [optional] 
**Pkg** | Pointer to **bool** | Pkg is true when the npm registry answered its ping. | [optional] 
**PkgHost** | Pointer to **string** | PkgHost is the npm registry host. | [optional] 
**Realm** | Pointer to **string** | Realm is the token endpoint the OCI registry&#39;s challenge advertises, present only when the registry is reachable and auth-gated. | [optional] 
**Service** | Pointer to **string** | Service is the token service name from the same challenge. | [optional] 

## Methods

### NewRegistryStatus

`func NewRegistryStatus() *RegistryStatus`

NewRegistryStatus instantiates a new RegistryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryStatusWithDefaults

`func NewRegistryStatusWithDefaults() *RegistryStatus`

NewRegistryStatusWithDefaults instantiates a new RegistryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RegistryStatus) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RegistryStatus) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RegistryStatus) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RegistryStatus) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetOci

`func (o *RegistryStatus) GetOci() bool`

GetOci returns the Oci field if non-nil, zero value otherwise.

### GetOciOk

`func (o *RegistryStatus) GetOciOk() (*bool, bool)`

GetOciOk returns a tuple with the Oci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOci

`func (o *RegistryStatus) SetOci(v bool)`

SetOci sets Oci field to given value.

### HasOci

`func (o *RegistryStatus) HasOci() bool`

HasOci returns a boolean if a field has been set.

### GetPkg

`func (o *RegistryStatus) GetPkg() bool`

GetPkg returns the Pkg field if non-nil, zero value otherwise.

### GetPkgOk

`func (o *RegistryStatus) GetPkgOk() (*bool, bool)`

GetPkgOk returns a tuple with the Pkg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkg

`func (o *RegistryStatus) SetPkg(v bool)`

SetPkg sets Pkg field to given value.

### HasPkg

`func (o *RegistryStatus) HasPkg() bool`

HasPkg returns a boolean if a field has been set.

### GetPkgHost

`func (o *RegistryStatus) GetPkgHost() string`

GetPkgHost returns the PkgHost field if non-nil, zero value otherwise.

### GetPkgHostOk

`func (o *RegistryStatus) GetPkgHostOk() (*string, bool)`

GetPkgHostOk returns a tuple with the PkgHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgHost

`func (o *RegistryStatus) SetPkgHost(v string)`

SetPkgHost sets PkgHost field to given value.

### HasPkgHost

`func (o *RegistryStatus) HasPkgHost() bool`

HasPkgHost returns a boolean if a field has been set.

### GetRealm

`func (o *RegistryStatus) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *RegistryStatus) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *RegistryStatus) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *RegistryStatus) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetService

`func (o *RegistryStatus) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *RegistryStatus) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *RegistryStatus) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *RegistryStatus) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


