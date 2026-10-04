# SecurityScanView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the scan ran, in Unix milliseconds. | [optional] 
**Critical** | Pointer to **int64** | Critical is how many findings carry the highest severity. | [optional] 
**Files** | Pointer to **int64** | Files is how many files the scan read. | [optional] 
**Findings** | Pointer to **int64** | Findings is how many secrets fired across them. | [optional] 
**High** | Pointer to **int64** | High is how many findings rank high. | [optional] 
**Id** | Pointer to **string** | ID addresses this scan and every finding on it. | [optional] 
**Low** | Pointer to **int64** | Low is how many findings rank low. | [optional] 
**Medium** | Pointer to **int64** | Medium is how many findings rank medium. | [optional] 
**Project** | Pointer to **string** | Project is the sub-scope the scan was filed under. | [optional] 

## Methods

### NewSecurityScanView

`func NewSecurityScanView() *SecurityScanView`

NewSecurityScanView instantiates a new SecurityScanView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanViewWithDefaults

`func NewSecurityScanViewWithDefaults() *SecurityScanView`

NewSecurityScanViewWithDefaults instantiates a new SecurityScanView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SecurityScanView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityScanView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityScanView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityScanView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCritical

`func (o *SecurityScanView) GetCritical() int64`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *SecurityScanView) GetCriticalOk() (*int64, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *SecurityScanView) SetCritical(v int64)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *SecurityScanView) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetFiles

`func (o *SecurityScanView) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SecurityScanView) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SecurityScanView) SetFiles(v int64)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SecurityScanView) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFindings

`func (o *SecurityScanView) GetFindings() int64`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *SecurityScanView) GetFindingsOk() (*int64, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *SecurityScanView) SetFindings(v int64)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *SecurityScanView) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### GetHigh

`func (o *SecurityScanView) GetHigh() int64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *SecurityScanView) GetHighOk() (*int64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *SecurityScanView) SetHigh(v int64)`

SetHigh sets High field to given value.

### HasHigh

`func (o *SecurityScanView) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetId

`func (o *SecurityScanView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityScanView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityScanView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityScanView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLow

`func (o *SecurityScanView) GetLow() int64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *SecurityScanView) GetLowOk() (*int64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *SecurityScanView) SetLow(v int64)`

SetLow sets Low field to given value.

### HasLow

`func (o *SecurityScanView) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetMedium

`func (o *SecurityScanView) GetMedium() int64`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *SecurityScanView) GetMediumOk() (*int64, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *SecurityScanView) SetMedium(v int64)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *SecurityScanView) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetProject

`func (o *SecurityScanView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SecurityScanView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SecurityScanView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SecurityScanView) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


