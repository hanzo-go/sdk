# PlatformDomainView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is the unix second the custom claim was made. | [optional] 
**Detail** | Pointer to **string** | Detail says why a claim is still pending, in the resolver&#39;s own words. | [optional] 
**Host** | Pointer to **string** | Host is the hostname itself. | [optional] 
**Kind** | Pointer to **string** | Kind is &#x60;default&#x60;, &#x60;subtree&#x60; or &#x60;custom&#x60; — how the org came to own it. | [optional] 
**Primary** | Pointer to **bool** | Primary marks the app&#39;s permanent default host. | [optional] 
**Records** | Pointer to [**[]PlatformRecord**](PlatformRecord.md) | Records are the DNS records to publish while a custom claim is pending. | [optional] 
**Status** | Pointer to **string** | Status is &#x60;live&#x60;, &#x60;provisioning&#x60;, &#x60;pending_deploy&#x60; or &#x60;pending&#x60;, derived from the operator CR and never fabricated. | [optional] 
**Url** | Pointer to **string** | URL is the host as an HTTPS address. | [optional] 
**Verified** | Pointer to **bool** | Verified is whether ownership is settled — always true for a host the org structurally owns. | [optional] 

## Methods

### NewPlatformDomainView

`func NewPlatformDomainView() *PlatformDomainView`

NewPlatformDomainView instantiates a new PlatformDomainView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDomainViewWithDefaults

`func NewPlatformDomainViewWithDefaults() *PlatformDomainView`

NewPlatformDomainViewWithDefaults instantiates a new PlatformDomainView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PlatformDomainView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlatformDomainView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlatformDomainView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlatformDomainView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDetail

`func (o *PlatformDomainView) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *PlatformDomainView) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *PlatformDomainView) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *PlatformDomainView) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetHost

`func (o *PlatformDomainView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PlatformDomainView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PlatformDomainView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PlatformDomainView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *PlatformDomainView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PlatformDomainView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PlatformDomainView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PlatformDomainView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPrimary

`func (o *PlatformDomainView) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PlatformDomainView) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PlatformDomainView) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *PlatformDomainView) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetRecords

`func (o *PlatformDomainView) GetRecords() []PlatformRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *PlatformDomainView) GetRecordsOk() (*[]PlatformRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *PlatformDomainView) SetRecords(v []PlatformRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *PlatformDomainView) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformDomainView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformDomainView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformDomainView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformDomainView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *PlatformDomainView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PlatformDomainView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PlatformDomainView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PlatformDomainView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVerified

`func (o *PlatformDomainView) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PlatformDomainView) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PlatformDomainView) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *PlatformDomainView) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


