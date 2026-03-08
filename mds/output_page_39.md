33

The resistor is far and away the simplest circuit element. In a resistor, the voltage is proportional to the
current, with the constant of proportionality R , known as the resistance .

v ( t ) = Ri ( t )

Resistance has units of ohms, denoted by , named for the German electrical scientist Georg Ohm . Some-
times, the v-i relation for the resistor is written i = Gv , with G , the conductance , equal to
. Conductance
has units of Siemens (S), and is named for the German electronics industrialist Werner von Siemens .
When resistance is positive, as it is in most cases, a resistor consumes power. A resistor's instantaneous
power consumption can be written one of two ways.

p ( t ) = Ri ( t ) =
1
v ( t )

As the resistance approaches in nity, we have what is known as an open circuit : No current ows but a
non-zero voltage can appear across the open circuit. As the resistance becomes zero, the voltage goes to zero
for a non-zero current ow. This situation corresponds to a short circuit . A superconductor physically
realizes a short circuit.

3.2.2 Capacitor

Figure 3.3: Capacitor. i = C
d
v ( t )

The capacitor stores charge and the relationship between the charge stored and the resultant voltage is
q = Cv . The constant of proportionality, the capacitance, has units of farads (F), and is named for the
English experimental physicist Michael Faraday . As current is the rate of change of charge, the v-i relation
can be expressed in di erential or integral form.

i ( t ) = C
d
v ( t ) or v ( t ) =
1

Z

i ( ) d (3.1)

If the voltage across a capacitor is constant, then the current owing into it equals zero. In this situation,
the capacitor is equivalent to an open circuit. The power consumed/produced by a voltage applied to a
capacitor depends on the product of the voltage and its derivative.

p ( t ) = Cv ( t )
d
v ( t )

This result means that a capacitor's total energy expenditure up to time t is concisely given by

E ( t ) =
1
Cv ( t )

This expression presumes the fundamental assumption of circuit theory: all voltages and currents in
any circuit were zero in the far distant past ( t = 1 ).
http://www-groups.dcs.st-and.ac.uk/~history/Biographies/Ohm.html
https://www.siemens.com/history/en/personalities/founder_generation.htm
https://en.wikipedia.org/wiki/Michael_Faraday
