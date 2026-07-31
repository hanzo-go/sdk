# SecurityScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Files** | Pointer to **int32** |  | [optional] 
**Findings** | Pointer to **int32** |  | [optional] 
**Critical** | Pointer to **int32** |  | [optional] 
**High** | Pointer to **int32** |  | [optional] 
**Medium** | Pointer to **int32** |  | [optional] 
**Low** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int64** | Unix milliseconds | [optional] 

## Methods

### NewSecurityScan

`func NewSecurityScan() *SecurityScan`

NewSecurityScan instantiates a new SecurityScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanWithDefaults

`func NewSecurityScanWithDefaults() *SecurityScan`

NewSecurityScanWithDefaults instantiates a new SecurityScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityScan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityScan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityScan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityScan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *SecurityScan) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SecurityScan) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SecurityScan) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SecurityScan) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetFiles

`func (o *SecurityScan) GetFiles() int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SecurityScan) GetFilesOk() (*int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SecurityScan) SetFiles(v int32)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SecurityScan) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFindings

`func (o *SecurityScan) GetFindings() int32`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *SecurityScan) GetFindingsOk() (*int32, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *SecurityScan) SetFindings(v int32)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *SecurityScan) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### GetCritical

`func (o *SecurityScan) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *SecurityScan) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *SecurityScan) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *SecurityScan) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetHigh

`func (o *SecurityScan) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *SecurityScan) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *SecurityScan) SetHigh(v int32)`

SetHigh sets High field to given value.

### HasHigh

`func (o *SecurityScan) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetMedium

`func (o *SecurityScan) GetMedium() int32`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *SecurityScan) GetMediumOk() (*int32, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *SecurityScan) SetMedium(v int32)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *SecurityScan) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetLow

`func (o *SecurityScan) GetLow() int32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *SecurityScan) GetLowOk() (*int32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *SecurityScan) SetLow(v int32)`

SetLow sets Low field to given value.

### HasLow

`func (o *SecurityScan) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SecurityScan) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityScan) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityScan) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityScan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


