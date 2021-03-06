package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/DragosPancescu/SD-Tema1/client"
	"github.com/DragosPancescu/SD-Tema1/utils"
)

// Parses a message sent from a user to the server
// It takes the first element as the command name
// The rest of the elements as the arguments
func Parse_message(message string) (string, []string) {

	// Remove newline, carriage return and accidental space characters
	message = strings.TrimRight(message, " \r\n")

	message_list := strings.Split(message, " ")

	return message_list[0], message_list[1:]
}

// ---------------------------------------COMMANDS----------------------------------------------//

// Send the help panel to the user
func command_help(file_path string) string {

	fmt.Println("The server is processing the data...")
	// Opens the help file
	help_file, err := os.Open(file_path)

	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	defer help_file.Close()

	// Reads the file conntents
	file_contents, _ := ioutil.ReadAll(help_file)

	return string(file_contents)
}

// Requirement 1 from the homework
func command1(input []string) ([]string, bool) {

	// Check if elements have the same length
	if !utils.Check_elem_len(input) {
		return nil, false
	}

	var output []string

	for i := 0; i < len(input[0]); i++ {
		var aux_string string
		for j := 0; j < len(input); j++ {
			aux_string += string(input[j][i])
		}
		output = append(output, aux_string)
	}

	return output, true
}

// Requirement 2 from the homework
func command2(input []string) int {
	counter := 0

	for i := 0; i < len(input); i++ {
		// Extract the number from the string element
		number, extracted_any := utils.Extract_number(input[i])

		// Check if it is a perfect square
		if extracted_any {
			if utils.Check_perfect_square(float64(number)) {
				counter++
			}
		}
	}

	return counter
}

// Requirement 3 from the homework
func command3(input []string) (int, bool) {
	output := 0

	for i := 0; i < len(input); i++ {
		reversed_number, good_format := utils.Reverse_number(input[i])

		if good_format {
			output += reversed_number
		} else {
			return -1, false
		}
	}

	return output, true
}

// Requirement from the homework
func command4(input []string) (float64, bool) {

	// Checks the first 2 arguments
	a, err_a := strconv.Atoi(input[0])
	b, err_b := strconv.Atoi(input[1])

	if err_a != nil || err_b != nil {
		return -1, false
	}

	sum := 0
	counter := 0
	for i := 2; i < len(input); i++ {
		number, err := strconv.Atoi(input[i])
		if err != nil {
			return -1, false
		}

		if utils.Digits_sum(number) >= a && utils.Digits_sum(number) <= b {
			sum += number
			counter++
		}
	}

	if counter == 0 {
		return 0, true
	}
	return float64(sum / counter), true
}

// Requirement 5 from homework
func command5(input []string) []string {
	var output []string

	for i := 0; i < len(input); i++ {
		if utils.Check_binary(input[i]) {
			decimal, err := strconv.ParseInt(input[i], 2, 64)
			if err != nil {
				fmt.Println("Error: " + err.Error())
				return nil
			}

			output = append(output, strconv.Itoa(int(decimal)))
		}
	}

	return output
}

// ---------------------------------------COMMANDS HANDLERS----------------------------------------------//
// All of this might need refactoring as there is a lot of duplicate code
// Did not have the time tho :(
// TODO: Refactor this

func Handle_help_command(client client.Client, file_path string) {
	fmt.Println(Color_string((client.Name + " sent a help request."), client.Color))
	help_string := command_help(file_path)

	fmt.Println(Color_string(("Server is sending a response to " + client.Name + "."), client.Color))
	client.Connection.Write([]byte(Color_string(help_string, client.Color)))
}

func Handle_exit_command(client client.Client) {
	client.Connection.Write([]byte(Color_string(("See you soon " + client.Name + ".\n"), client.Color)))
	fmt.Println(Color_string((client.Name + " is leaving the server."), client.Color))
	client.Connection.Close()
}

func Handle_command1(client client.Client, args []string) {
	fmt.Println(Color_string((client.Name + " sent a command1 request."), client.Color))

	if len(args) < 1 {
		client.Connection.Write([]byte(Color_string("No arguments provided, please try again.\n", client.Color)))
		fmt.Println(Color_string((client.Name + " - the command1 request was invalid."), client.Color))
	} else {
		output, good_format := command1(args)

		fmt.Println(Color_string(("Server is sending a response to " + client.Name + "."), client.Color))
		if !good_format {
			client.Connection.Write([]byte(Color_string("Incorrect data format, please try again.\n", client.Color)))
		} else {
			client.Connection.Write([]byte(Color_string("Response: "+strings.Join(output, " ")+"\n", client.Color)))
		}
	}
}

func Handle_command2(client client.Client, args []string) {
	fmt.Println(Color_string((client.Name + " sent a command2 request."), client.Color))

	if len(args) < 1 {
		client.Connection.Write([]byte(Color_string("No arguments provided, please try again.\n", client.Color)))
		fmt.Println(Color_string((client.Name + " - the command1 request was invalid."), client.Color))
	} else {
		output := command2(args)

		fmt.Println(Color_string(("Server is sending a response to " + client.Name + "."), client.Color))
		client.Connection.Write([]byte(Color_string("Response: "+strconv.Itoa(output)+"\n", client.Color)))
	}
}

func Handle_command3(client client.Client, args []string) {
	fmt.Println(Color_string((client.Name + " sent a command3 request."), client.Color))

	if len(args) < 1 {
		client.Connection.Write([]byte(Color_string("No arguments provided, please try again.\n", client.Color)))
		fmt.Println(Color_string((client.Name + " - the command1 request was invalid."), client.Color))
	} else {
		output, good_format := command3(args)

		fmt.Println(Color_string(("Server is sending a response to " + client.Name + "."), client.Color))
		if !good_format {
			client.Connection.Write([]byte(Color_string("Incorrect data format, please try again.\n", client.Color)))
		} else {
			client.Connection.Write([]byte(Color_string("Response: "+strconv.Itoa(output)+"\n", client.Color)))
		}
	}
}

func Handle_command4(client client.Client, args []string) {
	fmt.Println(Color_string((client.Name + " sent a command4 request."), client.Color))

	if len(args) < 2 {
		client.Connection.Write([]byte(Color_string("The command need at least 2 arguments\n", client.Color)))
		fmt.Println(Color_string((client.Name + " - the command1 request was invalid."), client.Color))
	} else {
		output, good_format := command4(args)

		fmt.Println(Color_string(("Server is sending a response to " + client.Name + "."), client.Color))
		if !good_format {
			client.Connection.Write([]byte(Color_string("Incorrect data format, please try again.\n", client.Color)))
		} else {
			client.Connection.Write([]byte(Color_string("Response: "+fmt.Sprint(output)+"\n", client.Color)))
		}
	}
}

func Handle_command5(client client.Client, args []string) {
	fmt.Println(Color_string((client.Name + " sent a command1 request."), client.Color))

	if len(args) < 1 {
		client.Connection.Write([]byte(Color_string("No arguments provided, please try again.\n", client.Color)))
		fmt.Println(Color_string((client.Name + " - the command5 request was invalid."), client.Color))
	} else {
		output := command5(args)

		fmt.Println(Color_string(("Server is sending a response to " + client.Name + "."), client.Color))
		client.Connection.Write([]byte(Color_string("Response: "+strings.Join(output, " ")+"\n", client.Color)))
	}
}

func Handle_unknown_command(client client.Client, args []string) {
	fmt.Println(Color_string((client.Name + " sent an unknown request."), client.Color))
	client.Connection.Write([]byte(Color_string("That request is not valid, please try again.\n", client.Color)))
}
