# Task: Sets Calculator

Write a sets calculator. It should calculate union, intersection and difference of sets of integers for given expression. Grammar of calculator is given:
```
expression := “[“ operator sets “]”
sets := set | set sets
set := file | expression
operator := “SUM” | “INT” | “DIF”
```

Each file contains sorted integers, one integer in a line.
Meaning of operators:
`SUM` - returns union of all sets
`INT` - returns intersection of all sets
`DIF` - returns difference of first set and the rest ones

Program should print result on standard output: sorted integers, one integer in a line.

Solution should include: source code, building script (if building is needed). Final program should be able to be run on Linux. Solution should be delivered in tar archive file: “solution.tgz”.

The task will be assessed taking into consideration the following criteria (sorted by severity):
* compliance with specification and correctness
* good programming practices (clean code)
* clearness and readability of implemented algorithm
* computational complexity

## Example:

`$ cat a.txt`
```
1
2
3
```

`$ cat b.txt`
```
2
3
4
```

`$ cat c.txt`
```
3
4
5
```


`./scalc [ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]`
```
1
3
4
```
