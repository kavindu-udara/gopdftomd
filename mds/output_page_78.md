72 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.47 cannot use impedances nor series/parallel combination rules
to analyze circuits containing them. The reliable node method can always be used; it only relies on KVL for
its application, and KVL is a statement about voltage drops around a closed path regardless of whether
the elements are linear or not. Thus, for this simple circuit we have

v
= I e
1 (3.33)

This equation cannot be solved in closed form. We must understand what is going on from basic principles,
using computational and graphical aids. As an approximation, when v is positive, current ows through
the diode so long as the voltage v is smaller than v (so the diode is forward biased). If the source is
negative or v tries to be bigger than v , the diode is reverse-biased, and the reverse-bias current ows
through the diode. Thus, at this level of analysis, positive input voltages result in positive output voltages
with negative ones resulting in v = RI .

Figure 3.48

We need to detail the exponential nonlinearity to determine how the circuit distorts the input voltage
waveform. We can of course numerically solve the circuit shown in Figure 3.47 to determine the output
voltage when the input is a sinusoid. To learn more, let's express this equation graphically. We plot each
term as a function of v for various values of the input voltage v ; where they intersect gives us the output
voltage. The left side, the current through the output resistor, does not vary itself with v , and thus we
have a xed straight line. As for the right side, which expresses the diode's v-i relation, the point at which
the curve crosses the v axis gives us the value of v . Clearly, the two curves will always intersect just once
for any value of v , and for positive v the intersection occurs at a value for v smaller than v . This
reduction is smaller if the straight line has a shallower slope, which corresponds to using a bigger output
resistor. For negative v , the diode is reverse-biased and the output voltage equals RI .
What utility might this simple circuit have? The diode's nonlinearity cannot be escaped here, and the
clearly evident distortion must have some practical application if the circuit were to be useful. This circuit,
