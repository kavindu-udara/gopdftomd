54 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.26
already found. When we short the terminals, the capacitor no longer has any e ect on the circuit,

and the short-circuit current I equals
V
. The equivalent impedance can be found by setting

the source to zero, and nding the impedance using series and parallel combination rules. In
our case, the resistor and capacitor are in parallel once the voltage source is removed (setting it

to zero amounts to replacing it with a short-circuit). Thus, Z = R k
1 2 fC
=
R j 2 fRC
.

Consequently, we have

V =
1 j 2 fRC
V

I =
1
V

Z =
R j 2 fRC

Again, we should check the units of our answer. Note in particular that j 2 fRC must be dimen-
sionless. Is it?

### 3.13 Transfer Functions

The ratio of the output and input amplitudes for Figure 3.27, known as the transfer function or the
frequency response , is given by
V
= H ( f )

=
1 2 fRC + 1

(3.15)

Implicit in using the transfer function is that the input is a complex exponential, and the output is also a
complex exponential having the same frequency. The transfer function reveals how the circuit modi es the
input amplitude in creating the output amplitude. Thus, the transfer function completely describes how
the circuit processes the input complex exponential to produce the output complex exponential. The circuit's
function is thus summarized by the transfer function. In fact, circuits are often designed to meet transfer
function speci cations. Because transfer functions are complex-valued, frequency-dependent quantities, we
can better appreciate a circuit's function by examining the magnitude and phase of its transfer function
(Figure 3.28).
This transfer function has many important properties and provides all the insights needed to determine
how the circuit functions. First of all, note that we can compute the frequency response for both positive and
negative frequencies. Recall that sinusoids consist of the sum of two complex exponentials, one having the
negative frequency of the other. We will consider how the circuit acts on a sinusoid soon. Do note that the
This content is available online at http://cnx.org/content/m0028/2.19/ .
