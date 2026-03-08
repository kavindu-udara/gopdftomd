40 CHAPTER 3. ANALOG SIGNAL PROCESSING

node consists of the entire upper interconnection section. The KCL equation is i i i = 0 . Using the
v-i relations, we nd that

i =
R + R
i

Exercise 3.6 (Solution on p. 94.)
Suppose that you replaced the current source in Figure 3.9 by a voltage source. How would i be
related to the source voltage? Based on this result, what purpose does this revised circuit have?

This circuit highlights some important properties of parallel circuits. You can easily show that the parallel

combination of R and R has the v-i relation of a resistor having resistance
1
+
1
=
R R + R
.

A shorthand notation for this quantity is ( R k R ) . As the reciprocal of resistance is conductance, we can
say that for a parallel combination of resistors, the equivalent conductance is the sum of the
conductances .

Figure 3.10

Similar to voltage divider (p. 38) for series resistances, we have current divider for parallel resistances.
The current through a resistor in parallel with another is the ratio of the conductance of the rst to the

sum of the conductances. Thus, for the depicted circuit, i =
G + G
i . Expressed in terms of resistances,

current divider takes the form of the resistance of the other resistor divided by the sum of resistances:

i =
R + R
i .

Figure 3.11
