68 CHAPTER 3. ANALOG SIGNAL PROCESSING

Under these conditions, we obtain the classic input-output relationship for the op-amp-based inverting am-
pli er.

v =
R
v (3.30)

Consequently, the gain provided by our circuit is entirely determined by our choice of the feedback resistor
R and the input resistor R . It is always negative, and can be less than one or greater than one in
magnitude. It cannot exceed the op-amp's inherent gain and should not produce such large outputs that
distortion results (remember the power supply!). Interestingly, note that this relationship does not depend
on the load resistance. This e ect occurs because we use load resistances large compared to the op-amp's
output resistance. Thus observation means that, if careful, we can place op-amp circuits in cascade, without
incurring the e ect of succeeding circuits changing the behavior (transfer function) of previous ones; see this
problem (Problem 3.43).

3.19.2 Active Filters

As long as design requirements are met, the input-output relation for the inverting ampli er also applies
when the feedback and input circuit elements are impedances (resistors, capacitors, and inductors).

Figure 3.42:
V
=
Z

Example 3.7
Let's design an op-amp circuit that functions as a lowpass lter. We want the transfer function
between the output and input voltage to be

H ( f ) =
K

where K equals the passband gain and f is the cuto frequency. Let's assume that the inversion
(negative gain) does not matter. With the transfer function of the above op-amp circuit in mind,
let's consider some choices.

Z = K , Z = 1 +
jf
. This choice means the feedback impedance is a resistor and that the

input impedance is a series combination of an inductor and a resistor. In circuit design, we
try to avoid inductors because they are physically bulkier than capacitors.

Z =
1

, Z =
1
. Consider the reciprocal of the feedback impedance (its admittance):

Z = 1 +
jf
. Since this admittance is a sum of admittances, this expression suggests

the parallel combination of a resistor (value = 1 ) and a capacitor (value =
1
F). We

have the right idea, but the values (like 1 ) are not right. Consider the general RC parallel
