71

Let's try this analysis technique on a simple extension of the inverting ampli er con guration shown
in Figure 3.45. If either of the source-resistor combinations were not present, the inverting ampli er
remains, and we know that transfer function. By superposition, we know that the input-output
relation is

v =
R
v
R
v (3.31)

When we start from scratch, the node joining the three resistors is at the same potential as the
reference, e 0 , and the sum of currents owing into that node is zero. Thus, the current i owing

in the resistor R equals
v
+
v
. Because the feedback resistor is essentially in parallel with

the load resistor, the voltages must satisfy v = v . In this way, we obtain the input-output
relation given above.
What utility does this circuit have? Can the basic notion of the circuit be extended without
bound?

### 3.20 The Diode

The resistor, capacitor, and inductor are linear circuit elements in that their v-i relations are linear in the
mathematical sense. Voltage and current sources are (technically) nonlinear devices: stated simply, doubling
the current through a voltage source does not double the voltage. A more blatant, and very useful, nonlinear
circuit element is the diode (learn more ). Its input-output relation has an exponential form.

i ( t ) = I e

v ( t )
1 (3.32)

Here, the quantity q represents the charge of a single electron in coulombs, k is Boltzmann's constant, and

Figure 3.46: v-i relation and schematic symbol for the diode. Here, the diode parameters were room
temperature and I = 1 A .

T is the diode's temperature in K . At room temperature, the ratio
kT
= 25 mV . The constant I is the

leakage current, and is usually very small. Viewing this v-i relation in Figure 3.46, the nonlinearity becomes
obvious. When the voltage is positive, current ows easily through the diode. This situation is known as
forward biasing . When we apply a negative voltage, the current is quite small, and equals I , known as
the leakage or reverse-bias current. A less detailed model for the diode has any positive current owing
through the diode when it is forward biased, and no current when negative biased. Note that the diode's
schematic symbol looks like an arrowhead; the direction of current ow corresponds to the direction the
arrowhead points.
This content is available online at http://cnx.org/content/m0037/2.16/ .
P-N Junction: Part II http://cnx.org/content/m1004/latest/
