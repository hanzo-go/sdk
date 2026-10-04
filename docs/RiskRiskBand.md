# RiskRiskBand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **time.Time** | Day is the day the band covers. | [optional] 
**Dim** | Pointer to **string** | Dim is the dimension, named as this API publishes it. | [optional] 
**Kind** | Pointer to **string** | Kind is the subject kind it was computed over. | [optional] 
**N** | Pointer to **int32** | N is how many subject-days went into it. | [optional] 
**Orgs** | Pointer to **int32** | Orgs is how many organisations contributed, each weighted exactly one vote whatever its size. It is published so a reader can judge the band rather than trust it. | [optional] 
**Q10** | Pointer to **float64** | Q10 is the quiet end of the network&#39;s day: a tenth of contributing organisations sit at or below it. | [optional] 
**Q50** | Pointer to **float64** | Q50 is the network&#39;s median day. | [optional] 
**Q90** | Pointer to **float64** | Q90 is the busy end: a tenth of contributing organisations sit at or above it. It is the highest level published. | [optional] 

## Methods

### NewRiskRiskBand

`func NewRiskRiskBand() *RiskRiskBand`

NewRiskRiskBand instantiates a new RiskRiskBand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskBandWithDefaults

`func NewRiskRiskBandWithDefaults() *RiskRiskBand`

NewRiskRiskBandWithDefaults instantiates a new RiskRiskBand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *RiskRiskBand) GetDay() time.Time`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *RiskRiskBand) GetDayOk() (*time.Time, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *RiskRiskBand) SetDay(v time.Time)`

SetDay sets Day field to given value.

### HasDay

`func (o *RiskRiskBand) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetDim

`func (o *RiskRiskBand) GetDim() string`

GetDim returns the Dim field if non-nil, zero value otherwise.

### GetDimOk

`func (o *RiskRiskBand) GetDimOk() (*string, bool)`

GetDimOk returns a tuple with the Dim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDim

`func (o *RiskRiskBand) SetDim(v string)`

SetDim sets Dim field to given value.

### HasDim

`func (o *RiskRiskBand) HasDim() bool`

HasDim returns a boolean if a field has been set.

### GetKind

`func (o *RiskRiskBand) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RiskRiskBand) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RiskRiskBand) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RiskRiskBand) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetN

`func (o *RiskRiskBand) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *RiskRiskBand) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *RiskRiskBand) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *RiskRiskBand) HasN() bool`

HasN returns a boolean if a field has been set.

### GetOrgs

`func (o *RiskRiskBand) GetOrgs() int32`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *RiskRiskBand) GetOrgsOk() (*int32, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *RiskRiskBand) SetOrgs(v int32)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *RiskRiskBand) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetQ10

`func (o *RiskRiskBand) GetQ10() float64`

GetQ10 returns the Q10 field if non-nil, zero value otherwise.

### GetQ10Ok

`func (o *RiskRiskBand) GetQ10Ok() (*float64, bool)`

GetQ10Ok returns a tuple with the Q10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ10

`func (o *RiskRiskBand) SetQ10(v float64)`

SetQ10 sets Q10 field to given value.

### HasQ10

`func (o *RiskRiskBand) HasQ10() bool`

HasQ10 returns a boolean if a field has been set.

### GetQ50

`func (o *RiskRiskBand) GetQ50() float64`

GetQ50 returns the Q50 field if non-nil, zero value otherwise.

### GetQ50Ok

`func (o *RiskRiskBand) GetQ50Ok() (*float64, bool)`

GetQ50Ok returns a tuple with the Q50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ50

`func (o *RiskRiskBand) SetQ50(v float64)`

SetQ50 sets Q50 field to given value.

### HasQ50

`func (o *RiskRiskBand) HasQ50() bool`

HasQ50 returns a boolean if a field has been set.

### GetQ90

`func (o *RiskRiskBand) GetQ90() float64`

GetQ90 returns the Q90 field if non-nil, zero value otherwise.

### GetQ90Ok

`func (o *RiskRiskBand) GetQ90Ok() (*float64, bool)`

GetQ90Ok returns a tuple with the Q90 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ90

`func (o *RiskRiskBand) SetQ90(v float64)`

SetQ90 sets Q90 field to given value.

### HasQ90

`func (o *RiskRiskBand) HasQ90() bool`

HasQ90 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


