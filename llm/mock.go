package llm

import (
	"github.com/jwhenry28/LLMUtils/model"
)

type MockLLM struct {
	messages []model.Chat
}

func NewMockLLM() *MockLLM {
	return &MockLLM{}
}

func (llm *MockLLM) AddMessage(message model.Chat) {
	llm.messages = append(llm.messages, message)
}

func (llm *MockLLM) Type() string {
	return "mock"
}

func (llm *MockLLM) CompleteChat(_ []model.Chat) (string, error) {
	if len(llm.messages) == 0 {
		return "no messages available", nil
	}

	message := llm.messages[0]
	llm.messages = llm.messages[1:]
	return message.Content, nil
}
