package invgo

// ImplementedEndpoints when adding a new Invgate endpoint
// this must be updated so that the coverage script knows which
// endpoints Invgo supports calling.
var ImplementedEndpoints = map[string][]string{
	"/breakingnews":                   {"POST", "PUT", "GET"},
	"/breakingnews.all":               {"GET"},
	"/breakingnews.attributes.status": {"GET"},
	"/breakingnews.attributes.type":   {"GET"},
	"/breakingnews.status":            {"POST", "GET"},
	"/categories":                     {"GET"},
	"/helpesks":                       {"GET"},
	"/incident":                       {"POST", "PUT", "GET"},
	"/incident.attributes.status":     {"GET"},
	"/incident.attributes.type":       {"GET"},
	"/incidents":                      {"GET"},
	"/incidents.by.status":            {"GET"},
	"/sd.version":                     {"GET"},
}
