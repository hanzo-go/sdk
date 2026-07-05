# MqListSubjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subjects** | Pointer to [**[]MqSubjectInfo**](MqSubjectInfo.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewMqListSubjects200Response

`func NewMqListSubjects200Response() *MqListSubjects200Response`

NewMqListSubjects200Response instantiates a new MqListSubjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqListSubjects200ResponseWithDefaults

`func NewMqListSubjects200ResponseWithDefaults() *MqListSubjects200Response`

NewMqListSubjects200ResponseWithDefaults instantiates a new MqListSubjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjects

`func (o *MqListSubjects200Response) GetSubjects() []MqSubjectInfo`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *MqListSubjects200Response) GetSubjectsOk() (*[]MqSubjectInfo, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *MqListSubjects200Response) SetSubjects(v []MqSubjectInfo)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *MqListSubjects200Response) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetTotal

`func (o *MqListSubjects200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MqListSubjects200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MqListSubjects200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MqListSubjects200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


