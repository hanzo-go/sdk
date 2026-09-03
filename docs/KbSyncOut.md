# KbSyncOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingested** | Pointer to **int64** | Ingested is how many documents landed in the org&#39;s knowledge store. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector that was pulled. | [optional] 

## Methods

### NewKbSyncOut

`func NewKbSyncOut() *KbSyncOut`

NewKbSyncOut instantiates a new KbSyncOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKbSyncOutWithDefaults

`func NewKbSyncOutWithDefaults() *KbSyncOut`

NewKbSyncOutWithDefaults instantiates a new KbSyncOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngested

`func (o *KbSyncOut) GetIngested() int64`

GetIngested returns the Ingested field if non-nil, zero value otherwise.

### GetIngestedOk

`func (o *KbSyncOut) GetIngestedOk() (*int64, bool)`

GetIngestedOk returns a tuple with the Ingested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngested

`func (o *KbSyncOut) SetIngested(v int64)`

SetIngested sets Ingested field to given value.

### HasIngested

`func (o *KbSyncOut) HasIngested() bool`

HasIngested returns a boolean if a field has been set.

### GetProvider

`func (o *KbSyncOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KbSyncOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KbSyncOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KbSyncOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


