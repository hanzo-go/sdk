# PaasDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **string** |  | [optional] 
**Ssl** | Pointer to **bool** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPaasDomain

`func NewPaasDomain() *PaasDomain`

NewPaasDomain instantiates a new PaasDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasDomainWithDefaults

`func NewPaasDomainWithDefaults() *PaasDomain`

NewPaasDomainWithDefaults instantiates a new PaasDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaasDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaasDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaasDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaasDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *PaasDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PaasDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PaasDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PaasDomain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetContainerId

`func (o *PaasDomain) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *PaasDomain) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *PaasDomain) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *PaasDomain) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetSsl

`func (o *PaasDomain) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *PaasDomain) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *PaasDomain) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *PaasDomain) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetVerified

`func (o *PaasDomain) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PaasDomain) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PaasDomain) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *PaasDomain) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaasDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaasDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaasDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaasDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


