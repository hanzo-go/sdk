# DomainView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is the unix second the custom claim was made. | [optional] 
**Detail** | Pointer to **string** | Detail says why a claim is still pending, in the resolver&#39;s own words. | [optional] 
**Host** | Pointer to **string** | Host is the hostname itself. | [optional] 
**Kind** | Pointer to **string** | Kind is &#x60;default&#x60;, &#x60;subtree&#x60; or &#x60;custom&#x60; — how the org came to own it. | [optional] 
**Primary** | Pointer to **bool** | Primary marks the app&#39;s permanent default host. | [optional] 
**Records** | Pointer to [**[]Record**](Record.md) | Records are the DNS records to publish while a custom claim is pending. | [optional] 
**Status** | Pointer to **string** | Status is &#x60;live&#x60;, &#x60;provisioning&#x60;, &#x60;pending_deploy&#x60; or &#x60;pending&#x60;, derived from the operator CR and never fabricated. | [optional] 
**Url** | Pointer to **string** | URL is the host as an HTTPS address. | [optional] 
**Verified** | Pointer to **bool** | Verified is whether ownership is settled — always true for a host the org structurally owns. | [optional] 

## Methods

### NewDomainView

`func NewDomainView() *DomainView`

NewDomainView instantiates a new DomainView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainViewWithDefaults

`func NewDomainViewWithDefaults() *DomainView`

NewDomainViewWithDefaults instantiates a new DomainView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DomainView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDetail

`func (o *DomainView) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DomainView) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DomainView) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *DomainView) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetHost

`func (o *DomainView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DomainView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DomainView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DomainView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *DomainView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DomainView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DomainView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DomainView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPrimary

`func (o *DomainView) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *DomainView) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *DomainView) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *DomainView) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetRecords

`func (o *DomainView) GetRecords() []Record`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *DomainView) GetRecordsOk() (*[]Record, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *DomainView) SetRecords(v []Record)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *DomainView) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetStatus

`func (o *DomainView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *DomainView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DomainView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DomainView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DomainView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVerified

`func (o *DomainView) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *DomainView) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *DomainView) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *DomainView) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


