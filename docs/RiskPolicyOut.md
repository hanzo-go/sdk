# RiskPolicyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to **int64** | Changes is how many DISTINCT regimes may be adopted per Window. A restatement identical to the regime in force mints no version and is not counted against it. | [optional] 
**Disposed** | Pointer to **int64** | Disposed is how many versions retention has taken. It is NOT a silence: a history bounded on disk must say what it no longer holds, because a decision citing a disposed version can no longer be reconstructed from this record. | [optional] 
**History** | Pointer to [**[]RiskPolicyVersion**](RiskPolicyVersion.md) | History is the retained versions, newest first. | [optional] 
**Retained** | Pointer to **int64** | Retained is how many versions this organisation&#39;s history holds at most, derived from the byte budget its rows are a multiple of. | [optional] 
**Version** | Pointer to **int64** | Version is the version in force — the one every score currently cites. Zero means no regime has ever been stated and the default posture, shadow, is in force. | [optional] 
**Window** | Pointer to **string** | Window is the period Changes is measured over. | [optional] 

## Methods

### NewRiskPolicyOut

`func NewRiskPolicyOut() *RiskPolicyOut`

NewRiskPolicyOut instantiates a new RiskPolicyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicyOutWithDefaults

`func NewRiskPolicyOutWithDefaults() *RiskPolicyOut`

NewRiskPolicyOutWithDefaults instantiates a new RiskPolicyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *RiskPolicyOut) GetChanges() int64`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *RiskPolicyOut) GetChangesOk() (*int64, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *RiskPolicyOut) SetChanges(v int64)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *RiskPolicyOut) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetDisposed

`func (o *RiskPolicyOut) GetDisposed() int64`

GetDisposed returns the Disposed field if non-nil, zero value otherwise.

### GetDisposedOk

`func (o *RiskPolicyOut) GetDisposedOk() (*int64, bool)`

GetDisposedOk returns a tuple with the Disposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposed

`func (o *RiskPolicyOut) SetDisposed(v int64)`

SetDisposed sets Disposed field to given value.

### HasDisposed

`func (o *RiskPolicyOut) HasDisposed() bool`

HasDisposed returns a boolean if a field has been set.

### GetHistory

`func (o *RiskPolicyOut) GetHistory() []RiskPolicyVersion`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *RiskPolicyOut) GetHistoryOk() (*[]RiskPolicyVersion, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *RiskPolicyOut) SetHistory(v []RiskPolicyVersion)`

SetHistory sets History field to given value.

### HasHistory

`func (o *RiskPolicyOut) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetRetained

`func (o *RiskPolicyOut) GetRetained() int64`

GetRetained returns the Retained field if non-nil, zero value otherwise.

### GetRetainedOk

`func (o *RiskPolicyOut) GetRetainedOk() (*int64, bool)`

GetRetainedOk returns a tuple with the Retained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetained

`func (o *RiskPolicyOut) SetRetained(v int64)`

SetRetained sets Retained field to given value.

### HasRetained

`func (o *RiskPolicyOut) HasRetained() bool`

HasRetained returns a boolean if a field has been set.

### GetVersion

`func (o *RiskPolicyOut) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RiskPolicyOut) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RiskPolicyOut) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RiskPolicyOut) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWindow

`func (o *RiskPolicyOut) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *RiskPolicyOut) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *RiskPolicyOut) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *RiskPolicyOut) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


