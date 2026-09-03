# TrustCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controls** | Pointer to [**TrustTally**](TrustTally.md) | Controls is how the controls stand, independent of any framework. | [optional] 
**Frameworks** | Pointer to [**[]CoverRow**](CoverRow.md) | Frameworks is the per-framework counts. | [optional] 
**Generated** | Pointer to **int64** | Generated is when this was computed, unix milliseconds. | [optional] 
**Version** | Pointer to **string** | Version is the embedded inventory&#39;s version. | [optional] 

## Methods

### NewTrustCoverage

`func NewTrustCoverage() *TrustCoverage`

NewTrustCoverage instantiates a new TrustCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustCoverageWithDefaults

`func NewTrustCoverageWithDefaults() *TrustCoverage`

NewTrustCoverageWithDefaults instantiates a new TrustCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControls

`func (o *TrustCoverage) GetControls() TrustTally`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *TrustCoverage) GetControlsOk() (*TrustTally, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *TrustCoverage) SetControls(v TrustTally)`

SetControls sets Controls field to given value.

### HasControls

`func (o *TrustCoverage) HasControls() bool`

HasControls returns a boolean if a field has been set.

### GetFrameworks

`func (o *TrustCoverage) GetFrameworks() []CoverRow`

GetFrameworks returns the Frameworks field if non-nil, zero value otherwise.

### GetFrameworksOk

`func (o *TrustCoverage) GetFrameworksOk() (*[]CoverRow, bool)`

GetFrameworksOk returns a tuple with the Frameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworks

`func (o *TrustCoverage) SetFrameworks(v []CoverRow)`

SetFrameworks sets Frameworks field to given value.

### HasFrameworks

`func (o *TrustCoverage) HasFrameworks() bool`

HasFrameworks returns a boolean if a field has been set.

### GetGenerated

`func (o *TrustCoverage) GetGenerated() int64`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *TrustCoverage) GetGeneratedOk() (*int64, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *TrustCoverage) SetGenerated(v int64)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *TrustCoverage) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetVersion

`func (o *TrustCoverage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrustCoverage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrustCoverage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TrustCoverage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


