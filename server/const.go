package server

import "time"

const (
	ApplicationName       = "dockerhub-webhook"                // Application name.
	version               = "1.0.0"                            // Application version.
	DefaultHostname       = "0.0.0.0"                          // The hostname or address of the server.
	DefaultPort           = 8080                               // Port to receive requests: see IANA Port Numbers.
	DefaultProfPort       = 6060                               // Profiler port to receive requests.*
	DefaultMaxConnections = 0                                  // Maximum number of connections allowed.*
	DefaultMaxProcs       = 0                                  // Maximum number of computer processors to utilize.*
	DefaultNamespace      = "development"                      // Default namespace to process notification into.
	DefaultAlivePath      = "/v1.0/alive"                      // Default path to check health of server.
	DefaultNotifyPath     = "/v1.0/notify"                     // Default path to handle notify events.
	DefaultStatusPath     = "/v1.0/status"                     // Default path to check server status information.
	DefaultTargetHost     = "jenkins"                          // Server hostname or IP to forward to.
	DefaultTargetPort     = 8080                               // Server port to forward to.
	DefaultTargetPath     = "/generic-webhook-trigger/invoke/" // Path to generic jenkins webhook
	DefaultTargetToken    = ""                                 // Security token to send to target

	// * zeros = no change or no limitations or not enabled.

	// Listener and connections.
	TCPKeepAliveTimeout = 3 * time.Minute
	TCPReadTimeout      = 10 * time.Second
	TCPWriteTimeout     = 10 * time.Second

	httpGet    = "GET"
	httpPost   = "POST"
	httpPut    = "PUT"
	httpDelete = "DELETE"
	httpHead   = "HEAD"
	httpTrace  = "TRACE"
	httpPatch  = "PATCH"

	// Error messages.
	InvalidMediaType     = "Invalid Content-Type or Accept header value."
	InvalidMethod        = "Invalid Method for this route."
	InvalidBody          = "Invalid body of text in request."
	InvalidJSONText      = "Invalid JSON format in text of body in request."
	CouldNotProcess      = "Could not process this request."
	InvalidAuthorization = "Invalid authorization."
)
