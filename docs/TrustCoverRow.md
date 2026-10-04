# TrustCoverRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automated** | Pointer to **int64** | Automated is how many clauses have an automated control behind them that something can fail on behalf of. | [optional] 
**Edition** | Pointer to **string** | Edition is which edition the clause list is taken from. | [optional] 
**Framework** | Pointer to **string** | Framework is the framework id — \&quot;soc2\&quot;, \&quot;iso27001\&quot;, \&quot;nist80053\&quot;. | [optional] 
**Name** | Pointer to **string** | Name is the published standard&#39;s name. | [optional] 
**None** | Pointer to **int64** | None is how many have nothing behind them. It stays visible rather than dropping out of the fraction. | [optional] 
**Note** | Pointer to **string** | Note is what the clause list itself is scoped to, when the framework&#39;s catalog says something a count alone would misrepresent. | [optional] 
**Partial** | Pointer to **int64** | Partial is how many are answered in part. | [optional] 
**Publisher** | Pointer to **string** | Publisher is who publishes it — AICPA, ISO/IEC, NIST. | [optional] 
**Statement** | Pointer to **string** | Statement is the counts as one sentence, carrying the unit. | [optional] 
**Total** | Pointer to **int64** | Total is the framework&#39;s WHOLE published clause list — the denominator. Counting only the clauses some control happened to name would report 100% every time. | [optional] 
**Unit** | Pointer to **string** | Unit is what ONE clause is — \&quot;criterion\&quot;, \&quot;control\&quot;, \&quot;family\&quot;. A count without its unit is not a fact, so it travels with every number here. | [optional] 
**Units** | Pointer to **string** | Units is the plural of Unit, for rendering a sentence. | [optional] 

## Methods

### NewTrustCoverRow

`func NewTrustCoverRow() *TrustCoverRow`

NewTrustCoverRow instantiates a new TrustCoverRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustCoverRowWithDefaults

`func NewTrustCoverRowWithDefaults() *TrustCoverRow`

NewTrustCoverRowWithDefaults instantiates a new TrustCoverRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomated

`func (o *TrustCoverRow) GetAutomated() int64`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *TrustCoverRow) GetAutomatedOk() (*int64, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *TrustCoverRow) SetAutomated(v int64)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *TrustCoverRow) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetEdition

`func (o *TrustCoverRow) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *TrustCoverRow) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *TrustCoverRow) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *TrustCoverRow) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetFramework

`func (o *TrustCoverRow) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *TrustCoverRow) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *TrustCoverRow) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *TrustCoverRow) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetName

`func (o *TrustCoverRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustCoverRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustCoverRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrustCoverRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNone

`func (o *TrustCoverRow) GetNone() int64`

GetNone returns the None field if non-nil, zero value otherwise.

### GetNoneOk

`func (o *TrustCoverRow) GetNoneOk() (*int64, bool)`

GetNoneOk returns a tuple with the None field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNone

`func (o *TrustCoverRow) SetNone(v int64)`

SetNone sets None field to given value.

### HasNone

`func (o *TrustCoverRow) HasNone() bool`

HasNone returns a boolean if a field has been set.

### GetNote

`func (o *TrustCoverRow) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *TrustCoverRow) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *TrustCoverRow) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *TrustCoverRow) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPartial

`func (o *TrustCoverRow) GetPartial() int64`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *TrustCoverRow) GetPartialOk() (*int64, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *TrustCoverRow) SetPartial(v int64)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *TrustCoverRow) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetPublisher

`func (o *TrustCoverRow) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *TrustCoverRow) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *TrustCoverRow) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *TrustCoverRow) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetStatement

`func (o *TrustCoverRow) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *TrustCoverRow) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *TrustCoverRow) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *TrustCoverRow) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetTotal

`func (o *TrustCoverRow) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TrustCoverRow) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TrustCoverRow) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TrustCoverRow) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnit

`func (o *TrustCoverRow) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TrustCoverRow) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TrustCoverRow) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TrustCoverRow) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnits

`func (o *TrustCoverRow) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *TrustCoverRow) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *TrustCoverRow) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *TrustCoverRow) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


