# CloudEnrollResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyEnrolled** | Pointer to **int32** | AlreadyEnrolled is how many this sequence had already taken and were left alone. | [optional] 
**Enrolled** | Pointer to **int32** | Enrolled is how many started a walk on this call. | [optional] 
**EnrollmentId** | Pointer to **string** | EnrollmentID names the walk, and is present ONLY for a single-address enroll — a fan-out has many, and reporting one of them would be a lie. | [optional] 
**Resolved** | Pointer to **int32** | Resolved is how many addresses the request named — 1 for an address, the audience&#39;s deliverable count for an audience. | [optional] 

## Methods

### NewCloudEnrollResult

`func NewCloudEnrollResult() *CloudEnrollResult`

NewCloudEnrollResult instantiates a new CloudEnrollResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEnrollResultWithDefaults

`func NewCloudEnrollResultWithDefaults() *CloudEnrollResult`

NewCloudEnrollResultWithDefaults instantiates a new CloudEnrollResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyEnrolled

`func (o *CloudEnrollResult) GetAlreadyEnrolled() int32`

GetAlreadyEnrolled returns the AlreadyEnrolled field if non-nil, zero value otherwise.

### GetAlreadyEnrolledOk

`func (o *CloudEnrollResult) GetAlreadyEnrolledOk() (*int32, bool)`

GetAlreadyEnrolledOk returns a tuple with the AlreadyEnrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyEnrolled

`func (o *CloudEnrollResult) SetAlreadyEnrolled(v int32)`

SetAlreadyEnrolled sets AlreadyEnrolled field to given value.

### HasAlreadyEnrolled

`func (o *CloudEnrollResult) HasAlreadyEnrolled() bool`

HasAlreadyEnrolled returns a boolean if a field has been set.

### GetEnrolled

`func (o *CloudEnrollResult) GetEnrolled() int32`

GetEnrolled returns the Enrolled field if non-nil, zero value otherwise.

### GetEnrolledOk

`func (o *CloudEnrollResult) GetEnrolledOk() (*int32, bool)`

GetEnrolledOk returns a tuple with the Enrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolled

`func (o *CloudEnrollResult) SetEnrolled(v int32)`

SetEnrolled sets Enrolled field to given value.

### HasEnrolled

`func (o *CloudEnrollResult) HasEnrolled() bool`

HasEnrolled returns a boolean if a field has been set.

### GetEnrollmentId

`func (o *CloudEnrollResult) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CloudEnrollResult) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CloudEnrollResult) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CloudEnrollResult) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### GetResolved

`func (o *CloudEnrollResult) GetResolved() int32`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *CloudEnrollResult) GetResolvedOk() (*int32, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *CloudEnrollResult) SetResolved(v int32)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *CloudEnrollResult) HasResolved() bool`

HasResolved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


