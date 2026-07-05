# IntegrationsSlackEventEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | e.g. url_verification | event_callback | [optional] 
**Challenge** | Pointer to **string** | Present on url_verification | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewIntegrationsSlackEventEnvelope

`func NewIntegrationsSlackEventEnvelope() *IntegrationsSlackEventEnvelope`

NewIntegrationsSlackEventEnvelope instantiates a new IntegrationsSlackEventEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsSlackEventEnvelopeWithDefaults

`func NewIntegrationsSlackEventEnvelopeWithDefaults() *IntegrationsSlackEventEnvelope`

NewIntegrationsSlackEventEnvelopeWithDefaults instantiates a new IntegrationsSlackEventEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IntegrationsSlackEventEnvelope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationsSlackEventEnvelope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationsSlackEventEnvelope) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationsSlackEventEnvelope) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChallenge

`func (o *IntegrationsSlackEventEnvelope) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *IntegrationsSlackEventEnvelope) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *IntegrationsSlackEventEnvelope) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *IntegrationsSlackEventEnvelope) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetTeamId

`func (o *IntegrationsSlackEventEnvelope) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *IntegrationsSlackEventEnvelope) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *IntegrationsSlackEventEnvelope) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *IntegrationsSlackEventEnvelope) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetEvent

`func (o *IntegrationsSlackEventEnvelope) GetEvent() map[string]interface{}`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *IntegrationsSlackEventEnvelope) GetEventOk() (*map[string]interface{}, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *IntegrationsSlackEventEnvelope) SetEvent(v map[string]interface{})`

SetEvent sets Event field to given value.

### HasEvent

`func (o *IntegrationsSlackEventEnvelope) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


