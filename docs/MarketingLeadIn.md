# MarketingLeadIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**MarketingAddress**](MarketingAddress.md) | Address is where an investor lives, or where the entity they invest through is: street, city and country required, each at most 200 characters. | [optional] 
**Amount** | Pointer to **interface{}** |  | [optional] 
**Answers** | Pointer to [**[]MarketingAnswer**](MarketingAnswer.md) | Answers is the application question by question, as the form asked it: at most 40, each question clipped at 200 characters and each answer at 2 KiB. Recorded on the lead as they came; optional. | [optional] 
**Campaign** | Pointer to [**MarketingTags**](MarketingTags.md) | Campaign is the utm_* tags of the link the visit arrived on, recorded on the lead. Optional. | [optional] 
**Company** | Pointer to **string** | Company is the person&#39;s company. Optional; clipped at 200 characters. | [optional] 
**Consent** | Pointer to **bool** | Consent is the applicant ticking \&quot;Keep me posted about roles and news at Hanzo. Unsubscribe any time.\&quot; Unticked, they hear nothing but the reply to their application. Read only with Job. | [optional] 
**Country** | Pointer to **string** | Country is where the person is. Optional; clipped at 200 characters. | [optional] 
**Coupon** | Pointer to **string** | Coupon is the offer code the visitor&#39;s page held: letters, digits, dash and underscore, at most 32. Any other value is recorded as none. Optional. | [optional] 
**Department** | Pointer to **string** | Department is the opening&#39;s department. Read only with Job. | [optional] 
**Email** | Pointer to **string** | Email is where sales replies: the person&#39;s work email. Required, and it must look like an address. It is not verified. | [optional] 
**Entity** | Pointer to **string** | Entity is the legal name of the entity an investor invests through; empty, they invest in their own name. At most 200 characters. Read only with program investor, as are Title, Amount, Phone and Address. | [optional] 
**Job** | Pointer to **string** | Job is the opening a careers form applies to, by its id on the careers page (senior-backend-engineer). A body naming one is a job application: it is filed as a candidate, never as a lead, so it is offered no call and no code, no ad platform hears of it, and the brand&#39;s careers mailbox is told instead of sales. Name, Email, Role (the opening&#39;s title), Country (where the person is), Need (what they wrote about themselves), Source and Answers carry over. Optional; clipped at 200 characters. | [optional] 
**Links** | Pointer to **[]string** | Links are https addresses the applicant gave: a portfolio, a profile, a site. At most five, each at most 500 characters; any other value is dropped. Read only with Job. | [optional] 
**Name** | Pointer to **string** | Name is the person&#39;s name. Optional, and clipped at 200 characters, except for an investor, who must give it, in at most 200. | [optional] 
**Need** | Pointer to **string** | Need is what the person wants from sales, in their own words. Optional; clipped at 4 KiB. | [optional] 
**Payment** | Pointer to **string** | Payment is how an order the form files would be paid: \&quot;card\&quot; or \&quot;invoice\&quot; (a bank transfer against an invoice sales sends). Empty is none; anything else is 400. An Enterprise inquiry asking to pay by invoice is stated on the event plane as invoice_requested beside its generate_lead. | [optional] 
**Phone** | Pointer to **string** | Phone is an investor&#39;s phone number: 3 to 20 digits, an optional leading +, and any spaces, dots, dashes or brackets between them. Optional. | [optional] 
**Program** | Pointer to **string** | Program is the program the person applies to: beta, research or partner. Any other value is recorded as none. An application is mailed a code for 50% off the first three months, once per address (offer.go). Optional. \&quot;enterprise\&quot; is a sales inquiry about the Enterprise plan: recorded as the lead&#39;s program, offered the intro call as every inquiry is, and mailed no code. \&quot;investor\&quot; is an investor&#39;s subscription (investor.go): it is mailed no code, it is sent a SAFE, and Name, Amount and Address are required. | [optional] 
**Resume** | Pointer to **string** | Resume is an https address the applicant gave for their resume. Read only with Job; any other value is recorded as none. | [optional] 
**ResumeFile** | Pointer to **string** | ResumeFile is the key POST /v1/marketing/leads/resume handed out for the resume the applicant uploaded. Any other value is recorded as none. Read only with Job. | [optional] 
**Role** | Pointer to **string** | Role is the person&#39;s job title. Optional; clipped at 200 characters. | [optional] 
**Size** | Pointer to **string** | Size is the company&#39;s headcount band: 1-49, 50-249, 250-999, 1000-4999 or 5000+. Any other value is recorded as unstated rather than refused. | [optional] 
**Source** | Pointer to **string** | Source is the page the form was filed from, such as hanzo.ai/contact-sales. Optional; clipped at 200 characters. | [optional] 
**Title** | Pointer to **string** | Title is the investor&#39;s title at Entity, as the SAFE&#39;s signature block names it. At most 200 characters; read only with Entity. | [optional] 
**Upgrade** | Pointer to **string** | Upgrade is the handle of the organization an order upgrades. It is filed only for a person signed in, acting in that organization (X-Org-Id) as its owner or an admin, which the identity boundary attests from the token; any other caller naming one is 403. Empty is a new organization, or none. | [optional] 
**Visit** | Pointer to [**MarketingVisit**](MarketingVisit.md) | Visit is what the visitor&#39;s page knows about the ad click that brought them and the consent they gave, so the lead reaches the ad platforms matched to the click and under the visitor&#39;s own choice. Optional. | [optional] 

