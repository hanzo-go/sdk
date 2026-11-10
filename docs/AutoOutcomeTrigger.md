# AutoOutcomeTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the event, for example improvement.reviewed. | [optional] 
**Source** | Pointer to **string** | Source is the producing capability, for example research. | [optional] 

## Methods

### NewAutoOutcomeTrigger

`func NewAutoOutcomeTrigger() *AutoOutcomeTrigger`

NewAutoOutcomeTrigger instantiates a new AutoOutcomeTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoOutcomeTriggerWithDefaults

`func NewAutoOutcomeTriggerWithDefaults() *AutoOutcomeTrigger`

NewAutoOutcomeTriggerWithDefaults instantiates a new AutoOutcomeTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutoOutcomeTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoOutcomeTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoOutcomeTrigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoOutcomeTrigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *AutoOutcomeTrigger) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AutoOutcomeTrigger) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AutoOutcomeTrigger) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AutoOutcomeTrigger) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


