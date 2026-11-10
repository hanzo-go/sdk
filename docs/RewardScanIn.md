# RewardScanIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | Handle is your handle there; on Mastodon user@instance. | [optional] 
**Network** | Pointer to **string** | Network is bluesky, mastodon, reddit, hn or devto. | [optional] 

## Methods

### NewRewardScanIn

`func NewRewardScanIn() *RewardScanIn`

NewRewardScanIn instantiates a new RewardScanIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardScanInWithDefaults

`func NewRewardScanInWithDefaults() *RewardScanIn`

NewRewardScanInWithDefaults instantiates a new RewardScanIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *RewardScanIn) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *RewardScanIn) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *RewardScanIn) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *RewardScanIn) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetNetwork

`func (o *RewardScanIn) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RewardScanIn) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RewardScanIn) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RewardScanIn) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


