# UnlinkedView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unlinked** | Pointer to **bool** | Unlinked is always true. Unlinking is idempotent: an account this org does not hold answers the same, so a repeated call is not an error and is not an existence oracle either. | [optional] 

## Methods

### NewUnlinkedView

`func NewUnlinkedView() *UnlinkedView`

NewUnlinkedView instantiates a new UnlinkedView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkedViewWithDefaults

`func NewUnlinkedViewWithDefaults() *UnlinkedView`

NewUnlinkedViewWithDefaults instantiates a new UnlinkedView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnlinked

`func (o *UnlinkedView) GetUnlinked() bool`

GetUnlinked returns the Unlinked field if non-nil, zero value otherwise.

### GetUnlinkedOk

`func (o *UnlinkedView) GetUnlinkedOk() (*bool, bool)`

GetUnlinkedOk returns a tuple with the Unlinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlinked

`func (o *UnlinkedView) SetUnlinked(v bool)`

SetUnlinked sets Unlinked field to given value.

### HasUnlinked

`func (o *UnlinkedView) HasUnlinked() bool`

HasUnlinked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


