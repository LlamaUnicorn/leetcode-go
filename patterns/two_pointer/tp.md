 # Key Concepts of Two-Pointer Pattern:

* Opposite Direction (Examples 1, 2, 4): Start pointers at both ends and move them toward each other based on some condition.
* Same Direction / Fast-Slow (Example 3): Both pointers move in the same direction but at different speeds. Often used for in-place operations.
* Fixed Window with Iteration (Example 5): Fix one element and use two pointers on the remaining array.

# When to Use Two-Pointers:

* The array is sorted (or you can sort it first)
* You need to find pairs or triplets with certain properties
* You want to avoid nested loops (reduce from O(n²) to O(n))
* In-place array operations
* Palindrome or symmetry checking

# Pattern Recognition Tips:

* Look for words like "pair", "sum", "sorted array"
* Problems asking for O(1) space complexity
* Problems involving comparisons from both ends
* Removing duplicates or modifying arrays in-place

# The examples progress from simple to more complex:

1. Two Sum - The classic introduction
2. Palindrome - Simple opposite direction pointers
3. Remove Duplicates - Fast/slow pointer technique
4. Container With Most Water - Optimization problem
5. Three Sum - Combining techniques (sorting + two-pointer)