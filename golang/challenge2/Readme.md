# add two numbers 

topic

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their 
nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8

Problem solving ideas

(2 -> 4 -> 3)yes 342

(5 -> 6 -> 4)yes 465

(7 -> 0 -> 8)yes 807

342 + 465 = 807

Therefore, the original intention of the question is to realize its addition after changing an integer expression.

When designing a program, the points that need to be dealt with are
   - Addition on the bit, need to deal with the carry problem
   -  How to enter the next operation
   -  After the bitwise addition is over, the carry problem also needs to be dealt with.

to sum up

- After reading the meaning of the question, follow the steps to achieve the requirements of the question.
