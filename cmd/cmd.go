package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/philistino/lapjv"
	"github.com/spf13/cobra"
)

// FillType in an alias used to identify the way we want to fill our matrix in the generator.
type FillType int

const (
	// Random is a FillType in which we use rand.Intn(MaxValue) to fill the matrix.
	Random FillType = iota
	// Constant is a FillType in which we use i*j to fill the matrix.
	Constant FillType = iota
)

var (
	filename    string
	size        int
	constness   string
	interactive bool
)

// RootCmd is the main command displayed by Cobra with no argument
var RootCmd = &cobra.Command{
	Use:   "LAPJV Algorithm - Golang implementation",
	Short: "An implementation of the LAPJV Algorithm working in Golang.",
	Long:  "The Linear Assignment Problem Solver by Jonker-Volgenant Algorithm - with benchmark and test.",
}

var generatorCmd = &cobra.Command{
	Use:   "generator",
	Short: "Generate a JSON file that describes the matrix with given parameters.",
	Long:  "Use this command and generate a JSON file that describes the matrix you want to resolve - will be saved in 'resources' folder.",
	RunE:  runGenerator,
}

// runGenerator function will be called in order to generate a matrix and save it in a file.
// This function create the file and run a function between MatrixGeneratorInteractive and MatrixGeneratorManual following the CLI flags.
func runGenerator(cmd *cobra.Command, args []string) error {
	if filename == "" {
		filename = "example.json"
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	var m *MatrixGenerator

	if interactive {
		var err error
		m, err = NewInteractiveMatrixGenerator()
		if err != nil {
			return err
		}
	} else {
		t := Random
		if constness == "constant" {
			t = Constant
		}

		m = NewManualMatrixGenerator(size, t)
	}
	m.Run()
	return m.Save(f)
}

var solverCmd = &cobra.Command{
	Use:   "solver",
	Short: "Solve a matrix described in the JSON file given as parameter",
	Long:  "Use this command to solve a matrix you described in the JSON file before. Response will be printed in stdout",
	RunE:  runSolver,
}

// runSolver function will be called in order to solve the matrix using a file or a generated matrix using the generator.
// This function opens the file and reads the content. Once this step done, it will calls the MatrixSolver itself.
func runSolver(cmd *cobra.Command, args []string) error {
	var matrix [][]int

	if filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		e := json.NewDecoder(f)
		if err := e.Decode(&matrix); err != nil {
			return err
		}
	} else if interactive {
		var err error
		m, err := NewInteractiveMatrixGenerator()
		if err != nil {
			return err
		}
		m.Run()
		matrix = m.Matrix
	} else if cmd.Flag("size").Changed {
		t := Random
		if constness == "constant" {
			t = Constant
		}

		m := NewManualMatrixGenerator(size, t)
		m.Run()
		matrix = m.Matrix
	} else {
		// Stdin will be used.
		e := json.NewDecoder(os.Stdin)
		if err := e.Decode(&matrix); err != nil {
			return err
		}
	}

	s := lapjv.Lapjv(matrix)
	fmt.Printf("Matrix resolution : \n\tCost : \t\t%f\n\tRow solution : \t%v\n\tCol solution : \t%v\n", s.Cost, s.InRow, s.InCol)

	return nil
}

func init() {
	// Set flags to the program CLI Commands
	RootCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "file in which the matrix will be stored")
	RootCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Set the value to true in order to run the generator in interactive mode")
	RootCmd.PersistentFlags().StringVarP(&constness, "type", "t", "random", "Set the value to 'constant' in order to fill the matrix with constant values (choices are random / constant)")
	RootCmd.PersistentFlags().IntVarP(&size, "size", "s", 10, "size of the matrix.")

	// Add commands to the program CLI
	RootCmd.AddCommand(generatorCmd, solverCmd)
}

// Function used by Cobra to execute the command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
