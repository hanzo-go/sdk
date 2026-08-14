# Enrollment

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

### NewEnrollment

`func NewEnrollment() *Enrollment`

NewEnrollment instantiates a new Enrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentWithDefaults

`func NewEnrollmentWithDefaults() *Enrollment`

NewEnrollmentWithDefaults instantiates a new Enrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Enrollment) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Enrollment) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Enrollment) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Enrollment) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetChannel

`func (o *Enrollment) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Enrollment) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Enrollment) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Enrollment) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCurrentStep

`func (o *Enrollment) GetCurrentStep() int32`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *Enrollment) GetCurrentStepOk() (*int32, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *Enrollment) SetCurrentStep(v int32)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *Enrollment) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### GetEnrolledAt

`func (o *Enrollment) GetEnrolledAt() int32`

GetEnrolledAt returns the EnrolledAt field if non-nil, zero value otherwise.

### GetEnrolledAtOk

`func (o *Enrollment) GetEnrolledAtOk() (*int32, bool)`

GetEnrolledAtOk returns a tuple with the EnrolledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledAt

`func (o *Enrollment) SetEnrolledAt(v int32)`

SetEnrolledAt sets EnrolledAt field to given value.

### HasEnrolledAt

`func (o *Enrollment) HasEnrolledAt() bool`

HasEnrolledAt returns a boolean if a field has been set.

### GetId

`func (o *Enrollment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Enrollment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Enrollment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Enrollment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNextRunAt

`func (o *Enrollment) GetNextRunAt() int32`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *Enrollment) GetNextRunAtOk() (*int32, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *Enrollment) SetNextRunAt(v int32)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *Enrollment) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### GetSequenceId

`func (o *Enrollment) GetSequenceId() string`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *Enrollment) GetSequenceIdOk() (*string, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *Enrollment) SetSequenceId(v string)`

SetSequenceId sets SequenceId field to given value.

### HasSequenceId

`func (o *Enrollment) HasSequenceId() bool`

HasSequenceId returns a boolean if a field has been set.

### GetStatus

`func (o *Enrollment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Enrollment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Enrollment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Enrollment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Enrollment) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Enrollment) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Enrollment) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Enrollment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


