# ReferenceReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** | AsOf is when the load happened, RFC 3339. Absent is dated on arrival, which can only make the list look older than it is. | [optional] 
**Keys** | Pointer to **int64** | Keys is how many designations that load carried. Zero from a publisher who designates somebody is a failed load wearing a successful one&#39;s clothes, and belongs in Refusal instead. | [optional] 
**Refusal** | Pointer to **string** | Refusal is why the load failed, when it did. | [optional] 
**Source** | Pointer to **string** | Source is the publisher this receipt is for. | [optional] 
**Version** | Pointer to **string** | Version is the digest of what that publisher supplied, so a refresh that changed nothing can be told from a refresh that did not run. | [optional] 

## Methods

### NewReferenceReceipt

`func NewReferenceReceipt() *ReferenceReceipt`

NewReferenceReceipt instantiates a new ReferenceReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceReceiptWithDefaults

`func NewReferenceReceiptWithDefaults() *ReferenceReceipt`

NewReferenceReceiptWithDefaults instantiates a new ReferenceReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *ReferenceReceipt) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ReferenceReceipt) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ReferenceReceipt) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *ReferenceReceipt) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetKeys

`func (o *ReferenceReceipt) GetKeys() int64`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ReferenceReceipt) GetKeysOk() (*int64, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ReferenceReceipt) SetKeys(v int64)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ReferenceReceipt) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetRefusal

`func (o *ReferenceReceipt) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *ReferenceReceipt) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *ReferenceReceipt) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *ReferenceReceipt) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetSource

`func (o *ReferenceReceipt) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReferenceReceipt) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReferenceReceipt) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReferenceReceipt) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *ReferenceReceipt) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReferenceReceipt) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReferenceReceipt) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReferenceReceipt) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


