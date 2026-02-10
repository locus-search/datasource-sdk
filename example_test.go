package datasource_test

import (
	"errors"
	"testing"

	datasource "github.com/locus-search/datasource-sdk"
)

// ExampleDataSource demonstrates a minimal implementation
type ExampleDataSource struct {
	Name string
}

func (ds *ExampleDataSource) Init() error {
	if ds.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (ds *ExampleDataSource) CheckAvailability() bool {
	return true
}

func (ds *ExampleDataSource) FetchTopics(count int, input datasource.NewQuestionInput) ([]datasource.DataSourceTopic, error) {
	if input.QuestionText == "" {
		return nil, errors.New("question text is required")
	}

	// Example: return a mock topic
	return []datasource.DataSourceTopic{
		{
			Topic:     "Example Topic for: " + input.QuestionText,
			SourceURL: "https://example.com/topic/1",
			TopicID:   1,
		},
	}, nil
}

func (ds *ExampleDataSource) FetchData(count int, topicID int64) ([]datasource.DataSourceData, error) {
	if topicID <= 0 {
		return nil, errors.New("invalid topic ID")
	}

	// Example: return mock data
	return []datasource.DataSourceData{
		{
			DataText:  "This is example data for the topic",
			SourceURL: "https://example.com/topic/1#data-1",
			AnswerID:  1,
		},
	}, nil
}

// Verify ExampleDataSource implements DataSource interface
var _ datasource.DataSource = (*ExampleDataSource)(nil)

func TestExampleDataSourceImplementation(t *testing.T) {
	ds := &ExampleDataSource{Name: "test"}

	// Test Init
	if err := ds.Init(); err != nil {
		t.Errorf("Init failed: %v", err)
	}

	// Test CheckAvailability
	if !ds.CheckAvailability() {
		t.Error("CheckAvailability should return true")
	}

	// Test FetchTopics
	input := datasource.NewQuestionInput{
		QuestionText: "test question",
	}
	topics, err := ds.FetchTopics(5, input)
	if err != nil {
		t.Errorf("FetchTopics failed: %v", err)
	}
	if len(topics) == 0 {
		t.Error("Expected at least one topic")
	}

	// Test FetchData
	data, err := ds.FetchData(3, 1)
	if err != nil {
		t.Errorf("FetchData failed: %v", err)
	}
	if len(data) == 0 {
		t.Error("Expected at least one data item")
	}
}

func TestExampleDataSourceValidation(t *testing.T) {
	ds := &ExampleDataSource{}

	// Test Init with missing name
	if err := ds.Init(); err == nil {
		t.Error("Expected error when name is missing")
	}

	ds.Name = "test"
	ds.Init()

	// Test FetchTopics with empty question
	_, err := ds.FetchTopics(5, datasource.NewQuestionInput{})
	if err == nil {
		t.Error("Expected error with empty question text")
	}

	// Test FetchData with invalid ID
	_, err = ds.FetchData(3, 0)
	if err == nil {
		t.Error("Expected error with invalid topic ID")
	}
}
