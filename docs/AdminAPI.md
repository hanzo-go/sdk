# \AdminAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAIMetrics**](AdminAPI.md#AdminAIMetrics) | **Get** /v1/admin/aimetrics | Is the fleet AI board: LLM generations over gen_ai spans (count, cost, avg/p95 latency, per-model), per-model usage from the live cloud_usage ledger, and the eval plane (traces, scores, score names, runs, and the average-score trend).
[**AdminAnalytics**](AdminAPI.md#AdminAnalytics) | **Get** /v1/admin/analytics | Is the SaaS product-analytics board over the caller&#39;s tenant window: active customers, new and churned, retention, MRR, ARPU, the usage trend and the top customers by spend — every number folded from the commerce ledger, not sampled.
[**AdminApplications**](AdminAPI.md#AdminApplications) | **Get** /v1/admin/applications | Lists IAM applications for one owner org, forwarded VERBATIM from IAM&#39;s get-applications.
[**AdminAudit**](AdminAPI.md#AdminAudit) | **Get** /v1/admin/audit | Reads one chain of cloud&#39;s tamper-evident audit trail, newest first, with that chain&#39;s live integrity attached so a listing can be badged as verified.
[**AdminAuditVerify**](AdminAPI.md#AdminAuditVerify) | **Get** /v1/admin/audit/verify | Walks EVERY hash chain this deployment keeps and reports each one: which chains were checked, how many records each holds, the head hash to pin externally against tail-truncation, and — when a chain is broken — the seq of the first bad record and why.
[**AdminBases**](AdminAPI.md#AdminBases) | **Get** /v1/admin/bases | Lists the tenant Base instances in the caller&#39;s window — a SuperAdmin sees every tenant&#39;s, anyone else only their own subtree&#39;s.
[**AdminCaps**](AdminAPI.md#AdminCaps) | **Get** /v1/admin/caps | Reads one org&#39;s usage caps: its spend alerts plus the derived period spend, over/warn state and reset time.
[**AdminCompute**](AdminAPI.md#AdminCompute) | **Get** /v1/admin/compute | Rolls the fleet&#39;s compute usage up to one row per (org, app, project, kind): how many distinct machines ran in the window, how many are still active, what they billed, and when each group last emitted an event.
[**AdminCordonNode**](AdminAPI.md#AdminCordonNode) | **Post** /v1/admin/infra/nodes/{id}/cordon | Marks one cluster node unschedulable — or schedulable again — and can drain the pods already on it.
[**AdminCreateCap**](AdminAPI.md#AdminCreateCap) | **Post** /v1/admin/caps | Sets a usage cap on one org — a platform override of a customer budget, written to the customer&#39;s own spend-alert rows.
[**AdminCustomer**](AdminAPI.md#AdminCustomer) | **Get** /v1/admin/customers/{org} | Answers GET /v1/admin/customers/:org.
[**AdminCustomers**](AdminAPI.md#AdminCustomers) | **Get** /v1/admin/customers | Lists every customer org at a glance, sorted by slug: owner email, plan, suspend status, member count, balance, month-to-date spend and MRR.
[**AdminDeleteCap**](AdminAPI.md#AdminDeleteCap) | **Delete** /v1/admin/caps/{id} | Removes one cap by id, lifting the ceiling entirely.
[**AdminDeleteDroplet**](AdminAPI.md#AdminDeleteDroplet) | **Delete** /v1/admin/infra/droplets/{id} | Destroys a droplet the board has just proven is NOT a DOKS node.
[**AdminDeleteLoadBalancer**](AdminAPI.md#AdminDeleteLoadBalancer) | **Delete** /v1/admin/infra/loadbalancers/{id} | Destroys a load balancer the board has just proven no live type&#x3D;LoadBalancer Service in any cluster targets.
[**AdminDeleteVolume**](AdminAPI.md#AdminDeleteVolume) | **Delete** /v1/admin/infra/volumes/{id} | Destroys a volume the board has just proven no PersistentVolume in any cluster references.
[**AdminDisablePlugin**](AdminAPI.md#AdminDisablePlugin) | **Post** /v1/admin/plugins/{name}/disable | Stops the plugin.
[**AdminEnablePlugin**](AdminAPI.md#AdminEnablePlugin) | **Post** /v1/admin/plugins/{name}/enable | Brings a stopped or disabled plugin back on the artifact it already has: the zero Plugin names no new artifact, so Reload reuses the loaded spec and clears the disabled flag.
[**AdminFinance**](AdminAPI.md#AdminFinance) | **Get** /v1/admin/finance | Answers GET /v1/admin/finance.
[**AdminFinanceBackfill**](AdminAPI.md#AdminFinanceBackfill) | **Post** /v1/admin/finance/backfill | Carries ONE org&#39;s current commerce prepaid balance into the native finance wallet — the one-time cutover between the two ledgers.
[**AdminFlags**](AdminAPI.md#AdminFlags) | **Get** /v1/admin/flags | Reads the platform control-plane board: every runtime launch/release switch (waitlist, public signup, subsystem activation, gateway limits, network ids) with its LIVE value and where that value came from — a stored definition or the compiled-in default.
[**AdminGrantCredit**](AdminAPI.md#AdminGrantCredit) | **Post** /v1/admin/customers/{org}/credit | Issues a staff credit grant to the org named in the path — a comp, refund or promo — through the ONE credit-write path core.ApplyGrant, which validates the amount against the per-grant cap, checks the org exists, moves the money and records the tamper-evident audit row.
[**AdminGrants**](AdminAPI.md#AdminGrants) | **Get** /v1/admin/grants | Reads the credit-grant ledger across ALL orgs, newest first — who granted what to whom, when, and from which money bucket.
[**AdminInfra**](AdminAPI.md#AdminInfra) | **Get** /v1/admin/infra | Serves the whole DigitalOcean infrastructure board: droplets, volumes, DOKS clusters and load balancers, each cross-referenced against every cluster&#39;s live Kubernetes state so the board can say what is safe to destroy and what is not.
[**AdminInvoices**](AdminAPI.md#AdminInvoices) | **Get** /v1/admin/invoices | Answers GET /v1/admin/invoices.
[**AdminIssueGrant**](AdminAPI.md#AdminIssueGrant) | **Post** /v1/admin/grants | Issues a credit grant to any org from the operator Grants view, with the target named in the body.
[**AdminMe**](AdminAPI.md#AdminMe) | **Get** /v1/admin/me | Answers with the validated operator identity — who the console is signed in as, which tier they are, and how wide their tenant window is.
[**AdminMetrics**](AdminAPI.md#AdminMetrics) | **Get** /v1/admin/metrics | Answers GET /v1/admin/metrics by aggregating commerce.events directly (fleet-wide, no per-org fan-out).
[**AdminMoney**](AdminAPI.md#AdminMoney) | **Get** /v1/admin/money | moneyBoardHandler answers GET /v1/admin/money.
[**AdminO11y**](AdminAPI.md#AdminO11y) | **Get** /v1/admin/o11y | Is the fleet-wide observability board: LLM usage (requests, tokens, cost, errors, top orgs, top models), trace RED metrics (count, p50/p95/p99 latency in ms, error rate, top services), fleet log volume, and the O11yAI generation rollup — all aggregated across EVERY tenant, with no org filter applied.
[**AdminOrgs**](AdminAPI.md#AdminOrgs) | **Get** /v1/admin/orgs | Lists the tenant directory one row per org, sorted by slug: member count and the org&#39;s month-to-date spend and credit balance, read live from IAM and commerce.
[**AdminOverview**](AdminAPI.md#AdminOverview) | **Get** /v1/admin/overview | Is the Platform Overview tiles: how many orgs and users are in the caller&#39;s tenant window, the fleet workload counts, and month-to-date spend and credits.
[**AdminPlugins**](AdminAPI.md#AdminPlugins) | **Get** /v1/admin/plugins | Reports what each host is actually running: every loaded plugin with its version, pid, uptime, reload and restart counts, and its measured CPU, RSS, thread and fd cost — read from the kernel, which is only answerable at all because a plugin is a process.
[**AdminProducts**](AdminAPI.md#AdminProducts) | **Get** /v1/admin/products | Lists the fleet workload registry: every operator App CR across the platform namespaces with its declared vs running image tag, reconciled health/phase and drift verdict.
[**AdminPromo**](AdminAPI.md#AdminPromo) | **Get** /v1/admin/promos | Reads the current platform plan promo — the singleton discount offer, e.g.
[**AdminProvidersCredit**](AdminAPI.md#AdminProvidersCredit) | **Get** /v1/admin/providers/credit | Serves GET /v1/admin/providers/credit — the per-provider upstream credit ledger.
[**AdminReactivateCustomer**](AdminAPI.md#AdminReactivateCustomer) | **Post** /v1/admin/customers/{org}/reactivate | Restores access for every member of the org, undoing a suspend.
[**AdminReloadPlugin**](AdminAPI.md#AdminReloadPlugin) | **Post** /v1/admin/plugins/{name}/reload | Swaps a plugin for another build without dropping a request.
[**AdminResizeDroplet**](AdminAPI.md#AdminResizeDroplet) | **Post** /v1/admin/infra/droplets/{id}/resize | Changes a droplet&#39;s plan.
[**AdminResizeVolume**](AdminAPI.md#AdminResizeVolume) | **Post** /v1/admin/infra/volumes/{id}/resize | Grows a volume.
[**AdminRevenue**](AdminAPI.md#AdminRevenue) | **Get** /v1/admin/revenue | Is the fleet money board: total prepaid balances held, total realized spend, MRR, ARPU, a per-customer table sorted highest-revenue first, and a real 30-day spend trend from the usage ledger.
[**AdminRoles**](AdminAPI.md#AdminRoles) | **Get** /v1/admin/roles | Lists IAM roles for one owner org, forwarded VERBATIM from IAM&#39;s get-roles.
[**AdminScaleNodePool**](AdminAPI.md#AdminScaleNodePool) | **Post** /v1/admin/infra/clusters/{id}/nodepools/{pool}/scale | Sets a node pool&#39;s node count — the ONE correct way to change how many nodes a DOKS cluster has.
[**AdminServices**](AdminAPI.md#AdminServices) | **Get** /v1/admin/services | Reads the launch board: every hosted service in the registry with its LIVE waitlist mode, evaluated through the flag engine.
[**AdminSetFlag**](AdminAPI.md#AdminSetFlag) | **Put** /v1/admin/flags/{key} | Stores or overwrites ONE platform switch&#39;s definition and answers with the whole board as it now stands.
[**AdminSetPromo**](AdminAPI.md#AdminSetPromo) | **Put** /v1/admin/promos | Upserts the platform plan promo — the ONE place the offer is configured.
[**AdminSetServiceMode**](AdminAPI.md#AdminSetServiceMode) | **Post** /v1/admin/services/{service}/mode | Flips ONE service&#39;s waitlist switch — the launch lever.
[**AdminSnapshotVolume**](AdminAPI.md#AdminSnapshotVolume) | **Post** /v1/admin/infra/volumes/{id}/snapshot | Takes a point-in-time snapshot of one volume — the undo a delete relies on, available on its own so an operator can take one before any risky change.
[**AdminSubscriptions**](AdminAPI.md#AdminSubscriptions) | **Get** /v1/admin/subscriptions | Answers GET /v1/admin/subscriptions.
[**AdminSubsystems**](AdminAPI.md#AdminSubsystems) | **Get** /v1/admin/subsystems | subsystems answers GET /v1/admin/subsystems.
[**AdminSuspendCustomer**](AdminAPI.md#AdminSuspendCustomer) | **Post** /v1/admin/customers/{org}/suspend | Cuts off every member of the org: IAM refuses a forbidden user at login AND at token issuance, so a suspended customer can neither sign in nor mint a fresh token.
[**AdminSync**](AdminAPI.md#AdminSync) | **Post** /v1/admin/sync | Answers the operator&#39;s \&quot;Sync now\&quot; button.
[**AdminUpdateCap**](AdminAPI.md#AdminUpdateCap) | **Patch** /v1/admin/caps/{id} | Edits one cap by id — raise or lower the ceiling, flip enforcement.
[**AdminUpsertService**](AdminAPI.md#AdminUpsertService) | **Post** /v1/admin/services | Onboards a hosted service, or edits one, so a new host comes under the launch gate WITHOUT a redeploy.
[**AdminUsage**](AdminAPI.md#AdminUsage) | **Get** /v1/admin/usage | Returns the trailing 30 days of AI usage: one org&#39;s when org names one, else the whole fleet&#39;s — the spend, the tokens and the requests, the daily curve behind them, and the split by model.
[**AdminUsageFunding**](AdminAPI.md#AdminUsageFunding) | **Get** /v1/admin/usage/funding | Splits our upstream AI usage by how it was FUNDED: one row per (provider, model) over the window, tagged credit (provider grant still remaining), paid (grant exhausted) or paid_only (no grant at all).
[**AdminUsers**](AdminAPI.md#AdminUsers) | **Get** /v1/admin/users | Lists the user directory across the caller&#39;s tenant window, one page at a time.
[**AdminVolumes**](AdminAPI.md#AdminVolumes) | **Get** /v1/admin/volumes | Returns the realtime block-storage board: the DigitalOcean volume fleet (count, capacity, monthly list cost, per-volume region and attachment) plus the analytics datastore&#39;s OWN fill, read from its system.disks.
[**AdminWaitlist**](AdminAPI.md#AdminWaitlist) | **Get** /v1/admin/waitlist | Reads one waitlist&#39;s leaderboard from the Hanzo waitlist engine — position, points and referral standing per entry — proxied server-authed with the engine secret, never a client credential.
[**AdminWaitlistBoost**](AdminAPI.md#AdminWaitlistBoost) | **Post** /v1/admin/waitlist/boost | Grants a user waitlist points, moving them up toward the access cutoff.
[**GetAdminAffiliates**](AdminAPI.md#GetAdminAffiliates) | **Get** /v1/admin/affiliates | Lists every affiliate across the fleet with its ORG exposed, plus a fleet summary of lifetime accrued, still-pending and paid commission in integer cents.
[**GetAdminAuthors**](AdminAPI.md#GetAdminAuthors) | **Get** /v1/admin/authors | Returns the platform&#39;s whole author program — every org&#39;s author record, not the caller&#39;s — with each one&#39;s repository and deploy counts and a fleet roll-up of the money accrued, pending and paid.
[**GetAdminAuthorsByIdBasis**](AdminAPI.md#GetAdminAuthorsByIdBasis) | **Get** /v1/admin/authors/{id}/basis | Returns the audit trail behind ONE author&#39;s royalty — the same payload the author reads at /v1/authors/basis, from the same builder, so support sees exactly what the author sees rather than a parallel view free to drift.
[**GetAdminCatalog**](AdminAPI.md#GetAdminCatalog) | **Get** /v1/admin/catalog | Returns the full model and provider catalog annotated with each entry&#39;s enablement state, for the operator console.
[**GetAdminEnablement**](AdminAPI.md#GetAdminEnablement) | **Get** /v1/admin/enablement | Returns every item an operator has set an enablement state on — its global state (off, beta or ga) and the orgs granted its beta.
[**GetAdminReferrals**](AdminAPI.md#GetAdminReferrals) | **Get** /v1/admin/referrals | Answers the referral board: the top referrers by lifetime commission, the funnel conversion rate (referred orgs that have actually produced commission, over all referred orgs), and the accrual LIABILITY the platform owes, broken out by upline level.
[**GetAdminReferralsBonuses**](AdminAPI.md#GetAdminReferralsBonuses) | **Get** /v1/admin/referrals/bonuses | Returns every referral edge in the directory with a fleet summary.
[**GetAdminTreasury**](AdminAPI.md#GetAdminTreasury) | **Get** /v1/admin/treasury | Returns the whole treasury board for a SuperAdmin: the reserve fund report, the recent double-entry journal, and the Hanzo L1 anchor status of the ledger root.
[**PatchAdminCatalogModelsByWildcard1**](AdminAPI.md#PatchAdminCatalogModelsByWildcard1) | **Patch** /v1/admin/catalog/models/{wildcard1} | Turn one model off, into beta for named orgs, or generally available
[**PatchAdminCatalogProvidersByName**](AdminAPI.md#PatchAdminCatalogProvidersByName) | **Patch** /v1/admin/catalog/providers/{name} | Sets one provider&#39;s availability overlay.
[**PostAdminAffiliatesByIdApprove**](AdminAPI.md#PostAdminAffiliatesByIdApprove) | **Post** /v1/admin/affiliates/{id}/approve | Approves an affiliate and MINTS its referral code — the moment the partner has a working share link and starts accruing.
[**PostAdminAffiliatesByIdPayout**](AdminAPI.md#PostAdminAffiliatesByIdPayout) | **Post** /v1/admin/affiliates/{id}/payout | Pays out accrued commission and answers the payout row with the affiliate&#39;s updated balances.
[**PostAdminAffiliatesByIdRate**](AdminAPI.md#PostAdminAffiliatesByIdRate) | **Post** /v1/admin/affiliates/{id}/rate | Sets one affiliate&#39;s DIRECT commission rate, in basis points of Hanzo&#39;s margin.
[**PostAdminAffiliatesByIdSuspend**](AdminAPI.md#PostAdminAffiliatesByIdSuspend) | **Post** /v1/admin/affiliates/{id}/suspend | Suspends an affiliate: it stops accruing on the next sweep, and its code stops resolving for new attributions.
[**PostAdminAffiliatesSweep**](AdminAPI.md#PostAdminAffiliatesSweep) | **Post** /v1/admin/affiliates/sweep | Runs the accrual: for each referred org it reads that org&#39;s metered spend for the current period and accrues commission to every affiliate up its referral chain, then answers how many sources were swept and how many NEW accruals landed.
[**PostAdminAuthorsByIdApprove**](AdminAPI.md#PostAdminAuthorsByIdApprove) | **Post** /v1/admin/authors/{id}/approve | Admits one author to EARNING, optionally on a negotiated royalty share.
[**PostAdminAuthorsByIdPayout**](AdminAPI.md#PostAdminAuthorsByIdPayout) | **Post** /v1/admin/authors/{id}/payout | Records a payout of accrued royalty and settles it.
[**PostAdminAuthorsByIdSuspend**](AdminAPI.md#PostAdminAuthorsByIdSuspend) | **Post** /v1/admin/authors/{id}/suspend | Stops one author earning.
[**PostAdminAuthorsSweep**](AdminAPI.md#PostAdminAuthorsSweep) | **Post** /v1/admin/authors/sweep | Runs the accrual sweep across every approved author: for each of their deploying orgs it computes this period&#39;s royalty from that org&#39;s metered spend and latches it at most once per period.
[**PostAdminReferralsSweep**](AdminAPI.md#PostAdminReferralsSweep) | **Post** /v1/admin/referrals/sweep | Qualify-checks every pending referral and advances the ones that now qualify.
[**PostAdminTreasuryAnchor**](AdminAPI.md#PostAdminTreasuryAnchor) | **Post** /v1/admin/treasury/anchor | Commits the current ledger root to Hanzo L1, making the books tamper-evident on chain, and returns the anchoring status.
[**PostAdminTreasuryPolicy**](AdminAPI.md#PostAdminTreasuryPolicy) | **Post** /v1/admin/treasury/policy | Sets the revenue-share basis points a sweep accrues into the reserve fund and returns the stored policy.
[**PostAdminTreasurySeed**](AdminAPI.md#PostAdminTreasurySeed) | **Post** /v1/admin/treasury/seed | Injects bootstrap capital into the reserve fund so backed payouts can begin before the first revenue-share sweep, and returns the journal entry it wrote.
[**PostAdminTreasurySweep**](AdminAPI.md#PostAdminTreasurySweep) | **Post** /v1/admin/treasury/sweep | Posts the revenue-share accrual for one period — revenue into the reserve fund, at the current policy&#39;s basis points — and returns what it moved.
[**PutAdminEnablement**](AdminAPI.md#PutAdminEnablement) | **Put** /v1/admin/enablement | Sets one item&#39;s global enablement state — off, beta or ga — and optionally replaces the list of orgs granted its beta.
[**PutAdminTreasuryAnchorSigner**](AdminAPI.md#PutAdminTreasuryAnchorSigner) | **Put** /v1/admin/treasury/anchor/signer | Installs the reserve&#39;s threshold MPC wallet as the signer for on-chain anchors, and returns its EVM address so an operator can fund it for gas.



## AdminAIMetrics

> AimetricsOut AdminAIMetrics(ctx).Range_(range_).Execute()

Is the fleet AI board: LLM generations over gen_ai spans (count, cost, avg/p95 latency, per-model), per-model usage from the live cloud_usage ledger, and the eval plane (traces, scores, score names, runs, and the average-score trend).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	range_ := "7d" // string | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board's own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminAIMetrics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminAIMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAIMetrics`: AimetricsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminAIMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAIMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board&#39;s own default. | 

### Return type

[**AimetricsOut**](AimetricsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAnalytics

> AnalyticsOut AdminAnalytics(ctx).Range_(range_).Execute()

Is the SaaS product-analytics board over the caller's tenant window: active customers, new and churned, retention, MRR, ARPU, the usage trend and the top customers by spend — every number folded from the commerce ledger, not sampled.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	range_ := "30d" // string | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board's own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminAnalytics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAnalytics`: AnalyticsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board&#39;s own default. | 

### Return type

[**AnalyticsOut**](AnalyticsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminApplications

> IamRowsOut AdminApplications(ctx).Owner(owner).P(p).PageSize(pageSize).Execute()

Lists IAM applications for one owner org, forwarded VERBATIM from IAM's get-applications.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "admin" // string | Owner is the org whose rows to read. Defaults to the admin org, which owns the platform's roles and applications. (optional)
	p := "1" // string | Page is the 1-based page number. Forwarded only when set — IAM applies its own default otherwise. (optional)
	pageSize := "50" // string | PageSize is rows per page. Forwarded only when set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminApplications(context.Background()).Owner(owner).P(p).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminApplications`: IamRowsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Owner is the org whose rows to read. Defaults to the admin org, which owns the platform&#39;s roles and applications. | 
 **p** | **string** | Page is the 1-based page number. Forwarded only when set — IAM applies its own default otherwise. | 
 **pageSize** | **string** | PageSize is rows per page. Forwarded only when set. | 

### Return type

[**IamRowsOut**](IamRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAudit

> RecordsOut AdminAudit(ctx).Org(org).Sub(sub).Action(action).Resource(resource).ResourceId(resourceId).Result(result).Since(since).Until(until).PageSize(pageSize).P(p).Execute()

Reads one chain of cloud's tamper-evident audit trail, newest first, with that chain's live integrity attached so a listing can be badged as verified.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	org := "acme" // string | Org restricts the trail to one tenant. (optional)
	sub := "sub_example" // string | Sub restricts it to one actor (the validated subject that made the request). (optional)
	action := "admin.waitlist.grant" // string | Action restricts it to one action name, e.g. \"admin.waitlist.grant\". (optional)
	resource := "resource_example" // string | Resource restricts it to one resource kind, e.g. \"credit-grant\". (optional)
	resourceId := "resourceId_example" // string | ResourceID restricts it to one resource instance. (optional)
	result := "result_example" // string | Result restricts it to \"success\" or \"error\". (optional)
	since := "2026-07-01T00:00:00Z" // string | Since is the inclusive lower time bound, RFC3339. An unparseable value is ignored rather than refused — one malformed filter must not hide the trail. (optional)
	until := "until_example" // string | Until is the upper time bound, RFC3339, with the same tolerance. (optional)
	pageSize := "50" // string | PageSize is rows per page, default 100. (optional)
	p := "p_example" // string | Page is the 1-based page number, driving the offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminAudit(context.Background()).Org(org).Sub(sub).Action(action).Resource(resource).ResourceId(resourceId).Result(result).Since(since).Until(until).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAudit`: RecordsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org restricts the trail to one tenant. | 
 **sub** | **string** | Sub restricts it to one actor (the validated subject that made the request). | 
 **action** | **string** | Action restricts it to one action name, e.g. \&quot;admin.waitlist.grant\&quot;. | 
 **resource** | **string** | Resource restricts it to one resource kind, e.g. \&quot;credit-grant\&quot;. | 
 **resourceId** | **string** | ResourceID restricts it to one resource instance. | 
 **result** | **string** | Result restricts it to \&quot;success\&quot; or \&quot;error\&quot;. | 
 **since** | **string** | Since is the inclusive lower time bound, RFC3339. An unparseable value is ignored rather than refused — one malformed filter must not hide the trail. | 
 **until** | **string** | Until is the upper time bound, RFC3339, with the same tolerance. | 
 **pageSize** | **string** | PageSize is rows per page, default 100. | 
 **p** | **string** | Page is the 1-based page number, driving the offset. | 

### Return type

[**RecordsOut**](RecordsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAuditVerify

> VerifyOut AdminAuditVerify(ctx).Execute()

Walks EVERY hash chain this deployment keeps and reports each one: which chains were checked, how many records each holds, the head hash to pin externally against tail-truncation, and — when a chain is broken — the seq of the first bad record and why.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminAuditVerify(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminAuditVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAuditVerify`: VerifyOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminAuditVerify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAuditVerifyRequest struct via the builder pattern


### Return type

[**VerifyOut**](VerifyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminBases

> BasesOut AdminBases(ctx).Execute()

Lists the tenant Base instances in the caller's window — a SuperAdmin sees every tenant's, anyone else only their own subtree's.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminBases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminBases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminBases`: BasesOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminBases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminBasesRequest struct via the builder pattern


### Return type

[**BasesOut**](BasesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminCaps

> RawOut AdminCaps(ctx).Org(org).Id(id).Execute()

Reads one org's usage caps: its spend alerts plus the derived period spend, over/warn state and reset time.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	org := "acme" // string | Org is the tenant to act on. Required for a SuperAdmin — they must name their target; ignored for a white-label admin, who always acts on their own org. (optional)
	id := "id_example" // string | ID is the cap to edit or remove, from the path. Unused by the list and create ops. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminCaps(context.Background()).Org(org).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminCaps`: RawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminCaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org is the tenant to act on. Required for a SuperAdmin — they must name their target; ignored for a white-label admin, who always acts on their own org. | 
 **id** | **string** | ID is the cap to edit or remove, from the path. Unused by the list and create ops. | 

### Return type

[**RawOut**](RawOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminCompute

> ComputeOut AdminCompute(ctx).Kind(kind).Org(org).Range_(range_).Execute()

Rolls the fleet's compute usage up to one row per (org, app, project, kind): how many distinct machines ran in the window, how many are still active, what they billed, and when each group last emitted an event.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	kind := "bot" // string | Kind narrows to one workload class (bot | machine | cluster | nodepool | container | function | …). An OPEN spectrum matched as a plain string, lowercased to the warehouse's convention; empty means every kind. (optional)
	org := "acme" // string | Org narrows to one tenant. Empty means every tenant — this board is cross-tenant by nature. (optional)
	range_ := "7d" // string | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as 30d. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminCompute(context.Background()).Kind(kind).Org(org).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminCompute`: ComputeOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminCompute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminComputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Kind narrows to one workload class (bot | machine | cluster | nodepool | container | function | …). An OPEN spectrum matched as a plain string, lowercased to the warehouse&#39;s convention; empty means every kind. | 
 **org** | **string** | Org narrows to one tenant. Empty means every tenant — this board is cross-tenant by nature. | 
 **range_** | **string** | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as 30d. | 

### Return type

[**ComputeOut**](ComputeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminCordonNode

> MutationOut AdminCordonNode(ctx, id).CordonIn(cordonIn).Execute()

Marks one cluster node unschedulable — or schedulable again — and can drain the pods already on it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the node's droplet id, from the path.
	cordonIn := *openapiclient.NewCordonIn() // CordonIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminCordonNode(context.Background(), id).CordonIn(cordonIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCordonNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminCordonNode`: MutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminCordonNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the node&#39;s droplet id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminCordonNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cordonIn** | [**CordonIn**](CordonIn.md) |  | 

### Return type

[**MutationOut**](MutationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminCreateCap

> RawOut AdminCreateCap(ctx).CapIn(capIn).Execute()

Sets a usage cap on one org — a platform override of a customer budget, written to the customer's own spend-alert rows.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	capIn := *openapiclient.NewCapIn() // CapIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminCreateCap(context.Background()).CapIn(capIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCreateCap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminCreateCap`: RawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminCreateCap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateCapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **capIn** | [**CapIn**](CapIn.md) |  | 

### Return type

[**RawOut**](RawOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminCustomer

> CustomerDetailOut AdminCustomer(ctx, org).Execute()

Answers GET /v1/admin/customers/:org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	org := "org_example" // string | Org is the tenant slug from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminCustomer(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminCustomer`: CustomerDetailOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the tenant slug from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerDetailOut**](CustomerDetailOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminCustomers

> CustomersOut AdminCustomers(ctx).Execute()

Lists every customer org at a glance, sorted by slug: owner email, plan, suspend status, member count, balance, month-to-date spend and MRR.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminCustomers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminCustomers`: CustomersOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminCustomers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminCustomersRequest struct via the builder pattern


### Return type

[**CustomersOut**](CustomersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminDeleteCap

> RawOut AdminDeleteCap(ctx, id).Org(org).Execute()

Removes one cap by id, lifting the ceiling entirely.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "cap_1" // string | ID is the cap to edit or remove, from the path. Unused by the list and create ops.
	org := "acme" // string | Org is the tenant to act on. Required for a SuperAdmin — they must name their target; ignored for a white-label admin, who always acts on their own org. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminDeleteCap(context.Background(), id).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminDeleteCap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminDeleteCap`: RawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminDeleteCap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cap to edit or remove, from the path. Unused by the list and create ops. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteCapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **org** | **string** | Org is the tenant to act on. Required for a SuperAdmin — they must name their target; ignored for a white-label admin, who always acts on their own org. | 

### Return type

[**RawOut**](RawOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminDeleteDroplet

> MutationOut AdminDeleteDroplet(ctx, id).Size(size).Disk(disk).Execute()

Destroys a droplet the board has just proven is NOT a DOKS node.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the DO droplet id, from the path. Numeric.
	size := "size_example" // string | Size is the target DigitalOcean size slug on resize, e.g. \"s-4vcpu-8gb\". (optional)
	disk := true // bool | Disk requests a PERMANENT resize that grows the disk. DO can never resize such a droplet down again, so it defaults false — a CPU/RAM-only change, reversible. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminDeleteDroplet(context.Background(), id).Size(size).Disk(disk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminDeleteDroplet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminDeleteDroplet`: MutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminDeleteDroplet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO droplet id, from the path. Numeric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteDropletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **string** | Size is the target DigitalOcean size slug on resize, e.g. \&quot;s-4vcpu-8gb\&quot;. | 
 **disk** | **bool** | Disk requests a PERMANENT resize that grows the disk. DO can never resize such a droplet down again, so it defaults false — a CPU/RAM-only change, reversible. | 

### Return type

[**MutationOut**](MutationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminDeleteLoadBalancer

> MutationOut AdminDeleteLoadBalancer(ctx, id).Execute()

Destroys a load balancer the board has just proven no live type=LoadBalancer Service in any cluster targets.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the DO load balancer id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminDeleteLoadBalancer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminDeleteLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminDeleteLoadBalancer`: MutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminDeleteLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO load balancer id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MutationOut**](MutationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminDeleteVolume

> MutationOut AdminDeleteVolume(ctx, id).Snapshot(snapshot).Name(name).SizeGiB(sizeGiB).Execute()

Destroys a volume the board has just proven no PersistentVolume in any cluster references.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the DO volume id, from the path.
	snapshot := "snapshot_example" // string | Snapshot is the snapshot-first switch on DELETE. Anything other than the literal \"false\" snapshots before destroying — the snapshot IS the undo, so waiving it is deliberate and explicit. (optional)
	name := "name_example" // string | Name is the snapshot name on the snapshot action. Blank gets a deterministic \"<volume>-predelete-<unix>\" so the undo is findable in the DO console. (optional)
	sizeGiB := int32(56) // int32 | SizeGiB is the target size on the resize action. A volume only ever grows — ExpandTo is the verdict that refuses a shrink, so this is not validated here. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminDeleteVolume(context.Background(), id).Snapshot(snapshot).Name(name).SizeGiB(sizeGiB).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminDeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminDeleteVolume`: MutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminDeleteVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO volume id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshot** | **string** | Snapshot is the snapshot-first switch on DELETE. Anything other than the literal \&quot;false\&quot; snapshots before destroying — the snapshot IS the undo, so waiving it is deliberate and explicit. | 
 **name** | **string** | Name is the snapshot name on the snapshot action. Blank gets a deterministic \&quot;&lt;volume&gt;-predelete-&lt;unix&gt;\&quot; so the undo is findable in the DO console. | 
 **sizeGiB** | **int32** | SizeGiB is the target size on the resize action. A volume only ever grows — ExpandTo is the verdict that refuses a shrink, so this is not validated here. | 

### Return type

[**MutationOut**](MutationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminDisablePlugin

> ActionOut AdminDisablePlugin(ctx, name).NameIn(nameIn).Execute()

Stops the plugin.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "billing" // string | Name is the app, from the path.
	nameIn := *openapiclient.NewNameIn() // NameIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminDisablePlugin(context.Background(), name).NameIn(nameIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminDisablePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminDisablePlugin`: ActionOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminDisablePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the app, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDisablePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nameIn** | [**NameIn**](NameIn.md) |  | 

### Return type

[**ActionOut**](ActionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminEnablePlugin

> ActionOut AdminEnablePlugin(ctx, name).NameIn(nameIn).Execute()

Brings a stopped or disabled plugin back on the artifact it already has: the zero Plugin names no new artifact, so Reload reuses the loaded spec and clears the disabled flag.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "billing" // string | Name is the app, from the path.
	nameIn := *openapiclient.NewNameIn() // NameIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminEnablePlugin(context.Background(), name).NameIn(nameIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminEnablePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminEnablePlugin`: ActionOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminEnablePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the app, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminEnablePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nameIn** | [**NameIn**](NameIn.md) |  | 

### Return type

[**ActionOut**](ActionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminFinance

> FinanceOut AdminFinance(ctx).Execute()

Answers GET /v1/admin/finance.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminFinance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminFinance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminFinance`: FinanceOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminFinance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFinanceRequest struct via the builder pattern


### Return type

[**FinanceOut**](FinanceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminFinanceBackfill

> BackfillOut AdminFinanceBackfill(ctx).BackfillIn(backfillIn).Execute()

Carries ONE org's current commerce prepaid balance into the native finance wallet — the one-time cutover between the two ledgers.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	backfillIn := *openapiclient.NewBackfillIn() // BackfillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminFinanceBackfill(context.Background()).BackfillIn(backfillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminFinanceBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminFinanceBackfill`: BackfillOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminFinanceBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminFinanceBackfillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backfillIn** | [**BackfillIn**](BackfillIn.md) |  | 

### Return type

[**BackfillOut**](BackfillOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminFlags

> FlagsOut AdminFlags(ctx).Execute()

Reads the platform control-plane board: every runtime launch/release switch (waitlist, public signup, subsystem activation, gateway limits, network ids) with its LIVE value and where that value came from — a stored definition or the compiled-in default.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminFlags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminFlags`: FlagsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminFlags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFlagsRequest struct via the builder pattern


### Return type

[**FlagsOut**](FlagsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGrantCredit

> GrantOut AdminGrantCredit(ctx, org).GrantIn(grantIn).Execute()

Issues a staff credit grant to the org named in the path — a comp, refund or promo — through the ONE credit-write path core.ApplyGrant, which validates the amount against the per-grant cap, checks the org exists, moves the money and records the tamper-evident audit row.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	org := "org_example" // string | Org is the tenant to credit. Required.
	grantIn := *openapiclient.NewGrantIn() // GrantIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminGrantCredit(context.Background(), org).GrantIn(grantIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminGrantCredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGrantCredit`: GrantOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminGrantCredit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the tenant to credit. Required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGrantCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **grantIn** | [**GrantIn**](GrantIn.md) |  | 

### Return type

[**GrantOut**](GrantOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGrants

> GrantsOut AdminGrants(ctx).Org(org).Result(result).Limit(limit).Execute()

Reads the credit-grant ledger across ALL orgs, newest first — who granted what to whom, when, and from which money bucket.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	org := "org_example" // string | Org filters by the ACTOR's org (the staff org that issued the grant), which is rarely what a reader wants — the target org is a row field, not a filter. (optional)
	result := "success" // string | Result filters by outcome: \"success\" or \"error\". Empty returns both, which is the point of this view — a refused grant is as interesting as a granted one. (optional)
	limit := "50" // string | Limit caps the rows returned. Default 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminGrants(context.Background()).Org(org).Result(result).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGrants`: GrantsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org filters by the ACTOR&#39;s org (the staff org that issued the grant), which is rarely what a reader wants — the target org is a row field, not a filter. | 
 **result** | **string** | Result filters by outcome: \&quot;success\&quot; or \&quot;error\&quot;. Empty returns both, which is the point of this view — a refused grant is as interesting as a granted one. | 
 **limit** | **string** | Limit caps the rows returned. Default 200. | 

### Return type

[**GrantsOut**](GrantsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminInfra

> ReadOut AdminInfra(ctx).Refresh(refresh).Execute()

Serves the whole DigitalOcean infrastructure board: droplets, volumes, DOKS clusters and load balancers, each cross-referenced against every cluster's live Kubernetes state so the board can say what is safe to destroy and what is not.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	refresh := "1" // string | Refresh, when present, forces a full re-scan instead of serving the cached snapshot. Every MUTATION re-scans regardless — this is only for the reader. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminInfra(context.Background()).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminInfra``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminInfra`: ReadOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminInfra`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminInfraRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **string** | Refresh, when present, forces a full re-scan instead of serving the cached snapshot. Every MUTATION re-scans regardless — this is only for the reader. | 

### Return type

[**ReadOut**](ReadOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminInvoices

> InvoicesOut AdminInvoices(ctx).Status(status).Org(org).Limit(limit).Execute()

Answers GET /v1/admin/invoices.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	status := "status_example" // string | Status filters on the invoice's LATEST lifecycle status (paid, open, void, …), matched case-insensitively. (optional)
	org := "org_example" // string | Org filters to one tenant, matched exactly. (optional)
	limit := "limit_example" // string | Limit caps the rows returned. total still reports the full match count. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminInvoices(context.Background()).Status(status).Org(org).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminInvoices`: InvoicesOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status filters on the invoice&#39;s LATEST lifecycle status (paid, open, void, …), matched case-insensitively. | 
 **org** | **string** | Org filters to one tenant, matched exactly. | 
 **limit** | **string** | Limit caps the rows returned. total still reports the full match count. | 

### Return type

[**InvoicesOut**](InvoicesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminIssueGrant

> GrantOut AdminIssueGrant(ctx).GrantIn(grantIn).Execute()

Issues a credit grant to any org from the operator Grants view, with the target named in the body.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	grantIn := *openapiclient.NewGrantIn() // GrantIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminIssueGrant(context.Background()).GrantIn(grantIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminIssueGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminIssueGrant`: GrantOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminIssueGrant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminIssueGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantIn** | [**GrantIn**](GrantIn.md) |  | 

### Return type

[**GrantOut**](GrantOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminMe

> MeOut AdminMe(ctx).Execute()

Answers with the validated operator identity — who the console is signed in as, which tier they are, and how wide their tenant window is.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminMe`: MeOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminMeRequest struct via the builder pattern


### Return type

[**MeOut**](MeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminMetrics

> MetricsOut AdminMetrics(ctx).Window(window).Limit(limit).Execute()

Answers GET /v1/admin/metrics by aggregating commerce.events directly (fleet-wide, no per-org fan-out).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	window := "window_example" // string | Window is the movement window the new/churned MRR and the recent feed are measured over. Anything unrecognised falls back to the board default. (optional)
	limit := "limit_example" // string | Limit caps the top-customers table. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminMetrics(context.Background()).Window(window).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminMetrics`: MetricsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **window** | **string** | Window is the movement window the new/churned MRR and the recent feed are measured over. Anything unrecognised falls back to the board default. | 
 **limit** | **string** | Limit caps the top-customers table. | 

### Return type

[**MetricsOut**](MetricsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminMoney

> MoneyOut AdminMoney(ctx).Execute()

moneyBoardHandler answers GET /v1/admin/money.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminMoney(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminMoney``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminMoney`: MoneyOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminMoney`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminMoneyRequest struct via the builder pattern


### Return type

[**MoneyOut**](MoneyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminO11y

> O11yOut AdminO11y(ctx).Range_(range_).Execute()

Is the fleet-wide observability board: LLM usage (requests, tokens, cost, errors, top orgs, top models), trace RED metrics (count, p50/p95/p99 latency in ms, error rate, top services), fleet log volume, and the O11yAI generation rollup — all aggregated across EVERY tenant, with no org filter applied.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	range_ := "7d" // string | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board's own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminO11y(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminO11y``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminO11y`: O11yOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminO11y`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminO11yRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board&#39;s own default. | 

### Return type

[**O11yOut**](O11yOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminOrgs

> OrgsOut AdminOrgs(ctx).Execute()

Lists the tenant directory one row per org, sorted by slug: member count and the org's month-to-date spend and credit balance, read live from IAM and commerce.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminOrgs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminOrgs`: OrgsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminOrgs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminOrgsRequest struct via the builder pattern


### Return type

[**OrgsOut**](OrgsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminOverview

> OverviewOut AdminOverview(ctx).Execute()

Is the Platform Overview tiles: how many orgs and users are in the caller's tenant window, the fleet workload counts, and month-to-date spend and credits.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminOverview(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminOverview`: OverviewOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminOverviewRequest struct via the builder pattern


### Return type

[**OverviewOut**](OverviewOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminPlugins

> ListOut AdminPlugins(ctx).Scope(scope).Execute()

Reports what each host is actually running: every loaded plugin with its version, pid, uptime, reload and restart counts, and its measured CPU, RSS, thread and fd cost — read from the kernel, which is only answerable at all because a plugin is a process.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	scope := "fleet" // string | Scope \"host\" answers for THIS host only. Default \"fleet\" fans out to every live peer. A peer answers a host-scoped read, which is what stops the fan-out recursing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminPlugins(context.Background()).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminPlugins`: ListOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope \&quot;host\&quot; answers for THIS host only. Default \&quot;fleet\&quot; fans out to every live peer. A peer answers a host-scoped read, which is what stops the fan-out recursing. | 

### Return type

[**ListOut**](ListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminProducts

> ProductsOut AdminProducts(ctx).Kind(kind).Tier(tier).Env(env).Execute()

Lists the fleet workload registry: every operator App CR across the platform namespaces with its declared vs running image tag, reconciled health/phase and drift verdict.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	kind := "kind_example" // string | Kind matches the operator App CR's declared spec.role (sql|kv|generic|ingress). (optional)
	tier := "data" // string | Tier matches the derived infra grouping (cloud|data|edge|daemon|paas|app). (optional)
	env := "main" // string | Env matches the lifecycle namespace (main|test|dev). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminProducts(context.Background()).Kind(kind).Tier(tier).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminProducts`: ProductsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Kind matches the operator App CR&#39;s declared spec.role (sql|kv|generic|ingress). | 
 **tier** | **string** | Tier matches the derived infra grouping (cloud|data|edge|daemon|paas|app). | 
 **env** | **string** | Env matches the lifecycle namespace (main|test|dev). | 

### Return type

[**ProductsOut**](ProductsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminPromo

> RawOut AdminPromo(ctx).Execute()

Reads the current platform plan promo — the singleton discount offer, e.g.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminPromo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminPromo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminPromo`: RawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminPromo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminPromoRequest struct via the builder pattern


### Return type

[**RawOut**](RawOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminProvidersCredit

> ProvidersCreditOut AdminProvidersCredit(ctx).Execute()

Serves GET /v1/admin/providers/credit — the per-provider upstream credit ledger.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminProvidersCredit(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminProvidersCredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminProvidersCredit`: ProvidersCreditOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminProvidersCredit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminProvidersCreditRequest struct via the builder pattern


### Return type

[**ProvidersCreditOut**](ProvidersCreditOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReactivateCustomer

> AccessOut AdminReactivateCustomer(ctx, org).Execute()

Restores access for every member of the org, undoing a suspend.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	org := "org_example" // string | Org is the tenant slug from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminReactivateCustomer(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminReactivateCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminReactivateCustomer`: AccessOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminReactivateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the tenant slug from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReactivateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessOut**](AccessOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReloadPlugin

> ActionOut AdminReloadPlugin(ctx, name).ReloadIn(reloadIn).Execute()

Swaps a plugin for another build without dropping a request.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "billing" // string | Name is the app, from the path. It must be one the manifest declares.
	reloadIn := *openapiclient.NewReloadIn() // ReloadIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminReloadPlugin(context.Background(), name).ReloadIn(reloadIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminReloadPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminReloadPlugin`: ActionOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminReloadPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the app, from the path. It must be one the manifest declares. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReloadPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reloadIn** | [**ReloadIn**](ReloadIn.md) |  | 

### Return type

[**ActionOut**](ActionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminResizeDroplet

> MutationOut AdminResizeDroplet(ctx, id).DropletIn(dropletIn).Execute()

Changes a droplet's plan.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the DO droplet id, from the path. Numeric.
	dropletIn := *openapiclient.NewDropletIn() // DropletIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminResizeDroplet(context.Background(), id).DropletIn(dropletIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminResizeDroplet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminResizeDroplet`: MutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminResizeDroplet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO droplet id, from the path. Numeric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminResizeDropletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dropletIn** | [**DropletIn**](DropletIn.md) |  | 

### Return type

[**MutationOut**](MutationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminResizeVolume

> MutationOut AdminResizeVolume(ctx, id).VolumeIn(volumeIn).Execute()

Grows a volume.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the DO volume id, from the path.
	volumeIn := *openapiclient.NewVolumeIn() // VolumeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminResizeVolume(context.Background(), id).VolumeIn(volumeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminResizeVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminResizeVolume`: MutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminResizeVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO volume id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminResizeVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeIn** | [**VolumeIn**](VolumeIn.md) |  | 

### Return type

[**MutationOut**](MutationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRevenue

> RevenueOut AdminRevenue(ctx).Execute()

Is the fleet money board: total prepaid balances held, total realized spend, MRR, ARPU, a per-customer table sorted highest-revenue first, and a real 30-day spend trend from the usage ledger.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminRevenue(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminRevenue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRevenue`: RevenueOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminRevenue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRevenueRequest struct via the builder pattern


### Return type

[**RevenueOut**](RevenueOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRoles

> IamRowsOut AdminRoles(ctx).Owner(owner).P(p).PageSize(pageSize).Execute()

Lists IAM roles for one owner org, forwarded VERBATIM from IAM's get-roles.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "admin" // string | Owner is the org whose rows to read. Defaults to the admin org, which owns the platform's roles and applications. (optional)
	p := "1" // string | Page is the 1-based page number. Forwarded only when set — IAM applies its own default otherwise. (optional)
	pageSize := "50" // string | PageSize is rows per page. Forwarded only when set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminRoles(context.Background()).Owner(owner).P(p).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRoles`: IamRowsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Owner is the org whose rows to read. Defaults to the admin org, which owns the platform&#39;s roles and applications. | 
 **p** | **string** | Page is the 1-based page number. Forwarded only when set — IAM applies its own default otherwise. | 
 **pageSize** | **string** | PageSize is rows per page. Forwarded only when set. | 

### Return type

[**IamRowsOut**](IamRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminScaleNodePool

> MutationOut AdminScaleNodePool(ctx, id, pool).ScaleIn(scaleIn).Execute()

Sets a node pool's node count — the ONE correct way to change how many nodes a DOKS cluster has.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the DOKS cluster id, from the path.
	pool := "pool_example" // string | Pool is the node pool, from the path. Its DO id or its name — both are unique within a cluster, and an operator reads the name off the board.
	scaleIn := *openapiclient.NewScaleIn() // ScaleIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminScaleNodePool(context.Background(), id, pool).ScaleIn(scaleIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminScaleNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminScaleNodePool`: MutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminScaleNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DOKS cluster id, from the path. | 
**pool** | **string** | Pool is the node pool, from the path. Its DO id or its name — both are unique within a cluster, and an operator reads the name off the board. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminScaleNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scaleIn** | [**ScaleIn**](ScaleIn.md) |  | 

### Return type

[**MutationOut**](MutationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServices

> ServicesOut AdminServices(ctx).Execute()

Reads the launch board: every hosted service in the registry with its LIVE waitlist mode, evaluated through the flag engine.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminServices`: ServicesOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServicesRequest struct via the builder pattern


### Return type

[**ServicesOut**](ServicesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSetFlag

> FlagsOut AdminSetFlag(ctx, key).SetFlagIn(setFlagIn).Execute()

Stores or overwrites ONE platform switch's definition and answers with the whole board as it now stands.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	key := "key_example" // string | Key is the switch to write, taken from the path (e.g. \"waitlist.chat\").
	setFlagIn := *openapiclient.NewSetFlagIn() // SetFlagIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminSetFlag(context.Background(), key).SetFlagIn(setFlagIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminSetFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminSetFlag`: FlagsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminSetFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the switch to write, taken from the path (e.g. \&quot;waitlist.chat\&quot;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSetFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setFlagIn** | [**SetFlagIn**](SetFlagIn.md) |  | 

### Return type

[**FlagsOut**](FlagsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSetPromo

> RawOut AdminSetPromo(ctx).PromoIn(promoIn).Execute()

Upserts the platform plan promo — the ONE place the offer is configured.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	promoIn := *openapiclient.NewPromoIn() // PromoIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminSetPromo(context.Background()).PromoIn(promoIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminSetPromo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminSetPromo`: RawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminSetPromo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminSetPromoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promoIn** | [**PromoIn**](PromoIn.md) |  | 

### Return type

[**RawOut**](RawOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSetServiceMode

> ServiceOut AdminSetServiceMode(ctx, service).ServiceModeIn(serviceModeIn).Execute()

Flips ONE service's waitlist switch — the launch lever.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	service := "service_example" // string | Service is the slug to flip, taken from the path.
	serviceModeIn := *openapiclient.NewServiceModeIn() // ServiceModeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminSetServiceMode(context.Background(), service).ServiceModeIn(serviceModeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminSetServiceMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminSetServiceMode`: ServiceOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminSetServiceMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** | Service is the slug to flip, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSetServiceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceModeIn** | [**ServiceModeIn**](ServiceModeIn.md) |  | 

### Return type

[**ServiceOut**](ServiceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSnapshotVolume

> VolumeSnapshotOut AdminSnapshotVolume(ctx, id).VolumeIn(volumeIn).Execute()

Takes a point-in-time snapshot of one volume — the undo a delete relies on, available on its own so an operator can take one before any risky change.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the DO volume id, from the path.
	volumeIn := *openapiclient.NewVolumeIn() // VolumeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminSnapshotVolume(context.Background(), id).VolumeIn(volumeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminSnapshotVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminSnapshotVolume`: VolumeSnapshotOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminSnapshotVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO volume id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSnapshotVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeIn** | [**VolumeIn**](VolumeIn.md) |  | 

### Return type

[**VolumeSnapshotOut**](VolumeSnapshotOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptions

> SubscriptionsOut AdminSubscriptions(ctx).Status(status).Org(org).Limit(limit).Execute()

Answers GET /v1/admin/subscriptions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	status := "status_example" // string | Status filters on the subscription's LATEST lifecycle status (active, trialing, canceled, …), matched case-insensitively. (optional)
	org := "org_example" // string | Org filters to one tenant, matched exactly. (optional)
	limit := "limit_example" // string | Limit caps the rows returned. total still reports the full match count. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminSubscriptions(context.Background()).Status(status).Org(org).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminSubscriptions`: SubscriptionsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status filters on the subscription&#39;s LATEST lifecycle status (active, trialing, canceled, …), matched case-insensitively. | 
 **org** | **string** | Org filters to one tenant, matched exactly. | 
 **limit** | **string** | Limit caps the rows returned. total still reports the full match count. | 

### Return type

[**SubscriptionsOut**](SubscriptionsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubsystems

> SubsystemsOut AdminSubsystems(ctx).Range_(range_).Execute()

subsystems answers GET /v1/admin/subsystems.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	range_ := "range__example" // string | Range bounds the telemetry window: 24h, 7d or 30d. Anything else, including empty, resolves to the default through the same o11yRange the o11y board uses. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminSubsystems(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminSubsystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminSubsystems`: SubsystemsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminSubsystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubsystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range bounds the telemetry window: 24h, 7d or 30d. Anything else, including empty, resolves to the default through the same o11yRange the o11y board uses. | 

### Return type

[**SubsystemsOut**](SubsystemsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSuspendCustomer

> AccessOut AdminSuspendCustomer(ctx, org).Execute()

Cuts off every member of the org: IAM refuses a forbidden user at login AND at token issuance, so a suspended customer can neither sign in nor mint a fresh token.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	org := "org_example" // string | Org is the tenant slug from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminSuspendCustomer(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminSuspendCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminSuspendCustomer`: AccessOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminSuspendCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the tenant slug from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSuspendCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessOut**](AccessOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSync

> SyncOut AdminSync(ctx).Execute()

Answers the operator's \"Sync now\" button.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminSync`: SyncOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSyncRequest struct via the builder pattern


### Return type

[**SyncOut**](SyncOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUpdateCap

> RawOut AdminUpdateCap(ctx, id).CapIn(capIn).Execute()

Edits one cap by id — raise or lower the ceiling, flip enforcement.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the cap to edit or remove, from the path. Unused by the list and create ops.
	capIn := *openapiclient.NewCapIn() // CapIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminUpdateCap(context.Background(), id).CapIn(capIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpdateCap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminUpdateCap`: RawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminUpdateCap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cap to edit or remove, from the path. Unused by the list and create ops. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateCapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **capIn** | [**CapIn**](CapIn.md) |  | 

### Return type

[**RawOut**](RawOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUpsertService

> ServiceOut AdminUpsertService(ctx).ServiceInput(serviceInput).Execute()

Onboards a hosted service, or edits one, so a new host comes under the launch gate WITHOUT a redeploy.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	serviceInput := *openapiclient.NewServiceInput() // ServiceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminUpsertService(context.Background()).ServiceInput(serviceInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpsertService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminUpsertService`: ServiceOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminUpsertService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpsertServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceInput** | [**ServiceInput**](ServiceInput.md) |  | 

### Return type

[**ServiceOut**](ServiceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUsage

> UsageOut AdminUsage(ctx).Org(org).Execute()

Returns the trailing 30 days of AI usage: one org's when org names one, else the whole fleet's — the spend, the tokens and the requests, the daily curve behind them, and the split by model.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	org := "acme" // string | Org reads ONE tenant's trailing-30-day total instead of the fleet sum. Honoured for a SuperAdmin only — a white-label admin always reads their own org.  The window is the one core.OrgMoney returns, and it is what the operator board beside this already labelled (\"Daily, last 30 days\"). The wire used to say month-to-date while that UI said 30 days; they agree now. This comment is REGENERATED into plugin/admin/openapi.json and openapi.yaml as the ?org parameter description, so a stale word here ships as a contradiction inside one spec file — which is the drift this whole change set exists to remove. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminUsage(context.Background()).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminUsage`: UsageOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org reads ONE tenant&#39;s trailing-30-day total instead of the fleet sum. Honoured for a SuperAdmin only — a white-label admin always reads their own org.  The window is the one core.OrgMoney returns, and it is what the operator board beside this already labelled (\&quot;Daily, last 30 days\&quot;). The wire used to say month-to-date while that UI said 30 days; they agree now. This comment is REGENERATED into plugin/admin/openapi.json and openapi.yaml as the ?org parameter description, so a stale word here ships as a contradiction inside one spec file — which is the drift this whole change set exists to remove. | 

### Return type

[**UsageOut**](UsageOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUsageFunding

> UsageFundingOut AdminUsageFunding(ctx).From(from).To(to).Execute()

Splits our upstream AI usage by how it was FUNDED: one row per (provider, model) over the window, tagged credit (provider grant still remaining), paid (grant exhausted) or paid_only (no grant at all).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	from := "2026-07-01T00:00:00Z" // string | From is the inclusive start of the window. Unparseable or absent, together with To, falls back to the last 30 days. (optional)
	to := "2026-07-27T00:00:00Z" // string | To is the exclusive end of the window. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminUsageFunding(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUsageFunding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminUsageFunding`: UsageFundingOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminUsageFunding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsageFundingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | From is the inclusive start of the window. Unparseable or absent, together with To, falls back to the last 30 days. | 
 **to** | **string** | To is the exclusive end of the window. | 

### Return type

[**UsageFundingOut**](UsageFundingOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUsers

> UsersOut AdminUsers(ctx).Org(org).Q(q).P(p).PageSize(pageSize).Execute()

Lists the user directory across the caller's tenant window, one page at a time.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	org := "acme" // string | Org narrows the directory to ONE tenant. Honoured for a SuperAdmin only — a white-label admin is pinned to their own org and this is ignored. (optional)
	q := "ada" // string | Query is a free-text filter, matched by IAM as a \"contains\" over the user name. (optional)
	p := "1" // string | Page is the 1-based page number. Defaults to \"1\"; IAM returns zero rows AND a zero total when it is unset, so this layer never leaves it empty. (optional)
	pageSize := "50" // string | PageSize is rows per page. Defaults to \"200\", the shared admin page size. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminUsers(context.Background()).Org(org).Q(q).P(p).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminUsers`: UsersOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org narrows the directory to ONE tenant. Honoured for a SuperAdmin only — a white-label admin is pinned to their own org and this is ignored. | 
 **q** | **string** | Query is a free-text filter, matched by IAM as a \&quot;contains\&quot; over the user name. | 
 **p** | **string** | Page is the 1-based page number. Defaults to \&quot;1\&quot;; IAM returns zero rows AND a zero total when it is unset, so this layer never leaves it empty. | 
 **pageSize** | **string** | PageSize is rows per page. Defaults to \&quot;200\&quot;, the shared admin page size. | 

### Return type

[**UsersOut**](UsersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminVolumes

> VolumesOut AdminVolumes(ctx).Execute()

Returns the realtime block-storage board: the DigitalOcean volume fleet (count, capacity, monthly list cost, per-volume region and attachment) plus the analytics datastore's OWN fill, read from its system.disks.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminVolumes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminVolumes`: VolumesOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminVolumes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminVolumesRequest struct via the builder pattern


### Return type

[**VolumesOut**](VolumesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminWaitlist

> RawOut AdminWaitlist(ctx).Waitlist(waitlist).Page(page).PageSize(pageSize).Execute()

Reads one waitlist's leaderboard from the Hanzo waitlist engine — position, points and referral standing per entry — proxied server-authed with the engine secret, never a client credential.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	waitlist := "chat" // string | Waitlist is the waitlist slug to read (e.g. \"chat\"). The engine decides what an empty slug means. (optional)
	page := "1" // string | Page is the 1-based page number. (optional)
	pageSize := "50" // string | PageSize is entries per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminWaitlist(context.Background()).Waitlist(waitlist).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminWaitlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminWaitlist`: RawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminWaitlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminWaitlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **waitlist** | **string** | Waitlist is the waitlist slug to read (e.g. \&quot;chat\&quot;). The engine decides what an empty slug means. | 
 **page** | **string** | Page is the 1-based page number. | 
 **pageSize** | **string** | PageSize is entries per page. | 

### Return type

[**RawOut**](RawOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminWaitlistBoost

> RawOut AdminWaitlistBoost(ctx).WaitlistBoostRequest(waitlistBoostRequest).Execute()

Grants a user waitlist points, moving them up toward the access cutoff.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	waitlistBoostRequest := *openapiclient.NewWaitlistBoostRequest() // WaitlistBoostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminWaitlistBoost(context.Background()).WaitlistBoostRequest(waitlistBoostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminWaitlistBoost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminWaitlistBoost`: RawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminWaitlistBoost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminWaitlistBoostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **waitlistBoostRequest** | [**WaitlistBoostRequest**](WaitlistBoostRequest.md) |  | 

### Return type

[**RawOut**](RawOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminAffiliates

> DirectoryOut GetAdminAffiliates(ctx).Limit(limit).Execute()

Lists every affiliate across the fleet with its ORG exposed, plus a fleet summary of lifetime accrued, still-pending and paid commission in integer cents.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	limit := int32(56) // int32 | Limit caps the rows returned. Absent or non-positive means the default of 500; anything above 1000 is clamped to 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminAffiliates(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminAffiliates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminAffiliates`: DirectoryOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminAffiliates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminAffiliatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned. Absent or non-positive means the default of 500; anything above 1000 is clamped to 1000. | 

### Return type

[**DirectoryOut**](DirectoryOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminAuthors

> AdminBook GetAdminAuthors(ctx).Limit(limit).Execute()

Returns the platform's whole author program — every org's author record, not the caller's — with each one's repository and deploy counts and a fleet roll-up of the money accrued, pending and paid.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	limit := int32(56) // int32 | Limit bounds the page. 0 or less means the default of 500; anything above 1000 is clamped to 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminAuthors(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminAuthors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminAuthors`: AdminBook
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminAuthors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminAuthorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit bounds the page. 0 or less means the default of 500; anything above 1000 is clamped to 1000. | 

### Return type

[**AdminBook**](AdminBook.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminAuthorsByIdBasis

> BasisResult GetAdminAuthorsByIdBasis(ctx, id).Period(period).Execute()

Returns the audit trail behind ONE author's royalty — the same payload the author reads at /v1/authors/basis, from the same builder, so support sees exactly what the author sees rather than a parallel view free to drift.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "aut_1f…" // string | ID is the author record's handle, from the path.
	period := "2026-07" // string | Period is the UTC accrual month, YYYY-MM. Empty means every period; any other shape is refused with 400. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminAuthorsByIdBasis(context.Background(), id).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminAuthorsByIdBasis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminAuthorsByIdBasis`: BasisResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminAuthorsByIdBasis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the author record&#39;s handle, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminAuthorsByIdBasisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Period is the UTC accrual month, YYYY-MM. Empty means every period; any other shape is refused with 400. | 

### Return type

[**BasisResult**](BasisResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminCatalog

> AdminCatalogOut GetAdminCatalog(ctx).Execute()

Returns the full model and provider catalog annotated with each entry's enablement state, for the operator console.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminCatalog`: AdminCatalogOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminCatalogRequest struct via the builder pattern


### Return type

[**AdminCatalogOut**](AdminCatalogOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminEnablement

> AdminEnablementBoard GetAdminEnablement(ctx).Execute()

Returns every item an operator has set an enablement state on — its global state (off, beta or ga) and the orgs granted its beta.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminEnablement(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminEnablement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminEnablement`: AdminEnablementBoard
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminEnablement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminEnablementRequest struct via the builder pattern


### Return type

[**AdminEnablementBoard**](AdminEnablementBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminReferrals

> ReferralsOut GetAdminReferrals(ctx).Execute()

Answers the referral board: the top referrers by lifetime commission, the funnel conversion rate (referred orgs that have actually produced commission, over all referred orgs), and the accrual LIABILITY the platform owes, broken out by upline level.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminReferrals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminReferrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminReferrals`: ReferralsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminReferrals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminReferralsRequest struct via the builder pattern


### Return type

[**ReferralsOut**](ReferralsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminReferralsBonuses

> AdminBonusesEnvelope GetAdminReferralsBonuses(ctx).Limit(limit).Execute()

Returns every referral edge in the directory with a fleet summary.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	limit := "limit_example" // string | Limit is how many referrals to return, as a decimal string in the `?limit=` query. Absent, unparseable or non-positive means 500; over 1000 is clamped to 1000. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminReferralsBonuses(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminReferralsBonuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminReferralsBonuses`: AdminBonusesEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminReferralsBonuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminReferralsBonusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | Limit is how many referrals to return, as a decimal string in the &#x60;?limit&#x3D;&#x60; query. Absent, unparseable or non-positive means 500; over 1000 is clamped to 1000. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. | 

### Return type

[**AdminBonusesEnvelope**](AdminBonusesEnvelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminTreasury

> AdminReportOut GetAdminTreasury(ctx).Limit(limit).Execute()

Returns the whole treasury board for a SuperAdmin: the reserve fund report, the recent double-entry journal, and the Hanzo L1 anchor status of the ledger root.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	limit := int32(56) // int32 | Limit caps the journal entries returned. Out of range or unparseable takes the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminTreasury(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminTreasury``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminTreasury`: AdminReportOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminTreasury`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminTreasuryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the journal entries returned. Out of range or unparseable takes the default. | 

### Return type

[**AdminReportOut**](AdminReportOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAdminCatalogModelsByWildcard1

> PatchAdminCatalogModelsByWildcard1(ctx, wildcard1).Execute()

Turn one model off, into beta for named orgs, or generally available



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.PatchAdminCatalogModelsByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PatchAdminCatalogModelsByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAdminCatalogModelsByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAdminCatalogProvidersByName

> Overlay PatchAdminCatalogProvidersByName(ctx, name).ProviderPatchIn(providerPatchIn).Execute()

Sets one provider's availability overlay.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "name_example" // string | Name is the provider the overlay belongs to, from the URL.
	providerPatchIn := *openapiclient.NewProviderPatchIn() // ProviderPatchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PatchAdminCatalogProvidersByName(context.Background(), name).ProviderPatchIn(providerPatchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PatchAdminCatalogProvidersByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAdminCatalogProvidersByName`: Overlay
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PatchAdminCatalogProvidersByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the provider the overlay belongs to, from the URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAdminCatalogProvidersByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerPatchIn** | [**ProviderPatchIn**](ProviderPatchIn.md) |  | 

### Return type

[**Overlay**](Overlay.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminAffiliatesByIdApprove

> AffiliateOut PostAdminAffiliatesByIdApprove(ctx, id).Approval(approval).Execute()

Approves an affiliate and MINTS its referral code — the moment the partner has a working share link and starts accruing.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the affiliate to approve, from the path.
	approval := *openapiclient.NewApproval() // Approval | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminAffiliatesByIdApprove(context.Background(), id).Approval(approval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminAffiliatesByIdApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminAffiliatesByIdApprove`: AffiliateOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminAffiliatesByIdApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the affiliate to approve, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminAffiliatesByIdApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approval** | [**Approval**](Approval.md) |  | 

### Return type

[**AffiliateOut**](AffiliateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminAffiliatesByIdPayout

> PayoutOut PostAdminAffiliatesByIdPayout(ctx, id).Disbursal(disbursal).Execute()

Pays out accrued commission and answers the payout row with the affiliate's updated balances.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the affiliate to pay, from the path.
	disbursal := *openapiclient.NewDisbursal() // Disbursal | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminAffiliatesByIdPayout(context.Background(), id).Disbursal(disbursal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminAffiliatesByIdPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminAffiliatesByIdPayout`: PayoutOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminAffiliatesByIdPayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the affiliate to pay, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminAffiliatesByIdPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disbursal** | [**Disbursal**](Disbursal.md) |  | 

### Return type

[**PayoutOut**](PayoutOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminAffiliatesByIdRate

> AffiliateOut PostAdminAffiliatesByIdRate(ctx, id).RateSet(rateSet).Execute()

Sets one affiliate's DIRECT commission rate, in basis points of Hanzo's margin.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the affiliate whose direct rate moves, from the path.
	rateSet := *openapiclient.NewRateSet() // RateSet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminAffiliatesByIdRate(context.Background(), id).RateSet(rateSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminAffiliatesByIdRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminAffiliatesByIdRate`: AffiliateOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminAffiliatesByIdRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the affiliate whose direct rate moves, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminAffiliatesByIdRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rateSet** | [**RateSet**](RateSet.md) |  | 

### Return type

[**AffiliateOut**](AffiliateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminAffiliatesByIdSuspend

> AffiliateOut PostAdminAffiliatesByIdSuspend(ctx, id).Execute()

Suspends an affiliate: it stops accruing on the next sweep, and its code stops resolving for new attributions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the affiliate's server-minted handle, \"aff_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminAffiliatesByIdSuspend(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminAffiliatesByIdSuspend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminAffiliatesByIdSuspend`: AffiliateOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminAffiliatesByIdSuspend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the affiliate&#39;s server-minted handle, \&quot;aff_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminAffiliatesByIdSuspendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AffiliateOut**](AffiliateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminAffiliatesSweep

> AccrualsOut PostAdminAffiliatesSweep(ctx).Execute()

Runs the accrual: for each referred org it reads that org's metered spend for the current period and accrues commission to every affiliate up its referral chain, then answers how many sources were swept and how many NEW accruals landed.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminAffiliatesSweep(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminAffiliatesSweep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminAffiliatesSweep`: AccrualsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminAffiliatesSweep`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminAffiliatesSweepRequest struct via the builder pattern


### Return type

[**AccrualsOut**](AccrualsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminAuthorsByIdApprove

> AuthorResult PostAdminAuthorsByIdApprove(ctx, id).ApproveRequest(approveRequest).Execute()

Admits one author to EARNING, optionally on a negotiated royalty share.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "aut_1f…" // string | ID is the author to approve, from the path.
	approveRequest := *openapiclient.NewApproveRequest() // ApproveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminAuthorsByIdApprove(context.Background(), id).ApproveRequest(approveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminAuthorsByIdApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminAuthorsByIdApprove`: AuthorResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminAuthorsByIdApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the author to approve, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminAuthorsByIdApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approveRequest** | [**ApproveRequest**](ApproveRequest.md) |  | 

### Return type

[**AuthorResult**](AuthorResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminAuthorsByIdPayout

> PayoutResult PostAdminAuthorsByIdPayout(ctx, id).PayoutRequest(payoutRequest).Execute()

Records a payout of accrued royalty and settles it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "aut_1f…" // string | ID is the author to pay, from the path.
	payoutRequest := *openapiclient.NewPayoutRequest() // PayoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminAuthorsByIdPayout(context.Background(), id).PayoutRequest(payoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminAuthorsByIdPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminAuthorsByIdPayout`: PayoutResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminAuthorsByIdPayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the author to pay, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminAuthorsByIdPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payoutRequest** | [**PayoutRequest**](PayoutRequest.md) |  | 

### Return type

[**PayoutResult**](PayoutResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminAuthorsByIdSuspend

> AuthorResult PostAdminAuthorsByIdSuspend(ctx, id).Execute()

Stops one author earning.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the author record's handle, \"aut_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminAuthorsByIdSuspend(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminAuthorsByIdSuspend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminAuthorsByIdSuspend`: AuthorResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminAuthorsByIdSuspend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the author record&#39;s handle, \&quot;aut_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminAuthorsByIdSuspendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorResult**](AuthorResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminAuthorsSweep

> AuthorSweepResult PostAdminAuthorsSweep(ctx).Execute()

Runs the accrual sweep across every approved author: for each of their deploying orgs it computes this period's royalty from that org's metered spend and latches it at most once per period.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminAuthorsSweep(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminAuthorsSweep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminAuthorsSweep`: AuthorSweepResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminAuthorsSweep`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminAuthorsSweepRequest struct via the builder pattern


### Return type

[**AuthorSweepResult**](AuthorSweepResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminReferralsSweep

> SweepEnvelope PostAdminReferralsSweep(ctx).Execute()

Qualify-checks every pending referral and advances the ones that now qualify.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminReferralsSweep(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminReferralsSweep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminReferralsSweep`: SweepEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminReferralsSweep`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminReferralsSweepRequest struct via the builder pattern


### Return type

[**SweepEnvelope**](SweepEnvelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminTreasuryAnchor

> AnchorOut PostAdminTreasuryAnchor(ctx).Execute()

Commits the current ledger root to Hanzo L1, making the books tamper-evident on chain, and returns the anchoring status.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminTreasuryAnchor(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminTreasuryAnchor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminTreasuryAnchor`: AnchorOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminTreasuryAnchor`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminTreasuryAnchorRequest struct via the builder pattern


### Return type

[**AnchorOut**](AnchorOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminTreasuryPolicy

> PolicyOut PostAdminTreasuryPolicy(ctx).PolicyRequest(policyRequest).Execute()

Sets the revenue-share basis points a sweep accrues into the reserve fund and returns the stored policy.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	policyRequest := *openapiclient.NewPolicyRequest() // PolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminTreasuryPolicy(context.Background()).PolicyRequest(policyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminTreasuryPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminTreasuryPolicy`: PolicyOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminTreasuryPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminTreasuryPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyRequest** | [**PolicyRequest**](PolicyRequest.md) |  | 

### Return type

[**PolicyOut**](PolicyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminTreasurySeed

> SeedOut PostAdminTreasurySeed(ctx).SeedRequest(seedRequest).Execute()

Injects bootstrap capital into the reserve fund so backed payouts can begin before the first revenue-share sweep, and returns the journal entry it wrote.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	seedRequest := *openapiclient.NewSeedRequest() // SeedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminTreasurySeed(context.Background()).SeedRequest(seedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminTreasurySeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminTreasurySeed`: SeedOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminTreasurySeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminTreasurySeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seedRequest** | [**SeedRequest**](SeedRequest.md) |  | 

### Return type

[**SeedOut**](SeedOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminTreasurySweep

> SweepOut PostAdminTreasurySweep(ctx).SweepRequest(sweepRequest).Execute()

Posts the revenue-share accrual for one period — revenue into the reserve fund, at the current policy's basis points — and returns what it moved.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	sweepRequest := *openapiclient.NewSweepRequest() // SweepRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminTreasurySweep(context.Background()).SweepRequest(sweepRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminTreasurySweep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminTreasurySweep`: SweepOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminTreasurySweep`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminTreasurySweepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sweepRequest** | [**SweepRequest**](SweepRequest.md) |  | 

### Return type

[**SweepOut**](SweepOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAdminEnablement

> AdminEnablementItem PutAdminEnablement(ctx).SetEnablementBody(setEnablementBody).Execute()

Sets one item's global enablement state — off, beta or ga — and optionally replaces the list of orgs granted its beta.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	setEnablementBody := *openapiclient.NewSetEnablementBody() // SetEnablementBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PutAdminEnablement(context.Background()).SetEnablementBody(setEnablementBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PutAdminEnablement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAdminEnablement`: AdminEnablementItem
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PutAdminEnablement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAdminEnablementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setEnablementBody** | [**SetEnablementBody**](SetEnablementBody.md) |  | 

### Return type

[**AdminEnablementItem**](AdminEnablementItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAdminTreasuryAnchorSigner

> SignerOut PutAdminTreasuryAnchorSigner(ctx).Execute()

Installs the reserve's threshold MPC wallet as the signer for on-chain anchors, and returns its EVM address so an operator can fund it for gas.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PutAdminTreasuryAnchorSigner(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PutAdminTreasuryAnchorSigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAdminTreasuryAnchorSigner`: SignerOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PutAdminTreasuryAnchorSigner`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutAdminTreasuryAnchorSignerRequest struct via the builder pattern


### Return type

[**SignerOut**](SignerOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

