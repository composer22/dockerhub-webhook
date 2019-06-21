package server

import "encoding/json"

// Info provides basic information about the running server.
type Info struct {
	Version    string `json:"version"`        // Version of the server.
	Name       string `json:"name"`           // The name of the server.
	Hostname   string `json:"hostname"`       // The hostname of the server.
	UUID       string `json:"UUID"`           // Unique ID of the server.
	Port       int    `json:"port"`           // Port the server is listening on.
	ProfPort   int    `json:"profPort"`       // Profiler port the server is listening on.
	MaxConn    int    `json:"maxConnections"` // The maximum concurrent connections accepted.
	Debug      bool   `json:"debugEnabled"`   // Is debugging enabled on the server.
	Namespace  string `json:"namespace"`      // Namespace to process image.
	AlivePath  string `json:"alivePath"`      // Path to set within server for handling health checks.
	NotifyPath string `json:"notifyPath"`     // Path to set within server for handling duckerhub requests.
	StatusPath string `json:"statusPath"`     // Path to set within server for handling status requests
	TargetHost string `json:"targetHost"`     // Host redirects are sent forward.
	TargetPort int    `json:"targetPort"`     // Port of host redirects.
	TargetPath string `json:"targetPath"`     // Path of action for redirects.
}

// InfoNew is a factory function that returns a new instance of Info.
// options is an optional list of functions that initialize the structure
func InfoNew(options ...func(*Info)) *Info {
	inf := &Info{
		Version: version,
		UUID:    createV4UUID(),
	}
	for _, option := range options {
		option(inf)
	}
	return inf
}

// String is an implentation of the Stringer interface so the structure is returned as a
// string to fmt.Print() etc.
func (i *Info) String() string {
	b, _ := json.Marshal(i)
	return string(b)
}
