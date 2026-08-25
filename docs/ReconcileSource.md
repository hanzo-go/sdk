# ReconcileSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path is the directory WITHIN the repository that is rendered — everything outside it is not this plane&#39;s desired state and is never applied. | [optional] 
**Ref** | Pointer to **string** | Ref is the branch or tag the revision was resolved from. | [optional] 
**Repo** | Pointer to **string** | Repo is the clone URL of the repository holding the desired state. | [optional] 

## Methods

### NewReconcileSource

`func NewReconcileSource() *ReconcileSource`

NewReconcileSource instantiates a new ReconcileSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconcileSourceWithDefaults

`func NewReconcileSourceWithDefaults() *ReconcileSource`

NewReconcileSourceWithDefaults instantiates a new ReconcileSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ReconcileSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReconcileSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReconcileSource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ReconcileSource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRef

`func (o *ReconcileSource) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ReconcileSource) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ReconcileSource) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ReconcileSource) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRepo

`func (o *ReconcileSource) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ReconcileSource) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ReconcileSource) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ReconcileSource) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


