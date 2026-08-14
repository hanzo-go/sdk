# EsignOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsignRef** | Pointer to **string** | EsignRef is the provider&#39;s reference for the signature request. | [optional] 
**Formation** | Pointer to [**Formation**](Formation.md) | Formation is the org&#39;s incorporation record with the reference recorded on it. | [optional] 
**Provider** | Pointer to **string** | Provider is the wired e-signature provider&#39;s name. | [optional] 

## Methods

### NewEsignOut

`func NewEsignOut() *EsignOut`

NewEsignOut instantiates a new EsignOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignOutWithDefaults

`func NewEsignOutWithDefaults() *EsignOut`

NewEsignOutWithDefaults instantiates a new EsignOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsignRef

`func (o *EsignOut) GetEsignRef() string`

GetEsignRef returns the EsignRef field if non-nil, zero value otherwise.

### GetEsignRefOk

`func (o *EsignOut) GetEsignRefOk() (*string, bool)`

GetEsignRefOk returns a tuple with the EsignRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsignRef

`func (o *EsignOut) SetEsignRef(v string)`

SetEsignRef sets EsignRef field to given value.

### HasEsignRef

`func (o *EsignOut) HasEsignRef() bool`

HasEsignRef returns a boolean if a field has been set.

### GetFormation

`func (o *EsignOut) GetFormation() Formation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *EsignOut) GetFormationOk() (*Formation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *EsignOut) SetFormation(v Formation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *EsignOut) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetProvider

`func (o *EsignOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EsignOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EsignOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EsignOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


