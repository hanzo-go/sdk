# GraphGraphEraseIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | **string** | Entity is the key to erase. Every assertion ABOUT it and every edge pointing AT it is removed, under every spelling it may have been filed under — its NFC key, that key decomposed, and the bytes sent — so a subject filed before keys were folded goes with the rest; a property whose scalar happens to spell it is not a reference and stays. Required, 512 bytes at most. | 
**Reason** | **string** | Reason is why, as the audit trail will keep it — a request number, a legal basis. It outlives the erasure, so it may not contain the entity: refused when it does. Required, 512 bytes at most. | 

## Methods

### NewGraphGraphEraseIn

`func NewGraphGraphEraseIn(entity string, reason string, ) *GraphGraphEraseIn`

NewGraphGraphEraseIn instantiates a new GraphGraphEraseIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphEraseInWithDefaults

`func NewGraphGraphEraseInWithDefaults() *GraphGraphEraseIn`

NewGraphGraphEraseInWithDefaults instantiates a new GraphGraphEraseIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *GraphGraphEraseIn) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphGraphEraseIn) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphGraphEraseIn) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetReason

`func (o *GraphGraphEraseIn) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GraphGraphEraseIn) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GraphGraphEraseIn) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


