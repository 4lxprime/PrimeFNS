package responses

import (
	"encoding/json"
	"time"
)

type ContentpagesLangs struct {
	De    string `json:"de"`
	Ru    string `json:"ru"`
	Ko    string `json:"ko"`
	En    string `json:"en"`
	It    string `json:"it"`
	Fr    string `json:"fr"`
	Es    string `json:"es"`
	Ar    string `json:"ar"`
	Ja    string `json:"ja"`
	Pl    string `json:"pl"`
	Es419 string `json:"es-419"`
	Tr    string `json:"tr"`
}

type ContentpagesSubGameSDField struct {
	Type    string `json:"_type"`
	Message struct {
		Image       string            `json:"image"`
		Hidden      bool              `json:"hidden"`
		Messagetype string            `json:"messagetype"`
		Type        string            `json:"_type"`
		Title       ContentpagesLangs `json:"title"`
		Body        ContentpagesLangs `json:"body"`
		Spotlight   bool              `json:"spotlight"`
	} `json:"message"`
}

type ContentpagesResponse struct {
	Title        string    `json:"_title"`
	ActiveDate   time.Time `json:"_activeDate"`
	LastModified time.Time `json:"lastModified"`
	Locale       string    `json:"_locale"`
	Loginmessage struct {
		Title        string `json:"_title"`
		Loginmessage struct {
			Type    string `json:"_type"`
			Message struct {
				Type  string `json:"_type"`
				Title string `json:"title"`
				Body  string `json:"body"`
			} `json:"message"`
		} `json:"loginmessage"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"loginmessage"`
	Survivalmessage struct {
		Title               string `json:"_title"`
		Overrideablemessage struct {
			Type    string `json:"_type"`
			Message struct {
				Type  string `json:"_type"`
				Title string `json:"title"`
				Body  string `json:"body"`
			} `json:"message"`
		} `json:"overrideablemessage"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"survivalmessage"`
	Athenamessage struct {
		Title               string `json:"_title"`
		Overrideablemessage struct {
			Type    string `json:"_type"`
			Message struct {
				Image string `json:"image"`
				Type  string `json:"_type"`
				Title string `json:"title"`
				Body  string `json:"body"`
			} `json:"message"`
		} `json:"overrideablemessage"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"athenamessage"`
	Subgameselectdata struct {
		Title               string    `json:"_title"`
		ActiveDate          time.Time `json:"_activeDate"`
		LastModified        time.Time `json:"lastModified"`
		Locale              string    `json:"_locale"`
		SaveTheWorldUnowned struct {
			Type    string `json:"_type"`
			Message struct {
				Image       string            `json:"image"`
				Hidden      bool              `json:"hidden"`
				Messagetype string            `json:"messagetype"`
				Type        string            `json:"_type"`
				Title       ContentpagesLangs `json:"title"`
				Body        ContentpagesLangs `json:"body"`
				Spotlight   bool              `json:"spotlight"`
			} `json:"message"`
		} `json:"saveTheWorldUnowned"`
		BattleRoyale struct {
			Type    string `json:"_type"`
			Message struct {
				Image       string            `json:"image"`
				Hidden      bool              `json:"hidden"`
				Messagetype string            `json:"messagetype"`
				Type        string            `json:"_type"`
				Title       ContentpagesLangs `json:"title"`
				Body        ContentpagesLangs `json:"body"`
				Spotlight   bool              `json:"spotlight"`
			} `json:"message"`
		} `json:"battleRoyale"`
		Creative struct {
			Type    string `json:"_type"`
			Message struct {
				Image       string            `json:"image"`
				Hidden      bool              `json:"hidden"`
				Messagetype string            `json:"messagetype"`
				Type        string            `json:"_type"`
				Title       ContentpagesLangs `json:"title"`
				Body        ContentpagesLangs `json:"body"`
				Spotlight   bool              `json:"spotlight"`
			} `json:"message"`
		} `json:"creative"`
		SaveTheWorld struct {
			Type    string `json:"_type"`
			Message struct {
				Image       string            `json:"image"`
				Hidden      bool              `json:"hidden"`
				Messagetype string            `json:"messagetype"`
				Type        string            `json:"_type"`
				Title       ContentpagesLangs `json:"title"`
				Body        ContentpagesLangs `json:"body"`
				Spotlight   bool              `json:"spotlight"`
			} `json:"message"`
		} `json:"saveTheWorld"`
	} `json:"subgameselectdata"`
	Savetheworldnews struct {
		News struct {
			Type     string `json:"_type"`
			Messages []struct {
				Image     string `json:"image"`
				Hidden    bool   `json:"hidden"`
				Type      string `json:"_type"`
				Adspace   string `json:"adspace"`
				Title     string `json:"title"`
				Body      string `json:"body"`
				Spotlight bool   `json:"spotlight"`
			} `json:"messages"`
		} `json:"news"`
		Title        string    `json:"_title"`
		NoIndex      bool      `json:"_noIndex"`
		AlwaysShow   bool      `json:"alwaysShow"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"savetheworldnews"`
	Battlepassaboutmessages struct {
		News struct {
			Type     string `json:"_type"`
			Messages []struct {
				Layout    string `json:"layout,omitempty"`
				Image     string `json:"image"`
				Hidden    bool   `json:"hidden"`
				Type      string `json:"_type"`
				Title     string `json:"title"`
				Body      string `json:"body"`
				Spotlight bool   `json:"spotlight"`
			} `json:"messages"`
		} `json:"news"`
		Title        string    `json:"_title"`
		NoIndex      bool      `json:"_noIndex"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"battlepassaboutmessages"`
	Playlistinformation struct {
		FrontendMatchmakingHeaderStyle string `json:"frontend_matchmaking_header_style"`
		Title                          string `json:"_title"`
		FrontendMatchmakingHeaderText  string `json:"frontend_matchmaking_header_text"`
		PlaylistInfo                   struct {
			Type      string `json:"_type"`
			Playlists []struct {
				Image         string `json:"image"`
				PlaylistName  string `json:"playlist_name"`
				Type          string `json:"_type"`
				Hidden        bool   `json:"hidden,omitempty"`
				SpecialBorder string `json:"special_border,omitempty"`
			} `json:"playlists"`
		} `json:"playlist_info"`
		NoIndex      bool      `json:"_noIndex"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"playlistinformation"`
	Playlistimages struct {
		Playlistimages struct {
			Images []struct {
				Image        string `json:"image"`
				Type         string `json:"_type"`
				Playlistname string `json:"playlistname"`
			} `json:"images"`
			Type string `json:"_type"`
		} `json:"playlistimages"`
		Title        string    `json:"_title"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"playlistimages"`
	Tournamentinformation struct {
		TournamentInfo struct {
			Tournaments []struct {
				TitleColor           string `json:"title_color"`
				LoadingScreenImage   string `json:"loading_screen_image"`
				BackgroundTextColor  string `json:"background_text_color"`
				BackgroundRightColor string `json:"background_right_color"`
				PosterBackImage      string `json:"poster_back_image"`
				Type                 string `json:"_type"`
				PinEarnedText        string `json:"pin_earned_text"`
				TournamentDisplayID  string `json:"tournament_display_id"`
				HighlightColor       string `json:"highlight_color"`
				ScheduleInfo         string `json:"schedule_info"`
				PrimaryColor         string `json:"primary_color"`
				FlavorDescription    string `json:"flavor_description"`
				PosterFrontImage     string `json:"poster_front_image"`
				ShortFormatTitle     string `json:"short_format_title"`
				TitleLine2           string `json:"title_line_2"`
				TitleLine1           string `json:"title_line_1"`
				ShadowColor          string `json:"shadow_color"`
				DetailsDescription   string `json:"details_description"`
				BackgroundLeftColor  string `json:"background_left_color"`
				LongFormatTitle      string `json:"long_format_title"`
				PosterFadeColor      string `json:"poster_fade_color"`
				SecondaryColor       string `json:"secondary_color"`
				PlaylistTileImage    string `json:"playlist_tile_image"`
				BaseColor            string `json:"base_color"`
			} `json:"tournaments"`
			Type string `json:"_type"`
		} `json:"tournament_info"`
		Title        string    `json:"_title"`
		NoIndex      bool      `json:"_noIndex"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"tournamentinformation"`
	Emergencynotice struct {
		News struct {
			Type     string `json:"_type"`
			Messages []struct {
				Hidden    bool   `json:"hidden"`
				Type      string `json:"_type"`
				Title     string `json:"title"`
				Body      string `json:"body"`
				Spotlight bool   `json:"spotlight"`
			} `json:"messages"`
		} `json:"news"`
		Title        string    `json:"_title"`
		NoIndex      bool      `json:"_noIndex"`
		AlwaysShow   bool      `json:"alwaysShow"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"emergencynotice"`
	Emergencynoticev2 struct {
		JcrIsCheckedOut  bool   `json:"jcr:isCheckedOut"`
		Title            string `json:"_title"`
		NoIndex          bool   `json:"_noIndex"`
		JcrBaseVersion   string `json:"jcr:baseVersion"`
		Emergencynotices struct {
			Type             string `json:"_type"`
			Emergencynotices []struct {
				Gamemodes []interface{} `json:"gamemodes"`
				Hidden    bool          `json:"hidden"`
				Type      string        `json:"_type"`
				Title     string        `json:"title"`
				Body      string        `json:"body"`
			} `json:"emergencynotices"`
		} `json:"emergencynotices"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"emergencynoticev2"`
	Koreancafe struct {
		Title    string `json:"_title"`
		CafeInfo struct {
			Cafes []struct {
				KoreanCafe            string `json:"korean_cafe"`
				KoreanCafeDescription string `json:"korean_cafe_description"`
				Type                  string `json:"_type"`
				KoreanCafeHeader      string `json:"korean_cafe_header"`
			} `json:"cafes"`
			Type string `json:"_type"`
		} `json:"cafe_info"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"koreancafe"`
	CreativeAds struct {
		AdInfo struct {
			Ads  []interface{} `json:"ads"`
			Type string        `json:"_type"`
		} `json:"ad_info"`
		Title        string    `json:"_title"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"creativeAds"`
	Playersurvey     struct{} `json:"playersurvey"`
	CreativeFeatures struct {
		AdInfo struct {
			Type string `json:"_type"`
		} `json:"ad_info"`
		Title        string    `json:"_title"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"creativeFeatures"`
	Specialoffervideo struct{} `json:"specialoffervideo"`
	Subgameinfo       struct {
		Battleroyale struct {
			Image                string            `json:"image"`
			Color                string            `json:"color"`
			Type                 string            `json:"_type"`
			Description          ContentpagesLangs `json:"description"`
			Subgame              string            `json:"subgame"`
			StandardMessageLine2 string            `json:"standardMessageLine2"`
			Title                ContentpagesLangs `json:"title"`
			StandardMessageLine1 string            `json:"standardMessageLine1"`
		} `json:"battleroyale"`
		Savetheworld struct {
			Image          string            `json:"image"`
			Color          string            `json:"color"`
			SpecialMessage string            `json:"specialMessage"`
			Type           string            `json:"_type"`
			Description    ContentpagesLangs `json:"description"`
			Subgame        string            `json:"subgame"`
			Title          ContentpagesLangs `json:"title"`
		} `json:"savetheworld"`
		Title    string `json:"_title"`
		NoIndex  bool   `json:"_noIndex"`
		Creative struct {
			Image                string            `json:"image"`
			Color                string            `json:"color"`
			Type                 string            `json:"_type"`
			Description          ContentpagesLangs `json:"description"`
			Subgame              string            `json:"subgame"`
			Title                ContentpagesLangs `json:"title"`
			StandardMessageLine1 string            `json:"standardMessageLine1"`
		} `json:"creative"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"subgameinfo"`
	Lobby struct {
		Backgroundimage string    `json:"backgroundimage"`
		Stage           string    `json:"stage"`
		Title           string    `json:"_title"`
		ActiveDate      time.Time `json:"_activeDate"`
		LastModified    time.Time `json:"lastModified"`
		Locale          string    `json:"_locale"`
	} `json:"lobby"`
	Battleroyalenews struct {
		News struct {
			Type  string `json:"_type"`
			Motds []struct {
				EntryType             string            `json:"entryType"`
				Image                 string            `json:"image"`
				TileImage             string            `json:"tileImage"`
				VideoMute             bool              `json:"videoMute"`
				Hidden                bool              `json:"hidden"`
				TabTitleOverride      string            `json:"tabTitleOverride"`
				Type                  string            `json:"_type"`
				Title                 ContentpagesLangs `json:"title"`
				Body                  ContentpagesLangs `json:"body"`
				OfferAction           string            `json:"offerAction"`
				VideoLoop             bool              `json:"videoLoop"`
				VideoStreamingEnabled bool              `json:"videoStreamingEnabled"`
				SortingPriority       int               `json:"sortingPriority"`
				WebsiteButtonText     string            `json:"websiteButtonText"`
				WebsiteURL            string            `json:"websiteURL"`
				ID                    string            `json:"id"`
				VideoAutoplay         bool              `json:"videoAutoplay"`
				VideoFullscreen       bool              `json:"videoFullscreen"`
				Spotlight             bool              `json:"spotlight"`
			} `json:"motds"`
			Messages []struct {
				Image     string `json:"image"`
				Hidden    bool   `json:"hidden"`
				Type      string `json:"_type"`
				Adspace   string `json:"adspace"`
				Title     string `json:"title"`
				Body      string `json:"body"`
				Spotlight bool   `json:"spotlight"`
			} `json:"messages"`
		} `json:"news"`
		Title        string    `json:"_title"`
		Header       string    `json:"header"`
		Style        string    `json:"style"`
		NoIndex      bool      `json:"_noIndex"`
		AlwaysShow   bool      `json:"alwaysShow"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"battleroyalenews"`
	Battleroyalenewsv2 struct{} `json:"battleroyalenewsv2"`
	Dynamicbackgrounds struct {
		Backgrounds struct {
			Backgrounds []struct {
				Stage string `json:"stage"`
				Type  string `json:"_type"`
				Key   string `json:"key"`
			} `json:"backgrounds"`
			Type string `json:"_type"`
		} `json:"backgrounds"`
		Title        string    `json:"_title"`
		NoIndex      bool      `json:"_noIndex"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
		TemplateName string    `json:"_templateName"`
	} `json:"dynamicbackgrounds"`
	Creativenews struct {
		News struct {
			Type     string `json:"_type"`
			Messages []struct {
				Image     string `json:"image"`
				Hidden    bool   `json:"hidden"`
				Type      string `json:"_type"`
				Adspace   string `json:"adspace"`
				Title     string `json:"title"`
				Body      string `json:"body"`
				Spotlight bool   `json:"spotlight"`
			} `json:"messages"`
		} `json:"news"`
		Title        string    `json:"_title"`
		Header       string    `json:"header"`
		Style        string    `json:"style"`
		NoIndex      bool      `json:"_noIndex"`
		AlwaysShow   bool      `json:"alwaysShow"`
		ActiveDate   time.Time `json:"_activeDate"`
		LastModified time.Time `json:"lastModified"`
		Locale       string    `json:"_locale"`
	} `json:"creativenews"`
	ShopSections      struct{}      `json:"shopSections"`
	SuggestedPrefetch []interface{} `json:"_suggestedPrefetch"`
}

func (contentpages *ContentpagesResponse) Parse(b []byte) error {
	return json.Unmarshal(b, contentpages)
}
