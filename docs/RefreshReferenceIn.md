# RefreshReferenceIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Force** | Pointer to **bool** | Force accepts a take whose size moved past the change bound. A publisher serving a tenth or ten times its previous list is refused by default and the previous version is left standing; this is the operator saying the change is real. It cannot make an empty, truncated or unparseable take land — those are errors, not magnitudes. | [optional] 
**Receipts** | Pointer to [**[]ReferenceReceipt**](ReferenceReceipt.md) | Receipts are supplied by the component that holds the membership, for a set of kind attest. They are refused on any other kind, and a set of kind attest is refused without them: this plane never invents a freshness it did not observe. | [optional] 
**Set** | Pointer to **string** | Set is the set to refresh. | [optional] 

## Methods

### NewRefreshReferenceIn

`func NewRefreshReferenceIn() *RefreshReferenceIn`

NewRefreshReferenceIn instantiates a new RefreshReferenceIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshReferenceInWithDefaults

`func NewRefreshReferenceInWithDefaults() *RefreshReferenceIn`

NewRefreshReferenceInWithDefaults instantiates a new RefreshReferenceIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForce

`func (o *RefreshReferenceIn) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *RefreshReferenceIn) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *RefreshReferenceIn) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *RefreshReferenceIn) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetReceipts

`func (o *RefreshReferenceIn) GetReceipts() []ReferenceReceipt`

GetReceipts returns the Receipts field if non-nil, zero value otherwise.

### GetReceiptsOk

`func (o *RefreshReferenceIn) GetReceiptsOk() (*[]ReferenceReceipt, bool)`

GetReceiptsOk returns a tuple with the Receipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipts

`func (o *RefreshReferenceIn) SetReceipts(v []ReferenceReceipt)`

SetReceipts sets Receipts field to given value.

### HasReceipts

`func (o *RefreshReferenceIn) HasReceipts() bool`

HasReceipts returns a boolean if a field has been set.

### GetSet

`func (o *RefreshReferenceIn) GetSet() string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *RefreshReferenceIn) GetSetOk() (*string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *RefreshReferenceIn) SetSet(v string)`

SetSet sets Set field to given value.

### HasSet

`func (o *RefreshReferenceIn) HasSet() bool`

HasSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


