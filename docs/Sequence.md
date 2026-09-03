# Sequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is unix seconds when the sequence was registered, server-assigned and never rewritten. | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned sequence id (\&quot;seq_\&quot; + 128 random bits). | [optional] 
**Name** | Pointer to **string** | Name is the sequence&#39;s label. Required, trimmed, capped at 1024 bytes. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle: draft, active or archived. Empty means draft, and ONLY an active sequence accepts enrollments. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is unix seconds of the last status flip, server-assigned, and the key the sequence list is ordered by (newest first). Adding a step or enrolling a contact does NOT touch it — only draft/active/archived does — so it tracks activation rather than activity. | [optional] 

## Methods

### NewSequence

`func NewSequence() *Sequence`

NewSequence instantiates a new Sequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSequenceWithDefaults

`func NewSequenceWithDefaults() *Sequence`

NewSequenceWithDefaults instantiates a new Sequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Sequence) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Sequence) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Sequence) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Sequence) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *Sequence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sequence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sequence) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Sequence) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Sequence) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sequence) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sequence) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sequence) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *Sequence) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Sequence) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Sequence) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Sequence) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Sequence) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Sequence) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Sequence) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Sequence) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


