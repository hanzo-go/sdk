# ProjectProjectsDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the host was claimed, as Unix seconds — not when it went live. | [optional] 
**Detail** | Pointer to **string** | Detail is what is holding the claim up, in words a person can act on. | [optional] 
**Host** | Pointer to **string** | Host is the custom hostname claimed for this site. | [optional] 
**Records** | Pointer to [**[]ProjectRecord**](ProjectRecord.md) | Records are EXACTLY the DNS records to publish to prove ownership and route the host. Present only while pending, because a live host has already proved it; absent is therefore \&quot;nothing left to do\&quot;, not \&quot;we cannot say what to do\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is &#x60;live&#x60; when the edge answers for this host now, &#x60;pending&#x60; while the claim is waiting on DNS proof of ownership. A pending host is claimed but serves nothing. | [optional] 
**Url** | Pointer to **string** | URL is where the host will serve once it is live — present on a pending claim too, so a console can show the destination before it works. | [optional] 
**Verified** | Pointer to **bool** | Verified is the same fact as a boolean, for a caller that only needs the yes or no. It cannot disagree with status. | [optional] 

## Methods

### NewProjectProjectsDomain

`func NewProjectProjectsDomain() *ProjectProjectsDomain`

NewProjectProjectsDomain instantiates a new ProjectProjectsDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsDomainWithDefaults

`func NewProjectProjectsDomainWithDefaults() *ProjectProjectsDomain`

NewProjectProjectsDomainWithDefaults instantiates a new ProjectProjectsDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ProjectProjectsDomain) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectProjectsDomain) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectProjectsDomain) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectProjectsDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDetail

`func (o *ProjectProjectsDomain) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ProjectProjectsDomain) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ProjectProjectsDomain) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ProjectProjectsDomain) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetHost

`func (o *ProjectProjectsDomain) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProjectProjectsDomain) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProjectProjectsDomain) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ProjectProjectsDomain) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetRecords

`func (o *ProjectProjectsDomain) GetRecords() []ProjectRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ProjectProjectsDomain) GetRecordsOk() (*[]ProjectRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ProjectProjectsDomain) SetRecords(v []ProjectRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ProjectProjectsDomain) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectProjectsDomain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectProjectsDomain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectProjectsDomain) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectProjectsDomain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *ProjectProjectsDomain) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectProjectsDomain) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectProjectsDomain) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectProjectsDomain) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVerified

`func (o *ProjectProjectsDomain) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ProjectProjectsDomain) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ProjectProjectsDomain) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ProjectProjectsDomain) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


