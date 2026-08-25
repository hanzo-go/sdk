# WorkerScriptPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to **interface{}** |  | [optional] 
**CompatibilityDate** | Pointer to **string** | CompatibilityDate pins which Workers runtime behaviour the script runs under, as a plain calendar date (\&quot;2024-01-01\&quot;). Absent leaves the account&#39;s own default in force. | [optional] 
**CompatibilityFlags** | Pointer to **[]string** | CompatibilityFlags turn individual runtime behaviours on or off around that date (\&quot;nodejs_compat\&quot;), in Cloudflare&#39;s own flag vocabulary. Absent means the date alone decides. | [optional] 
**MainModule** | Pointer to **string** | MainModule is the module file the runtime starts at. Absent means \&quot;worker.js\&quot;. | [optional] 
**Script** | Pointer to **string** | Script means two things on this route, and the document says so in both places it appears: the PATH segment names the Worker to publish, and the BODY field carries that Worker&#39;s ES-module source — the code itself, never a name or a URL. A blank or absent source is refused; there is no empty Worker. | [optional] 

## Methods

### NewWorkerScriptPut

`func NewWorkerScriptPut() *WorkerScriptPut`

NewWorkerScriptPut instantiates a new WorkerScriptPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerScriptPutWithDefaults

`func NewWorkerScriptPutWithDefaults() *WorkerScriptPut`

NewWorkerScriptPutWithDefaults instantiates a new WorkerScriptPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *WorkerScriptPut) GetBindings() interface{}`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *WorkerScriptPut) GetBindingsOk() (*interface{}, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *WorkerScriptPut) SetBindings(v interface{})`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *WorkerScriptPut) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### SetBindingsNil

`func (o *WorkerScriptPut) SetBindingsNil(b bool)`

 SetBindingsNil sets the value for Bindings to be an explicit nil

### UnsetBindings
`func (o *WorkerScriptPut) UnsetBindings()`

UnsetBindings ensures that no value is present for Bindings, not even an explicit nil
### GetCompatibilityDate

`func (o *WorkerScriptPut) GetCompatibilityDate() string`

GetCompatibilityDate returns the CompatibilityDate field if non-nil, zero value otherwise.

### GetCompatibilityDateOk

`func (o *WorkerScriptPut) GetCompatibilityDateOk() (*string, bool)`

GetCompatibilityDateOk returns a tuple with the CompatibilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityDate

`func (o *WorkerScriptPut) SetCompatibilityDate(v string)`

SetCompatibilityDate sets CompatibilityDate field to given value.

### HasCompatibilityDate

`func (o *WorkerScriptPut) HasCompatibilityDate() bool`

HasCompatibilityDate returns a boolean if a field has been set.

### GetCompatibilityFlags

`func (o *WorkerScriptPut) GetCompatibilityFlags() []string`

GetCompatibilityFlags returns the CompatibilityFlags field if non-nil, zero value otherwise.

### GetCompatibilityFlagsOk

`func (o *WorkerScriptPut) GetCompatibilityFlagsOk() (*[]string, bool)`

GetCompatibilityFlagsOk returns a tuple with the CompatibilityFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityFlags

`func (o *WorkerScriptPut) SetCompatibilityFlags(v []string)`

SetCompatibilityFlags sets CompatibilityFlags field to given value.

### HasCompatibilityFlags

`func (o *WorkerScriptPut) HasCompatibilityFlags() bool`

HasCompatibilityFlags returns a boolean if a field has been set.

### GetMainModule

`func (o *WorkerScriptPut) GetMainModule() string`

GetMainModule returns the MainModule field if non-nil, zero value otherwise.

### GetMainModuleOk

`func (o *WorkerScriptPut) GetMainModuleOk() (*string, bool)`

GetMainModuleOk returns a tuple with the MainModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainModule

`func (o *WorkerScriptPut) SetMainModule(v string)`

SetMainModule sets MainModule field to given value.

### HasMainModule

`func (o *WorkerScriptPut) HasMainModule() bool`

HasMainModule returns a boolean if a field has been set.

### GetScript

`func (o *WorkerScriptPut) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *WorkerScriptPut) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *WorkerScriptPut) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *WorkerScriptPut) HasScript() bool`

HasScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


