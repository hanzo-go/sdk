# RiskDatasetRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is the row&#39;s instant. | [optional] 
**Id** | Pointer to **string** | ID names the row forever. It is DERIVED from the row&#39;s own subject and instant, not allocated, so two materialisations of the same fact agree on it without coordinating. | [optional] 
**Kind** | Pointer to **string** | Kind is the subject kind: person, session or account. | [optional] 
**Point** | Pointer to **[]float32** | Point is the coordinates, in the order the version&#39;s spec names its dims. | [optional] 
**Split** | Pointer to **string** | Split is train, val or test. | [optional] 
**Subject** | Pointer to **string** | Subject is the identity within that kind — whose row this is. Every row of one subject is in ONE split, decided by that subject&#39;s earliest instant, so a subject is never on both sides of a cut. | [optional] 

## Methods

### NewRiskDatasetRow

`func NewRiskDatasetRow() *RiskDatasetRow`

NewRiskDatasetRow instantiates a new RiskDatasetRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskDatasetRowWithDefaults

`func NewRiskDatasetRowWithDefaults() *RiskDatasetRow`

NewRiskDatasetRowWithDefaults instantiates a new RiskDatasetRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *RiskDatasetRow) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *RiskDatasetRow) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *RiskDatasetRow) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *RiskDatasetRow) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetId

`func (o *RiskDatasetRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskDatasetRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskDatasetRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskDatasetRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *RiskDatasetRow) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RiskDatasetRow) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RiskDatasetRow) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RiskDatasetRow) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPoint

`func (o *RiskDatasetRow) GetPoint() []float32`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *RiskDatasetRow) GetPointOk() (*[]float32, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *RiskDatasetRow) SetPoint(v []float32)`

SetPoint sets Point field to given value.

### HasPoint

`func (o *RiskDatasetRow) HasPoint() bool`

HasPoint returns a boolean if a field has been set.

### GetSplit

`func (o *RiskDatasetRow) GetSplit() string`

GetSplit returns the Split field if non-nil, zero value otherwise.

### GetSplitOk

`func (o *RiskDatasetRow) GetSplitOk() (*string, bool)`

GetSplitOk returns a tuple with the Split field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplit

`func (o *RiskDatasetRow) SetSplit(v string)`

SetSplit sets Split field to given value.

### HasSplit

`func (o *RiskDatasetRow) HasSplit() bool`

HasSplit returns a boolean if a field has been set.

### GetSubject

`func (o *RiskDatasetRow) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RiskDatasetRow) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RiskDatasetRow) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RiskDatasetRow) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


