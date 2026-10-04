# AutoStarter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**AutoSchedule**](AutoSchedule.md) |  | [optional] 

## Methods

### NewAutoStarter

`func NewAutoStarter() *AutoStarter`

NewAutoStarter instantiates a new AutoStarter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoStarterWithDefaults

`func NewAutoStarterWithDefaults() *AutoStarter`

NewAutoStarterWithDefaults instantiates a new AutoStarter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AutoStarter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoStarter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoStarter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoStarter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *AutoStarter) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AutoStarter) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AutoStarter) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AutoStarter) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetInstructions

`func (o *AutoStarter) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AutoStarter) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AutoStarter) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AutoStarter) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetKey

`func (o *AutoStarter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AutoStarter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AutoStarter) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AutoStarter) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *AutoStarter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoStarter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoStarter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoStarter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *AutoStarter) GetSchedule() AutoSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AutoStarter) GetScheduleOk() (*AutoSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AutoStarter) SetSchedule(v AutoSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AutoStarter) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


