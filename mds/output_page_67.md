61

Figure 3.33 (Solution on p. 95.)
What is the equivalent resistance seen by the voltage source in Figure 3.33?

The node method applies to RLC circuits, without signi cant modi cation from the methods used on simple
resistive circuits, if we use complex amplitudes. We rely on the fact that complex amplitudes satisfy KVL,
KCL, and impedance-based v-i relations. In the example circuit shown in Figure 3.34, we de ne complex
amplitudes for the input and output variables and for the node voltages. We need only one node voltage
here, and its KCL equation is
E V
+ Ej 2 fC +
E
= 0

with the result

E =
R + R + j 2 fR R C
V

To nd the transfer function between input and output voltages, we compute the ratio
E
. The transfer

function's magnitude and angle are

j H ( f ) j =
R R + R ) + (2 fR R C )

\ H ( f ) = arctan
2 fR R C + R

This circuit di ers from the one shown previously (Figure 3.27) in that the resistor R has been added across
the output. What e ect has it had on the transfer function, which in the original circuit was a lowpass lter

having cuto frequency f =
1 R C
? As shown in Figure 3.35, adding the second resistor has two e ects:

it lowers the gain in the passband (the range of frequencies for which the lter has little e ect on the input)
and increases the cuto frequency.

Figure 3.34: Modi cation of the circuit shown on the left to illustrate the node method and the e ect
of adding the resistor R .
