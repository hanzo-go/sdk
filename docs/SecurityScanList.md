# SecurityScanList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SecurityScanView**](SecurityScanView.md) | Data is the caller org&#39;s scans, newest first. | [optional] 

## Methods

### NewSecurityScanList

`func NewSecurityScanList() *SecurityScanList`

NewSecurityScanList instantiates a new SecurityScanList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanListWithDefaults

`func NewSecurityScanListWithDefaults() *SecurityScanList`

NewSecurityScanListWithDefaults instantiates a new SecurityScanList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SecurityScanList) GetData() []SecurityScanView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecurityScanList) GetDataOk() (*[]SecurityScanView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecurityScanList) SetData(v []SecurityScanView)`

SetData sets Data field to given value.

### HasData

`func (o *SecurityScanList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


