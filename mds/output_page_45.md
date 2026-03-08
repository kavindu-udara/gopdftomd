39

The input-output relationship for this system, found in this particular case by voltage divider, takes
the form of a ratio of the output voltage to the input voltage.

v
=
R + R

In this way, we express how the components used to build the system a ect the input-output relationship.
Because this analysis was made with ideal circuit elements, we might expect this relation to break down if
the input amplitude is too high (Will the circuit survive if the input changes from 1 volt to one million volts?)
or if the source's frequency becomes too high. In any case, this important way of expressing input-output
relationships as a ratio of output to input pervades circuit and system theory.
The current i is the current owing out of the voltage source. Because it equals i , we have that the
ratio of the source's voltage to the current owing out of it equals
= R + R . Consequently, from the
viewpoint of the source, it appears to be attached to a single resistor having resistance R + R .

Resistors in series: The series combination of two resistors acts, as far as the voltage source is
concerned, as a single resistor having a value equal to the sum of the two resistances.

This result is the rst of several equivalent circuit ideas: In many cases, a complicated circuit when viewed
from its terminals (the two places to which you might attach a source) appears to be a single circuit element
(at best) or a simple combination of elements at worst. Thus, the equivalent circuit for a series combination
of resistors is a single resistor having a resistance equal to the sum of its component resistances.

Figure 3.8: The resistor (on the right) is equivalent to the two resistors (on the left) and has a resistance
equal to the sum of the resistances of the other two resistors.

Thus, the circuit the voltage source feels (through the current drawn from it) is a single resistor having
resistance R + R . Note that in making this equivalent circuit, the output voltage can no longer be de ned:
The output resistor labeled R no longer appears. Thus, this equivalence is made strictly from the voltage
source's viewpoint.

Figure 3.9: A simple parallel circuit.

One interesting simple circuit (Figure 3.9) has two resistors connected side-by-side, what we will term a
parallel connection, rather than in series. Here, applying KVL reveals that all the voltages are identical:
v = v and v = v . This result typi es parallel connections. To write the KCL equation, note that the top
