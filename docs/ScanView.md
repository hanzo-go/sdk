# ScanView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when the scan ran, in Unix milliseconds. | [optional] 
**Critical** | Pointer to **int32** | Critical is how many findings carry the highest severity. | [optional] 
**Files** | Pointer to **int32** | Files is how many files the scan read. | [optional] 
**Findings** | Pointer to **int32** | Findings is how many secrets fired across them. | [optional] 
**High** | Pointer to **int32** | High is how many findings rank high. | [optional] 
**Id** | Pointer to **string** | ID addresses this scan and every finding on it. | [optional] 
**Low** | Pointer to **int32** | Low is how many findings rank low. | [optional] 
**Medium** | Pointer to **int32** | Medium is how many findings rank medium. | [optional] 
**Project** | Pointer to **string** | Project is the sub-scope the scan was filed under. | [optional] 

## Methods

### NewScanView

`func NewScanView() *ScanView`

NewScanView instantiates a new ScanView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanViewWithDefaults

`func NewScanViewWithDefaults() *ScanView`

NewScanViewWithDefaults instantiates a new ScanView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ScanView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScanView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScanView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ScanView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCritical

`func (o *ScanView) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *ScanView) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *ScanView) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *ScanView) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetFiles

`func (o *ScanView) GetFiles() int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ScanView) GetFilesOk() (*int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ScanView) SetFiles(v int32)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ScanView) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFindings

`func (o *ScanView) GetFindings() int32`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *ScanView) GetFindingsOk() (*int32, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *ScanView) SetFindings(v int32)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *ScanView) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### GetHigh

`func (o *ScanView) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *ScanView) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *ScanView) SetHigh(v int32)`

SetHigh sets High field to given value.

### HasHigh

`func (o *ScanView) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetId

`func (o *ScanView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScanView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScanView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScanView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLow

`func (o *ScanView) GetLow() int32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *ScanView) GetLowOk() (*int32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *ScanView) SetLow(v int32)`

SetLow sets Low field to given value.

### HasLow

`func (o *ScanView) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetMedium

`func (o *ScanView) GetMedium() int32`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *ScanView) GetMediumOk() (*int32, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *ScanView) SetMedium(v int32)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *ScanView) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetProject

`func (o *ScanView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ScanView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ScanView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ScanView) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


