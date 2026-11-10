# TrainAdaptation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alpha** | Pointer to **float64** | Alpha is an adapter&#39;s scale. | [optional] 
**Basis** | Pointer to **string** | Basis is a basis artifact, by sha256 or &#x60;basis://&lt;base&gt;/&lt;name&gt;&#x60;. | [optional] 
**Chose** | Pointer to [**TrainChose**](TrainChose.md) | Chose is what auto resolved to and why; absent unless auto was asked. | [optional] 
**Mode** | Pointer to **string** | Mode is full, readout (the head alone, over the frozen base), lora, qlora, basis or auto. | [optional] 
**Rank** | Pointer to **int64** | Rank is an adapter&#39;s rank, or how many of a basis&#39;s directions are used. | [optional] 
**ResidualRank** | Pointer to **int64** | ResidualRank is the rank of an orthonormal residual learned beside a basis&#39;s coefficients; 0 trains the coefficients alone. | [optional] 
**Sources** | Pointer to **[]string** | Sources are lora artifacts (sha256) a basis is built from, in place of Basis. | [optional] 
**Targets** | Pointer to **[]string** | Targets are the modules an adapter or a basis attaches to. | [optional] 

## Methods

### NewTrainAdaptation

`func NewTrainAdaptation() *TrainAdaptation`

NewTrainAdaptation instantiates a new TrainAdaptation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainAdaptationWithDefaults

`func NewTrainAdaptationWithDefaults() *TrainAdaptation`

NewTrainAdaptationWithDefaults instantiates a new TrainAdaptation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpha

`func (o *TrainAdaptation) GetAlpha() float64`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *TrainAdaptation) GetAlphaOk() (*float64, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *TrainAdaptation) SetAlpha(v float64)`

SetAlpha sets Alpha field to given value.

### HasAlpha

`func (o *TrainAdaptation) HasAlpha() bool`

HasAlpha returns a boolean if a field has been set.

### GetBasis

`func (o *TrainAdaptation) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *TrainAdaptation) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *TrainAdaptation) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *TrainAdaptation) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetChose

`func (o *TrainAdaptation) GetChose() TrainChose`

GetChose returns the Chose field if non-nil, zero value otherwise.

### GetChoseOk

`func (o *TrainAdaptation) GetChoseOk() (*TrainChose, bool)`

GetChoseOk returns a tuple with the Chose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChose

`func (o *TrainAdaptation) SetChose(v TrainChose)`

SetChose sets Chose field to given value.

### HasChose

`func (o *TrainAdaptation) HasChose() bool`

HasChose returns a boolean if a field has been set.

### GetMode

`func (o *TrainAdaptation) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TrainAdaptation) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TrainAdaptation) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *TrainAdaptation) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRank

`func (o *TrainAdaptation) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *TrainAdaptation) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *TrainAdaptation) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *TrainAdaptation) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetResidualRank

`func (o *TrainAdaptation) GetResidualRank() int64`

GetResidualRank returns the ResidualRank field if non-nil, zero value otherwise.

### GetResidualRankOk

`func (o *TrainAdaptation) GetResidualRankOk() (*int64, bool)`

GetResidualRankOk returns a tuple with the ResidualRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidualRank

`func (o *TrainAdaptation) SetResidualRank(v int64)`

SetResidualRank sets ResidualRank field to given value.

### HasResidualRank

`func (o *TrainAdaptation) HasResidualRank() bool`

HasResidualRank returns a boolean if a field has been set.

### GetSources

`func (o *TrainAdaptation) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *TrainAdaptation) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *TrainAdaptation) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *TrainAdaptation) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTargets

`func (o *TrainAdaptation) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *TrainAdaptation) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *TrainAdaptation) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *TrainAdaptation) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


