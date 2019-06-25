# Chapter 2 - Algorithms Analysis

## Introduction

By looking into various problem sloving algorithms or problem solving techniques we begin to develop a patter that will help us in solving solving similar problemas when they come in front us.

## Algorithms

An algorithm is a set of step to accomplish a task. Or an algorithms is a computer program in wich a set of steps applied over a set of input to produce a set of output. The most important properties of an algorithm are:

1. Correctness: The algorithm should be correct. It should be able to process all the given inputs and provide correct output.

2. Efficiency: The algorithm should be efficient in solving problems. Efficiency is measured in two parameters. First, is Time-Complexity, how fast result is porvided by an algorithm. And second is Space-Complexity, how much RAM that an an algorithm is going to consume.

Time-Complexity us represented by the function ```T(n)```, time versus input size n.

Space-Complecity is represented by the function```S(n)```, memory used versus the input size n.

## Asymptotic analysis

Is used to compare the efficiency of algorithm independtyly of any particular data set or PL.

### Big-O Notation

```f(n)``` is big-O of ```f(n)```. Or ```f(n) = O(g(n))```

### Omega Notation

### Theta Notation

## Complexity analysis of algorithms

1. Wrost Case Complexity: It is the complexity of solving the problem for the worst input of size n. It provides the upper bound for the algorithm

2. Avarage Case Complexity: It is the complexity of solving the problem on an avarage. We calculate the time for all the possible inputs and then take an average of it.

3. Best Case Coimplexity: It is the complexity of solving the problem for the best input of size.

## Time Complexity Order

Constant => O(1)

- An algorithms said to run in constant time if the output is produced in constant time regardless of the input size. Examples:

1. Accessing n element of a list

2. Push and pop of a stack

3. Enqueue and remove of a queue

4. Accessing an element of Hash-Table

Logarithmatic = O(log n)

- An algorithm is said to run in logarithmic time if the execution time of the algorithm is porportional to the logarithm of the input size.

Linear => O(n)

- An algorith is said to run in linear time if the execution time of the algorithm is directly proportonial to the input size. Examples:

1. List operations like search element, find min, find max.

2. Linked list operations like traversal, find min, find max.

N-LogN => O(n log n)

- An algorithm is said to run in logarithmic time if the execution time of an algorithm is proportional of the product of input size and logarithm of the input size. Example:

1. Merge-Sort

2. Quick-Sort

3. Heap-Sort

Quadratic => O(n^2)

- An algorithm is said to run in quadratic time if the execution time of an algorithm is proportional to the square of the input size. Example:

1. Bubble-sort

2. Selection-Sort

3. Insertion-Sort

Polynomial => O(c^2) - c is constant and c>

Polynomial => O(c^m) - c is constant and c>1

Factorial or N-power-N => O(n!)or O(n^n)

## Deriving the Runtime Function of an Algorithm 

- Constant: Each statement takes a constant time to run. Time-Complexity is O(1)

- Loops: The running time of a loop is a product of running time of the satement inside a loop and number of iterations in the loop. Time-Complexity is O(n)

- Nested Loop: The runnin time of a nested loop is a product of time of the statement inside loop multiplied by a product of the size of all the loops. Time-Complexity is O(n^c), where c is a number of loops. For two loops, it will be O(n^n)

- Consecutive Statements: Just add the runnin time of all the consecutive statements

- If-Else Statement: Consider the running time of the larger of if block or else block. Moreover, ignore the other one.

- Logarithmic statement: If each iteration the input size is decreases by a constant factos. Time-Complexity O(log n)
