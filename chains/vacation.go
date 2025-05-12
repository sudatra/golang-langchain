package chains

import (
	"log"
	"github.com/google/uuid"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

var Vacations []*Vacation

func GetVacationFromDB(id uuid.UUID) (Vacation, error) {

}

func GenerateVacationIdeaChange(id uuid.UUID, budget int, season string, hobbies []string) {
	log.Printf("Generating New Vacation ID: %d", id);

	v := &Vacation{Id: id, Completed: false, Idea: ""};
	Vacations = append(Vacations, v);
}