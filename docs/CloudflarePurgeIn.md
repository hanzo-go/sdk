# CloudflarePurgeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to **[]string** | Files purges exactly the listed URLs — at most 30, Cloudflare&#39;s per-request cap. | [optional] 
**PurgeEverything** | Pointer to **bool** | Everything drops the zone&#39;s entire edge cache. | [optional] 
**Zone** | Pointer to **string** | Zone is the 32-hex Cloudflare zone id, from the path. | [optional] 

## Methods

### NewCloudflarePurgeIn

`func NewCloudflarePurgeIn() *CloudflarePurgeIn`

NewCloudflarePurgeIn instantiates a new CloudflarePurgeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudflarePurgeInWithDefaults

`func NewCloudflarePurgeInWithDefaults() *CloudflarePurgeIn`

NewCloudflarePurgeInWithDefaults instantiates a new CloudflarePurgeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *CloudflarePurgeIn) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CloudflarePurgeIn) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CloudflarePurgeIn) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CloudflarePurgeIn) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPurgeEverything

`func (o *CloudflarePurgeIn) GetPurgeEverything() bool`

GetPurgeEverything returns the PurgeEverything field if non-nil, zero value otherwise.

### GetPurgeEverythingOk

`func (o *CloudflarePurgeIn) GetPurgeEverythingOk() (*bool, bool)`

GetPurgeEverythingOk returns a tuple with the PurgeEverything field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeEverything

`func (o *CloudflarePurgeIn) SetPurgeEverything(v bool)`

SetPurgeEverything sets PurgeEverything field to given value.

### HasPurgeEverything

`func (o *CloudflarePurgeIn) HasPurgeEverything() bool`

HasPurgeEverything returns a boolean if a field has been set.

### GetZone

`func (o *CloudflarePurgeIn) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CloudflarePurgeIn) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CloudflarePurgeIn) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CloudflarePurgeIn) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


