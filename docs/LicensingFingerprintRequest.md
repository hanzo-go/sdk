# LicensingFingerprintRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signals** | Pointer to [**LicensingDeviceSignals**](LicensingDeviceSignals.md) | Signals is the host material the client agent collected. Which fields actually participate in the binding is deliberately unspecified — send everything available and let the server decide. | [optional] 

## Methods

### NewLicensingFingerprintRequest

`func NewLicensingFingerprintRequest() *LicensingFingerprintRequest`

NewLicensingFingerprintRequest instantiates a new LicensingFingerprintRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingFingerprintRequestWithDefaults

`func NewLicensingFingerprintRequestWithDefaults() *LicensingFingerprintRequest`

NewLicensingFingerprintRequestWithDefaults instantiates a new LicensingFingerprintRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignals

`func (o *LicensingFingerprintRequest) GetSignals() LicensingDeviceSignals`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *LicensingFingerprintRequest) GetSignalsOk() (*LicensingDeviceSignals, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *LicensingFingerprintRequest) SetSignals(v LicensingDeviceSignals)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *LicensingFingerprintRequest) HasSignals() bool`

HasSignals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


