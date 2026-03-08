82 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.66 When nothing is attached to the right-hand terminals, a voltage of v = 1 V is measured. What circuit
could produce this output?
(c) When a current source is attached so that i = 2 A, the voltage v is now 3 V. What resistor circuit
would be consistent with this and the previous part?

Problem 3.26: More Circuit Detective Work
The left terminal pair of a two terminal-pair circuit is attached to a testing circuit as shown in Figure 3.67.
The test source v ( t ) equals sin ( t ) .
We make the following measurements.

With nothing attached to the terminals on the right, the voltage v ( t ) equals
1
cos ( t + = 4) .

When a wire is placed across the terminals on the right, the current i ( t ) was sin ( t ) .

Figure 3.67

(a) What is the impedance seen from the terminals on the right?
(b) Find the voltage v ( t ) if a current source is attached to the terminals on the right so that i ( t ) = sin ( t ) .

Problem 3.27: Linear, Time-Invariant Systems
For a system to be completely characterized by a transfer function, it needs not only be linear, but also
to be time-invariant. A system is said to be time-invariant if delaying the input delays the output by the
same amount. Mathematically, if S [ x ( t )] = y ( t ) , meaning y ( t ) is the output of a system S [ ] when x ( t )
is the input, S [ ] is the time-invariant if S [ x ( t )] = y ( t ) for all delays and all inputs x ( t ) . Note
that both linear and nonlinear systems have this property. For example, a system that squares its input is
time-invariant.

(a) Show that if a circuit has xed circuit elements (their values don't change over time), its input-output
relationship is time-invariant. Hint : Consider the di erential equation that describes a circuit's input-
output relationship. What is its general form? Examine the derivative(s) of delayed signals.
(b) Show that impedances cannot characterize time-varying circuit elements (R, L, and C). Consequently,
show that linear, time-varying systems do not have a transfer function.
(c) Determine the linearity and time-invariance of the following. Find the transfer function of the linear,
time-invariant (LTI) one(s).
