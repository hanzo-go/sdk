# KycStartOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formation** | Pointer to [**Formation**](Formation.md) | Formation is the org&#39;s incorporation record, with each founder&#39;s session reference and status recorded on it. | [optional] 
**Provider** | Pointer to **string** | Provider is the wired identity-verification provider&#39;s name. | [optional] 
**Sessions** | Pointer to [**[]KycSession**](KycSession.md) | Sessions is one entry per founder, in the order the founders are recorded. | [optional] 

## Methods

### NewKycStartOut

`func NewKycStartOut() *KycStartOut`

NewKycStartOut instantiates a new KycStartOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKycStartOutWithDefaults

`func NewKycStartOutWithDefaults() *KycStartOut`

NewKycStartOutWithDefaults instantiates a new KycStartOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormation

`func (o *KycStartOut) GetFormation() Formation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *KycStartOut) GetFormationOk() (*Formation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *KycStartOut) SetFormation(v Formation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *KycStartOut) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetProvider

`func (o *KycStartOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KycStartOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KycStartOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KycStartOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSessions

`func (o *KycStartOut) GetSessions() []KycSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *KycStartOut) GetSessionsOk() (*[]KycSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *KycStartOut) SetSessions(v []KycSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *KycStartOut) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


