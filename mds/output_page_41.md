35

(a)

(b)

(c)

Figure 3.6: The circuit shown in the top two gures is perhaps the simplest circuit that performs a
signal processing function. On the bottom is the block diagram that corresponds to the circuit. The
input is provided by the voltage source v and the output is the voltage v across the resistor label R .
As shown in the middle, we analyze the circuit understand what it accomplishes by de ning currents
and voltages for all circuit elements, and then solving the circuit and element equations. resistor. First of all, physical devices
are manufactured to close tolerances (the tighter the tolerance, the more money you pay), but never have
exactly their advertised values. The fourth band on resistors speci es their tolerance; 10% is common. More
pertinent to the current discussion is another deviation from the ideal: If a sinusoidal voltage is placed across
a physical resistor, the current will not be exactly proportional to it as frequency becomes high, say above
1 MHz. At very high frequencies, the way the resistor is constructed introduces inductance and capacitance
e ects. Thus, the smart engineer must be aware of the frequency ranges over which his ideal models match
reality well.
On the other hand, physical circuit elements can be readily found that well approximate the ideal, but
they will always deviate from the ideal in some way. For example, a ashlight battery, like a C-cell, roughly
corresponds to a 1.5 V voltage source. However, it ceases to be modeled by a voltage source capable of
supplying any current (that's what ideal ones can do!) when the resistance of the light bulb is too small.

### 3.4 Electric Circuits and Interconnection Laws

A circuit connects circuit elements together in a speci c con guration designed to transform the source
signal (originating from a voltage or current source) into another signal the output that corresponds to
the current or voltage de ned for a particular circuit element. A simple resistive circuit is shown in Figure 3.6.
This circuit is the electrical embodiment of a system having its input provided by a source system producing
v ( t ) .
To understand what this circuit accomplishes, we want to determine the voltage across the resistor labeled
by its value R . Recasting this problem mathematically, we need to solve some set of equations so that we
relate the output voltage v to the source voltage. It would be simple a little too simple at this point if
we could instantly write down the one equation that relates these two voltages. Until we have more knowledge
about how circuits work, we must write a set of equations that allow us to nd all the voltages and currents
that can be de ned for every circuit element. Because we have a three-element circuit, we have a total of six
voltages and currents that must be either speci ed or determined. You can de ne the directions for current
ow and positive voltage drop any way you like . When two people solve a circuit their own ways, the
signs of their variables may not agree, but current ow and voltage drop values for each element will agree.
Do recall in de ning your voltage and current variables that the v-i relations for the elements presume that
This content is available online at http://cnx.org/content/m0014/2.27/ .
