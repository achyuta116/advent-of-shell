package solutions

import (
	"fmt"
	"os"
	"os/exec"
)

// Return a string solution from a function
// The string solution is generated from the input

func GenerateSolution(problem, subproblem, year, seed uint) (string, error) {
	cmd := exec.Command(fmt.Sprintf("./problem_%d_%d_%d.sh", problem, subproblem, year))
	cmd.Env = append(os.Environ(), fmt.Sprintf("SEED_VALUE=%d", seed))

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running executable for problem: %d_%d, seed: %d, year: %d\n", problem, subproblem, seed, year)
		return "", err
	}

	return string(output), nil
}

func GetSolution(problem uint, subproblem uint, year uint, seed uint) (string, error) {
	return GenerateSolution(problem, subproblem, year, seed)
}
