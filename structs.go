package weverse

type GetAccountStatusResult struct {
	EmailVerified bool `json:"emailVerified"`
	HasPassword   bool `json:"hasPassword"`
}

type AccountTerms struct {
	Terms []struct {
		TermsCode        string   `json:"termsCode"`
		TermsDocumentID  string   `json:"termsDocumentId"`
		Title            string   `json:"title"`
		Summary          string   `json:"summary"`
		URL              string   `json:"url"`
		URLType          string   `json:"urlType"`
		Required         bool     `json:"required"`
		Tags             []string `json:"tags"`
	} `json:"terms"`
}

type AccountInfo struct {
	UserID						string `json:"userId"`
	Email 						string `json:"email"`
	Nickname 					string `json:"nickname"`
	JoinCountry 				string `json:"joinCountry"`
	HasPassword 				bool   `json:"hasPassword"`
	ServiceConnected 			bool   `json:"serviceConnected"`
	ProfileUpdateRequired 		bool   `json:"profileUpdateRequired"`
	SMSVerified 				bool   `json:"smsVerified"`
	AgreementUpdateRequired 	bool   `json:"agreementUpdateRequired"`
	AdAdgreementUpdateRequired 	bool   `json:"adAgreementUpdateRequired"`
	TimeZoneUpdateRequired 		bool   `json:"timeZoneUpdateRequired"`
}

type Banner struct {
	BannerID int `json:"bannerId"`
	StartDate int `json:"startDate"`
	EndDate int `json:"endDate"`
	ContentType string `json:"contentType"`
	Content struct {
		ImageURL string `json:"imageUrl"`
		FirstTitle string `json:"firstTitle"`
		SecondTitle string `json:"secondTitle"`
		SubTitle string `json:"subTitle"`
		SecondSubTitle string `json:"secondSubTitle"`
		TextColor string `json:"textColor"`
	} `json:"content"`
	LandingUrl string `json:"landingUrl"`
	LandingUrlType string `json:"landingUrlType"`
	CommunityName string `json:"communityName"`
}

type Community struct {
	CommunityID int `json:"communityId"`
	CommunityName string `json:"communityName"`
	CommunityAlias string `json:"communityAlias"`
	UrlPath string `json:"urlPath"`
	LogoImage string `json:"logoImage"`
	HomeHeaderImage string `json:"homeHeaderImage"`
	MemberCount int `json:"memberCount"`
	LastArtistContentPublishedAt int `json:"lastArtistContentPublishedAt"`
	Recommended bool `json:"recommended"`
}

type Ad struct {
	AdType string `json:"adType"`
	AdUnitID string `json:"adUnitId"`
	Custom struct {
		Hl string `json:"hl"`
	}
	Width int `json:"width"`
	Height int `json:"height"`
}

type PageResult[T any] struct {
	Paging struct {
		PageNo    int `json:"pageNo"`
		MaxPageNo int `json:"maxPageNo"`
		Limit     int `json:"limit"`
	} `json:"paging"`
	Data []T `json:"data"`
	TotalCount int `json:"totalCount"`
}