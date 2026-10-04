# MarketingEnrollResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyEnrolled** | Pointer to **int64** | AlreadyEnrolled is how many this sequence had already taken and were left alone. | [optional] 
**Enrolled** | Pointer to **int64** | Enrolled is how many started a walk on this call. | [optional] 
**EnrollmentId** | Pointer to **string** | EnrollmentID names the walk, and is present ONLY for a single-address enroll — a fan-out has many, and reporting one of them would be a lie. | [optional] 
**Resolved** | Pointer to **int64** | Resolved is how many addresses the request named — 1 for an address, the audience&#39;s deliverable count for an audience. | [optional] 

## Methods

### NewMarketingEnrollResult

`func NewMarketingEnrollResult() *MarketingEnrollResult`

NewMarketingEnrollResult instantiates a new MarketingEnrollResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEnrollResultWithDefaults

`func NewMarketingEnrollResultWithDefaults() *MarketingEnrollResult`

NewMarketingEnrollResultWithDefaults instantiates a new MarketingEnrollResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyEnrolled

`func (o *MarketingEnrollResult) GetAlreadyEnrolled() int64`

GetAlreadyEnrolled returns the AlreadyEnrolled field if non-nil, zero value otherwise.

### GetAlreadyEnrolledOk

`func (o *MarketingEnrollResult) GetAlreadyEnrolledOk() (*int64, bool)`

GetAlreadyEnrolledOk returns a tuple with the AlreadyEnrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyEnrolled

`func (o *MarketingEnrollResult) SetAlreadyEnrolled(v int64)`

SetAlreadyEnrolled sets AlreadyEnrolled field to given value.

### HasAlreadyEnrolled

`func (o *MarketingEnrollResult) HasAlreadyEnrolled() bool`

HasAlreadyEnrolled returns a boolean if a field has been set.

### GetEnrolled

`func (o *MarketingEnrollResult) GetEnrolled() int64`

GetEnrolled returns the Enrolled field if non-nil, zero value otherwise.

### GetEnrolledOk

`func (o *MarketingEnrollResult) GetEnrolledOk() (*int64, bool)`

GetEnrolledOk returns a tuple with the Enrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolled

`func (o *MarketingEnrollResult) SetEnrolled(v int64)`

SetEnrolled sets Enrolled field to given value.

### HasEnrolled

`func (o *MarketingEnrollResult) HasEnrolled() bool`

HasEnrolled returns a boolean if a field has been set.

### GetEnrollmentId

`func (o *MarketingEnrollResult) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *MarketingEnrollResult) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *MarketingEnrollResult) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *MarketingEnrollResult) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### GetResolved

`func (o *MarketingEnrollResult) GetResolved() int64`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *MarketingEnrollResult) GetResolvedOk() (*int64, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *MarketingEnrollResult) SetResolved(v int64)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *MarketingEnrollResult) HasResolved() bool`

HasResolved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


