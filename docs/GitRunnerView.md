# GitRunnerView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ephemeral** | Pointer to **bool** | Ephemeral means the daemon takes ONE job and exits, so it is not expected back and its absence is not a fault. | [optional] 
**Labels** | Pointer to **[]string** | Labels are what the daemon last said it can do. | [optional] 
**LastActive** | Pointer to **int64** | LastActive is when it was last EXECUTING a job, in unix seconds. It lags LastOnline on an idle daemon, which is how the two differ. | [optional] 
**LastOnline** | Pointer to **int64** | LastOnline is when it last called at all, in unix seconds. | [optional] 
**Name** | Pointer to **string** | Name is what the daemon called itself when it registered. | [optional] 
**Pool** | Pointer to **string** | Pool is the declared capacity it entered. | [optional] 
**Version** | Pointer to **string** | Version is the runner build the daemon last reported. | [optional] 

## Methods

### NewGitRunnerView

`func NewGitRunnerView() *GitRunnerView`

NewGitRunnerView instantiates a new GitRunnerView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRunnerViewWithDefaults

`func NewGitRunnerViewWithDefaults() *GitRunnerView`

NewGitRunnerViewWithDefaults instantiates a new GitRunnerView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEphemeral

`func (o *GitRunnerView) GetEphemeral() bool`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *GitRunnerView) GetEphemeralOk() (*bool, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *GitRunnerView) SetEphemeral(v bool)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *GitRunnerView) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetLabels

`func (o *GitRunnerView) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GitRunnerView) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GitRunnerView) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GitRunnerView) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastActive

`func (o *GitRunnerView) GetLastActive() int64`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *GitRunnerView) GetLastActiveOk() (*int64, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *GitRunnerView) SetLastActive(v int64)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *GitRunnerView) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetLastOnline

`func (o *GitRunnerView) GetLastOnline() int64`

GetLastOnline returns the LastOnline field if non-nil, zero value otherwise.

### GetLastOnlineOk

`func (o *GitRunnerView) GetLastOnlineOk() (*int64, bool)`

GetLastOnlineOk returns a tuple with the LastOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnline

`func (o *GitRunnerView) SetLastOnline(v int64)`

SetLastOnline sets LastOnline field to given value.

### HasLastOnline

`func (o *GitRunnerView) HasLastOnline() bool`

HasLastOnline returns a boolean if a field has been set.

### GetName

`func (o *GitRunnerView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitRunnerView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitRunnerView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitRunnerView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPool

`func (o *GitRunnerView) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GitRunnerView) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GitRunnerView) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *GitRunnerView) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetVersion

`func (o *GitRunnerView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GitRunnerView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GitRunnerView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GitRunnerView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


