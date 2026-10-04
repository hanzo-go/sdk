# WorldWorldIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | Product is the product&#39;s name as customers know it. | [optional] 
**Summary** | Pointer to **string** | Summary is one sentence naming what this surface serves. | [optional] 
**Wires** | Pointer to [**[]WorldWorldWire**](WorldWorldWire.md) | Wires is every protocol entry point onto World, REST first. It is deliberately NOT a list of REST operations: GET /v1/openapi.json is the one enumeration of those, and a second copy here would be a second thing to keep true. | [optional] 

## Methods

### NewWorldWorldIndex

`func NewWorldWorldIndex() *WorldWorldIndex`

NewWorldWorldIndex instantiates a new WorldWorldIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldWorldIndexWithDefaults

`func NewWorldWorldIndexWithDefaults() *WorldWorldIndex`

NewWorldWorldIndexWithDefaults instantiates a new WorldWorldIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *WorldWorldIndex) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *WorldWorldIndex) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *WorldWorldIndex) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *WorldWorldIndex) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSummary

`func (o *WorldWorldIndex) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *WorldWorldIndex) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *WorldWorldIndex) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *WorldWorldIndex) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetWires

`func (o *WorldWorldIndex) GetWires() []WorldWorldWire`

GetWires returns the Wires field if non-nil, zero value otherwise.

### GetWiresOk

`func (o *WorldWorldIndex) GetWiresOk() (*[]WorldWorldWire, bool)`

GetWiresOk returns a tuple with the Wires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWires

`func (o *WorldWorldIndex) SetWires(v []WorldWorldWire)`

SetWires sets Wires field to given value.

### HasWires

`func (o *WorldWorldIndex) HasWires() bool`

HasWires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


