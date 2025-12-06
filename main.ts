/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-06
 * @fileoverview This program asks the user for a number of marks,then use loop to read in each number and then calculate the average
 */

// get user input
let numMarks: number = Number(prompt("How many marks will you enter today? "));

// set variables
let total: number = 0;
let mark: number = 0;

// read in each mark
for (let i: number = 1; i <= numMarks; i = i + 1) {
  mark = Number(prompt("Enter mark " + i + ": "));
  total = total + mark;
}

// calculate average
let average: number = total / numMarks;

// display results
console.log("You have entered " + numMarks + " marks. The student's average is " + average.toFixed(1) + "%.");

// give feedback
if (average <= 49) {
  console.log("The student is failing.");
} else if (average >= 50 && average <= 69) {
  console.log("The student's performance is just under average.");
} else if (average >= 70 && average <= 79) {
  console.log("The student's performance is average.");
} else if (average >= 80) {
  console.log("The student is on the honour roll.");
}