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

type CommunityDetail struct {
	CommunityID int `json:"communityId"`
	CommunityName string `json:"communityName"`
	UrlPath string `json:"urlPath"`
	AgencyProfile struct {
		ProfileImageUrl string `json:"profileImageUrl"`
		ProfileName string `json:"profileName"`
		ProfileCoverImageUrl string `json:"profileCoverImageUrl"`
	} `json:"agencyProfile"`
	LogoImage string `json:"logoImage"`
	HomeHeaderImage string `json:"homeHeaderImage"`
	ArtistCode string `json:"artistCode"`
	FanEventUrl string `json:"fanEventUrl"`
	HomeGradationColor struct {
		UpLeftColorCode string `json:"upLeftColorCode"`
		UpRightColorCode string `json:"upRightColorCode"`
		DownLeftColorCode string `json:"downLeftColorCode"`
		DownRightColorCode string `json:"downRightColorCode"`
	} `json:"homeGradationColor"`
	GradationColor struct {
		MainLeftColorCode string `json:"mainLeftColorCode"`
		MainRightColorCode string `json:"mainRightColorCode"`
		SubLeftColorCode string `json:"subLeftColorCode"`
		SubRightColorCode string `json:"subRightColorCode"`
	} `json:"gradationColor"`
	CommunityColor string `json:"communityColor"`
	HasMembershipProduct bool `json:"hasMembershipProduct"`
	MemberCountString string `json:"memberCountString"`
	AvailableActions []string `json:"availableActions"`
	BirthdayArtists []struct {
		Name string `json:"name"`
		Birthday string `json:"birthday"`
	} `json:"birthdayArtists"`
	HasOnAirLivePost bool `json:"hasOnAirLivePost"`
	HasOnAirPartyPost bool `json:"hasOnAirPartyPost"`
	Membership struct {
		IsOnSale bool `json:"isOnSale"`
		IsActivated bool `json:"isActivated"`
		Data []struct {
			Type string `json:"type"`
			IsPurchased bool `json:"isPurchased"`
			AdditionalInfo struct {
				ArtistCode string `json:"artistCode"`
			} `json:"additionalInfo"`
		} `json:"data"`
	} `json:"membership"`
	ShopUrl string `json:"shopUrl"`
	Tabs []struct {
		TabKey string `json:"tabKey"`
		ChildTabs []struct {
			TabKey string `json:"tabKey"`
		} `json:"childTabs,omitempty"`
		Data struct {
			LastStickerUpdatedAt int `json:"lastStickerUpdatedAt"`
			LastUpdatedAt int64 `json:"lastUpdatedAt"`
		} `json:"data,omitempty"`
	} `json:"tabs"`
	CommunityProducts struct {
		MEMBERSHIP struct {
			Enabled bool `json:"enabled"`
			Data struct {
				IsOnSale bool `json:"isOnSale"`
				IsActivated bool `json:"isActivated"`
			} `json:"data"`
		} `json:"MEMBERSHIP"`
		WDM struct {
			Enabled bool `json:"enabled"`
		} `json:"WDM"`
		FAN_LETTER struct {
			Enabled bool `json:"enabled"`
			Data struct {
				Enable bool `json:"enable"`
				LastUpdatedAt int64 `json:"lastUpdatedAt"`
				LastStickerUpdatedAt int `json:"lastStickerUpdatedAt"`
			} `json:"data"`
		} `json:"FAN_LETTER"`
	} `json:"communityProducts"`
	Config struct {} `json:"config"`
}