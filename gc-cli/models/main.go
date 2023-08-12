package models

type ErrorEvents struct {
	ErrorEvent    []ErrorEvent `json:"errorEvents"`
	NextPageToken string       `json:"nextPageToken"`
}

type ErrorEvent struct {
	Message        string `json:"message"`
	EventTime      string `json:"eventTime"`
	ServiceContext struct {
		Service string `json:"service"`
		Version string `json:"version"` // what is this
	} `json:"serviceContext"`
}

type GroupStats struct {
	ErrorGroupStats []GroupStat `json:"errorGroupStats"`
	TimeRangeBegin  string      `json:"timeRangeBegin"`
	NextPageToken   string      `json:"nextPageToken"`
}

type AffectedService struct {
	Service string `json:"service"`
}
type GroupStat struct {
	Group struct {
		GroupId          string `json:"groupId"`
		ResolutionStatus string `json:"resolutionStatus"` // acknowledged | open | etc
	} `json:"group"`
	AffectedServices []AffectedService `json:"affectedServices"`
	Count            string            `json:"count"`
	FirstSeenTime    string            `json:"firstSeenTime"`
	LastSeenTime     string            `json:"lastSeenTime"`
	Representative   struct {
		Message        string `json:"message"`
		ServiceContext struct {
			Service string `json:"service"`
		} `json:"serviceContext"`
	} `json:"representative"`
	// lastSeenTime
	// firstSeenTime
	// affectedServices
	// numAffectedServices
}
