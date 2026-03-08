63

Figure 3.36 we conclude that the sum of element powers must equal zero in any
circuit regardless of the elements used to construct the circuit .

X
v i = 0

The simplicity and generality with which we proved this results generalizes to other situations as well.
In particular, note that the complex amplitudes of voltages and currents obey KVL and KCL, respectively.
Consequently, we have that
P
V I = 0 . Furthermore, the complex-conjugate of currents also satis es
KCL, which means we also have
P
V I = 0 . And nally, we know that evaluating the real-part of an
expression is linear. Finding the real-part of this power conservation gives the result that average power
is also conserved in any circuit .
X
1
Re [ V I ] = 0

note: This proof of power conservation can be generalized in another very interesting way. All we
need is a set of voltages that obey KVL and a set of currents that obey KCL. Thus, for a given
circuit topology (the speci c way elements are interconnected), the voltages and currents can be
measured at di erent times and the sum of v-i products is zero.

X
v ( t ) i ( t ) = 0

Even more interesting is the fact that the elements don't matter. We can take a circuit and measure
all the voltages. We can then make element-for-element replacements and, if the topology has not
changed, we can measure a set of currents. The sum of the product of element voltages and currents
will also be zero!

### 3.17 Electronics

So far we have analyzed electrical circuits: The source signal has more power than the output variable,
be it a voltage or a current. Power has not been explicitly de ned, but no matter. Resistors, inductors,
and capacitors as individual elements certainly provide no power gain, and circuits built of them will not
magically do so either. Such circuits are termed electrical in distinction to those that do provide power
gain: electronic circuits . Providing power gain, such as your stereo reading a CD and producing sound, is
accomplished by semiconductor circuits that contain transistors. The basic idea of the transistor is to let the
weak input signal modulate a strong current provided by a source of electrical power the power supply to
produce a more powerful signal. A physical analogy is a water faucet: By turning the faucet back and
forth, the water ow varies accordingly, and has much more power than expended in turning the handle.
The waterpower results from the static pressure of the water in your plumbing created by the water utility
pumping the water up to your local water tower. The power supply is like the water tower, and the faucet
This content is available online at http://cnx.org/content/m0035/2.8/ .
