#Variables

echo "This is Variables part"
echo "--------------------------------------------------------------------"

name="Bhargav"
echo "Hello $name, welcome to DevOps!".

#Taking User Input

echo "Enter employee name"
read employeename

echo "Enter the competency of the empoloyee"
read competency

echo "Enter the company name"
read company

echo "==============================="

echo "Employee Name: $employeename"
echo "Competancy: $competency"
echo "Comapny: $company"

echo "==============================="



echo "This is Conditonal Statements Part"
echo "--------------------------------------------------------------------"

#!/bin/bash

echo "Enter Student Name"
read studentName

echo "Enter your score"
read score

if [ $score -ge 90 ]; then
    echo "$studentName got Grade A"
elif [ $score -ge 70 ] && [ $score -le 89 ]; then
    echo "$studentName got Grade B"
elif [ $score -ge 50 ] && [ $score -le 69 ]; then
    echo "$studentName got Grade C"
else
    echo "$studentName has failed in the examination"
fi



echo "This is Switch Case Part"
echo "--------------------------------------------------------------------"

echo "Enter Employee Name"
read employeeName
echo "Enter Employee Id"
read employeeId
echo "Choose any technology (DevOps, Java, Python, ASP.NET)"
read technology

case $technology in
"DevOps")
echo "$employeeName ($employeeId) is assigned to $technology"
;;
"Java")
echo "$employeeName ($employeeId) is assigned to $technology"
;;
"Python")
echo "$employeeName ($employeeId) is assigned to $technology"
;;
"ASP.NET")
echo "$employeeName ($employeeId) is assigned to $technology"
;;
*)
echo "Invalid choice choose from the menu"
;;
esac



echo "This is Loops Part"
echo "--------------------------------------------------------------------"
echo "Enter your name"
read name
echo "Enter how many times you want to repeate the name (enter in numbers (Example: 3)"
read num
for i in $(seq 1 $num)
do
echo "Your name: $name"
done
