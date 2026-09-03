# CoverRow

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

### NewCoverRow

`func NewCoverRow() *CoverRow`

NewCoverRow instantiates a new CoverRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoverRowWithDefaults

`func NewCoverRowWithDefaults() *CoverRow`

NewCoverRowWithDefaults instantiates a new CoverRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomated

`func (o *CoverRow) GetAutomated() int64`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *CoverRow) GetAutomatedOk() (*int64, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *CoverRow) SetAutomated(v int64)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *CoverRow) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetEdition

`func (o *CoverRow) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *CoverRow) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *CoverRow) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *CoverRow) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetFramework

`func (o *CoverRow) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *CoverRow) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *CoverRow) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *CoverRow) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetName

`func (o *CoverRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoverRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoverRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoverRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNone

`func (o *CoverRow) GetNone() int64`

GetNone returns the None field if non-nil, zero value otherwise.

### GetNoneOk

`func (o *CoverRow) GetNoneOk() (*int64, bool)`

GetNoneOk returns a tuple with the None field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNone

`func (o *CoverRow) SetNone(v int64)`

SetNone sets None field to given value.

### HasNone

`func (o *CoverRow) HasNone() bool`

HasNone returns a boolean if a field has been set.

### GetNote

`func (o *CoverRow) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CoverRow) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CoverRow) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CoverRow) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPartial

`func (o *CoverRow) GetPartial() int64`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *CoverRow) GetPartialOk() (*int64, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *CoverRow) SetPartial(v int64)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *CoverRow) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetPublisher

`func (o *CoverRow) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *CoverRow) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *CoverRow) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *CoverRow) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetStatement

`func (o *CoverRow) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *CoverRow) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *CoverRow) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *CoverRow) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetTotal

`func (o *CoverRow) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CoverRow) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CoverRow) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CoverRow) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnit

`func (o *CoverRow) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CoverRow) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CoverRow) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CoverRow) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnits

`func (o *CoverRow) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CoverRow) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CoverRow) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *CoverRow) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


