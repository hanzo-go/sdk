# PluginPluginStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** | Prefix is the FIRST subtree this plugin answers — the one a log line names it by. Prefixes is every subtree, and a plugin may own several. Reporting only the first would understate the blast radius of taking this plugin down, which is the question a fleet view exists to answer.  | [optional] 
**Prefixes** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **string** | Source is where the binary came from: \&quot;embedded\&quot;, \&quot;path\&quot;, \&quot;url\&quot;, or \&quot;remote\&quot; for an instance this host did not start.  | [optional] 
**Version** | Pointer to **string** | Version is the artifact&#39;s SHA-256 when it was installed from a URL — the only version identifier that cannot drift from the bits actually running, since it IS the bits. Empty for the other sources.  | [optional] 
**Addr** | Pointer to **string** | Addr is the socket or address serving it. | [optional] 
**Pid** | Pointer to **int32** | PID is the child process, or 0 when this host did not start it. | [optional] 
**Running** | Pointer to **bool** | Running is false after Unload, or after a child exited and no Reload has replaced it. Its routes stay registered and answer 503, so a false here is the difference between \&quot;not deployed\&quot; and \&quot;deployed but down\&quot;.  | [optional] 
**Disabled** | Pointer to **bool** | Disabled is true when Unload stopped it deliberately, as opposed to it having crashed. Both answer 503, so without this an operator cannot tell a maintenance window from an outage — and would page for the former.  | [optional] 
**Since** | Pointer to **time.Time** | Since is when the CURRENT instance started — it resets on Reload, so it reports the age of what is running, not of the mount.  | [optional] 
**Reloads** | Pointer to **int32** | Reloads counts successful swaps since Load. A climbing number on one host and not its peers is the signal that a rollout is uneven.  | [optional] 
**Restarts** | Pointer to **int32** | Restarts counts times the supervisor brought this plugin back after it died on its own. Distinct from Reloads, which are deliberate: a nonzero Restarts is a plugin crashing, and a climbing one is a crash loop.  | [optional] 
**Usage** | Pointer to [**PluginPluginUsage**](PluginPluginUsage.md) |  | [optional] 

## Methods

### NewPluginPluginStatus

`func NewPluginPluginStatus() *PluginPluginStatus`

NewPluginPluginStatus instantiates a new PluginPluginStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginPluginStatusWithDefaults

`func NewPluginPluginStatusWithDefaults() *PluginPluginStatus`

NewPluginPluginStatusWithDefaults instantiates a new PluginPluginStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PluginPluginStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginPluginStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginPluginStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginPluginStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *PluginPluginStatus) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PluginPluginStatus) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PluginPluginStatus) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PluginPluginStatus) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPrefixes

`func (o *PluginPluginStatus) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *PluginPluginStatus) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *PluginPluginStatus) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.

### HasPrefixes

`func (o *PluginPluginStatus) HasPrefixes() bool`

HasPrefixes returns a boolean if a field has been set.

### GetSource

`func (o *PluginPluginStatus) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PluginPluginStatus) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PluginPluginStatus) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PluginPluginStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *PluginPluginStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginPluginStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginPluginStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PluginPluginStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAddr

`func (o *PluginPluginStatus) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *PluginPluginStatus) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *PluginPluginStatus) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *PluginPluginStatus) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetPid

`func (o *PluginPluginStatus) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *PluginPluginStatus) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *PluginPluginStatus) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *PluginPluginStatus) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetRunning

`func (o *PluginPluginStatus) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *PluginPluginStatus) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *PluginPluginStatus) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *PluginPluginStatus) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetDisabled

`func (o *PluginPluginStatus) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PluginPluginStatus) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PluginPluginStatus) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PluginPluginStatus) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetSince

`func (o *PluginPluginStatus) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *PluginPluginStatus) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *PluginPluginStatus) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *PluginPluginStatus) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetReloads

`func (o *PluginPluginStatus) GetReloads() int32`

GetReloads returns the Reloads field if non-nil, zero value otherwise.

### GetReloadsOk

`func (o *PluginPluginStatus) GetReloadsOk() (*int32, bool)`

GetReloadsOk returns a tuple with the Reloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReloads

`func (o *PluginPluginStatus) SetReloads(v int32)`

SetReloads sets Reloads field to given value.

### HasReloads

`func (o *PluginPluginStatus) HasReloads() bool`

HasReloads returns a boolean if a field has been set.

### GetRestarts

`func (o *PluginPluginStatus) GetRestarts() int32`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *PluginPluginStatus) GetRestartsOk() (*int32, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *PluginPluginStatus) SetRestarts(v int32)`

SetRestarts sets Restarts field to given value.

### HasRestarts

`func (o *PluginPluginStatus) HasRestarts() bool`

HasRestarts returns a boolean if a field has been set.

### GetUsage

`func (o *PluginPluginStatus) GetUsage() PluginPluginUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PluginPluginStatus) GetUsageOk() (*PluginPluginUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PluginPluginStatus) SetUsage(v PluginPluginUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *PluginPluginStatus) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


