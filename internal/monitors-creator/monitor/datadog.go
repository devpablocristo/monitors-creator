package monitor

import (
	"time"
)

type DatadogResponse []struct {
	ID      int      `json:"id"`
	OrgID   int      `json:"org_id"`
	Type    string   `json:"type"`
	Name    string   `json:"name"`
	Message string   `json:"message"`
	Tags    []string `json:"tags"`
	Query   string   `json:"query"`
	Options struct {
		Thresholds struct {
			Critical float64 `json:"critical"`
			Warning  float64 `json:"warning"`
		} `json:"thresholds"`
		NotifyAudit   bool     `json:"notify_audit"`
		IncludeTags   bool     `json:"include_tags"`
		NotifyBy      []string `json:"notify_by"`
		NewGroupDelay int      `json:"new_group_delay"`
		NotifyNoData  bool     `json:"notify_no_data"`
		Silenced      struct {
		} `json:"silenced"`
	} `json:"options"`
	Multi                bool        `json:"multi"`
	CreatedAt            int64       `json:"created_at"`
	Created              time.Time   `json:"created"`
	Modified             time.Time   `json:"modified"`
	Deleted              interface{} `json:"deleted"`
	RestrictedRoles      interface{} `json:"restricted_roles"`
	Priority             int         `json:"priority"`
	OverallStateModified time.Time   `json:"overall_state_modified"`
	OverallState         string      `json:"overall_state"`
	State                struct {
		Groups struct {
			HostColima struct {
				Name            string      `json:"name"`
				Status          string      `json:"status"`
				LastTriggeredTs int         `json:"last_triggered_ts"`
				LastNotifiedTs  interface{} `json:"last_notified_ts"`
				LastNodataTs    interface{} `json:"last_nodata_ts"`
				LastResolvedTs  int         `json:"last_resolved_ts"`
			} `json:"host:colima"`
		} `json:"groups"`
	} `json:"state"`
	Creator struct {
		Name   string `json:"name"`
		Email  string `json:"email"`
		Handle string `json:"handle"`
		ID     int    `json:"id"`
	} `json:"creator"`
	MatchingDowntimes []interface{} `json:"matching_downtimes"`
}

func GetAllMonitors() (DatadogResponse, error) {
	return nil, nil
}