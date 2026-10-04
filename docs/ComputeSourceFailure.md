# ComputeSourceFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Reason is a terse, log-safe summary — never the upstream&#39;s response body. | [optional] 
**Source** | Pointer to **string** | Source is the dependency that failed, named as an operator names it. | [optional] 

## Methods

### NewComputeSourceFailure

`func NewComputeSourceFailure() *ComputeSourceFailure`

NewComputeSourceFailure instantiates a new ComputeSourceFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeSourceFailureWithDefaults

`func NewComputeSourceFailureWithDefaults() *ComputeSourceFailure`

NewComputeSourceFailureWithDefaults instantiates a new ComputeSourceFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ComputeSourceFailure) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ComputeSourceFailure) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ComputeSourceFailure) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ComputeSourceFailure) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSource

`func (o *ComputeSourceFailure) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ComputeSourceFailure) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ComputeSourceFailure) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ComputeSourceFailure) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


