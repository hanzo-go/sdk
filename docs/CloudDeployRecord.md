# CloudDeployRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Created reports whether this call recorded a new attribution edge (201) or found an existing one (200). Absent when nothing was recorded. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the edge was first recorded, in unix seconds. Absent when nothing was recorded. | [optional] 
**DeployId** | Pointer to **string** | DeployID is the attribution edge&#39;s handle. Absent when nothing was recorded. | [optional] 
**Reason** | Pointer to **string** | Reason says why nothing was attributed. Present only when recorded is false. | [optional] 
**Recorded** | Pointer to **bool** | Recorded reports whether the deploy was attributed to an author at all. False is the ordinary answer for a project built from no repository, or from one no author has verified — never an error, so a deploy path can fire this unconditionally. | [optional] 
**Self** | Pointer to **bool** | Self reports that the deploying org IS the author&#39;s org. Such a deploy is recorded for provenance but excluded from accrual. Absent when nothing was recorded. | [optional] 

## Methods

### NewCloudDeployRecord

`func NewCloudDeployRecord() *CloudDeployRecord`

NewCloudDeployRecord instantiates a new CloudDeployRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDeployRecordWithDefaults

`func NewCloudDeployRecordWithDefaults() *CloudDeployRecord`

NewCloudDeployRecordWithDefaults instantiates a new CloudDeployRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CloudDeployRecord) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudDeployRecord) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudDeployRecord) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudDeployRecord) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudDeployRecord) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudDeployRecord) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudDeployRecord) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudDeployRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeployId

`func (o *CloudDeployRecord) GetDeployId() string`

GetDeployId returns the DeployId field if non-nil, zero value otherwise.

### GetDeployIdOk

`func (o *CloudDeployRecord) GetDeployIdOk() (*string, bool)`

GetDeployIdOk returns a tuple with the DeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployId

`func (o *CloudDeployRecord) SetDeployId(v string)`

SetDeployId sets DeployId field to given value.

### HasDeployId

`func (o *CloudDeployRecord) HasDeployId() bool`

HasDeployId returns a boolean if a field has been set.

### GetReason

`func (o *CloudDeployRecord) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudDeployRecord) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudDeployRecord) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudDeployRecord) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRecorded

`func (o *CloudDeployRecord) GetRecorded() bool`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *CloudDeployRecord) GetRecordedOk() (*bool, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *CloudDeployRecord) SetRecorded(v bool)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *CloudDeployRecord) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetSelf

`func (o *CloudDeployRecord) GetSelf() bool`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CloudDeployRecord) GetSelfOk() (*bool, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CloudDeployRecord) SetSelf(v bool)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *CloudDeployRecord) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


