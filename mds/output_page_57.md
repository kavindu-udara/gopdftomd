51

(a)

(b)

Figure 3.24: (a) A simple RC circuit. (b) The impedance counterpart for the RC circuit. Note that
the source and output voltage are now complex amplitudes.

Using impedances, the complex amplitude of the output voltage V can be found using voltage
divider:

V =
Z + Z
V

=

+ R
V

=
1 2 fRC + 1
V

Referring to the di erential equation for this circuit (shown in Section 3.8, page 47, to be
RC
v + v = v ), when we let the output and input voltages be complex exponentials, we obtain
the same relationship between their complex amplitudes. Thus, using impedances is equivalent to using the
di erential equation and solving it when the source(s) is a complex exponential.
In fact, we can nd the di erential equation directly using impedances. If we cross-multiply the relation
between input and output amplitudes,

V ( j 2 fRC + 1) = V

and then put the complex exponentials back in, we have

RCj 2 fV e + V e = V e

In the process of de ning impedances, note that the factor j 2 f arises from the derivative of a complex
exponential. We can reverse the impedance process, and revert back to the di erential equation.

RC
d
v + v = v

This is the same equation that was derived much more tediously in Section 3.8. Finding the di erential
equation relating output to input is far simpler when we use impedances than with any other technique.

Exercise 3.11 (Solution on p. 94.)
Suppose you had an expression where a complex amplitude was divided by j 2 f . What time-domain
operation corresponds to this division?

### 3.11 Power in the Frequency Domain

Recalling that the instantaneous power consumed by a circuit element or an equivalent circuit that represents
a collection of elements equals the voltage times the current entering the positive-voltage terminal, p ( t ) =
This content is available online at http://cnx.org/content/m17308/1.2/ .
