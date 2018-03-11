package logging

// SourceKey is the map key used for logging sources (typically the name of the application; see NewLog).
const SourceKey = "_source"

// DomainKey is the map key used for logging domains. Domains divide a source into sub-sources (see Logger.NewDomainLogger).
const DomainKey = "_domain"

// HostKey is the map key used for logging the machine hostname (via os.Hostname; see NewLog).
const HostKey = "_host"
