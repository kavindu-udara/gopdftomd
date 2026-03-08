38 CHAPTER 3. ANALOG SIGNAL PROCESSING

(a)

(b)

Figure 3.7: The circuit shown is perhaps the simplest circuit that performs a signal processing function.
The input is provided by the voltage source v and the output is the voltage v across the resistor
labelled R . A physical wire has a resistance and hence dissipates power (it gets warm just like a resistor
in a circuit). In fact, the resistance of a wire of length L and cross-sectional area A is given by

R =
L

The quantity is known as the resistivity and presents the resistance of a unit-length unit cross-
sectional area material constituting the wire. Resistivity has units of ohm-meters. Most materials
have a positive value for , which means the longer the wire, the greater the resistance and thus
the power dissipated. The thicker the wire, the smaller the resistance. Superconductors have zero
resistivity and hence do not dissipate power. If a room-temperature superconductor could be found,
electric power could be sent through power lines without loss!

Exercise 3.4 (Solution on p. 94.)
Calculate the power consumed/created by the resistor R in our simple circuit example.

We conclude that both resistors in our example circuit consume power, which points to the voltage source
as the producer of power. The current owing into the source's positive terminal is i . Consequently,
the power calculation for the source yields

v i =
1 + R
v

We conclude that the source provides the power consumed by the resistors, no more, no less.

Exercise 3.5 (Solution on p. 94.)
Con rm that the source produces exactly the total power consumed by both resistors.

This result is quite general: sources produce power and the circuit elements, especially resistors, consume it.
But where do sources get their power? Again, circuit theory does not model how sources are constructed,
but the theory decrees that all sources must be provided energy to work.

### 3.6 Series and Parallel Circuits

The results shown in Section 3.4 with regard to this circuit (Figure 3.7), and the values of other currents
and voltages in this circuit as well, have profound implications.
Resistors connected in such a way that current from one must ow only into another currents in all
resistors connected this way have the same magnitude are said to be connected in series . For the two
series-connected resistors in the example, the voltage across one resistor equals the ratio of that
resistor's value and the sum of resistances times the voltage across the series combination .
This concept is so pervasive it has a name: voltage divider .
This content is available online at http://cnx.org/content/m10674/2.8/ .
