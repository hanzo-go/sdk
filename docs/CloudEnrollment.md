# CloudEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the normalized (lower-cased, trimmed) recipient. | [optional] 
**Channel** | Pointer to **string** | Channel is the delivery surface the steps go out on. | [optional] 
**CurrentStep** | Pointer to **int32** | CurrentStep is the index of the step that sends next. | [optional] 
**EnrolledAt** | Pointer to **int32** | EnrolledAt and UpdatedAt are unix seconds. | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned enrollment id (\&quot;enr_\&quot; + 128 random bits). | [optional] 
**NextRunAt** | Pointer to **int32** | NextRunAt is the unix time the current step comes due; 0 once the walk has ended. It IS the schedule — durable in SQLite, so it survives restarts. | [optional] 
**SequenceId** | Pointer to **string** | SequenceID is the sequence being walked. | [optional] 
**Status** | Pointer to **string** | Status is active, completed or canceled. | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudEnrollment

`func NewCloudEnrollment() *CloudEnrollment`

NewCloudEnrollment instantiates a new CloudEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEnrollmentWithDefaults

`func NewCloudEnrollmentWithDefaults() *CloudEnrollment`

NewCloudEnrollmentWithDefaults instantiates a new CloudEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CloudEnrollment) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CloudEnrollment) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CloudEnrollment) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CloudEnrollment) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetChannel

`func (o *CloudEnrollment) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CloudEnrollment) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CloudEnrollment) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CloudEnrollment) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCurrentStep

`func (o *CloudEnrollment) GetCurrentStep() int32`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *CloudEnrollment) GetCurrentStepOk() (*int32, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *CloudEnrollment) SetCurrentStep(v int32)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *CloudEnrollment) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### GetEnrolledAt

`func (o *CloudEnrollment) GetEnrolledAt() int32`

GetEnrolledAt returns the EnrolledAt field if non-nil, zero value otherwise.

### GetEnrolledAtOk

`func (o *CloudEnrollment) GetEnrolledAtOk() (*int32, bool)`

GetEnrolledAtOk returns a tuple with the EnrolledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledAt

`func (o *CloudEnrollment) SetEnrolledAt(v int32)`

SetEnrolledAt sets EnrolledAt field to given value.

### HasEnrolledAt

`func (o *CloudEnrollment) HasEnrolledAt() bool`

HasEnrolledAt returns a boolean if a field has been set.

### GetId

`func (o *CloudEnrollment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudEnrollment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudEnrollment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudEnrollment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNextRunAt

`func (o *CloudEnrollment) GetNextRunAt() int32`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *CloudEnrollment) GetNextRunAtOk() (*int32, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *CloudEnrollment) SetNextRunAt(v int32)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *CloudEnrollment) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### GetSequenceId

`func (o *CloudEnrollment) GetSequenceId() string`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *CloudEnrollment) GetSequenceIdOk() (*string, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *CloudEnrollment) SetSequenceId(v string)`

SetSequenceId sets SequenceId field to given value.

### HasSequenceId

`func (o *CloudEnrollment) HasSequenceId() bool`

HasSequenceId returns a boolean if a field has been set.

### GetStatus

`func (o *CloudEnrollment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudEnrollment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudEnrollment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudEnrollment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudEnrollment) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudEnrollment) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudEnrollment) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudEnrollment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


