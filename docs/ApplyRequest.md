# ApplyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedCode** | Pointer to **string** | RequestedCode is the vanity code the applicant asks for; approval may mint a different one if it is taken. Body-only: the URL cannot supply it. | [optional] 

## Methods

### NewApplyRequest

`func NewApplyRequest() *ApplyRequest`

NewApplyRequest instantiates a new ApplyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyRequestWithDefaults

`func NewApplyRequestWithDefaults() *ApplyRequest`

NewApplyRequestWithDefaults instantiates a new ApplyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedCode

`func (o *ApplyRequest) GetRequestedCode() string`

GetRequestedCode returns the RequestedCode field if non-nil, zero value otherwise.

### GetRequestedCodeOk

`func (o *ApplyRequest) GetRequestedCodeOk() (*string, bool)`

GetRequestedCodeOk returns a tuple with the RequestedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCode

`func (o *ApplyRequest) SetRequestedCode(v string)`

SetRequestedCode sets RequestedCode field to given value.

### HasRequestedCode

`func (o *ApplyRequest) HasRequestedCode() bool`

HasRequestedCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


