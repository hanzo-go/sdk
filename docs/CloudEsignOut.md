# CloudEsignOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsignRef** | Pointer to **string** | EsignRef is the provider&#39;s reference for the signature request. | [optional] 
**Formation** | Pointer to [**CloudFormation**](CloudFormation.md) | Formation is the org&#39;s incorporation record with the reference recorded on it. | [optional] 
**Provider** | Pointer to **string** | Provider is the wired e-signature provider&#39;s name. | [optional] 

## Methods

### NewCloudEsignOut

`func NewCloudEsignOut() *CloudEsignOut`

NewCloudEsignOut instantiates a new CloudEsignOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEsignOutWithDefaults

`func NewCloudEsignOutWithDefaults() *CloudEsignOut`

NewCloudEsignOutWithDefaults instantiates a new CloudEsignOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsignRef

`func (o *CloudEsignOut) GetEsignRef() string`

GetEsignRef returns the EsignRef field if non-nil, zero value otherwise.

### GetEsignRefOk

`func (o *CloudEsignOut) GetEsignRefOk() (*string, bool)`

GetEsignRefOk returns a tuple with the EsignRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsignRef

`func (o *CloudEsignOut) SetEsignRef(v string)`

SetEsignRef sets EsignRef field to given value.

### HasEsignRef

`func (o *CloudEsignOut) HasEsignRef() bool`

HasEsignRef returns a boolean if a field has been set.

### GetFormation

`func (o *CloudEsignOut) GetFormation() CloudFormation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *CloudEsignOut) GetFormationOk() (*CloudFormation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *CloudEsignOut) SetFormation(v CloudFormation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *CloudEsignOut) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetProvider

`func (o *CloudEsignOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudEsignOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudEsignOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudEsignOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


