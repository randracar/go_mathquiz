package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)

var played_questions uint = 0
var points uint = 0

func process_difficulty(dif int){
	if dif == 1 {
		generate_questions_easy()
	} else if dif == 2 { 
		generate_questions_medium()
	} else if dif == 3 {
		generate_questions_hard()
	}
}

func generate_questions_easy(){
	fmt.Println("Generating easy questions for you...")
	for played_questions < 10 {
		rand.Seed(time.Now().UnixNano())
		var operation int = rand.Intn(3) + 1
		if operation == 1 {
			rand.Seed(time.Now().UnixNano())
			var num1 int = rand.Intn(100)+1
			var num2 int = rand.Intn(100)+1
			var result int = num1 * num2
			var user_input int

			fmt.Printf("What is the result of %d * %d?: ", num1, num2)
			fmt.Scanln(&user_input)
			if user_input == result{
				points = points + 2
				fmt.Println("Correct! Adding 2 points, your current score: ", points)
			} else {
				fmt.Println("Wrong :(. Your score remains the same: ", points)
			}
			played_questions++
		} else if operation == 2 {
			rand.Seed(time.Now().UnixNano())
			var num1 int = rand.Intn(100)+1
			var num2 int = rand.Intn(100)+1
			var result int = num1 + num2
			var user_input int

			fmt.Printf("What is the result of %d + %d?: ", num1, num2)
			fmt.Scanln(&user_input)
			if user_input == result{
				points = points + 2
				fmt.Println("Correct! Adding 2 points, your current score: ", points)
			} else {
				fmt.Println("Wrong :(. Your score remains the same: ", points)
			}
			played_questions++
		} else if operation == 3 {
			rand.Seed(time.Now().UnixNano())
			var num1 int = rand.Intn(100)+1
			var num2 int = rand.Intn(100)+1
			var result int = num1 - num2
			var user_input int

			fmt.Printf("What is the result of %d - %d?: ", num1, num2)
			fmt.Scanln(&user_input)
			if user_input == result{
				points = points + 2
				fmt.Println("Correct! Adding 2 points, your current score: ", points)
			} else {
				fmt.Println("Wrong :(. Your score remains the same: ", points)
			}
			played_questions++
		}
	}
	process_points(points)
}


func generate_questions_medium(){
	fmt.Println("Generating medium questions for you...")
	for played_questions < 10 {
		rand.Seed(time.Now().UnixNano())
		var operation int = rand.Intn(3) + 1
		if operation == 1 { 
			rand.Seed(time.Now().UnixNano())
			var num1 int = rand.Intn(10)+1
			var num2 int = rand.Intn(10)+1
			var result int = num1 / num2
			var user_input int

			fmt.Printf("What is the result of %d / %d?: ", num1, num2)
			fmt.Scanln(&user_input)

			if user_input == result {
				points = points + 2
				fmt.Println("Correct! Adding 2 points, your current score: ", points)
			} else {
				fmt.Println("Wrong :(. Your score remains the same: ", points)
			}
		} else if operation == 2 {
			rand.Seed(time.Now().UnixNano())
			var num1 int = rand.Intn(10)+1
			var num2 int = rand.Intn(10)+1
			var result int = num1 * num2
			var user_input int

			fmt.Printf("What is the result of %d * %d?: ", num1, num2)
			fmt.Scanln(&user_input)
			
			if user_input == result {
				points = points + 2
				fmt.Println("Correct! Adding 2 points, your current score: ", points)
			} else {
				fmt.Println("Wrong :(. Your score remains the same: ", points)
			} 
		} else if operation == 3 {
			rand.Seed(time.Now().UnixNano())
			var num1int int = rand.Intn(10)+1
			var num2int int = 2
			var num1 float64 = float64(num1int)
			var num2 float64 = float64(num2int)
			var result float64 = math.Pow(num1, num2)
			var user_input_int int

			fmt.Printf("What is the result of %f ^ 2?: ", num1)
			fmt.Scanln(&user_input_int)

			var user_input float64 = float64(user_input_int)
			
			if user_input == result {
				points = points + 2
				fmt.Println("Correct! Adding 2 points, your current score: ", points)
			} else {
				fmt.Println("Wrong :(. Your score remains the same: ", points)
			}
		}
		played_questions++
	}
	process_points(points)
}

func generate_questions_hard(){
	fmt.Println("Generating hard questions for you...")
	for played_questions < 10 {
		rand.Seed(time.Now().UnixNano())
		var operation int = rand.Intn(3)+1
		if operation == 1 {
			rand.Seed(time.Now().UnixNano())
			var num1 int = rand.Intn(10)+1
			var num2 int = rand.Intn(10)+1
			var result int = num1 / num2
			var user_input int

			fmt.Printf("What is the result of %d / %d?: ", num1, num2)
			fmt.Scanln(&user_input)

			if user_input == result {
				points = points + 2
				fmt.Println("Correct! Adding 2 points, your current score: ", points)
			} else {
				fmt.Println("Wrong :(. Your score remains the same: ", points)
			}
		} else if operation == 2 {
			rand.Seed(time.Now().UnixNano())
			var num1int int = rand.Intn(10)+1
			var num2int int = rand.Intn(10)+1
			var num1 float64 = float64(num1int)
			var num2 float64 = float64(num2int)
			var result float64 = math.Pow(num1, num2)
			var user_input_int int

			fmt.Printf("What is the result of %f ^ %f?: ", num1, num2)
			fmt.Scanln(&user_input_int)

			var user_input float64 = float64(user_input_int)
			
			if user_input == result {
				points = points + 2
				fmt.Println("Correct! Adding 2 points, your current score: ", points)
			} else {
				fmt.Println("Wrong :(. Your score remains the same: ", points)
			}
		} else if operation == 3 {
			rand.Seed(time.Now().UnixNano())
			var num1int int = rand.Intn(10)+1
			var num2int int = 2
			var num1 float64 = float64(num1int)
			var num2 float64 = float64(num2int)
			var result float64 = math.Pow(num1, num2)
			var user_input_int int

			fmt.Printf("What is the result of %f ^ 2?: ", num1)
			fmt.Scanln(&user_input_int)

			var user_input float64 = float64(user_input_int)

			if user_input == result {
				points = points + 2
				fmt.Println("Correct! Adding 2 points, your current score: ", points)
			} else {
				fmt.Println("Wrong :(. Your score remains the same: ", points)
			}
		}
	}
	process_points(points)
}

func process_points(score uint){
	fmt.Printf("Your total score was: %d\n", score)
	if score == 20 {
		fmt.Println("You got the full score! Congrats :)")
	} else if score >= 18 && score < 20 {
		fmt.Println("You almost got the full score! Anyways, that's still pretty good! Good job! :)")
	} else if score >= 10 && score < 18 {
		fmt.Println("You did a good job, but maybe if you try a little bit harder you can do better! :)")
	} else if score < 10 {
		fmt.Println("You didn't do great, but don't lose hope, try harder! :)")
	} else {
		fmt.Println("Error in the score")
	}
}

func main(){

	var name string
	var dif int

	fmt.Println("Welcome to my math quiz game!")
	fmt.Printf("Please enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello %s! Which gamemode would you like to play?\n", name)
	fmt.Println("1 - Easy")
	fmt.Println("2 - Medium")
	fmt.Println("3 - Hard")

	fmt.Printf("Please enter your choice: ")

	fmt.Scanln(&dif)

	process_difficulty(dif)
}