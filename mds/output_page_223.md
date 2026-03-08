217

### 6.18 Digital Communication System Properties

Results from Section 6.17 reveal several properties about digital communication systems.

As the received signal becomes increasingly noisy, whether due to increased distance from the transmit-
ter (smaller ) or to increased noise in the channel (larger N ), the probability the receiver makes an
error approaches 1 = 2 . In such situations, the receiver performs only slightly better than the receiver
that ignores what was transmitted and merely guesses what bit was transmitted. Consequently, it
becomes almost impossible to communicate information when digital channels become noisy.
As the signal-to-noise ratio increases, performance gains smaller probability of error p can be easily
obtained. At a signal-to-noise ratio of 12 dB, the probability the receiver makes an error equals 10 .
In words, one out of one hundred million bits will, on the average, be in error.
Once the signal-to-noise ratio exceeds about 5 dB, the error probability decreases dramatically. Adding
1 dB improvement in signal-to-noise ratio can result in a factor of ten smaller p .
Signal set choice can make a signi cant di erence in performance. All BPSK signal sets, baseband or
modulated, yield the same performance for the same bit energy. The BPSK signal set does perform
much better than the FSK signal set once the signal-to-noise ratio exceeds about 5 dB.

Exercise 6.20 (Solution on p. 256.)
Derive the expression for the probability of error that would result if the FSK signal set were used.

The matched- lter receiver provides impressive performance once adequate signal-to-noise ratios occur. You
might wonder whether another receiver might be better. The answer is that the matched- lter receiver is
optimal: No other receiver can provide a smaller probability of error than the matched lter
regardless of the SNR . Furthermore, no signal set can provide better performance than the BPSK signal
set, where the signal representing a bit is the negative of the signal representing the other bit. The reason
for this result rests in the dependence of probability of error p on the di erence between the noise-free
integrator outputs: For a given E , no other signal set provides a greater di erence.
How small should the error probability be? Out of N transmitted bits, on the average Np bits will be
received in error. Do note the phrase on the average here: Errors occur randomly because of the noise
introduced by the channel, and we can only predict the probability of occurrence. Since bits are transmitted
at a rate R , errors occur at an average frequency of Rp . Suppose the error probability is an impressively
small number like 10 . Data on a computer network like Ethernet is transmitted at a rate R = 100Mbps ,
which means that errors would occur at a rate of roughly 100 per second. This error rate is very high,
requiring a much smaller p to achieve a more acceptable average occurrence rate for errors occurring.
Because Ethernet is a wireline channel, which means the channel noise is small and the attenuation low,
obtaining very small error probabilities is not di cult. We do have some tricks up our sleeves, however, that
can essentially reduce the error rate to zero without resorting to expending a large amount of energy at
the transmitter. We need to understand digital channels and what implications Shannon's Noisy Channel
Coding Theorem has for them.

### 6.19 Digital Channels

Let's review how digital communication systems work within the Fundamental Model of Communication
(Figure 1.3). As depicted in Figure 6.17, the message is a single bit. The analog components of a digital
transmission/reception system can be lumped into a single system known as the digital channel .
Digital channels are described by transition diagrams , which indicate the output alphabet symbols
that result for each possible transmitted symbol and the probabilities of the various reception possibilities.
The probabilities on transitions coming from the same symbol must sum to one. For the matched- lter
receiver and the signal sets we have seen, the depicted transition diagram, known as a binary symmetric
channel , captures how transmitted bits are received. The probability of error p is the sole parameter of
the digital channel, and it encapsulates signal set choice, channel properties, and the matched- lter receiver.
With this simple but entirely accurate model, we can concentrate on how bits are received.
This content is available online at http://cnx.org/content/m10282/2.9/ .
This content is available online at http://cnx.org/content/m0102/2.14/ .
