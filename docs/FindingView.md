# FindingView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when the finding was recorded, in Unix milliseconds. | [optional] 
**Fingerprint** | Pointer to **string** | Fingerprint is the SHA-256 of the raw secret. It is what makes the same secret recognisable across scans and after rotation without the secret ever being written down. | [optional] 
**Id** | Pointer to **string** | ID addresses this finding. | [optional] 
**Line** | Pointer to **int32** | Line is where in that file. | [optional] 
**Path** | Pointer to **string** | Path is the file the secret was found in. | [optional] 
**Preview** | Pointer to **string** | Preview is the secret MASKED — first and last characters kept, the middle starred — so a reviewer can recognise it without it being disclosed. | [optional] 
**RuleId** | Pointer to **string** | RuleID is the detection rule that fired. | [optional] 
**RuleName** | Pointer to **string** | RuleName is that rule&#39;s human name. | [optional] 
**ScanId** | Pointer to **string** | ScanID is the scan that produced it. | [optional] 
**Severity** | Pointer to **string** | Severity ranks the finding: critical, high, medium or low. | [optional] 

## Methods

### NewFindingView

`func NewFindingView() *FindingView`

NewFindingView instantiates a new FindingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingViewWithDefaults

`func NewFindingViewWithDefaults() *FindingView`

NewFindingViewWithDefaults instantiates a new FindingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindingView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindingView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindingView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindingView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFingerprint

`func (o *FindingView) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *FindingView) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *FindingView) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *FindingView) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetId

`func (o *FindingView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindingView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindingView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindingView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLine

`func (o *FindingView) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *FindingView) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *FindingView) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *FindingView) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetPath

`func (o *FindingView) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FindingView) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FindingView) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FindingView) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPreview

`func (o *FindingView) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *FindingView) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *FindingView) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *FindingView) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetRuleId

`func (o *FindingView) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *FindingView) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *FindingView) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *FindingView) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *FindingView) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *FindingView) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *FindingView) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *FindingView) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetScanId

`func (o *FindingView) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *FindingView) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *FindingView) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *FindingView) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetSeverity

`func (o *FindingView) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FindingView) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FindingView) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *FindingView) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


