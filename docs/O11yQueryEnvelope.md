# O11yQueryEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to **map[string]interface{}** | Spec is the deferred decoding of the query if any. | [optional] 
**Type** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewO11yQueryEnvelope

`func NewO11yQueryEnvelope() *O11yQueryEnvelope`

NewO11yQueryEnvelope instantiates a new O11yQueryEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yQueryEnvelopeWithDefaults

`func NewO11yQueryEnvelopeWithDefaults() *O11yQueryEnvelope`

NewO11yQueryEnvelopeWithDefaults instantiates a new O11yQueryEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *O11yQueryEnvelope) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *O11yQueryEnvelope) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *O11yQueryEnvelope) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *O11yQueryEnvelope) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetType

`func (o *O11yQueryEnvelope) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yQueryEnvelope) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yQueryEnvelope) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *O11yQueryEnvelope) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *O11yQueryEnvelope) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *O11yQueryEnvelope) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


