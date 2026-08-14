# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | Pointer to **string** | Addr is the socket or address serving it. | [optional] 
**Disabled** | Pointer to **bool** | Disabled is true when Unload stopped it deliberately, as opposed to it having crashed. Both answer 503, so without this an operator cannot tell a maintenance window from an outage — and would page for the former. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **int32** | PID is the child process, or 0 when this host did not start it. | [optional] 
**Prefix** | Pointer to **string** | Prefix is the FIRST subtree this plugin answers — the one a log line names it by. Prefixes is every subtree, and a plugin may own several. Reporting only the first would understate the blast radius of taking this plugin down, which is the question a fleet view exists to answer. | [optional] 
**Prefixes** | Pointer to **[]string** |  | [optional] 
**Reloads** | Pointer to **int32** | Reloads counts successful swaps since Load. A climbing number on one host and not its peers is the signal that a rollout is uneven. | [optional] 
**Restarts** | Pointer to **int32** | Restarts counts times the supervisor brought this plugin back after it died on its own. Distinct from Reloads, which are deliberate: a nonzero Restarts is a plugin crashing, and a climbing one is a crash loop. | [optional] 
**Running** | Pointer to **bool** | Running is false after Unload, or after a child exited and no Reload has replaced it. Its routes stay registered and answer 503, so a false here is the difference between \&quot;not deployed\&quot; and \&quot;deployed but down\&quot;. | [optional] 
**Since** | Pointer to **time.Time** | Since is when the CURRENT instance started — it resets on Reload, so it reports the age of what is running, not of the mount. | [optional] 
**Source** | Pointer to **string** | Source is where the binary came from: \&quot;embedded\&quot;, \&quot;path\&quot;, \&quot;url\&quot;, or \&quot;remote\&quot; for an instance this host did not start. | [optional] 
**Usage** | Pointer to [**Usage**](Usage.md) | Usage is what this plugin costs right now, read from the kernel. | [optional] 
**Version** | Pointer to **string** | Version is the artifact&#39;s SHA-256 when it was installed from a URL — the only version identifier that cannot drift from the bits actually running, since it IS the bits. Empty for the other sources. | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *Status) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *Status) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *Status) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *Status) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetDisabled

`func (o *Status) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Status) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Status) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Status) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetName

`func (o *Status) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Status) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Status) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Status) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPid

`func (o *Status) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *Status) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *Status) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *Status) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPrefix

`func (o *Status) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *Status) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *Status) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *Status) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPrefixes

`func (o *Status) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *Status) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *Status) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.

### HasPrefixes

`func (o *Status) HasPrefixes() bool`

HasPrefixes returns a boolean if a field has been set.

### GetReloads

`func (o *Status) GetReloads() int32`

GetReloads returns the Reloads field if non-nil, zero value otherwise.

### GetReloadsOk

`func (o *Status) GetReloadsOk() (*int32, bool)`

GetReloadsOk returns a tuple with the Reloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReloads

`func (o *Status) SetReloads(v int32)`

SetReloads sets Reloads field to given value.

### HasReloads

`func (o *Status) HasReloads() bool`

HasReloads returns a boolean if a field has been set.

### GetRestarts

`func (o *Status) GetRestarts() int32`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *Status) GetRestartsOk() (*int32, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *Status) SetRestarts(v int32)`

SetRestarts sets Restarts field to given value.

### HasRestarts

`func (o *Status) HasRestarts() bool`

HasRestarts returns a boolean if a field has been set.

### GetRunning

`func (o *Status) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *Status) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *Status) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *Status) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSince

`func (o *Status) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *Status) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *Status) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *Status) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetSource

`func (o *Status) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Status) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Status) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Status) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUsage

`func (o *Status) GetUsage() Usage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Status) GetUsageOk() (*Usage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Status) SetUsage(v Usage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Status) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetVersion

`func (o *Status) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Status) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Status) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Status) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


