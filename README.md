### Package Dispatcher

#### Objective

The objective of this program is to sort packages into their expected categories based on their volume and mass.

#### Rules

In an automated factory setting, packages are dispatched to different stacks depending on their characteristics:

- **Standard**: Packages that are neither bulky nor heavy are handled normally.
- **Special**: Packages that are either heavy or bulky can't be handled automatically.
- **Rejected**: Packages that are both heavy and bulky are rejected.

#### Implementation

The `solve(width, height, length, mass)` method is implemented to determine the appropriate stack for a given package. It returns a string representing the name of the stack.

#### Constraints

- Dimensions (width, height, length): 20 ≤ value ≤ 200 (centimeters)
- Mass: 10 ≤ value ≤ 1000 (kilograms)

#### Tests

```go
solve(160, 70, 70, 10)   // SPECIAL
solve(120, 100, 100, 10) // SPECIAL
solve(90, 90, 118, 10)   // STANDARD
solve(120, 100, 50, 30)  // SPECIAL
solve(80, 110, 80, 70)   // SPECIAL
solve(70, 80, 90, 10)    // STANDARD
solve(100, 120, 60, 19)  // STANDARD
solve(150, 70, 70, 10)   // SPECIAL
solve(120, 100, 100, 10) // SPECIAL
solve(90, 90, 118, 10)   // STANDARD
solve(120, 100, 110, 20) // REJECTED
solve(80, 110, 80, 70)   // SPECIAL
solve(70, 80, 90, 10)    // STANDARD
solve(100, 150, 60, 30)  // REJECTED
```
