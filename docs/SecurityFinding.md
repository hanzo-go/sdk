# SecurityFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ScanId** | Pointer to **string** |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**RuleName** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Line** | Pointer to **int32** |  | [optional] 
**Preview** | Pointer to **string** | Redacted preview — never the raw secret | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** | Unix milliseconds | [optional] 

## Methods

### NewSecurityFinding

`func NewSecurityFinding() *SecurityFinding`

NewSecurityFinding instantiates a new SecurityFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityFindingWithDefaults

`func NewSecurityFindingWithDefaults() *SecurityFinding`

NewSecurityFindingWithDefaults instantiates a new SecurityFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityFinding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityFinding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityFinding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityFinding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScanId

`func (o *SecurityFinding) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *SecurityFinding) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *SecurityFinding) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *SecurityFinding) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetRuleId

`func (o *SecurityFinding) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *SecurityFinding) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *SecurityFinding) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *SecurityFinding) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *SecurityFinding) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *SecurityFinding) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *SecurityFinding) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *SecurityFinding) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetSeverity

`func (o *SecurityFinding) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SecurityFinding) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SecurityFinding) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SecurityFinding) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetPath

`func (o *SecurityFinding) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SecurityFinding) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SecurityFinding) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SecurityFinding) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetLine

`func (o *SecurityFinding) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *SecurityFinding) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *SecurityFinding) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *SecurityFinding) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetPreview

`func (o *SecurityFinding) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *SecurityFinding) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *SecurityFinding) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *SecurityFinding) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetFingerprint

`func (o *SecurityFinding) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SecurityFinding) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *SecurityFinding) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *SecurityFinding) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SecurityFinding) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityFinding) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityFinding) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityFinding) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


