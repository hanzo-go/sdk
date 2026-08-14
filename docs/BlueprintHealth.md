# BlueprintHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprints** | Pointer to **int32** | Blueprints is how many blueprints this build has embedded and priced. | [optional] 
**RateCard** | Pointer to [**RateCard**](RateCard.md) | RateCard is the rate card actually in force after the operator env overlay (CLOUD_BLUEPRINT_UCPU_HR / CLOUD_BLUEPRINT_UGB_HR), not the shipped default. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem answering — always \&quot;blueprint\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot;; the route answers 200 whenever the subsystem is mounted. | [optional] 

## Methods

### NewBlueprintHealth

`func NewBlueprintHealth() *BlueprintHealth`

NewBlueprintHealth instantiates a new BlueprintHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintHealthWithDefaults

`func NewBlueprintHealthWithDefaults() *BlueprintHealth`

NewBlueprintHealthWithDefaults instantiates a new BlueprintHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprints

`func (o *BlueprintHealth) GetBlueprints() int32`

GetBlueprints returns the Blueprints field if non-nil, zero value otherwise.

### GetBlueprintsOk

`func (o *BlueprintHealth) GetBlueprintsOk() (*int32, bool)`

GetBlueprintsOk returns a tuple with the Blueprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprints

`func (o *BlueprintHealth) SetBlueprints(v int32)`

SetBlueprints sets Blueprints field to given value.

### HasBlueprints

`func (o *BlueprintHealth) HasBlueprints() bool`

HasBlueprints returns a boolean if a field has been set.

### GetRateCard

`func (o *BlueprintHealth) GetRateCard() RateCard`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *BlueprintHealth) GetRateCardOk() (*RateCard, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *BlueprintHealth) SetRateCard(v RateCard)`

SetRateCard sets RateCard field to given value.

### HasRateCard

`func (o *BlueprintHealth) HasRateCard() bool`

HasRateCard returns a boolean if a field has been set.

### GetService

`func (o *BlueprintHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BlueprintHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BlueprintHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BlueprintHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *BlueprintHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlueprintHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlueprintHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlueprintHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


