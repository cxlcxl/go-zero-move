package statements

type CampaignList struct {
	CommentCount                      int     `json:"comment_count"`
	AchievementUnlockedCost           string  `json:"achievement_unlocked_cost"`
	CouponCount                       int     `json:"coupon_count"`
	EffectiveLeadsOnlineCount         int     `json:"effective_leads_online_count"`
	ClickCount                        int     `json:"click_count"`
	ConsultOnlineCost                 string  `json:"consult_online_cost"`
	SevenDayRetainCount               int     `json:"seven_day_retain_count"`
	ConsultOnlineCount                int     `json:"consult_online_count"`
	GamePackageClaimingCost           string  `json:"game_package_claiming_cost"`
	EffectiveLeadsPhoneCount          int     `json:"effective_leads_phone_count"`
	FollowCount                       int     `json:"follow_count"`
	FirstPurchaseMemberCardCost       string  `json:"first_purchase_membercard_cost"`
	AttributionIncomeIaa              string  `json:"attribution_income_iaa"`
	VoteCost                          string  `json:"vote_cost"`
	RetainCostNormalized              string  `json:"retain_cost_normalized"`
	UpdateCost                        string  `json:"update_cost"`
	PotentialCustomerOnlineCount      int     `json:"potential_customer_online_count"`
	ForwardCost                       string  `json:"forward_cost"`
	ThreeDayRetainCount               int     `json:"three_day_retain_count"`
	TravelBookingCost                 string  `json:"travel_booking_cost"`
	PotentialCustomerFormCount        int     `json:"potential_customer_form_count"`
	HighQualityActiveCost             string  `json:"high_quality_active_cost"`
	AddToWishlistCount                int     `json:"add_to_wishlist_count"`
	LevelAchievedCount                int     `json:"level_achieved_count"`
	FormSubmitCount                   int     `json:"form_submit_count"`
	EffectiveBookCost                 string  `json:"effective_book_cost"`
	NavigateCount                     int     `json:"navigate_count"`
	FormSubmitCost                    string  `json:"form_submit_cost"`
	Country                           string  `json:"country"`
	LandingPageClickCost              string  `json:"landingpage_click_cost"`
	AppReservationCost                float64 `json:"app_reservation_cost"`
	LotteryCount                      int     `json:"lottery_count"`
	PotentialCustomerPhoneCount       int     `json:"potential_customer_phone_count"`
	InstallCount                      int     `json:"install_count"`
	PreCreditCost                     string  `json:"precredit_cost"`
	AppCustomCount                    int     `json:"app_custom_count"`
	ContentViewCount                  int     `json:"content_view_count"`
	FollowScanCost                    string  `json:"follow_scan_cost"`
	CampaignId                        string  `json:"campaign_id"`
	StatDatetime                      string  `json:"stat_datetime"`
	ReadCount                         int     `json:"read_count"`
	EffectiveBookCount                int     `json:"effective_book_count"`
	AddCartCount                      int     `json:"add_cart_count"`
	LikeCount                         int     `json:"like_count"`
	BrowseCost                        string  `json:"browse_cost"`
	ActiveCostNormalized              string  `json:"active_cost_normalized"`
	RegisterCount                     int     `json:"register_count"`
	EffectiveLeadsOnlineCost          string  `json:"effective_leads_online_cost"`
	ReEngageCost                      string  `json:"re_engage_cost"`
	CreditCost                        string  `json:"credit_cost"`
	InstallCost                       string  `json:"install_cost"`
	CollectionCount                   int     `json:"collection_count"`
	FollowCost                        string  `json:"follow_cost"`
	PayCostNormalized                 string  `json:"pay_cost_normalized"`
	AdIncomeThirtyDayRoi              string  `json:"ad_income_thirty_day_roi"`
	AdIncomeSevenDayLtvHms            string  `json:"ad_income_seven_day_ltv_hms"`
	SubscribeCost                     string  `json:"subscribe_cost"`
	FirstPurchaseMemberCardCount      int     `json:"first_purchase_membercard_count"`
	GamePackageRedemptionCost         string  `json:"game_package_redemption_cost"`
	CommentCost                       string  `json:"comment_cost"`
	AchievementUnlockedCount          int     `json:"achievement_unlocked_count"`
	PotentialCustomerPhoneCost        string  `json:"potential_customer_phone_cost"`
	TrackingConversionCount           int     `json:"tracking_conversion_count"`
	AdvertiserId                      string  `json:"advertiser_id"`
	GamePackageRedemptionCount        int     `json:"game_package_redemption_count"`
	LoanCompletionCost                string  `json:"loan_completion_cost"`
	PayAmountNormalized               string  `json:"pay_amount_normalized"`
	AuthorizeCount                    int     `json:"authorize_count"`
	ReEngageCount                     int     `json:"re_engage_count"`
	ReservationCost                   string  `json:"reservation_cost"`
	AdIncomeFifteenDayRoi             string  `json:"ad_income_fifteen_day_roi"`
	EffectiveLeadsFormCost            string  `json:"effective_leadsform_cost"`
	LevelAchievedCost                 string  `json:"level_achieved_cost"`
	HighQualityActiveCount            int     `json:"high_quality_active_count"`
	RegisterCost                      string  `json:"register_cost"`
	InviteCount                       int     `json:"invite_count"`
	CreditCount                       int     `json:"credit_count"`
	DownloadCost                      string  `json:"download_cost"`
	AttributionIncomeIapNormalized    string  `json:"attribution_income_iap_normalized"`
	PurchaseMemberCardCount           int     `json:"purchase_membercard_count"`
	PublisherRealPriceOneDay          string  `json:"publisher_real_price_one_day"`
	InitiatedCheckoutCost             string  `json:"initiated_checkout_cost"`
	ForwardCount                      int     `json:"forward_count"`
	PurchaseMemberCardCost            string  `json:"purchase_membercard_cost"`
	SearchCount                       int     `json:"search_count"`
	ReservationCount                  int     `json:"reservation_count"`
	PotentialCustomerOnlineCost       string  `json:"potential_customer_online_cost"`
	PhoneDialingCount                 int     `json:"phone_dialing_count"`
	AuthorizeCost                     string  `json:"authorize_cost"`
	AddPaymentInfoCount               int     `json:"add_payment_info_count"`
	InviteCost                        string  `json:"invite_cost"`
	AddCartCost                       string  `json:"add_cart_cost"`
	RateCost                          string  `json:"rate_cost"`
	LandingPageClickCount             int     `json:"landingpage_click_count"`
	AdIncomeThreeDayRoi               string  `json:"ad_income_three_day_roi"`
	LoginCost                         string  `json:"login_cost"`
	VenusFormSubmitCount              int     `json:"venus_form_submit_count"`
	TravelBookingCount                int     `json:"travel_booking_count"`
	ContentViewCost                   string  `json:"content_view_cost"`
	BookAmount                        int     `json:"book_amount"`
	EffectiveCustomerAcquisitionCost  string  `json:"effective_customer_acquisition_cost"`
	SubscribeCount                    int     `json:"subscribe_count"`
	AdIncomeOneDayLtvHms              string  `json:"ad_income_one_day_ltv_hms"`
	PlayOverCount                     int     `json:"play_over_count"`
	LikeCost                          string  `json:"like_cost"`
	CouponCost                        string  `json:"coupon_cost"`
	AddPaymentInfoCost                string  `json:"add_payment_info_cost"`
	WebCustomCount                    int     `json:"web_custom_count"`
	LeadsLotteryCount                 int     `json:"leads_lottery_count"`
	PreCreditCount                    int     `json:"precredit_count"`
	AdIncomeFifteenDayLtvHms          string  `json:"ad_income_fifteen_day_ltv_hms"`
	ShareCost                         string  `json:"share_cost"`
	ReadCost                          string  `json:"read_cost"`
	CreateRoleCost                    string  `json:"create_role_cost"`
	LoginCount                        int     `json:"login_count"`
	AddQuickAppCost                   string  `json:"add_quick_app_cost"`
	ActiveCountNormalized             int     `json:"active_count_normalized"`
	AdIncomeThreeDayLtvHms            string  `json:"ad_income_three_day_ltv_hms"`
	PreOrderCount                     int     `json:"pre_order_count"`
	PreOrderCost                      string  `json:"pre_order_cost"`
	BrowseCount                       int     `json:"browse_count"`
	VoteCount                         int     `json:"vote_count"`
	OpenedFromPushNotificationCost    string  `json:"opened_frompushnotification_cost"`
	OrderSigningCount                 int     `json:"order_signing_count"`
	SevenDayRetainCost                string  `json:"seven_day_retain_cost"`
	UpdateCount                       int     `json:"update_count"`
	RedirectCount                     int     `json:"redirect_count"`
	DownloadCount                     int     `json:"download_count"`
	AddQuickAppCount                  int     `json:"add_quick_app_count"`
	LoanCompletionCount               int     `json:"loan_completion_count"`
	ShowCount                         int     `json:"show_count"`
	TutorialCompletionCount           int     `json:"tutorial_completion_count"`
	TrackingConversionValue           int     `json:"tracking_conversion_value"`
	AppReservationCount               string  `json:"app_reservation_count"`
	StartTrialCount                   int     `json:"start_trial_count"`
	PhoneDialingCost                  string  `json:"phone_dialing_cost"`
	AdIncomeTwoDayLtvHms              string  `json:"ad_income_two_day_ltv_hms"`
	PotentialCustomerFormCost         string  `json:"potential_customer_form_cost"`
	Cpc                               string  `json:"cpc"`
	DeliverCost                       string  `json:"deliver_cost"`
	AdIncomeTwoDayRoi                 string  `json:"ad_income_two_day_roi"`
	CreateRoleCount                   int     `json:"create_role_count"`
	LotteryCost                       string  `json:"lottery_cost"`
	EffectiveConsultCost              string  `json:"effective_consult_cost"`
	PlayCount                         int     `json:"play_count"`
	OpenedFromPushNotificationCount   int     `json:"opened_frompushnotification_count"`
	GamePackageClaimingCount          int     `json:"game_package_claiming_count"`
	RateCount                         int     `json:"rate_count"`
	AdIncomeOneDayRoi                 string  `json:"ad_income_one_day_roi"`
	OrderSigningCost                  string  `json:"order_signing_cost"`
	PayCountNormalized                int     `json:"pay_count_normalized"`
	VenusFormSubmitCost               string  `json:"venus_form_submit_cost"`
	RedirectCost                      string  `json:"redirect_cost"`
	EffectiveLeadsPhoneCost           string  `json:"effective_leads_phone_cost"`
	AppCustomCost                     string  `json:"app_custom_cost"`
	NavigateCost                      string  `json:"navigate_cost"`
	SearchCost                        string  `json:"search_cost"`
	CampaignName                      string  `json:"campaign_name"`
	FollowScanCount                   int     `json:"follow_scan_count"`
	AdIncomeThirtyDayLtvHms           string  `json:"ad_income_thirty_day_ltv_hms"`
	LeadsLotteryCost                  string  `json:"leads_lottery_cost"`
	SpentCreditsCount                 int     `json:"spent_credits_count"`
	InitiatedCheckoutCount            int     `json:"initiated_checkout_count"`
	EffectiveLeadsFormCount           int     `json:"effective_leadsform_count"`
	EffectiveConsultCount             int     `json:"effective_consult_count,omitempty"`
	CollectionCost                    string  `json:"collection_cost"`
	SpentCredit                       string  `json:"spent_creditt,omitempty"`
	Cost                              string  `json:"cost"`
	AddToWishlistCost                 string  `json:"add_to_wishlist_cost"`
	WebCustomCost                     string  `json:"web_custom_cost"`
	StartTrialCost                    string  `json:"start_trial_cost"`
	ThreeDayRetainCost                string  `json:"three_day_retain_cost"`
	ShareCount                        int     `json:"share_count"`
	ThousandShowCost                  string  `json:"thousand_show_cost"`
	TutorialCompletionCost            string  `json:"tutorial_completion_cost"`
	RetainCountNormalized             int     `json:"retain_count_normalized"`
	EffectiveCustomerAcquisitionCount int     `json:"effective_customer_acquisition_count"`
	AdIncomeSevenDayRoi               string  `json:"ad_income_seven_day_roi"`
	DeliverCount                      int     `json:"deliver_count"`
	EffectiveConsultUnt               int     `json:"effective_consultunt,omitempty"`
	SpentCreditsCost                  string  `json:"spent_credits_cost,omitempty"`
}

