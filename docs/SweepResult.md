# SweepResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qualified** | Pointer to **int32** | Qualified is how many of those referrals qualified on this pass. | [optional] 
**Swept** | Pointer to **int32** | Swept is how many pending referrals were checked. | [optional] 

## Methods

### NewSweepResult

`func NewSweepResult() *SweepResult`

NewSweepResult instantiates a new SweepResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepResultWithDefaults

`func NewSweepResultWithDefaults() *SweepResult`

NewSweepResultWithDefaults instantiates a new SweepResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualified

`func (o *SweepResult) GetQualified() int32`

GetQualified returns the Qualified field if non-nil, zero value otherwise.

### GetQualifiedOk

`func (o *SweepResult) GetQualifiedOk() (*int32, bool)`

GetQualifiedOk returns a tuple with the Qualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualified

`func (o *SweepResult) SetQualified(v int32)`

SetQualified sets Qualified field to given value.

### HasQualified

`func (o *SweepResult) HasQualified() bool`

HasQualified returns a boolean if a field has been set.

### GetSwept

`func (o *SweepResult) GetSwept() int32`

GetSwept returns the Swept field if non-nil, zero value otherwise.

### GetSweptOk

`func (o *SweepResult) GetSweptOk() (*int32, bool)`

GetSweptOk returns a tuple with the Swept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwept

`func (o *SweepResult) SetSwept(v int32)`

SetSwept sets Swept field to given value.

### HasSwept

`func (o *SweepResult) HasSwept() bool`

HasSwept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


