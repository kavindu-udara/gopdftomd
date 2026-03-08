42 CHAPTER 3. ANALOG SIGNAL PROCESSING

Example 3.1

Figure 3.13

We want to nd the total resistance of the example circuit. To apply the series and parallel
combination rules, it is best to rst determine the circuit's structure: What is in series with what
and what is in parallel with what at both small- and large-scale views. We have R in parallel
with R ; this combination is in series with R . This series combination is in parallel with R . Note
that in determining this structure, we started away from the terminals, and worked toward them.
In most cases, this approach works well; try it rst. The total resistance expression mimics the
structure:

R = R k ( R k R + R )

=
R R R + R R R + R R R R + R R + R R + R R + R R

Such complicated expressions typify circuit simpli cations. A simple check for accuracy is the
units: Each component of the numerator should have the same units (here ) as well as in the
denominator ( ). The entire expression is to have units of resistance; thus, the ratio of the
numerator's and denominator's units should be ohms. Checking units does not guarantee accuracy,
but can catch many errors.

Another valuable lesson emerges from this example concerning the di erence between cascading systems and
cascading circuits. In system theory, systems can be cascaded without changing the input-output relation
of intermediate systems. In cascading circuits, this ideal is rarely true unless the circuits are so designed .
Design is in the hands of the engineer; he or she must recognize what have come to be known as loading
e ects. In our simple circuit, you might think that making the resistance R large enough would do the trick.
Because the resistors R and R can have virtually any value, you can never make the resistance of your
voltage measurement device big enough. Said another way, a circuit cannot be designed in isolation
that will work in cascade with all other circuits . Electrical engineers deal with this situation through
the notion of speci cations : Under what conditions will the circuit perform as designed? Thus, you will
nd that oscilloscopes and voltmeters have their internal resistances clearly stated, enabling you to determine
whether the voltage you measure closely equals what was present before they were attached to your circuit.
Furthermore, since our resistor circuit functions as an attenuator, with the attenuation (a fancy word for

gains less than one) depending only on the ratio of the two resistor values
= 1 +
, we can

select any values for the two resistances we want to achieve the desired attenuation. The designer of this
circuit must thus specify not only what the attenuation is, but also the resistance values employed so that
integrators people who put systems together from component systems can combine systems together and
have a chance of the combination working.
Figure 3.14 summarizes the series and parallel combination results. These results are easy to remember
and very useful. Keep in mind that for series combinations, voltage and resistance are the key quantities,
while for parallel combinations current and conductance are more important. In series combinations, the
currents through each element are the same; in parallel ones, the voltages are the same.
