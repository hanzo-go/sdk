# Health

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** | OK is true whenever this route answers at all: reaching the handler IS the proof that the routes are registered and dispatching. | [optional] 
**Subsystem** | Pointer to **string** | Subsystem names what answered, so a health response read out of context still says which surface it came from. | [optional] 

## Methods

### NewHealth

`func NewHealth() *Health`

NewHealth instantiates a new Health object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthWithDefaults

`func NewHealthWithDefaults() *Health`

NewHealthWithDefaults instantiates a new Health object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *Health) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *Health) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *Health) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *Health) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetSubsystem

`func (o *Health) GetSubsystem() string`

GetSubsystem returns the Subsystem field if non-nil, zero value otherwise.

### GetSubsystemOk

`func (o *Health) GetSubsystemOk() (*string, bool)`

GetSubsystemOk returns a tuple with the Subsystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystem

`func (o *Health) SetSubsystem(v string)`

SetSubsystem sets Subsystem field to given value.

### HasSubsystem

`func (o *Health) HasSubsystem() bool`

HasSubsystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


