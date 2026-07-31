# CloudSequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt and UpdatedAt are unix seconds, both server-assigned. | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned sequence id (\&quot;seq_\&quot; + 128 random bits). | [optional] 
**Name** | Pointer to **string** | Name is the sequence&#39;s label. Required, trimmed, capped at 1024 bytes. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle: draft, active or archived. Empty means draft, and ONLY an active sequence accepts enrollments. | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudSequence

`func NewCloudSequence() *CloudSequence`

NewCloudSequence instantiates a new CloudSequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSequenceWithDefaults

`func NewCloudSequenceWithDefaults() *CloudSequence`

NewCloudSequenceWithDefaults instantiates a new CloudSequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudSequence) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSequence) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSequence) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSequence) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudSequence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSequence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSequence) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSequence) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudSequence) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSequence) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSequence) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudSequence) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *CloudSequence) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSequence) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSequence) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudSequence) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudSequence) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudSequence) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudSequence) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudSequence) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


