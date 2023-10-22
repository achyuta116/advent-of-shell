package input

import (
	"fmt"
	"os"
	"os/exec"
)

// Generate inputs based on a seed value given
// Generate new input and store only if the input isn't present in the db
// for the given seed value and problem id
// Generate the input using different implementations each time
func GenerateInput(problem uint, seed uint, year uint) (string, error) {
	cmd := exec.Command(fmt.Sprintf("./problem_%d_%d", problem, year))
	cmd.Env = append(os.Environ(), fmt.Sprintf("SEED_VALUE=%d", seed))

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running executable for problem: %d, seed: %d, year: %d\n", problem, seed, year)
		return "", err
	}

	input := string(output)
	return input, nil
}

func GetInput(problem uint, seed uint, year uint) (string, error) {
    return GenerateInput(problem, seed, year)
}
