213

Figure 6.13: The depicted decomposition of the FSK-modulated alternating bit stream into its fre-
quency components simpli es the calculation of its bandwidth. etc. , and the second having the same structure but interleaved with the rst and containing s ( t )
(Figure 6.13).
Each component can be thought of as a xed-frequency sinusoid multiplied by a square wave of period
2 T that alternates between one and zero. This baseband square wave has the same Fourier spectrum as
our BPSK example, but with the addition of the constant term c . This quantity's presence changes the
number of Fourier series terms required for the 90% bandwidth: Now we need only include the zero and rst

harmonics to achieve it. The bandwidth thus equals, with f < f , f +
1 T
f
1 T
= f f +
1
.

If the two frequencies are harmonics of the bit-interval duration, f =
k
and f =
k
with k > k , the

bandwidth equals
k k + 1
. If the di erence between harmonic numbers is 1 , then the FSK bandwidth

is smaller than the BPSK bandwidth. If the di erence is 2 , the bandwidths are equal and larger di erences
produce a transmission bandwidth larger than that resulting from using a BPSK signal set.

### 6.16 Digital Communication Receivers

The receiver interested in the transmitted bit stream must perform two tasks when received waveform r ( t )
begins.

It must determine when bit boundaries occur: The receiver needs to synchronize with the transmitted
signal. Because transmitter and receiver are designed in concert, both use the same value for the bit
interval T . Synchronization can occur because the transmitter begins sending with a reference bit
sequence, known as the preamble . This reference bit sequence is usually the alternating sequence
as shown in the square wave example and in the FSK example (Figure 6.13). The receiver knows
what the preamble bit sequence is and uses it to determine when bit boundaries occur. This procedure
amounts to what in digital hardware as self-clocking signaling : The receiver of a bit stream must
derive the clock when bit boundaries occur from its input signal. Because the receiver usually
does not determine which bit was sent until synchronization occurs, it does not know when during the
preamble it obtained synchronization. The transmitter signals the end of the preamble by switching
to a second bit sequence. The second preamble phase informs the receiver that data bits are about to
come and that the preamble is almost over.
Once synchronized and data bits are transmitted, the receiver must then determine every T seconds
what bit was transmitted during the previous bit interval. We focus on this aspect of the digital
receiver because this strategy is also used in synchronization.
This content is available online at http://cnx.org/content/m0545/2.12/ .
This content is available online at http://cnx.org/content/m0520/2.18/ .
Transmission Bandwidth, Figure 1 http://cnx.org/content/m0544/latest/\#fig1003
