# CloudBotSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projected** | Pointer to **int32** | Projected is how many roster entries the reconcile touched. | [optional] 
**Synced** | Pointer to **bool** | Synced is true when the reconcile ran. | [optional] 

## Methods

### NewCloudBotSync

`func NewCloudBotSync() *CloudBotSync`

NewCloudBotSync instantiates a new CloudBotSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBotSyncWithDefaults

`func NewCloudBotSyncWithDefaults() *CloudBotSync`

NewCloudBotSyncWithDefaults instantiates a new CloudBotSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjected

`func (o *CloudBotSync) GetProjected() int32`

GetProjected returns the Projected field if non-nil, zero value otherwise.

### GetProjectedOk

`func (o *CloudBotSync) GetProjectedOk() (*int32, bool)`

GetProjectedOk returns a tuple with the Projected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjected

`func (o *CloudBotSync) SetProjected(v int32)`

SetProjected sets Projected field to given value.

### HasProjected

`func (o *CloudBotSync) HasProjected() bool`

HasProjected returns a boolean if a field has been set.

### GetSynced

`func (o *CloudBotSync) GetSynced() bool`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *CloudBotSync) GetSyncedOk() (*bool, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *CloudBotSync) SetSynced(v bool)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *CloudBotSync) HasSynced() bool`

HasSynced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


