# GpuView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the card&#39;s address: its host machine&#39;s id, \&quot;#\&quot;, and the card&#39;s ordinal within that machine (\&quot;gpu-1#0\&quot;). Stable for as long as the machine is, and the only id a single accelerator has — providers do not name cards. | [optional] 
**Location** | Pointer to **string** | Location is where the card physically sits, which for every source today is the same value Region carries — the console renders it in its own column. | [optional] 
**Machine** | Pointer to **string** | Machine is the id of the machine holding this card, addressable as-is on /v1/compute/machines/:id. | [optional] 
**Memory** | Pointer to **string** | Memory is the card&#39;s VRAM as its own tooling reported it (\&quot;122880 MiB\&quot;) — a display string in the reporter&#39;s units, not a byte count. BYO cards carry it (nvidia-smi); Visor&#39;s machine object states no VRAM, so a rented card leaves it empty and the console renders \&quot;—\&quot; rather than a fabricated 0. | [optional] 
**Model** | Pointer to **string** | Model is the accelerator: the model token read out of the size slug for a Visor GPU droplet (\&quot;H100\&quot;, \&quot;MI300X\&quot;), or the name nvidia-smi reported for a BYO card (\&quot;NVIDIA GB10\&quot;). | [optional] 
**Name** | Pointer to **string** | Name is the HOST MACHINE&#39;s display name, not the card&#39;s — every card in a gpu-h100x8 node repeats it. Model is what says which accelerator this is. | [optional] 
**Provider** | Pointer to **string** | Provider distinguishes a BYO accelerator (\&quot;byo\&quot;) from a Visor-provisioned one (the host machine&#39;s real provider). It is what tells a card the org owns from a card the org rents. | [optional] 
**Region** | Pointer to **string** | Region is the host machine&#39;s provider region slug; \&quot;on-prem\&quot; for a BYO card. | [optional] 
**Status** | Pointer to **string** | Status is the HOST MACHINE&#39;s lifecycle state, because nothing upstream reports a card&#39;s own health. A card reads running because its machine does. | [optional] 

## Methods

### NewGpuView

`func NewGpuView() *GpuView`

NewGpuView instantiates a new GpuView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuViewWithDefaults

`func NewGpuViewWithDefaults() *GpuView`

NewGpuViewWithDefaults instantiates a new GpuView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GpuView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GpuView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GpuView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GpuView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *GpuView) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GpuView) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GpuView) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GpuView) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMachine

`func (o *GpuView) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *GpuView) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *GpuView) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *GpuView) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMemory

`func (o *GpuView) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GpuView) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GpuView) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *GpuView) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetModel

`func (o *GpuView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GpuView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GpuView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GpuView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *GpuView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GpuView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GpuView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GpuView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *GpuView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GpuView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GpuView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GpuView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *GpuView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GpuView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GpuView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GpuView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *GpuView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GpuView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GpuView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GpuView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


