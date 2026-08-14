# AutoStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flow** | Pointer to **string** | Flow is the id of the flow to run. | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAutoStart

`func NewAutoStart() *AutoStart`

NewAutoStart instantiates a new AutoStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoStartWithDefaults

`func NewAutoStartWithDefaults() *AutoStart`

NewAutoStartWithDefaults instantiates a new AutoStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlow

`func (o *AutoStart) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *AutoStart) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *AutoStart) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *AutoStart) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetInput

`func (o *AutoStart) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AutoStart) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AutoStart) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *AutoStart) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *AutoStart) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *AutoStart) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


