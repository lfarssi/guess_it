Number Range Predictor
This program predicts a range for the next number based on a sequence of input numbers. It takes each input number and calculates a range within which the next number is likely to fall, leveraging basic statistical methods.

Objectives
Implement a program that reads a series of numbers as standard input.
For each input number, compute a range (lower limit and upper limit) for the next number.
The program aims to balance between a small range (better score) and accurate predictions.
Features
Accepts a continuous stream of numbers as input.
Outputs a range for the next input after each number.
The size of the range affects the score, encouraging optimal guesses.
Installation
Ensure you have the required programming environment set up for the language you choose (Python, Go, JavaScript, or Rust).
Clone the repository or download the project files.
Usage
Create a folder named student.

Place your program file (e.g., solution.py, solution.go, etc.) inside the student folder.

Create a script file named script.sh in the root directory with the following content (modify according to your programming language):

sh
Copier le code
#!/bin/sh
# Navigate to the student folder and run the program
python3 ./student/solution.py  # Replace with your program's command
Make the script executable:

sh
Copier le code
chmod +x script.sh
Run the script from the root directory:

sh
Copier le code
./script.sh
Example Input/Output
Copier le code
189
120 200
113
160 230
121
110 140
114
100 200
145
1 99
110
100 150
Testing
This program will be tested extensively using different datasets (Data 1, Data 2, Data 3). Ensure your program handles a variety of input sequences and produces accurate range predictions.

Learning Outcomes
By completing this project, you will enhance your skills in:

Statistical analysis and probability calculations.
Scripting and automation.
Problem-solving in programming.
Contribution
Feel free to contribute to this project by improving the algorithm or adding features. Please create a fork and submit a pull request for any changes you wish to propose.

License
This project is open-source. Feel free to use and modify it as per your needs.

For any questions or support, please contact [medlfarssi10@gmail.com].