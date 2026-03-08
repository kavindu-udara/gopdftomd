94 CHAPTER 3. ANALOG SIGNAL PROCESSING

### Solutions to Exercises in Chapter 3

Solution to Exercise 3.1 (p. 32)
One kilowatt-hour equals 3,600,000 watt-seconds, which indeed directly corresponds to 3,600,000 joules.
Solution to Exercise 3.2 (p. 36)
KCL says that the sum of currents entering or leaving a node must be zero. If we consider two nodes together
as a supernode, KCL applies as well to currents entering the combination. Since no currents enter an entire
circuit, the sum of currents must be zero. If we had a two-node circuit, the KCL equation of one must
be the negative of the other, We can combine all but one node in a circuit into a supernode; KCL for the
supernode must be the negative of the remaining node's KCL equation. Consequently, specifying n 1 KCL
equations always speci es the remaining one.
Solution to Exercise 3.3 (p. 37)

The circuit serves as an ampli er having a gain of
R + R
.

Solution to Exercise 3.4 (p. 38)
The power consumed by the resistor R can be expressed as

( v v ) i =
R R + R )
v

Solution to Exercise 3.5 (p. 38)

1 + R
v =
R R + R )
v + +
R R + R )
v

Solution to Exercise 3.6 (p. 40)
Replacing the current source by a voltage source does not change the fact that the voltages are identical.

Consequently, v = R i or i =
v
. This result does not depend on the resistor R , which means that

we simply have a resistor ( R ) across a voltage source. The two-resistor circuit has no apparent use.
Solution to Exercise 3.7 (p. 41)

R =
R R =R
. Thus, a 10% change means that the ratio
R
must be less than 0.1. A 1% change

means that
R
< 0 : 01 .

Solution to Exercise 3.8 (p. 42)
In a series combination of resistors, the current is the same in each; in a parallel combination, the voltage
is the same. For a series combination, the equivalent resistance is the sum of the resistances, which will be
larger than any component resistor's value; for a parallel combination, the equivalent conductance is the sum
of the component conductances, which is larger than any component conductance. The equivalent resistance
is therefore smaller than any component resistance.
Solution to Exercise 3.9 (p. 44)

v =
R + R
v and i =
v
(resistor R is shorted out in this case). Thus, v =
R + R
v and

R =
R R + R
.

Solution to Exercise 3.10 (p. 46)

i =
R + R
i and R = ( R k R + R ) .

Solution to Exercise 3.11 (p. 51)
Division by j 2 f arises from integrating a complex exponential. Consequently,

1 2 f
V ()

Z

V e dt
