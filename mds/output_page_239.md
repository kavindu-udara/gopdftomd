233

what this code might be; it has never been found. Codes such as the Hamming code work quite well in
practice to keep error rates low, but they remain greater than zero. Until the magic code is found, more
important in communication system design is the converse. It states that if your data rate exceeds capacity,
errors will overwhelm you no matter what channel coding you use. For this reason, capacity calculations are
made to understand the fundamental limits on transmission rates.

Exercise 6.34 (Solution on p. 259.)
The rst de nition of capacity applies only for binary symmetric channels, and represents the
number of bits/transmission. The second result states capacity more generally, having units of
bits/second. How would you convert the rst de nition's result into units of bits/second?

Example 6.5
The telephone channel has a bandwidth of 3 kHz and a signal-to-noise ratio exceeding 30 dB (at
least they promise this much). The maximum data rate a modem can produce for this wireline
channel and hope that errors will not become rampant is the capacity.

C = 3 10 log 1 + 10

= 29 : 901 kbps
(6.62)

Thus, the so-called 33 kbps modems operate right at the capacity limit.

Note that the data rate allowed by the capacity can exceed the bandwidth when the signal-to-noise ratio
exceeds 0 dB. Our results for BPSK and FSK indicated the bandwidth they require exceeds
. What kind
of signal sets might be used to achieve capacity? Modem signal sets send more than one bit/transmission
using a number, one of the most popular of which is multi-level signaling . Here, we can transmit several
bits during one transmission interval by representing bit by some signal's amplitude. For example, two bits

can be sent with a signal set comprised of a sinusoid with amplitudes of A and
A
.

### 6.32 Comparison of Analog and Digital Communication

Analog communication systems, amplitude modulation (AM) radio being a typifying example, can inexpen-
sively communicate a bandlimited analog signal from one location to another (point-to-point communication)
or from one point to many (broadcast). Although it is not shown here, the coherent receiver (Figure 6.6)
provides the largest possible signal-to-noise ratio for the demodulated message. An analysis of this receiver
thus indicates that some residual error will always be present in an analog system's output.
Although analog systems are less expensive in many cases than digital ones for the same application,
digital systems o er much more e ciency, better performance, and much greater exibility.

E ciency : The Source Coding Theorem allows quanti cation of just how complex a given message
source is and allows us to exploit that complexity by source coding (compression). In analog commu-
nication, the only parameters of interest are message bandwidth and amplitude. We cannot exploit
signal structure to achieve a more e cient communication system.
Performance : Because of the Noisy Channel Coding Theorem, we have a speci c criterion by which
to formulate error-correcting codes that can bring us as close to error-free transmission as we might
want. Even though we may send information by way of a noisy channel, digital schemes are capable
of error-free transmission while analog ones cannot overcome channel disturbances; see this problem
(Problem 6.15) for a comparison.
Flexibility : Digital communication systems can transmit real-valued discrete-time signals, which
could be analog ones obtained by analog-to-digital conversion, and symbolic-valued ones (computer
data, for example). Any signal that can be transmitted by analog means can be sent by digital means,
with the only issue being the number of bits used in A/D conversion (how accurately do we need
to represent signal amplitude). Images can be sent by analog means (commercial television), but
better communication performance occurs when we use digital systems (HDTV). In addition to digital
This content is available online at http://cnx.org/content/m0074/2.11/ .
