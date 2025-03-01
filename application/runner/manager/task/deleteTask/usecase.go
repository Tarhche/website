package deletetask

import (
	"github.com/khanzadimahdi/testproject/domain"
	"github.com/khanzadimahdi/testproject/domain/runner/task"
	"github.com/khanzadimahdi/testproject/domain/translator"
)

type UseCase struct {
	taskRepository task.Repository
	translator     translator.Translator
}

func NewUseCase(taskRepository task.Repository, translator translator.Translator) *UseCase {
	return &UseCase{
		taskRepository: taskRepository,
		translator:     translator,
	}
}

func (uc *UseCase) Execute(request *Request) (*Response, error) {
	t, err := uc.taskRepository.GetOne(request.UUID)
	if err != nil {
		return nil, err
	}

	if !task.IsTerminalState(t.State) {
		return &Response{
			ValidationErrors: domain.ValidationErrors{
				"task_id": uc.translator.Translate("task_is_not_terminal_state"),
			},
		}, nil
	}

	return nil, uc.taskRepository.Delete(request.UUID)
}
