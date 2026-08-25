# AppliedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message is the engine&#39;s own sentence about this object — the apiserver&#39;s refusal on a failure, and typically empty on success. It is for a human reading a failed run, not a value to branch on. | [optional] 
**Resource** | Pointer to **string** | Resource identifies the object as group/version/kind/namespace/name, the engine&#39;s own key. It is stable across runs, so two reports can be diffed on it. | [optional] 
**Status** | Pointer to **string** | Status is what happened to this object, from the engine&#39;s closed vocabulary: &#x60;Synced&#x60; (applied), &#x60;Pruned&#x60; (deleted because the source no longer declares it), &#x60;SyncFailed&#x60; (refused — read Message) and &#x60;PruneSkipped&#x60; (deletion was declined). It is the per-object detail behind the report&#39;s counts. | [optional] 

## Methods

### NewAppliedResource

`func NewAppliedResource() *AppliedResource`

NewAppliedResource instantiates a new AppliedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedResourceWithDefaults

`func NewAppliedResourceWithDefaults() *AppliedResource`

NewAppliedResourceWithDefaults instantiates a new AppliedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AppliedResource) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AppliedResource) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AppliedResource) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AppliedResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResource

`func (o *AppliedResource) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AppliedResource) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AppliedResource) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AppliedResource) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetStatus

`func (o *AppliedResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppliedResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppliedResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppliedResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


