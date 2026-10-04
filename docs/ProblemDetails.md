# ProblemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | A stable token a client can branch on, when the refusal names one. | [optional] 
**Detail** | **string** | What was refused, for a person to read. | 
**Status** | **int64** | The HTTP status. | 
**Title** | Pointer to **string** | The status&#39;s reason phrase. | [optional] 
**Type** | **string** | about:blank | 

## Methods

### NewProblemDetails

`func NewProblemDetails(detail string, status int64, type_ string, ) *ProblemDetails`

NewProblemDetails instantiates a new ProblemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemDetailsWithDefaults

`func NewProblemDetailsWithDefaults() *ProblemDetails`

NewProblemDetailsWithDefaults instantiates a new ProblemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ProblemDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProblemDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProblemDetails) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProblemDetails) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *ProblemDetails) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ProblemDetails) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ProblemDetails) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetStatus

`func (o *ProblemDetails) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProblemDetails) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProblemDetails) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *ProblemDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProblemDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProblemDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProblemDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *ProblemDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProblemDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProblemDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


