# EdgeDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**FunctionSlug** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TlsState** | Pointer to **string** |  | [optional] 
**Verification** | Pointer to [**EdgeDomainVerification**](EdgeDomainVerification.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEdgeDomain

`func NewEdgeDomain() *EdgeDomain`

NewEdgeDomain instantiates a new EdgeDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeDomainWithDefaults

`func NewEdgeDomainWithDefaults() *EdgeDomain`

NewEdgeDomainWithDefaults instantiates a new EdgeDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EdgeDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHostname

`func (o *EdgeDomain) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EdgeDomain) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EdgeDomain) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *EdgeDomain) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetFunctionSlug

`func (o *EdgeDomain) GetFunctionSlug() string`

GetFunctionSlug returns the FunctionSlug field if non-nil, zero value otherwise.

### GetFunctionSlugOk

`func (o *EdgeDomain) GetFunctionSlugOk() (*string, bool)`

GetFunctionSlugOk returns a tuple with the FunctionSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSlug

`func (o *EdgeDomain) SetFunctionSlug(v string)`

SetFunctionSlug sets FunctionSlug field to given value.

### HasFunctionSlug

`func (o *EdgeDomain) HasFunctionSlug() bool`

HasFunctionSlug returns a boolean if a field has been set.

### GetStatus

`func (o *EdgeDomain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EdgeDomain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EdgeDomain) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EdgeDomain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTlsState

`func (o *EdgeDomain) GetTlsState() string`

GetTlsState returns the TlsState field if non-nil, zero value otherwise.

### GetTlsStateOk

`func (o *EdgeDomain) GetTlsStateOk() (*string, bool)`

GetTlsStateOk returns a tuple with the TlsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsState

`func (o *EdgeDomain) SetTlsState(v string)`

SetTlsState sets TlsState field to given value.

### HasTlsState

`func (o *EdgeDomain) HasTlsState() bool`

HasTlsState returns a boolean if a field has been set.

### GetVerification

`func (o *EdgeDomain) GetVerification() EdgeDomainVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *EdgeDomain) GetVerificationOk() (*EdgeDomainVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *EdgeDomain) SetVerification(v EdgeDomainVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *EdgeDomain) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EdgeDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EdgeDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EdgeDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EdgeDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


