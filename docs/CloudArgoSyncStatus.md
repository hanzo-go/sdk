# CloudArgoSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudArgoSyncStatus

`func NewCloudArgoSyncStatus() *CloudArgoSyncStatus`

NewCloudArgoSyncStatus instantiates a new CloudArgoSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoSyncStatusWithDefaults

`func NewCloudArgoSyncStatusWithDefaults() *CloudArgoSyncStatus`

NewCloudArgoSyncStatusWithDefaults instantiates a new CloudArgoSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *CloudArgoSyncStatus) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CloudArgoSyncStatus) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CloudArgoSyncStatus) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CloudArgoSyncStatus) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStatus

`func (o *CloudArgoSyncStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudArgoSyncStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudArgoSyncStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudArgoSyncStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


