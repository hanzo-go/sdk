# StopIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the sandbox to interrupt, from an earlier lease. Every command running in it stops; the lease itself survives, so the checkout and the half-written files are still there to read. Use EndIn to give the computer back. | [optional] 

## Methods

### NewStopIn

`func NewStopIn() *StopIn`

NewStopIn instantiates a new StopIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopInWithDefaults

`func NewStopInWithDefaults() *StopIn`

NewStopInWithDefaults instantiates a new StopIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StopIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StopIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StopIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StopIn) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