## Methods

### NewMarketingLeadIn

`func NewMarketingLeadIn() *MarketingLeadIn`

NewMarketingLeadIn instantiates a new MarketingLeadIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingLeadInWithDefaults

`func NewMarketingLeadInWithDefaults() *MarketingLeadIn`

NewMarketingLeadInWithDefaults instantiates a new MarketingLeadIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MarketingLeadIn) GetAddress() MarketingAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MarketingLeadIn) GetAddressOk() (*MarketingAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MarketingLeadIn) SetAddress(v MarketingAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MarketingLeadIn) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAmount

`func (o *MarketingLeadIn) GetAmount() interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MarketingLeadIn) GetAmountOk() (*interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MarketingLeadIn) SetAmount(v interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MarketingLeadIn) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *MarketingLeadIn) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *MarketingLeadIn) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetAnswers

`func (o *MarketingLeadIn) GetAnswers() []MarketingAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *MarketingLeadIn) GetAnswersOk() (*[]MarketingAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *MarketingLeadIn) SetAnswers(v []MarketingAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *MarketingLeadIn) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetCampaign

`func (o *MarketingLeadIn) GetCampaign() MarketingTags`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *MarketingLeadIn) GetCampaignOk() (*MarketingTags, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *MarketingLeadIn) SetCampaign(v MarketingTags)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *MarketingLeadIn) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetCompany

`func (o *MarketingLeadIn) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *MarketingLeadIn) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *MarketingLeadIn) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *MarketingLeadIn) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetConsent

`func (o *MarketingLeadIn) GetConsent() bool`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *MarketingLeadIn) GetConsentOk() (*bool, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *MarketingLeadIn) SetConsent(v bool)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *MarketingLeadIn) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetCountry

`func (o *MarketingLeadIn) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MarketingLeadIn) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MarketingLeadIn) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MarketingLeadIn) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCoupon

`func (o *MarketingLeadIn) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *MarketingLeadIn) GetCouponOk() (*string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *MarketingLeadIn) SetCoupon(v string)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *MarketingLeadIn) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### GetDepartment

`func (o *MarketingLeadIn) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *MarketingLeadIn) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *MarketingLeadIn) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *MarketingLeadIn) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetEmail

