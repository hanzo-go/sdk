# BlueprintBlueprintHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprints** | Pointer to **int64** | Blueprints is how many blueprints this build has embedded and priced. | [optional] 
**RateCard** | Pointer to [**BlueprintRateCard**](BlueprintRateCard.md) | RateCard is the rate card actually in force after the operator env overlay (CLOUD_BLUEPRINT_UCPU_HR / CLOUD_BLUEPRINT_UGB_HR), not the shipped default. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem answering — always \&quot;blueprint\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot;; the route answers 200 whenever the subsystem is mounted. | [optional] 

## Methods

### NewBlueprintBlueprintHealth

`func NewBlueprintBlueprintHealth() *BlueprintBlueprintHealth`

NewBlueprintBlueprintHealth instantiates a new BlueprintBlueprintHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintBlueprintHealthWithDefaults

`func NewBlueprintBlueprintHealthWithDefaults() *BlueprintBlueprintHealth`

NewBlueprintBlueprintHealthWithDefaults instantiates a new BlueprintBlueprintHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprints

`func (o *BlueprintBlueprintHealth) GetBlueprints() int64`

GetBlueprints returns the Blueprints field if non-nil, zero value otherwise.

### GetBlueprintsOk

`func (o *BlueprintBlueprintHealth) GetBlueprintsOk() (*int64, bool)`

GetBlueprintsOk returns a tuple with the Blueprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprints

`func (o *BlueprintBlueprintHealth) SetBlueprints(v int64)`

SetBlueprints sets Blueprints field to given value.

### HasBlueprints

`func (o *BlueprintBlueprintHealth) HasBlueprints() bool`

HasBlueprints returns a boolean if a field has been set.

### GetRateCard

`func (o *BlueprintBlueprintHealth) GetRateCard() BlueprintRateCard`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *BlueprintBlueprintHealth) GetRateCardOk() (*BlueprintRateCard, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *BlueprintBlueprintHealth) SetRateCard(v BlueprintRateCard)`

SetRateCard sets RateCard field to given value.

### HasRateCard

`func (o *BlueprintBlueprintHealth) HasRateCard() bool`

HasRateCard returns a boolean if a field has been set.

### GetService

`func (o *BlueprintBlueprintHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BlueprintBlueprintHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BlueprintBlueprintHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BlueprintBlueprintHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *BlueprintBlueprintHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlueprintBlueprintHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlueprintBlueprintHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlueprintBlueprintHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


