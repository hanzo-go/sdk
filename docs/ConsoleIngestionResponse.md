# ConsoleIngestionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successes** | Pointer to [**[]ConsoleIngestionResponseSuccessesInner**](ConsoleIngestionResponseSuccessesInner.md) |  | [optional] 
**Errors** | Pointer to [**[]ConsoleIngestionResponseErrorsInner**](ConsoleIngestionResponseErrorsInner.md) |  | [optional] 

## Methods

### NewConsoleIngestionResponse

`func NewConsoleIngestionResponse() *ConsoleIngestionResponse`

NewConsoleIngestionResponse instantiates a new ConsoleIngestionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleIngestionResponseWithDefaults

`func NewConsoleIngestionResponseWithDefaults() *ConsoleIngestionResponse`

NewConsoleIngestionResponseWithDefaults instantiates a new ConsoleIngestionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccesses

`func (o *ConsoleIngestionResponse) GetSuccesses() []ConsoleIngestionResponseSuccessesInner`

GetSuccesses returns the Successes field if non-nil, zero value otherwise.

### GetSuccessesOk

`func (o *ConsoleIngestionResponse) GetSuccessesOk() (*[]ConsoleIngestionResponseSuccessesInner, bool)`

GetSuccessesOk returns a tuple with the Successes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesses

`func (o *ConsoleIngestionResponse) SetSuccesses(v []ConsoleIngestionResponseSuccessesInner)`

SetSuccesses sets Successes field to given value.

### HasSuccesses

`func (o *ConsoleIngestionResponse) HasSuccesses() bool`

HasSuccesses returns a boolean if a field has been set.

### GetErrors

`func (o *ConsoleIngestionResponse) GetErrors() []ConsoleIngestionResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ConsoleIngestionResponse) GetErrorsOk() (*[]ConsoleIngestionResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ConsoleIngestionResponse) SetErrors(v []ConsoleIngestionResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ConsoleIngestionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


