// Package datasource provides the core interface and types for implementing
// external data source integrations with Locus.
//
// A DataSource represents any external API or service that can provide topics
// (questions, articles, etc.) and associated data (answers, content) based on
// search queries. Common examples include Stack Exchange, Wikipedia, YouTube,
// or any custom knowledge base.
package datasource

// DataSource defines the contract for integrating external data sources.
// Implementations should handle API communication, rate limiting, and error
// handling internally.
type DataSource interface {
	// Init performs any heavy initialization required by the data source,
	// such as fetching configuration, warming caches, or establishing
	// persistent connections. This is called once at startup.
	//
	// Returns an error if initialization fails.
	Init() error

	// CheckAvailability performs a lightweight health check to verify the
	// external source is reachable and responsive. This should complete
	// quickly (typically under 5 seconds).
	//
	// Returns true if the service is available, false otherwise.
	CheckAvailability() bool

	// FetchTopics searches for relevant topics based on the provided query.
	// Topics represent high-level items like questions, articles, or videos
	// that may contain relevant data.
	//
	// Parameters:
	//   - count: Maximum number of topics to return (sorted by relevance)
	//   - input: Search query including question text and optional embedding
	//
	// Returns a slice of topics and an error if the operation fails.
	// An empty slice with nil error indicates no results were found.
	FetchTopics(count int, input NewQuestionInput) ([]DataSourceTopic, error)

	// FetchData retrieves detailed data items for a specific topic.
	// Data items typically represent answers, excerpts, or content chunks
	// associated with the topic.
	//
	// Parameters:
	//   - count: Maximum number of data items to return (sorted by relevance/votes)
	//   - topicID: Identifier of the topic to fetch data for
	//
	// Returns a slice of data items and an error if the operation fails.
	// An empty slice with nil error indicates the topic has no data.
	FetchData(count int, topicID int64) ([]DataSourceData, error)
}

// DataSourceTopic represents a high-level item from an external source that
// may contain relevant information (e.g., a question, article, or video).
type DataSourceTopic struct {
	// Topic is the title or main text of the topic
	Topic string `json:"topic"`

	// SourceURL is the canonical URL where this topic can be viewed
	SourceURL string `json:"source_url"`

	// Site identifies the specific site or subsection if the data source
	// supports multiple sites (e.g., "stackoverflow", "serverfault")
	// Optional - may be empty for single-site sources
	Site string `json:"site,omitempty"`

	// TopicID is the unique identifier for this topic in the external system
	// Used when calling FetchData to retrieve associated content
	TopicID int64 `json:"topic_id"`
}

// DataSourceData represents a specific piece of content associated with a topic
// (e.g., an answer to a question, a section of an article, or a transcript).
type DataSourceData struct {
	// DataText is the actual content text (may include HTML or markdown)
	DataText string `json:"data_text"`

	// SourceURL is the canonical URL where this specific data can be viewed
	SourceURL string `json:"source_url"`

	// Site identifies the specific site or subsection if applicable
	// Optional - may be empty for single-site sources
	Site string `json:"site,omitempty"`

	// AnswerID is the unique identifier for this data item in the external system
	// The name "AnswerID" is used for historical reasons but represents any
	// data item identifier (answer, excerpt, etc.)
	AnswerID int64 `json:"answer_id"`
}

// NewQuestionInput provides context for searching topics in a data source.
// It includes both the search query text and optional semantic information
// that advanced data sources can use for better matching.
type NewQuestionInput struct {
	// QuestionText is the search query or question being asked
	QuestionText string

	// Tags are optional topic tags that may help narrow the search
	Tags []string

	// AskedBy is the optional user ID of who is asking the question
	// May be nil if the query is anonymous
	AskedBy *int64

	// Embedding is an optional precomputed vector representation of the question
	// Advanced data sources can use this for semantic search or similarity matching
	// If nil or empty, the data source should fall back to text-based search
	Embedding []float64
}
