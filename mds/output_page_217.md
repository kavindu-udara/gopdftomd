211

The datarate R of a digital communication system is how frequently an information bit is transmitted.

In this example it equals the reciprocal of the bit interval: R =
1
. Thus, for a 1 Mbps (megabit per second)

transmission, we must have T = 1 s .
The choice of signals to represent bit values is arbitrary to some degree. Clearly, we do not want to
choose signal set members to be the same; we couldn't distinguish bits if we did so. We could also have
made the negative-amplitude pulse represent a 0 and the positive one a 1 . This choice is indeed arbitrary
and will have no e ect on performance assuming the receiver knows which signal represents which bit. As
in all communication systems, we design transmitter and receiver together.
A simple signal set for both wireless and wireline channels amounts to amplitude modulating a baseband
signal set (more appropriate for a wireline channel) by a carrier having a frequency harmonic with the bit
interval.

s ( t ) = Ap ( t ) sin
2 kt

s ( t ) = Ap ( t ) sin
2 kt

(6.40)

Figure 6.9

Exercise 6.13 (Solution on p. 256.)
What is the value of k in this example?

This signal set is also known as a BPSK signal set. We'll show later that indeed both signal sets provide
identical performance levels when the signal-to-noise ratios are equal.

Exercise 6.14 (Solution on p. 256.)
Write a formula, in the style of the baseband signal set, for the transmitted signal as shown in the
plot of the baseband signal set that emerges when we use this modulated signal.

What is the transmission bandwidth of these signal sets? We need only consider the baseband version
as the second is an amplitude-modulated version of the rst. The bandwidth is determined by the bit
sequence. If the bit sequence is constant always 0 or always 1 the transmitted signal is a constant, which
has zero bandwidth. The worst-case bandwidth consuming bit sequence is the alternating one shown in
Figure 6.10. In this case, the transmitted signal is a square wave having a period of 2 T .

Figure 6.10: Here we show the transmitted waveform corresponding to an alternating bit sequence.
Signal Sets, Figure 2 http://cnx.org/content/m0542/latest/\#fig1001
