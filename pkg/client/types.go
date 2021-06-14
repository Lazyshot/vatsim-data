package client

import "time"

type Response struct {
	General     GeneralData  `json:"general"`
	Pilots      []Pilot      `json:"pilots"`
	Controllers []Controller `json:"controllers"`
	ATIS        []ATIS       `json:"atis"`
	Servers     []Server     `json:"servers"`
	Prefiles    []Prefiled   `json:"prefiles"`
}

type GeneralData struct {
	Version          int       `json:"version"`
	Reload           int       `json:"reload"`
	UpdateString     string    `json:"update"`
	UpdateTimestamp  time.Time `json:"updated_timestamp"`
	ConnectedClients int       `json:"connected_clients"`
	UniqueUsers      int       `json:"unique_users"`
}

type Pilot struct {
	CID         int        `json:"cid"`
	Name        string     `json:"name"`
	Callsign    string     `json:"callsign"`
	Server      string     `json:"server"`
	PilotRating int        `json:"pilot_rating"`
	Latitude    float32    `json:"latitude"`
	Longitude   float32    `json:"longitude"`
	Altitude    int        `json:"altitude"`
	GroundSpeed int        `json:"groundspeed"`
	Transponder string     `json:"transponder"`
	Heading     int        `json:"heading"`
	QNHMercury  float32    `json:"qnh_i_hg"`
	QNHBars     int        `json:"qnh_mb"`
	FlightPlan  FlightPlan `json:"flight_plan"`
	LogonTime   time.Time  `json:"logon_time"`
	LastUpdated time.Time  `json:"last_updated"`
}

type FlightPlan struct {
	FlightRules   string `json:"flight_rules"`
	Aircraft      string `json:"aircraft"`
	AircraftFAA   string `json:"aircraft_faa"`
	AircraftShort string `json:"aircraft_short"`
	Departure     string `json:"departure"`
	Arrival       string `json:"arrival"`
	Alternate     string `json:"alternate"`
	Cruise        string `json:"cruise_tas"`
	Altitude      string `json:"altitude"`
	DepTime       string `json:"deptime"`
	EnrouteTime   string `json:"enroute_time"`
	FuelTime      string `json:"fuel_time"`
	Remarks       string `json:"remarks"`
	Route         string `json:"route"`
	RevisionID    int    `json:"revision_id"`
}

type Controller struct {
	CID         int       `json:"cid"`
	Name        string    `json:"name"`
	Callsign    string    `json:"callsign"`
	Frequency   string    `json:"frequency"`
	Facility    int       `json:"facility"`
	Rating      int       `json:"rating"`
	Server      string    `json:"server"`
	VisualRange int       `json:"visual_range"`
	TextATIS    []string  `json:"text_atis"`
	LogonTime   time.Time `json:"logon_time"`
	LastUpdated time.Time `json:"last_updated"`
}

type ATIS struct {
	CID         int       `json:"cid"`
	Name        string    `json:"name"`
	Callsign    string    `json:"callsign"`
	Frequency   string    `json:"frequency"`
	Facility    int       `json:"facility"`
	Rating      int       `json:"rating"`
	Server      string    `json:"server"`
	VisualRange int       `json:"visual_range"`
	ATISCode    string    `json:"atis_code"`
	TextATIS    []string  `json:"text_atis"`
	LogonTime   time.Time `json:"logon_time"`
	LastUpdated time.Time `json:"last_updated"`
}

type Server struct {
	Ident                   string `json:"ident"`
	Hostname                string `json:"hsotname_or_ip"`
	Location                string `json:"location"`
	Name                    string `json:"name"`
	ClientConnectionAllowed int    `json:"client_connection_allowed"`
}

type Prefiled struct {
	CID         int        `json:"cid"`
	Name        string     `json:"name"`
	Callsign    string     `json:"callsign"`
	FlightPlan  FlightPlan `json:"flight_plan"`
	LastUpdated time.Time  `json:"last_updated"`
}

type Facilities struct {
	ID    int    `json:"id"`
	Short string `json:"short"`
	Long  string `json:"long"`
}

type Ratings struct {
	ID    int    `json:"id"`
	Short string `json:"short"`
	Long  string `json:"long"`
}

type PilotRatings struct {
	ID    int    `json:"id"`
	Short string `json:"short"`
	Long  string `json:"long"`
}
