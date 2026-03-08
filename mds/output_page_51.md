45

Figure 3.17

Example 3.2

Figure 3.18

For the circuit depicted in Figure 3.18, let's derive its Thévenin equivalent two di erent ways.
Starting with the open/short-circuit approach, let's rst nd the open-circuit voltage v . We have
a current divider relationship as R is in parallel with the series combination of R and R . Thus,

v =
R R + R + R
i . When we short-circuit the terminals, no voltage appears across R , and

thus no current ows through it. In short, R does not a ect the short-circuit current, and can be

eliminated. We again have a current divider relationship: i =
R + R
i . Thus, the Thévenin

equivalent resistance is
R ( R + R ) + R + R
.

To verify, let's nd the equivalent resistance by reaching inside the circuit and setting the current
source to zero. Because the current is now zero, we can replace the current source by an open circuit.
From the viewpoint of the terminals, resistor R is now in parallel with the series combination of
R and R . Thus, R = ( R k R + R ) , and we obtain the same result.

As you might expect, equivalent circuits come in two forms: the voltage-source oriented Thévenin equiva-
lent and the current-source oriented Mayer-Norton equivalent (Figure 3.19). To derive the latter, the
v-i relation for the Thévenin equivalent can be written as

v = R i + v (3.9)

or
i =
v
i (3.10)

where i =
v
is the Mayer-Norton equivalent source. The Mayer-Norton equivalent shown in Figure 3.19

can be easily shown to have this v-i relation. Note that both variations have the same equivalent resistance.
The short-circuit current equals the negative of the Mayer-Norton equivalent source.
Finding Thévenin Equivalent Circuits http://cnx.org/content/m0021/latest/