`func (o *MarketingLeadIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MarketingLeadIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MarketingLeadIn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MarketingLeadIn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEntity

`func (o *MarketingLeadIn) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *MarketingLeadIn) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *MarketingLeadIn) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *MarketingLeadIn) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetJob

`func (o *MarketingLeadIn) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *MarketingLeadIn) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *MarketingLeadIn) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *MarketingLeadIn) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetLinks

`func (o *MarketingLeadIn) GetLinks() []string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MarketingLeadIn) GetLinksOk() (*[]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MarketingLeadIn) SetLinks(v []string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MarketingLeadIn) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *MarketingLeadIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingLeadIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingLeadIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingLeadIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeed

`func (o *MarketingLeadIn) GetNeed() string`

GetNeed returns the Need field if non-nil, zero value otherwise.

### GetNeedOk

`func (o *MarketingLeadIn) GetNeedOk() (*string, bool)`

GetNeedOk returns a tuple with the Need field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeed

`func (o *MarketingLeadIn) SetNeed(v string)`

SetNeed sets Need field to given value.

### HasNeed

`func (o *MarketingLeadIn) HasNeed() bool`

HasNeed returns a boolean if a field has been set.

### GetPayment

`func (o *MarketingLeadIn) GetPayment() string`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *MarketingLeadIn) GetPaymentOk() (*string, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *MarketingLeadIn) SetPayment(v string)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *MarketingLeadIn) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPhone

`func (o *MarketingLeadIn) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MarketingLeadIn) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MarketingLeadIn) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MarketingLeadIn) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetProgram

`func (o *MarketingLeadIn) GetProgram() string`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *MarketingLeadIn) GetProgramOk() (*string, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *MarketingLeadIn) SetProgram(v string)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *MarketingLeadIn) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetResume

`func (o *MarketingLeadIn) GetResume() string`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *MarketingLeadIn) GetResumeOk() (*string, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *MarketingLeadIn) SetResume(v string)`

SetResume sets Resume field to given value.

### HasResume

`func (o *MarketingLeadIn) HasResume() bool`

HasResume returns a boolean if a field has been set.

### GetResumeFile

`func (o *MarketingLeadIn) GetResumeFile() string`

GetResumeFile returns the ResumeFile field if non-nil, zero value otherwise.

### GetResumeFileOk

`func (o *MarketingLeadIn) GetResumeFileOk() (*string, bool)`

GetResumeFileOk returns a tuple with the ResumeFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeFile

`func (o *MarketingLeadIn) SetResumeFile(v string)`

SetResumeFile sets ResumeFile field to given value.

### HasResumeFile

`func (o *MarketingLeadIn) HasResumeFile() bool`

HasResumeFile returns a boolean if a field has been set.

### GetRole

`func (o *MarketingLeadIn) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MarketingLeadIn) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MarketingLeadIn) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MarketingLeadIn) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSize

`func (o *MarketingLeadIn) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MarketingLeadIn) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MarketingLeadIn) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *MarketingLeadIn) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *MarketingLeadIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketingLeadIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketingLeadIn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketingLeadIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *MarketingLeadIn) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarketingLeadIn) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarketingLeadIn) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MarketingLeadIn) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpgrade

`func (o *MarketingLeadIn) GetUpgrade() string`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *MarketingLeadIn) GetUpgradeOk() (*string, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *MarketingLeadIn) SetUpgrade(v string)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *MarketingLeadIn) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetVisit

`func (o *MarketingLeadIn) GetVisit() MarketingVisit`

GetVisit returns the Visit field if non-nil, zero value otherwise.

### GetVisitOk

`func (o *MarketingLeadIn) GetVisitOk() (*MarketingVisit, bool)`

GetVisitOk returns a tuple with the Visit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisit

`func (o *MarketingLeadIn) SetVisit(v MarketingVisit)`

SetVisit sets Visit field to given value.

### HasVisit

`func (o *MarketingLeadIn) HasVisit() bool`

HasVisit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


