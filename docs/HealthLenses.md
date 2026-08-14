# HealthLenses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**HealthLens**](HealthLens.md) | Events is the web/commerce lens (event.fact, signal&#x3D;&#39;act&#39;), honest-empty until the collector emits. | [optional] 
**Llm** | Pointer to [**HealthLens**](HealthLens.md) | LLM is the live per-org usage ledger lens (hanzo.cloud_usage). | [optional] 

## Methods

### NewHealthLenses

`func NewHealthLenses() *HealthLenses`

NewHealthLenses instantiates a new HealthLenses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthLensesWithDefaults

`func NewHealthLensesWithDefaults() *HealthLenses`

NewHealthLensesWithDefaults instantiates a new HealthLenses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *HealthLenses) GetEvents() HealthLens`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *HealthLenses) GetEventsOk() (*HealthLens, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *HealthLenses) SetEvents(v HealthLens)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *HealthLenses) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLlm

`func (o *HealthLenses) GetLlm() HealthLens`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *HealthLenses) GetLlmOk() (*HealthLens, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *HealthLenses) SetLlm(v HealthLens)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *HealthLenses) HasLlm() bool`

HasLlm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


