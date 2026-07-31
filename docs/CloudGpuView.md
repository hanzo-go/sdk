# CloudGpuView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Machine** | Pointer to **string** |  | [optional] 
**Memory** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** | Provider distinguishes a BYO accelerator (\&quot;byo\&quot;) from a Visor-provisioned one (the machine&#39;s real provider). Memory is VRAM when known (BYO reports it from nvidia-smi; Visor&#39;s machine object carries none, so it stays empty and the UI renders \&quot;—\&quot;). Both are additive + omitempty: existing rows are unaffected and the console normalizer ignores fields it does not read. | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudGpuView

`func NewCloudGpuView() *CloudGpuView`

NewCloudGpuView instantiates a new CloudGpuView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGpuViewWithDefaults

`func NewCloudGpuViewWithDefaults() *CloudGpuView`

NewCloudGpuViewWithDefaults instantiates a new CloudGpuView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudGpuView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudGpuView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudGpuView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudGpuView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *CloudGpuView) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudGpuView) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudGpuView) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CloudGpuView) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMachine

`func (o *CloudGpuView) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *CloudGpuView) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *CloudGpuView) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *CloudGpuView) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMemory

`func (o *CloudGpuView) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CloudGpuView) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CloudGpuView) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CloudGpuView) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetModel

`func (o *CloudGpuView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudGpuView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudGpuView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CloudGpuView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *CloudGpuView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudGpuView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudGpuView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudGpuView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *CloudGpuView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudGpuView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudGpuView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudGpuView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *CloudGpuView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudGpuView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudGpuView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudGpuView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *CloudGpuView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudGpuView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudGpuView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudGpuView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


