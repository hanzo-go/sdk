# TrustTrustCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controls** | Pointer to [**TrustTrustTally**](TrustTrustTally.md) | Controls is how the controls stand, independent of any framework. | [optional] 
**Frameworks** | Pointer to [**[]TrustCoverRow**](TrustCoverRow.md) | Frameworks is the per-framework counts. | [optional] 
**Generated** | Pointer to **int64** | Generated is when this was computed, unix milliseconds. | [optional] 
**Version** | Pointer to **string** | Version is the embedded inventory&#39;s version. | [optional] 

## Methods

### NewTrustTrustCoverage

`func NewTrustTrustCoverage() *TrustTrustCoverage`

NewTrustTrustCoverage instantiates a new TrustTrustCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustTrustCoverageWithDefaults

`func NewTrustTrustCoverageWithDefaults() *TrustTrustCoverage`

NewTrustTrustCoverageWithDefaults instantiates a new TrustTrustCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControls

`func (o *TrustTrustCoverage) GetControls() TrustTrustTally`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *TrustTrustCoverage) GetControlsOk() (*TrustTrustTally, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *TrustTrustCoverage) SetControls(v TrustTrustTally)`

SetControls sets Controls field to given value.

### HasControls

`func (o *TrustTrustCoverage) HasControls() bool`

HasControls returns a boolean if a field has been set.

### GetFrameworks

`func (o *TrustTrustCoverage) GetFrameworks() []TrustCoverRow`

GetFrameworks returns the Frameworks field if non-nil, zero value otherwise.

### GetFrameworksOk

`func (o *TrustTrustCoverage) GetFrameworksOk() (*[]TrustCoverRow, bool)`

GetFrameworksOk returns a tuple with the Frameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworks

`func (o *TrustTrustCoverage) SetFrameworks(v []TrustCoverRow)`

SetFrameworks sets Frameworks field to given value.

### HasFrameworks

`func (o *TrustTrustCoverage) HasFrameworks() bool`

HasFrameworks returns a boolean if a field has been set.

### GetGenerated

`func (o *TrustTrustCoverage) GetGenerated() int64`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *TrustTrustCoverage) GetGeneratedOk() (*int64, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *TrustTrustCoverage) SetGenerated(v int64)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *TrustTrustCoverage) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetVersion

`func (o *TrustTrustCoverage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrustTrustCoverage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrustTrustCoverage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TrustTrustCoverage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


