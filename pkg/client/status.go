package client

type Status struct {
	Data  DataStatus `json:"data"`
	User  []string   `json:"user"`
	Metar []string   `json:"metar"`
}

type DataStatus struct {
	V3           []string `json:"v3"`
	Transceivers []string `json:"transceivers"`
}

func GetStatus() (Status, error) {
	return Status{}, nil
}
