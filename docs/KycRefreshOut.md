# KycRefreshOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formation** | Pointer to [**Formation**](Formation.md) | Formation is the org&#39;s incorporation record with each founder&#39;s reconciled status. | [optional] 
**Provider** | Pointer to **string** | Provider is the identity-verification provider that was consulted. | [optional] 

## Methods

### NewKycRefreshOut

`func NewKycRefreshOut() *KycRefreshOut`

NewKycRefreshOut instantiates a new KycRefreshOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKycRefreshOutWithDefaults

`func NewKycRefreshOutWithDefaults() *KycRefreshOut`

NewKycRefreshOutWithDefaults instantiates a new KycRefreshOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormation

`func (o *KycRefreshOut) GetFormation() Formation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *KycRefreshOut) GetFormationOk() (*Formation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *KycRefreshOut) SetFormation(v Formation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *KycRefreshOut) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetProvider

`func (o *KycRefreshOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KycRefreshOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KycRefreshOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KycRefreshOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


