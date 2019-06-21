package server

import "encoding/json"

// Options represents parameters that are passed to the application to be used in constructing
// the run and the server (if server mode is indicated).
type Options struct {
	Name        string   `json:"name"`           // The name of the server.
	Hostname    string   `json:"hostname"`       // The hostname of the server.
	Port        int      `json:"port"`           // The default port of the server.
	ProfPort    int      `json:"profPort"`       // The profiler port of the server.
	MaxConn     int      `json:"maxConnections"` // The maximum concurrent connections accepted.
	MaxProcs    int      `json:"maxProcs"`       // The maximum number of processor cores available.
	Debug       bool     `json:"debugEnabled"`   // Is debugging enabled in the application or server.
	ValidTokens []string `json:"-"`              // Valid tokens to access this server.
	Namespace   string   `json:"namespace"`      // Namespace to process request image into.
	AlivePath   string   `json:"alivePath"`      // Path to set within server for handling health checks.
	NotifyPath  string   `json:"notifyPath"`     // Path to set within server for handling duckerhub requests.
	StatusPath  string   `json:"statusPath"`     // Path to set within server for handling status requests
	TargetHost  string   `json:"targetHost"`     // Host redirects are sent forward.
	TargetPort  int      `json:"targetPort"`     // Port of host redirects.
	TargetPath  string   `json:"targetPath"`     // Path of action for redirects.
	TargetToken string   `json:"-"`              // Token for redirects server.
}

func OptionsNew(options ...func(*Options)) *Options {
	o := &Options{
		Name:        ApplicationName,
		Hostname:    DefaultHostname,
		Port:        DefaultPort,
		ProfPort:    DefaultProfPort,
		MaxConn:     DefaultMaxConnections,
		MaxProcs:    DefaultMaxProcs,
		Debug:       false,
		ValidTokens: make([]string, 0),
		Namespace:   DefaultNamespace,
		AlivePath:   DefaultAlivePath,
		NotifyPath:  DefaultNotifyPath,
		StatusPath:  DefaultStatusPath,
		TargetHost:  DefaultTargetHost,
		TargetPort:  DefaultTargetPort,
		TargetPath:  DefaultTargetPath,
		TargetToken: DefaultTargetToken,
	}
	for _, option := range options {
		option(o)
	}
	return o
}

// String is an implentation of the Stringer interface so the structure is returned as a string
// to fmt.Print() etc.
func (o *Options) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}
