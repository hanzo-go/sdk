# PatentPatentAnchor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when the commitment was made, RFC3339, when there is one. | [optional] 
**Commitment** | Pointer to **string** | Commitment is the fold root the anchor would name. | [optional] 
**Reference** | Pointer to **string** | Reference is where a third party can check it — a block, an index, a published headline. Empty means nothing outside this response names it. | [optional] 
**Status** | Pointer to **string** | Status is unanchored or anchored. Nothing else: an anchored answer arrives only from a real commitment, and this app never invents a transaction id. | [optional] 

## Methods

### NewPatentPatentAnchor

`func NewPatentPatentAnchor() *PatentPatentAnchor`

NewPatentPatentAnchor instantiates a new PatentPatentAnchor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentAnchorWithDefaults

`func NewPatentPatentAnchorWithDefaults() *PatentPatentAnchor`

NewPatentPatentAnchorWithDefaults instantiates a new PatentPatentAnchor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *PatentPatentAnchor) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *PatentPatentAnchor) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *PatentPatentAnchor) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *PatentPatentAnchor) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetCommitment

`func (o *PatentPatentAnchor) GetCommitment() string`

GetCommitment returns the Commitment field if non-nil, zero value otherwise.

### GetCommitmentOk

`func (o *PatentPatentAnchor) GetCommitmentOk() (*string, bool)`

GetCommitmentOk returns a tuple with the Commitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitment

`func (o *PatentPatentAnchor) SetCommitment(v string)`

SetCommitment sets Commitment field to given value.

### HasCommitment

`func (o *PatentPatentAnchor) HasCommitment() bool`

HasCommitment returns a boolean if a field has been set.

### GetReference

`func (o *PatentPatentAnchor) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PatentPatentAnchor) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PatentPatentAnchor) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PatentPatentAnchor) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *PatentPatentAnchor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatentPatentAnchor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatentPatentAnchor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatentPatentAnchor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


