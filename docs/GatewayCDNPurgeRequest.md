# GatewayCDNPurgeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | Pointer to **[]string** | Specific URLs to purge | [optional] 
**Tags** | Pointer to **[]string** | Cache tags to purge | [optional] 
**PurgeAll** | Pointer to **bool** | Purge all cached content | [optional] [default to false]

## Methods

### NewGatewayCDNPurgeRequest

`func NewGatewayCDNPurgeRequest() *GatewayCDNPurgeRequest`

NewGatewayCDNPurgeRequest instantiates a new GatewayCDNPurgeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCDNPurgeRequestWithDefaults

`func NewGatewayCDNPurgeRequestWithDefaults() *GatewayCDNPurgeRequest`

NewGatewayCDNPurgeRequestWithDefaults instantiates a new GatewayCDNPurgeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *GatewayCDNPurgeRequest) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *GatewayCDNPurgeRequest) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *GatewayCDNPurgeRequest) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *GatewayCDNPurgeRequest) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCDNPurgeRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCDNPurgeRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCDNPurgeRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCDNPurgeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPurgeAll

`func (o *GatewayCDNPurgeRequest) GetPurgeAll() bool`

GetPurgeAll returns the PurgeAll field if non-nil, zero value otherwise.

### GetPurgeAllOk

`func (o *GatewayCDNPurgeRequest) GetPurgeAllOk() (*bool, bool)`

GetPurgeAllOk returns a tuple with the PurgeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeAll

`func (o *GatewayCDNPurgeRequest) SetPurgeAll(v bool)`

SetPurgeAll sets PurgeAll field to given value.

### HasPurgeAll

`func (o *GatewayCDNPurgeRequest) HasPurgeAll() bool`

HasPurgeAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


