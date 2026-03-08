171

R esponse) because their unit sample responses have nite duration. Plotting this response (Fig-

ure 5.19) shows that the unit-sample response is a pulse of width q and height
1
. This waveform

is also known as a boxcar, hence the name boxcar lter given to this system. We'll derive its
frequency response and develop its ltering interpretation in the next section. For now, note that
the di erence equation says that each output value equals the average of the input's current and
previous values. Thus, the output equals the running average of input's previous q values. Such a
system could be used to produce the average weekly temperature ( q = 7 ) that could be updated
daily.

### 5.13 Discrete-Time Systems in the Frequency Domain

As with analog linear systems, we need to nd the frequency response of discrete-time systems. We used
impedances to derive directly from the circuit's structure the frequency response. The only structure we have
so far for a discrete-time system is the di erence equation. We proceed as when we used impedances: let the
input be a complex exponential signal. When we have a linear, shift-invariant system, the output should also
be a complex exponential of the same frequency, changed in amplitude and phase. These amplitude and phase
changes comprise the frequency response we seek. The complex exponential input signal is x ( n ) = Xe .
Note that this input occurs for all values of n . No need to worry about initial conditions here. Assume the
output has a similar form: y ( n ) = Y e . Plugging these signals into the fundamental di erence equation
(5.42), we have

Y e = a Y e + + a Y e + b Xe + b Xe + + b Xe (5.47)

The assumed output does indeed satisfy the di erence equation if the output complex amplitude is related
to the input amplitude by

Y =
b + b e + + b e
a e a e
X

This relationship corresponds to the system's frequency response or, by another name, its transfer function.
We nd that any discrete-time system de ned by a di erence equation has a transfer function given by

H e =
b + b e + + b e
a e a e
(5.48)

In analogy with the results for analog signals, the transfer function should relate the discrete-time Fourier
transform of the system's output to the input's Fourier transform. To show this result, we evaluate the DTFT
of equation (5.42). Noting that the Fourier transform is linear and that y ( n n ) $ e Y e ,
we obtain

Y e = a e + + a e Y e + b + b e + + b e X e

Simpli cation gives use the expected result:

Y e = H e X e (5.49)

where H e is given by (5.48).
Example 5.5
The frequency response of the simple IIR system (di erence equation given in a previous example
(Example 5.3)) is given by

H e =
b ae
(5.50)

This Fourier transform occurred in a previous example; the exponential signal spectrum (Fig-
ure 5.10) portrays the magnitude and phase of this transfer function. When the lter coe cient a
This content is available online at http://cnx.org/content/m0510/2.14/ .
