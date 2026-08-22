# ProjectsDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when the host was claimed, as Unix seconds — not when it went live. | [optional] 
**Detail** | Pointer to **string** | Detail is what is holding the claim up, in words a person can act on. | [optional] 
**Host** | Pointer to **string** | Host is the custom hostname claimed for this site. | [optional] 
**Records** | Pointer to [**[]Record**](Record.md) | Records are EXACTLY the DNS records to publish to prove ownership and route the host. Present only while pending, because a live host has already proved it; absent is therefore \&quot;nothing left to do\&quot;, not \&quot;we cannot say what to do\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is &#x60;live&#x60; when the edge answers for this host now, &#x60;pending&#x60; while the claim is waiting on DNS proof of ownership. A pending host is claimed but serves nothing. | [optional] 
**Url** | Pointer to **string** | URL is where the host will serve once it is live — present on a pending claim too, so a console can show the destination before it works. | [optional] 
**Verified** | Pointer to **bool** | Verified is the same fact as a boolean, for a caller that only needs the yes or no. It cannot disagree with status. | [optional] 

## Methods

### NewProjectsDomain

`func NewProjectsDomain() *ProjectsDomain`

NewProjectsDomain instantiates a new ProjectsDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsDomainWithDefaults

`func NewProjectsDomainWithDefaults() *ProjectsDomain`

NewProjectsDomainWithDefaults instantiates a new ProjectsDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ProjectsDomain) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectsDomain) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectsDomain) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectsDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDetail

`func (o *ProjectsDomain) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ProjectsDomain) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ProjectsDomain) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ProjectsDomain) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetHost

`func (o *ProjectsDomain) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProjectsDomain) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProjectsDomain) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ProjectsDomain) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetRecords

`func (o *ProjectsDomain) GetRecords() []Record`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ProjectsDomain) GetRecordsOk() (*[]Record, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ProjectsDomain) SetRecords(v []Record)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ProjectsDomain) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectsDomain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectsDomain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectsDomain) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectsDomain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *ProjectsDomain) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectsDomain) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectsDomain) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectsDomain) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVerified

`func (o *ProjectsDomain) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ProjectsDomain) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ProjectsDomain) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ProjectsDomain) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