type CampaignRequest struct {
	AdvertiserId string            `json:"advertiser_id"`
	Page         int64             `json:"page"`
	PageSize     int64             `json:"page_size"`
	Filtering    CampaignFiltering `json:"filtering"`
}

type CampaignFiltering struct {
	CampaignIds      []string `json:"campaign_ids,optional,omitempty"`
	UpdatedBeginTime string   `json:"updated_begin_time,optional,omitempty"`
	UpdatedEndTime   string   `json:"updated_end_time,optional,omitempty"`
}

type CampaignResponse struct {
	BaseResponse
	Data struct {
		Data  CampaignQueue `json:"data"`
		Total int64         `json:"total"`
	} `json:"data"`
}

type CampaignQueue []*CampaignInfo

func (q CampaignQueue) GenerateMsg(fn func(interface{})) {
	for _, campaignData := range q {
		fn(campaignData)
	}
}

type CampaignInfo struct {
	CampaignName              string `json:"campaign_name"`
	CreatedTime               string `json:"created_time"`
	ProductType               string `json:"product_type"`
	UserBalanceStatus         string `json:"user_balance_status"`
	CampaignStatus            string `json:"campaign_status"`
	SyncFlowResourceSearchAd  string `json:"sync_flow_resource_searchad"`
	CampaignType              string `json:"campaign_type"`
	FlowResource              string `json:"flow_resource"`
	CampaignDailyBudgetStatus string `json:"campaign_daily_budget_status"`
	ShowStatus                string `json:"show_status"`
	CampaignId                string `json:"campaign_id"`
	TodayDailyBudget          int64  `json:"today_daily_budget"`
	TomorrowDailyBudget       int64  `json:"tomorrow_daily_budget"`
	MarketingGoal             string `json:"marketing_goal"`
}
