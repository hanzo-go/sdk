# SecurityFindingView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the finding was recorded, in Unix milliseconds. | [optional] 
**Fingerprint** | Pointer to **string** | Fingerprint is the SHA-256 of the raw secret. It is what makes the same secret recognisable across scans and after rotation without the secret ever being written down. | [optional] 
**Id** | Pointer to **string** | ID addresses this finding. | [optional] 
**Line** | Pointer to **int64** | Line is where in that file. | [optional] 
**Path** | Pointer to **string** | Path is the file the secret was found in. | [optional] 
**Preview** | Pointer to **string** | Preview is the secret MASKED — first and last characters kept, the middle starred — so a reviewer can recognise it without it being disclosed. | [optional] 
**RuleId** | Pointer to **string** | RuleID is the detection rule that fired. | [optional] 
**RuleName** | Pointer to **string** | RuleName is that rule&#39;s human name. | [optional] 
**ScanId** | Pointer to **string** | ScanID is the scan that produced it. | [optional] 
**Severity** | Pointer to **string** | Severity ranks the finding: critical, high, medium or low. | [optional] 

## Methods

### NewSecurityFindingView

`func NewSecurityFindingView() *SecurityFindingView`

NewSecurityFindingView instantiates a new SecurityFindingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityFindingViewWithDefaults

`func NewSecurityFindingViewWithDefaults() *SecurityFindingView`

NewSecurityFindingViewWithDefaults instantiates a new SecurityFindingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SecurityFindingView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityFindingView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityFindingView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityFindingView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFingerprint

`func (o *SecurityFindingView) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SecurityFindingView) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *SecurityFindingView) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *SecurityFindingView) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetId

`func (o *SecurityFindingView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityFindingView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityFindingView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityFindingView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLine

`func (o *SecurityFindingView) GetLine() int64`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *SecurityFindingView) GetLineOk() (*int64, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *SecurityFindingView) SetLine(v int64)`

SetLine sets Line field to given value.

### HasLine

`func (o *SecurityFindingView) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetPath

`func (o *SecurityFindingView) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SecurityFindingView) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SecurityFindingView) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SecurityFindingView) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPreview

`func (o *SecurityFindingView) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *SecurityFindingView) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *SecurityFindingView) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *SecurityFindingView) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetRuleId

`func (o *SecurityFindingView) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *SecurityFindingView) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *SecurityFindingView) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *SecurityFindingView) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *SecurityFindingView) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *SecurityFindingView) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *SecurityFindingView) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *SecurityFindingView) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetScanId

`func (o *SecurityFindingView) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *SecurityFindingView) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *SecurityFindingView) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *SecurityFindingView) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetSeverity

`func (o *SecurityFindingView) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SecurityFindingView) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SecurityFindingView) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SecurityFindingView) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


