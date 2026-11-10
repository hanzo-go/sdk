# EmailPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **time.Time** | At is the start of the day or hour, UTC. | [optional] 
**Tally** | Pointer to [**EmailTally**](EmailTally.md) | Tally is its mail counted. | [optional] 

## Methods

### NewEmailPoint

`func NewEmailPoint() *EmailPoint`

NewEmailPoint instantiates a new EmailPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailPointWithDefaults

`func NewEmailPointWithDefaults() *EmailPoint`

NewEmailPointWithDefaults instantiates a new EmailPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *EmailPoint) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *EmailPoint) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *EmailPoint) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *EmailPoint) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetTally

`func (o *EmailPoint) GetTally() EmailTally`

GetTally returns the Tally field if non-nil, zero value otherwise.

### GetTallyOk

`func (o *EmailPoint) GetTallyOk() (*EmailTally, bool)`

GetTallyOk returns a tuple with the Tally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTally

`func (o *EmailPoint) SetTally(v EmailTally)`

SetTally sets Tally field to given value.

### HasTally

`func (o *EmailPoint) HasTally() bool`

HasTally returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


