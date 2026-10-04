# AutoRunStarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Run** | Pointer to [**AutoRunHandle**](AutoRunHandle.md) | Run is the run this started, or the skipped start when one was going. | [optional] 

## Methods

### NewAutoRunStarted

`func NewAutoRunStarted() *AutoRunStarted`

NewAutoRunStarted instantiates a new AutoRunStarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRunStartedWithDefaults

`func NewAutoRunStartedWithDefaults() *AutoRunStarted`

NewAutoRunStartedWithDefaults instantiates a new AutoRunStarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRun

`func (o *AutoRunStarted) GetRun() AutoRunHandle`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *AutoRunStarted) GetRunOk() (*AutoRunHandle, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *AutoRunStarted) SetRun(v AutoRunHandle)`

SetRun sets Run field to given value.

### HasRun

`func (o *AutoRunStarted) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


