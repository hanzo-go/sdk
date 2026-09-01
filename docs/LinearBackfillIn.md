# LinearBackfillIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | State is the set of issues to walk: \&quot;open\&quot; (the default), \&quot;closed\&quot; or \&quot;all\&quot;. | [optional] 

## Methods

### NewLinearBackfillIn

`func NewLinearBackfillIn() *LinearBackfillIn`

NewLinearBackfillIn instantiates a new LinearBackfillIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinearBackfillInWithDefaults

`func NewLinearBackfillInWithDefaults() *LinearBackfillIn`

NewLinearBackfillInWithDefaults instantiates a new LinearBackfillIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *LinearBackfillIn) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LinearBackfillIn) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LinearBackfillIn) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LinearBackfillIn) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


