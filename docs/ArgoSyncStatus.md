# ArgoSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to **string** | Revision is what Status was reached against. For an App CR that is the declared IMAGE TAG, not a commit — the CR is image-pinned. For a CD row it is the commit CD last applied. | [optional] 
**Status** | Pointer to **string** | Status is the ArgoCD sync vocabulary, Capitalized: Synced, OutOfSync or Unknown. For an App CR it compares the tag the CR DECLARES against the tag the cluster&#39;s Deployment is RUNNING — equal is Synced, both known and different is OutOfSync, either unknown is Unknown. For a CD row it is CD&#39;s own git-versus-cluster verdict. | [optional] 

## Methods

### NewArgoSyncStatus

`func NewArgoSyncStatus() *ArgoSyncStatus`

NewArgoSyncStatus instantiates a new ArgoSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoSyncStatusWithDefaults

`func NewArgoSyncStatusWithDefaults() *ArgoSyncStatus`

NewArgoSyncStatusWithDefaults instantiates a new ArgoSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ArgoSyncStatus) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ArgoSyncStatus) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ArgoSyncStatus) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ArgoSyncStatus) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStatus

`func (o *ArgoSyncStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArgoSyncStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArgoSyncStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArgoSyncStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


