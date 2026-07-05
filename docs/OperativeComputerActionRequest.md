# OperativeComputerActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The computer action to perform | 
**Text** | Pointer to **string** | Text for key/type actions, modifier key for scroll | [optional] 
**Coordinate** | Pointer to **[]int32** | [x, y] pixel coordinates for mouse actions | [optional] 
**ScrollDirection** | Pointer to **string** | Scroll direction (for scroll action) | [optional] 
**ScrollAmount** | Pointer to **int32** | Number of scroll clicks | [optional] 
**Duration** | Pointer to **float32** | Duration in seconds (for hold_key, wait actions) | [optional] 
**Key** | Pointer to **string** | Modifier key to hold during click actions | [optional] 

## Methods

### NewOperativeComputerActionRequest

`func NewOperativeComputerActionRequest(action string, ) *OperativeComputerActionRequest`

NewOperativeComputerActionRequest instantiates a new OperativeComputerActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperativeComputerActionRequestWithDefaults

`func NewOperativeComputerActionRequestWithDefaults() *OperativeComputerActionRequest`

NewOperativeComputerActionRequestWithDefaults instantiates a new OperativeComputerActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OperativeComputerActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OperativeComputerActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OperativeComputerActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetText

`func (o *OperativeComputerActionRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OperativeComputerActionRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OperativeComputerActionRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *OperativeComputerActionRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetCoordinate

`func (o *OperativeComputerActionRequest) GetCoordinate() []int32`

GetCoordinate returns the Coordinate field if non-nil, zero value otherwise.

### GetCoordinateOk

`func (o *OperativeComputerActionRequest) GetCoordinateOk() (*[]int32, bool)`

GetCoordinateOk returns a tuple with the Coordinate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinate

`func (o *OperativeComputerActionRequest) SetCoordinate(v []int32)`

SetCoordinate sets Coordinate field to given value.

### HasCoordinate

`func (o *OperativeComputerActionRequest) HasCoordinate() bool`

HasCoordinate returns a boolean if a field has been set.

### GetScrollDirection

`func (o *OperativeComputerActionRequest) GetScrollDirection() string`

GetScrollDirection returns the ScrollDirection field if non-nil, zero value otherwise.

### GetScrollDirectionOk

`func (o *OperativeComputerActionRequest) GetScrollDirectionOk() (*string, bool)`

GetScrollDirectionOk returns a tuple with the ScrollDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollDirection

`func (o *OperativeComputerActionRequest) SetScrollDirection(v string)`

SetScrollDirection sets ScrollDirection field to given value.

### HasScrollDirection

`func (o *OperativeComputerActionRequest) HasScrollDirection() bool`

HasScrollDirection returns a boolean if a field has been set.

### GetScrollAmount

`func (o *OperativeComputerActionRequest) GetScrollAmount() int32`

GetScrollAmount returns the ScrollAmount field if non-nil, zero value otherwise.

### GetScrollAmountOk

`func (o *OperativeComputerActionRequest) GetScrollAmountOk() (*int32, bool)`

GetScrollAmountOk returns a tuple with the ScrollAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollAmount

`func (o *OperativeComputerActionRequest) SetScrollAmount(v int32)`

SetScrollAmount sets ScrollAmount field to given value.

### HasScrollAmount

`func (o *OperativeComputerActionRequest) HasScrollAmount() bool`

HasScrollAmount returns a boolean if a field has been set.

### GetDuration

`func (o *OperativeComputerActionRequest) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *OperativeComputerActionRequest) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *OperativeComputerActionRequest) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *OperativeComputerActionRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetKey

`func (o *OperativeComputerActionRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OperativeComputerActionRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OperativeComputerActionRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OperativeComputerActionRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


