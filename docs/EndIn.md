# EndIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the sandbox whose lease ends, from an earlier lease. | [optional] 
**Purge** | Pointer to **bool** | Purge deletes the project&#39;s DISK as well. It is opt-in because the disk holds the only copy of the checkout: ending a lease is cheap and reversible, deleting someone&#39;s uncommitted work is neither. | [optional] 

## Methods

### NewEndIn

`func NewEndIn() *EndIn`

NewEndIn instantiates a new EndIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndInWithDefaults

`func NewEndInWithDefaults() *EndIn`

NewEndInWithDefaults instantiates a new EndIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EndIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPurge

`func (o *EndIn) GetPurge() bool`

GetPurge returns the Purge field if non-nil, zero value otherwise.

### GetPurgeOk

`func (o *EndIn) GetPurgeOk() (*bool, bool)`

GetPurgeOk returns a tuple with the Purge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurge

`func (o *EndIn) SetPurge(v bool)`

SetPurge sets Purge field to given value.

### HasPurge

`func (o *EndIn) HasPurge() bool`

HasPurge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


