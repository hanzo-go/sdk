# CloudNewsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | Image is a lead-image URL when the upstream carried one. | [optional] 
**Lang** | Pointer to **string** | Lang is the article&#39;s language code when the upstream reported one. | [optional] 
**Link** | Pointer to **string** | Link is the article&#39;s URL at the outlet. | [optional] 
**PubDate** | Pointer to **string** | PubDate is when the outlet published it, RFC3339 UTC. Empty when the upstream gave no date this could parse — items with no date sort last. | [optional] 
**Source** | Pointer to **string** | Source is the outlet the item came from, as the upstream named it. | [optional] 
**Title** | Pointer to **string** | Title is the headline. | [optional] 
**Tone** | Pointer to **string** | Tone is GDELT&#39;s own sentiment score for the article, as text. Only GDELT items carry it. | [optional] 

## Methods

### NewCloudNewsItem

`func NewCloudNewsItem() *CloudNewsItem`

NewCloudNewsItem instantiates a new CloudNewsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNewsItemWithDefaults

`func NewCloudNewsItemWithDefaults() *CloudNewsItem`

NewCloudNewsItemWithDefaults instantiates a new CloudNewsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *CloudNewsItem) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudNewsItem) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudNewsItem) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudNewsItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLang

`func (o *CloudNewsItem) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *CloudNewsItem) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *CloudNewsItem) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *CloudNewsItem) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetLink

`func (o *CloudNewsItem) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CloudNewsItem) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CloudNewsItem) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *CloudNewsItem) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPubDate

`func (o *CloudNewsItem) GetPubDate() string`

GetPubDate returns the PubDate field if non-nil, zero value otherwise.

### GetPubDateOk

`func (o *CloudNewsItem) GetPubDateOk() (*string, bool)`

GetPubDateOk returns a tuple with the PubDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubDate

`func (o *CloudNewsItem) SetPubDate(v string)`

SetPubDate sets PubDate field to given value.

### HasPubDate

`func (o *CloudNewsItem) HasPubDate() bool`

HasPubDate returns a boolean if a field has been set.

### GetSource

`func (o *CloudNewsItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudNewsItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudNewsItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudNewsItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *CloudNewsItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudNewsItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudNewsItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudNewsItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTone

`func (o *CloudNewsItem) GetTone() string`

GetTone returns the Tone field if non-nil, zero value otherwise.

### GetToneOk

`func (o *CloudNewsItem) GetToneOk() (*string, bool)`

GetToneOk returns a tuple with the Tone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTone

`func (o *CloudNewsItem) SetTone(v string)`

SetTone sets Tone field to given value.

### HasTone

`func (o *CloudNewsItem) HasTone() bool`

HasTone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


