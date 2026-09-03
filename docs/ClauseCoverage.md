# ClauseCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automated** | Pointer to **int64** | Automated is how many clauses have an automated control behind them that something can fail on behalf of. | [optional] 
**Clauses** | Pointer to [**[]ClauseRow**](ClauseRow.md) | Clauses is every clause the standard publishes, with what stands behind it. | [optional] 
**Edition** | Pointer to **string** | Edition is which edition this clause list is taken from. | [optional] 
**Framework** | Pointer to **string** | Framework is the framework id — \&quot;soc2\&quot;, \&quot;iso27001\&quot;, \&quot;nist80053\&quot;. | [optional] 
**Generated** | Pointer to **int64** | Generated is when this was computed, unix milliseconds. | [optional] 
**Name** | Pointer to **string** | Name is the published standard&#39;s name. | [optional] 
**None** | Pointer to **int64** | None is how many have nothing behind them. | [optional] 
**Note** | Pointer to **string** | Note is what this clause list is scoped to, when a count alone would misrepresent it. | [optional] 
**Partial** | Pointer to **int64** | Partial is how many are answered in part. | [optional] 
**Publisher** | Pointer to **string** | Publisher is who publishes it — AICPA, ISO/IEC, NIST. | [optional] 
**Statement** | Pointer to **string** | Statement is the counts as one sentence, carrying the unit. | [optional] 
**Total** | Pointer to **int64** | Total is the framework&#39;s WHOLE published clause list — the denominator. | [optional] 
**Unit** | Pointer to **string** | Unit is what ONE clause is — \&quot;criterion\&quot;, \&quot;control\&quot;, \&quot;family\&quot;. | [optional] 
**Units** | Pointer to **string** | Units is the plural of Unit, for rendering a sentence. | [optional] 
**Version** | Pointer to **string** | Version is the embedded inventory&#39;s version. | [optional] 

## Methods

### NewClauseCoverage

`func NewClauseCoverage() *ClauseCoverage`

NewClauseCoverage instantiates a new ClauseCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClauseCoverageWithDefaults

`func NewClauseCoverageWithDefaults() *ClauseCoverage`

NewClauseCoverageWithDefaults instantiates a new ClauseCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomated

`func (o *ClauseCoverage) GetAutomated() int64`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *ClauseCoverage) GetAutomatedOk() (*int64, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *ClauseCoverage) SetAutomated(v int64)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *ClauseCoverage) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetClauses

`func (o *ClauseCoverage) GetClauses() []ClauseRow`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *ClauseCoverage) GetClausesOk() (*[]ClauseRow, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *ClauseCoverage) SetClauses(v []ClauseRow)`

SetClauses sets Clauses field to given value.

### HasClauses

`func (o *ClauseCoverage) HasClauses() bool`

HasClauses returns a boolean if a field has been set.

### GetEdition

`func (o *ClauseCoverage) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *ClauseCoverage) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *ClauseCoverage) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *ClauseCoverage) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetFramework

`func (o *ClauseCoverage) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *ClauseCoverage) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *ClauseCoverage) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *ClauseCoverage) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetGenerated

`func (o *ClauseCoverage) GetGenerated() int64`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *ClauseCoverage) GetGeneratedOk() (*int64, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *ClauseCoverage) SetGenerated(v int64)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *ClauseCoverage) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetName

`func (o *ClauseCoverage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClauseCoverage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClauseCoverage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClauseCoverage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNone

`func (o *ClauseCoverage) GetNone() int64`

GetNone returns the None field if non-nil, zero value otherwise.

### GetNoneOk

`func (o *ClauseCoverage) GetNoneOk() (*int64, bool)`

GetNoneOk returns a tuple with the None field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNone

`func (o *ClauseCoverage) SetNone(v int64)`

SetNone sets None field to given value.

### HasNone

`func (o *ClauseCoverage) HasNone() bool`

HasNone returns a boolean if a field has been set.

### GetNote

`func (o *ClauseCoverage) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ClauseCoverage) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ClauseCoverage) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ClauseCoverage) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPartial

`func (o *ClauseCoverage) GetPartial() int64`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *ClauseCoverage) GetPartialOk() (*int64, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *ClauseCoverage) SetPartial(v int64)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *ClauseCoverage) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetPublisher

`func (o *ClauseCoverage) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ClauseCoverage) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ClauseCoverage) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ClauseCoverage) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetStatement

`func (o *ClauseCoverage) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ClauseCoverage) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ClauseCoverage) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *ClauseCoverage) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetTotal

`func (o *ClauseCoverage) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ClauseCoverage) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ClauseCoverage) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ClauseCoverage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnit

`func (o *ClauseCoverage) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ClauseCoverage) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ClauseCoverage) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ClauseCoverage) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnits

`func (o *ClauseCoverage) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ClauseCoverage) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ClauseCoverage) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ClauseCoverage) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetVersion

`func (o *ClauseCoverage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClauseCoverage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClauseCoverage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClauseCoverage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


