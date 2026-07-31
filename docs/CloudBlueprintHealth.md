# CloudBlueprintHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprints** | Pointer to **int32** | Blueprints is how many blueprints this build has embedded and priced. | [optional] 
**RateCard** | Pointer to [**CloudRateCard**](CloudRateCard.md) | RateCard is the rate card actually in force after the operator env overlay (CLOUD_BLUEPRINT_UCPU_HR / CLOUD_BLUEPRINT_UGB_HR), not the shipped default. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem answering — always \&quot;blueprint\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot;; the route answers 200 whenever the subsystem is mounted. | [optional] 

## Methods

### NewCloudBlueprintHealth

`func NewCloudBlueprintHealth() *CloudBlueprintHealth`

NewCloudBlueprintHealth instantiates a new CloudBlueprintHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBlueprintHealthWithDefaults

`func NewCloudBlueprintHealthWithDefaults() *CloudBlueprintHealth`

NewCloudBlueprintHealthWithDefaults instantiates a new CloudBlueprintHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprints

`func (o *CloudBlueprintHealth) GetBlueprints() int32`

GetBlueprints returns the Blueprints field if non-nil, zero value otherwise.

### GetBlueprintsOk

`func (o *CloudBlueprintHealth) GetBlueprintsOk() (*int32, bool)`

GetBlueprintsOk returns a tuple with the Blueprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprints

`func (o *CloudBlueprintHealth) SetBlueprints(v int32)`

SetBlueprints sets Blueprints field to given value.

### HasBlueprints

`func (o *CloudBlueprintHealth) HasBlueprints() bool`

HasBlueprints returns a boolean if a field has been set.

### GetRateCard

`func (o *CloudBlueprintHealth) GetRateCard() CloudRateCard`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *CloudBlueprintHealth) GetRateCardOk() (*CloudRateCard, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *CloudBlueprintHealth) SetRateCard(v CloudRateCard)`

SetRateCard sets RateCard field to given value.

### HasRateCard

`func (o *CloudBlueprintHealth) HasRateCard() bool`

HasRateCard returns a boolean if a field has been set.

### GetService

`func (o *CloudBlueprintHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudBlueprintHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudBlueprintHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudBlueprintHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *CloudBlueprintHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudBlueprintHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudBlueprintHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudBlueprintHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


