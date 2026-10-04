# LabelRiskDisposeOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | Pointer to **string** | Before echoes the retention boundary that was applied, RFC 3339 in UTC, as this plane parsed it from the request. What was disposed of is every record WRITTEN strictly before it and not under litigation hold — written, measured against the server clock at the write, and not against the event or observation times the asserting caller supplies, because a tenant that could back-date could delete a compliance record on demand. A boundary younger than the platform floor of five years is refused before anything is removed. | [optional] 
**Disposed** | Pointer to **int64** | Disposed is how many whole records were removed. Records are disposed of whole, never redacted: a partially-erased compliance record is one nobody can attest to. | [optional] 
**Held** | Pointer to **int64** | Held is how many records inside the boundary were kept under litigation hold. | [optional] 
**Oldest** | Pointer to **string** | Oldest is the WRITE time of the oldest assertion this tenant still holds after the sweep, RFC 3339, and it is omitted exactly when nothing remains at all. Still older than Before means records survived on purpose and says which mechanism kept them: a litigation hold (Held), or the per-call bound with more to sweep on the next call (Remaining). | [optional] 
**Remaining** | Pointer to **int64** | Remaining is how many disposable records are still older than the boundary. A sweep is bounded per call, so a non-zero value here means call again rather than that something failed. | [optional] 
**Restored** | Pointer to **int64** | Restored is how many records this sweep had already removed from the derived columnar copy and then did NOT dispose of, because a litigation hold arrived between the identify and the delete — and which were therefore written back to the derived copy before this answered.  It is a NAMED state and not a silent repair. The copy is swept before the record so nothing is orphaned in the warehouse, which means a record the delete declines to remove is one the warehouse has already lost, with its seq behind the delivery cursor and no retry that can reach it. Non-zero here says the collision happened and was repaired; a non-zero that keeps recurring says retention and hold are racing on the same records, which is worth an operator&#39;s attention rather than a debug line. | [optional] 
**Total** | Pointer to **int64** | Total and Oldest describe what the tenant still holds afterwards, so a disposal that removed nothing is distinguishable from a tenant that had nothing. | [optional] 

## Methods

### NewLabelRiskDisposeOut

`func NewLabelRiskDisposeOut() *LabelRiskDisposeOut`

NewLabelRiskDisposeOut instantiates a new LabelRiskDisposeOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelRiskDisposeOutWithDefaults

`func NewLabelRiskDisposeOutWithDefaults() *LabelRiskDisposeOut`

NewLabelRiskDisposeOutWithDefaults instantiates a new LabelRiskDisposeOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *LabelRiskDisposeOut) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *LabelRiskDisposeOut) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *LabelRiskDisposeOut) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *LabelRiskDisposeOut) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetDisposed

`func (o *LabelRiskDisposeOut) GetDisposed() int64`

GetDisposed returns the Disposed field if non-nil, zero value otherwise.

### GetDisposedOk

`func (o *LabelRiskDisposeOut) GetDisposedOk() (*int64, bool)`

GetDisposedOk returns a tuple with the Disposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposed

`func (o *LabelRiskDisposeOut) SetDisposed(v int64)`

SetDisposed sets Disposed field to given value.

### HasDisposed

`func (o *LabelRiskDisposeOut) HasDisposed() bool`

HasDisposed returns a boolean if a field has been set.

### GetHeld

`func (o *LabelRiskDisposeOut) GetHeld() int64`

GetHeld returns the Held field if non-nil, zero value otherwise.

### GetHeldOk

`func (o *LabelRiskDisposeOut) GetHeldOk() (*int64, bool)`

GetHeldOk returns a tuple with the Held field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeld

`func (o *LabelRiskDisposeOut) SetHeld(v int64)`

SetHeld sets Held field to given value.

### HasHeld

`func (o *LabelRiskDisposeOut) HasHeld() bool`

HasHeld returns a boolean if a field has been set.

### GetOldest

`func (o *LabelRiskDisposeOut) GetOldest() string`

GetOldest returns the Oldest field if non-nil, zero value otherwise.

### GetOldestOk

`func (o *LabelRiskDisposeOut) GetOldestOk() (*string, bool)`

GetOldestOk returns a tuple with the Oldest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldest

`func (o *LabelRiskDisposeOut) SetOldest(v string)`

SetOldest sets Oldest field to given value.

### HasOldest

`func (o *LabelRiskDisposeOut) HasOldest() bool`

HasOldest returns a boolean if a field has been set.

### GetRemaining

`func (o *LabelRiskDisposeOut) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *LabelRiskDisposeOut) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *LabelRiskDisposeOut) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *LabelRiskDisposeOut) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetRestored

`func (o *LabelRiskDisposeOut) GetRestored() int64`

GetRestored returns the Restored field if non-nil, zero value otherwise.

### GetRestoredOk

`func (o *LabelRiskDisposeOut) GetRestoredOk() (*int64, bool)`

GetRestoredOk returns a tuple with the Restored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestored

`func (o *LabelRiskDisposeOut) SetRestored(v int64)`

SetRestored sets Restored field to given value.

### HasRestored

`func (o *LabelRiskDisposeOut) HasRestored() bool`

HasRestored returns a boolean if a field has been set.

### GetTotal

`func (o *LabelRiskDisposeOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LabelRiskDisposeOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LabelRiskDisposeOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *LabelRiskDisposeOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


