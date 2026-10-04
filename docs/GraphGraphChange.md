# GraphGraphChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Now** | Pointer to [**GraphWireFact**](GraphWireFact.md) | Now is the assertion in force at To: the claim that replaced Was, or, for a retraction, the retraction itself — which is who withdrew it and on what evidence. | [optional] 
**Was** | Pointer to [**GraphWireFact**](GraphWireFact.md) | Was is the assertion in force at From. | [optional] 

## Methods

### NewGraphGraphChange

`func NewGraphGraphChange() *GraphGraphChange`

NewGraphGraphChange instantiates a new GraphGraphChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphChangeWithDefaults

`func NewGraphGraphChangeWithDefaults() *GraphGraphChange`

NewGraphGraphChangeWithDefaults instantiates a new GraphGraphChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNow

`func (o *GraphGraphChange) GetNow() GraphWireFact`

GetNow returns the Now field if non-nil, zero value otherwise.

### GetNowOk

`func (o *GraphGraphChange) GetNowOk() (*GraphWireFact, bool)`

GetNowOk returns a tuple with the Now field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNow

`func (o *GraphGraphChange) SetNow(v GraphWireFact)`

SetNow sets Now field to given value.

### HasNow

`func (o *GraphGraphChange) HasNow() bool`

HasNow returns a boolean if a field has been set.

### GetWas

`func (o *GraphGraphChange) GetWas() GraphWireFact`

GetWas returns the Was field if non-nil, zero value otherwise.

### GetWasOk

`func (o *GraphGraphChange) GetWasOk() (*GraphWireFact, bool)`

GetWasOk returns a tuple with the Was field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWas

`func (o *GraphGraphChange) SetWas(v GraphWireFact)`

SetWas sets Was field to given value.

### HasWas

`func (o *GraphGraphChange) HasWas() bool`

HasWas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


