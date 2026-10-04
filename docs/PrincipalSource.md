# PrincipalSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App names it: company, tax, wallet or agent. | [optional] 
**Status** | Pointer to **string** | Status is read, or absent when this deployment does not run it. | [optional] 

## Methods

### NewPrincipalSource

`func NewPrincipalSource() *PrincipalSource`

NewPrincipalSource instantiates a new PrincipalSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalSourceWithDefaults

`func NewPrincipalSourceWithDefaults() *PrincipalSource`

NewPrincipalSourceWithDefaults instantiates a new PrincipalSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PrincipalSource) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PrincipalSource) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PrincipalSource) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PrincipalSource) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetStatus

`func (o *PrincipalSource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrincipalSource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrincipalSource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrincipalSource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


