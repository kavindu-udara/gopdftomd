169

a sum of signals can be the sum of the individual outputs occurs when the initial conditions in each case are
zero.

Exercise 5.25 (Solution on p. 193.)
The initial condition issue resolves making sense of the di erence equation for inputs that start at
some index. However, the program will not work because of a programming, not conceptual, error.
What is it? How can it be xed?

Example 5.3
Let's consider the simple system having p = 1 and q = 0 .

y ( n ) = ay ( n 1) + bx ( n ) (5.43)

To compute the output at some index, this di erence equation says we need to know what the
previous output y ( n 1) and what the input signal is at that moment of time. In more detail, let's
compute this system's output to a unit-sample input: x ( n ) = ( n ) . Because the input is zero for
negative indices, we start by trying to compute the output at n = 0 .

y (0) = ay ( 1) + b (5.44)

What is the value of y ( 1) ? Because we have used an input that is zero for all negative indices, it
is reasonable to assume that the output is also zero. Certainly, the di erence equation would not
describe a linear system if the input that is zero for all time did not produce a zero output. With
this assumption, y ( 1) = 0 , leaving y (0) = b . For n > 0 , the input unit-sample is zero, which
leaves us with the di erence equation y ( n ) = ay ( n 1) ; n > 0 . We can envision how the lter
responds to this input by making a table.

y ( n ) = ay ( n 1) + b ( n ) (5.45) ( n ) ( n ) 1

Coe cient values determine how the output behaves. The parameter b can be any value, and
serves as a gain. The e ect of the parameter a is more complicated (Table 5.2). If it equals zero,
the output simply equals the input times the gain b . For all non-zero values of a , the output
lasts forever; such systems are said to be IIR ( I n nite I mpulse R esponse). The reason for this
terminology is that the unit sample also known as the impulse (especially in analog situations), and
the system's response to the impulse lasts forever. If a is positive and less than one, the output
is a decaying exponential. When a = 1 , the output is a unit step. If a is negative and greater
than 1 , the output oscillates while decaying exponentially. When a = 1 , the output changes
sign forever, alternating between b and b . More dramatic e ects when j a j > 1 ; whether positive
or negative, the output signal becomes larger and larger, growing exponentially.
Positive values of a are used in population models to describe how population size increases
over time. Here, n might correspond to generation. The di erence equation says that the number
