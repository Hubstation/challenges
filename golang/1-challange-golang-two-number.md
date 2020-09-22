
# 1. Two Sum

## topic

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

Problem solving ideas
```
a + b = target
```
Can also be seen as
```
a = target - b
```
In map[Integer] The serial number of the integer, you can find the serial number of a. This eliminates the need to nest two for loops.
