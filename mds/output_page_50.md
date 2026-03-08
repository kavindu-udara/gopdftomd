44 CHAPTER 3. ANALOG SIGNAL PROCESSING

If the source were zero, it could be replaced by a short circuit, which would con rm that the circuit does
indeed function as a parallel combination of resistors. However, the source's presence means that the circuit
is not well modeled as a resistor.

Figure 3.16: The Thévenin equivalent circuit.

If we consider the simple circuit of Figure 3.16, we nd it has the v-i relation at its terminals of

v = R i + v (3.5)

Comparing the two v-i relations, we nd that they have the same form. In this case the Thévenin equiva-

lent resistance is R = ( R k R ) and the Thévenin equivalent source has voltage v =
R + R
v .

Thus, from viewpoint of the terminals, you cannot distinguish the two circuits. Because the equivalent
circuit has fewer elements, it is easier to analyze and understand than any other alternative.
For any circuit containing resistors and sources, the v-i relation will be of the form

v = R i + v (3.6)

and the Thévenin equivalent circuit for any such circuit is that of Figure 3.16. This equivalence applies no
matter how many sources or resistors may be present in the circuit. In the example (Example 3.2) below, we
know the circuit's construction and element values, and derive the equivalent source and resistance. Because
Thévenin's theorem applies in general, we should be able to make measurements or calculations only from
the terminals to determine the equivalent circuit.
To be more speci c, consider the equivalent circuit of Figure 3.16. Let the terminals be open-circuited,
which has the e ect of setting the current i to zero. Because no current ows through the resistor, the voltage
across it is zero (remember, Ohm's Law says that v = Ri ). Consequently, by applying KVL we have that
the so-called open-circuit voltage v equals the Thévenin equivalent voltage. Now consider the situation
when we set the terminal voltage to zero (short-circuit it) and measure the resulting current. Referring to
the equivalent circuit, the source voltage now appears entirely across the resistor, leaving the short-circuit

current to be i =
v
. From this property, we can determine the equivalent resistance.

v = v (3.7)

R =
v
(3.8)

Exercise 3.9 (Solution on p. 94.)
Use the open/short-circuit approach to derive the Thévenin equivalent of the circuit shown in
Figure 3.17.
