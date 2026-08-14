# PurgeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to **[]string** | Files purges exactly the listed URLs — at most 30, Cloudflare&#39;s per-request cap. | [optional] 
**PurgeEverything** | Pointer to **bool** | Everything drops the zone&#39;s entire edge cache. | [optional] 
**Zone** | Pointer to **string** | Zone is the 32-hex Cloudflare zone id, from the path. | [optional] 

## Methods

### NewPurgeIn

`func NewPurgeIn() *PurgeIn`

NewPurgeIn instantiates a new PurgeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeInWithDefaults

`func NewPurgeInWithDefaults() *PurgeIn`

NewPurgeInWithDefaults instantiates a new PurgeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *PurgeIn) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *PurgeIn) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *PurgeIn) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *PurgeIn) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPurgeEverything

`func (o *PurgeIn) GetPurgeEverything() bool`

GetPurgeEverything returns the PurgeEverything field if non-nil, zero value otherwise.

### GetPurgeEverythingOk

`func (o *PurgeIn) GetPurgeEverythingOk() (*bool, bool)`

GetPurgeEverythingOk returns a tuple with the PurgeEverything field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeEverything

`func (o *PurgeIn) SetPurgeEverything(v bool)`

SetPurgeEverything sets PurgeEverything field to given value.

### HasPurgeEverything

`func (o *PurgeIn) HasPurgeEverything() bool`

HasPurgeEverything returns a boolean if a field has been set.

### GetZone

`func (o *PurgeIn) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PurgeIn) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PurgeIn) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *PurgeIn) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


