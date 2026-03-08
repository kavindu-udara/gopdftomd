56 CHAPTER 3. ANALOG SIGNAL PROCESSING

depends only on the product of the resistance and the capacitance. Thus, a cuto frequency of 1 kHz occurs

when
1 RC
= 10 or RC =
10
= 1 : 59 10 . Thus resistance-capacitance combinations of 1.59 k and

100 nF or 10 and 1.59 F result in the same cuto frequency.
The phase shift caused by the circuit at the cuto frequency precisely equals = 4 . Thus, below the
cuto frequency, phase is little a ected, but at higher frequencies, the phase shift caused by the circuit
becomes = 2 . This phase shift corresponds to the di erence between a cosine and a sine.
We can use the transfer function to nd the output when the input voltage is a sinusoid for two reasons.
First of all, a sinusoid is the sum of two complex exponentials, each having a frequency equal to the negative
of the other. Secondly, because the circuit is linear, superposition applies. If the source is a sine wave, we
know that
v ( t ) = A sin (2 ft )

=
A j
e e
(3.16)

Since the input is the sum of two complex exponentials, we know that the output is also a sum of two similar
complex exponentials, the only di erence being that the complex amplitude of each is multiplied by the
transfer function evaluated at each exponential's frequency.

v ( t ) =
A j
H ( f ) e
A j
H ( f ) e (3.17)

As noted earlier, the transfer function is most conveniently expressed in polar form: H ( f ) = j H ( f ) j e .
Furthermore, j H ( f ) j = j H ( f ) j (even symmetry of the magnitude) and \ H ( f ) = \ H ( f ) (odd sym-
metry of the phase). The output voltage expression simpli es to

v ( t ) =
A j
j H ( f ) j e
A j
j H ( f ) j e

= A j H ( f ) j sin 2 ft + \ H ( f )

(3.18)

The circuit's output to a sinusoidal input is also a sinusoid, having a gain equal to the magnitude
of the circuit's transfer function evaluated at the source frequency and a phase equal to the
phase of the transfer function at the source frequency . It will turn out that this input-output relation
description applies to any linear circuit having a sinusoidal source.

Exercise 3.14 (Solution on p. 95.)
This input-output property is a special case of a more general result. Show that if the source can
be written as the imaginary part of a complex exponential v ( t ) = Im V e the output

is given by v ( t ) = Im V H ( f ) e . Show that a similar result also holds for the real part.

The notion of impedance arises when we assume the sources are complex exponentials. This assumption
may seem restrictive; what would we do if the source were a unit step? When we use impedances to nd the
transfer function between the source and the output variable, we can derive from it the di erential equation
that relates input and output. The di erential equation applies no matter what the source may be. As
we have argued, it is far simpler to use impedances to nd the di erential equation (because we can use
series and parallel combination rules) than any other method. In this sense, we have not lost anything by
temporarily pretending the source is a complex exponential.
In fact we can also solve the di erential equation using impedances! Thus, despite the apparent restric-
tiveness of impedances, assuming complex exponential sources is actually quite general.

### 3.14 Designing Transfer Functions

If the source consists of two (or more) signals, we know from linear system theory that the output voltage
equals the sum of the outputs produced by each signal alone. In short, linear circuits are a special case
This content is available online at http://cnx.org/content/m0031/2.21/ .
