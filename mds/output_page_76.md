70 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.44 The current i must be very small. The voltage produced by the dependent source is 10 times the

voltage v . Thus, the voltage v must be small, which means that i =
v
must be tiny. For example,

if the output is about 1 V, the voltage v = 10 V, making the current i = 10 A. Consequently,
we can ignore i in our calculations and assume it to be zero.
Because of this assumption essentially no current ow through R the voltage v must also be
essentially zero. This means that in op-amp circuits, the voltage across the op-amp's input is basically
zero.

Armed with these approximations, let's return to our original circuit as shown in Figure 3.44. The node
voltage e is essentially zero, meaning that it is essentially tied to the reference node. Thus, the current

through the resistor R equals
v
. Furthermore, the feedback resistor appears in parallel with the load

resistor. Because the current going into the op-amp is zero, all of the current owing through R ows

through the feedback resistor ( i = i )! The voltage across the feedback resistor v equals
v R
. Because

the left end of the feedback resistor is essentially attached to the reference node, the voltage across it equals

the negative of that across the output resistor: v = v =
v R
. Using this approach makes analyzing

new op-amp circuits much easier. When using this technique, check to make sure the results you obtain
are consistent with the assumptions of essentially zero current entering the op-amp and nearly zero voltage
across the op-amp's inputs.

Example 3.8

Figure 3.45: Two-source, single-output op-amp circuit example.

![Im86](../output/images/test_076_Im86.R11.tif)
