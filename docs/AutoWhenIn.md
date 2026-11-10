# AutoWhenIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Text is when it runs, as a person says it: \&quot;every monday at 9am\&quot;. | [optional] 
**Tz** | Pointer to **string** | TZ is the IANA zone the words are read in. Empty is UTC. | [optional] 

## Methods

### NewAutoWhenIn

`func NewAutoWhenIn() *AutoWhenIn`

NewAutoWhenIn instantiates a new AutoWhenIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoWhenInWithDefaults

`func NewAutoWhenInWithDefaults() *AutoWhenIn`

NewAutoWhenInWithDefaults instantiates a new AutoWhenIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *AutoWhenIn) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AutoWhenIn) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AutoWhenIn) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AutoWhenIn) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTz

`func (o *AutoWhenIn) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *AutoWhenIn) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *AutoWhenIn) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *AutoWhenIn) HasTz() bool`

HasTz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


