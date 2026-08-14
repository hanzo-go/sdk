# LicensingDeviceSignals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**Cpuid** | Pointer to **string** | CPUID is a CPU/board identifier string. | [optional] 
**DiskSerial** | Pointer to **string** | DiskSerial of the boot/root volume. | [optional] 
**Hostname** | Pointer to **string** | Hostname is a weak signal, used only as a tiebreaker. | [optional] 
**InstallId** | Pointer to **string** | InstallID is a per-install random the agent persists locally on first run. | [optional] 
**MachineId** | Pointer to **string** | MachineID is a stable per-host id (e.g. /etc/machine-id, IOPlatformUUID, MachineGuid). Strongest single signal where present. | [optional] 
**Macs** | Pointer to **[]string** | MAC addresses of stable interfaces (order-insensitive; we sort). | [optional] 
**Os** | Pointer to **string** | OS / Arch coarse platform tags. | [optional] 

## Methods

### NewLicensingDeviceSignals

`func NewLicensingDeviceSignals() *LicensingDeviceSignals`

NewLicensingDeviceSignals instantiates a new LicensingDeviceSignals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingDeviceSignalsWithDefaults

`func NewLicensingDeviceSignalsWithDefaults() *LicensingDeviceSignals`

NewLicensingDeviceSignalsWithDefaults instantiates a new LicensingDeviceSignals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *LicensingDeviceSignals) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *LicensingDeviceSignals) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *LicensingDeviceSignals) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *LicensingDeviceSignals) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCpuid

`func (o *LicensingDeviceSignals) GetCpuid() string`

GetCpuid returns the Cpuid field if non-nil, zero value otherwise.

### GetCpuidOk

`func (o *LicensingDeviceSignals) GetCpuidOk() (*string, bool)`

GetCpuidOk returns a tuple with the Cpuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuid

`func (o *LicensingDeviceSignals) SetCpuid(v string)`

SetCpuid sets Cpuid field to given value.

### HasCpuid

`func (o *LicensingDeviceSignals) HasCpuid() bool`

HasCpuid returns a boolean if a field has been set.

### GetDiskSerial

`func (o *LicensingDeviceSignals) GetDiskSerial() string`

GetDiskSerial returns the DiskSerial field if non-nil, zero value otherwise.

### GetDiskSerialOk

`func (o *LicensingDeviceSignals) GetDiskSerialOk() (*string, bool)`

GetDiskSerialOk returns a tuple with the DiskSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSerial

`func (o *LicensingDeviceSignals) SetDiskSerial(v string)`

SetDiskSerial sets DiskSerial field to given value.

### HasDiskSerial

`func (o *LicensingDeviceSignals) HasDiskSerial() bool`

HasDiskSerial returns a boolean if a field has been set.

### GetHostname

`func (o *LicensingDeviceSignals) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LicensingDeviceSignals) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LicensingDeviceSignals) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LicensingDeviceSignals) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInstallId

`func (o *LicensingDeviceSignals) GetInstallId() string`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *LicensingDeviceSignals) GetInstallIdOk() (*string, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *LicensingDeviceSignals) SetInstallId(v string)`

SetInstallId sets InstallId field to given value.

### HasInstallId

`func (o *LicensingDeviceSignals) HasInstallId() bool`

HasInstallId returns a boolean if a field has been set.

### GetMachineId

`func (o *LicensingDeviceSignals) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *LicensingDeviceSignals) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *LicensingDeviceSignals) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *LicensingDeviceSignals) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetMacs

`func (o *LicensingDeviceSignals) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *LicensingDeviceSignals) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *LicensingDeviceSignals) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *LicensingDeviceSignals) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetOs

`func (o *LicensingDeviceSignals) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *LicensingDeviceSignals) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *LicensingDeviceSignals) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *LicensingDeviceSignals) HasOs() bool`

HasOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


