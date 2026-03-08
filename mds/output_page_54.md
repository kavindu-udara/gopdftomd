48 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.21: A simple RC circuit.

(a)

(b)

(c)

Figure 3.22: (a) Resistor: Z = R (b) Capacitor: Z =
1 2 fC
(c) Inductor: Z = j 2 fL

Rather than solving the di erential equation that arises in circuits containing capacitors and inductors, let's
pretend that all sources in the circuit are complex exponentials having the same frequency. Although this
pretense can only be mathematically true, this ction will greatly ease solving the circuit no matter what
the source really is.
For the above example RC circuit (Figure 3.21), let v = V e . The complex amplitude V deter-
mines the size of the source and its phase. The critical consequence of assuming that sources have this form
is that all voltages and currents in the circuit are also complex exponentials, having amplitudes governed
by KVL, KCL, and the v-i relations and the same frequency as the source. To appreciate why this should
be true, let's investigate how each circuit element behaves when either the voltage or current is a complex

exponential. For the resistor, v = Ri . When v = V e ; then i =
V
e . Thus, if the resistor's voltage is

a complex exponential, so is the current, with an amplitude I =
V
(determined by the resistor's v-i relation)

and a frequency the same as the voltage. Clearly, if the current were assumed to be a complex exponential,
so would the voltage. For a capacitor, i = C
( v ) . Letting the voltage be a complex exponential, we have
i = CV j 2 fe . The amplitude of this complex exponential is I = CV j 2 f . Finally, for the inductor,
where v = L
( i ) , assuming the current to be a complex exponential results in the voltage having the form
v = LIj 2 fe , making its complex amplitude V = LIj 2 f .
The major consequence of assuming complex exponential voltage and currents is that the

ratio Z =
V
for each element does not depend on time, but does depend on source frequency .

This quantity is known as the element's impedance .
The impedance is, in general, a complex-valued, frequency-dependent quantity. For example, the magni-
tude of the capacitor's impedance is inversely related to frequency, and has a phase of = 2 . This observation
means that if the current is a complex exponential and has constant amplitude, the amplitude of the voltage
decreases with frequency.
Let's consider Kirchho 's circuit laws. When voltages around a loop are all complex exponentials of the
This content is available online at http://cnx.org/content/m0024/2.23/ .
