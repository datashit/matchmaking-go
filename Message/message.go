package matchmaking

// Request data struct
type Request struct {
	MessageType int
	UserID      string // User ID
	Command     string // Command
	Data        string // Command Data JSON
}

// Response data struct
type Response struct {
	MessageType int
	ServerID    string // Server ID
	Command     string // Command
	Data        string // Command Data JSON
}
