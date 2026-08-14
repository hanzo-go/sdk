# Leased

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Runtime** | Pointer to **string** | Runtime is the boundary this sandbox GOT, which need not be the one asked for — carried for the same reason Workdir is, that it is a fact only the owner knows and a caller assuming it would be holding a second copy. Empty is the node&#39;s default runtime, and a real answer. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Workdir** | Pointer to **string** |  | [optional] 

## Methods

### NewLeased

`func NewLeased() *Leased`

NewLeased instantiates a new Leased object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeasedWithDefaults

`func NewLeasedWithDefaults() *Leased`

NewLeasedWithDefaults instantiates a new Leased object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *Leased) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Leased) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Leased) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *Leased) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetId

`func (o *Leased) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Leased) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Leased) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Leased) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRuntime

`func (o *Leased) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Leased) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Leased) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *Leased) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetStatus

`func (o *Leased) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Leased) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Leased) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Leased) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkdir

`func (o *Leased) GetWorkdir() string`

GetWorkdir returns the Workdir field if non-nil, zero value otherwise.

### GetWorkdirOk

`func (o *Leased) GetWorkdirOk() (*string, bool)`

GetWorkdirOk returns a tuple with the Workdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkdir

`func (o *Leased) SetWorkdir(v string)`

SetWorkdir sets Workdir field to given value.

### HasWorkdir

`func (o *Leased) HasWorkdir() bool`

HasWorkdir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


