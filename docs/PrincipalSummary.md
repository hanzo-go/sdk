# PrincipalSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Missing** | Pointer to [**[]PrincipalStep**](PrincipalStep.md) | Missing are the steps it owes, each with the rule that asks for it. | [optional] 
**Ready** | Pointer to **bool** | Ready is true when it lacks nothing its own facts decide. | [optional] 

## Methods

### NewPrincipalSummary

`func NewPrincipalSummary() *PrincipalSummary`

NewPrincipalSummary instantiates a new PrincipalSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalSummaryWithDefaults

`func NewPrincipalSummaryWithDefaults() *PrincipalSummary`

NewPrincipalSummaryWithDefaults instantiates a new PrincipalSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMissing

`func (o *PrincipalSummary) GetMissing() []PrincipalStep`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *PrincipalSummary) GetMissingOk() (*[]PrincipalStep, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *PrincipalSummary) SetMissing(v []PrincipalStep)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *PrincipalSummary) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetReady

`func (o *PrincipalSummary) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *PrincipalSummary) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *PrincipalSummary) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *PrincipalSummary) HasReady() bool`

HasReady returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


