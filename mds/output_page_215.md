209

More transmitter power increasing A increases the signal-to-noise ratio proportionally.
The carrier frequency f has no e ect on SNR, but we have assumed that ( f W ) .
The signal bandwidth W enters the signal-to-noise expression in two places: implicitly through the
signal power and explicitly in the expression's denominator. If the signal spectrum had a constant
amplitude as we increased the bandwidth, signal power would increase proportionally. On the other
hand, our transmitter enforced the criterion that signal amplitude was constant (Section 6.7). Signal
amplitude essentially equals the integral of the magnitude of the signal's spectrum.

note: This result isn't exact, but we do know that m (0) =
R
M ( f ) df .

Enforcing the signal amplitude speci cation means that as the signal's bandwidth increases we must de-
crease the spectral amplitude, with the result that the signal power remains constant. Thus, increasing
signal bandwidth does indeed decrease the signal-to-noise ratio of the receiver's output.
Increasing channel attenuation moving the receiver farther from the transmitter decreases the
signal-to-noise ratio as the square. Thus, signal-to-noise ratio decreases as distance-squared between
transmitter and receiver.
Noise added by the channel adversely a ects the signal-to-noise ratio.

In summary, amplitude modulation provides an e ective means for sending a bandlimited signal from one
place to another. For wireline channels, using baseband or amplitude modulation makes little di erence in
terms of signal-to-noise ratio. For wireless channels, amplitude modulation is the only alternative. The one
AM parameter that does not a ect signal-to-noise ratio is the carrier frequency f : We can choose any value
we want so long as the transmitter and receiver use the same value. However, suppose someone else wants to
use AM and chooses the same carrier frequency. The two resulting transmissions will add, and both receivers
will produce the sum of the two signals. What we clearly need to do is talk to the other party, and agree to
use separate carrier frequencies. As more and more users wish to use radio, we need a forum for agreeing on
carrier frequencies and on signal bandwidth. On earth, this forum is the government. In the United States,
the Federal Communications Commission (FCC) strictly controls the use of the electromagnetic spectrum
for communications. Separate frequency bands are allocated for commercial AM, FM, cellular telephone
(the analog version of which is AM), short wave (also AM), and satellite communications.

Exercise 6.11 (Solution on p. 256.)
Suppose all users agree to use the same signal bandwidth. How closely can the carrier frequencies
be while avoiding communications crosstalk? What is the signal bandwidth for commercial AM?
How does this bandwidth compare to the speech bandwidth?

### 6.13 Digital Communication

E ective, error-free transmission of a sequence of bits a bit stream f b (0) ; b (1) ; : : : g is the goal here. We
found that analog schemes, as represented by amplitude modulation, always yield a received signal containing
noise as well as the message signal when the channel adds noise. Digital communication schemes are very
di erent. Once we decide how to represent bits by analog signals that can be transmitted over wireline
(like a computer network) or wireless (like digital cellular telephone) channels, we will then develop a way
of tacking on communication bits to the message bits that will reduce channel-induced errors greatly. In
theory, digital communication errors can be zero, even though the channel adds noise!
We represent a bit by associating one of two speci c analog signals with the bit's value. Thus, if b ( n ) = 0 ,
we transmit the signal s ( t ) ; if b ( n ) = 1 , send s ( t ) . These two signals comprise the signal set for digital
communication and are designed with the channel and bit stream in mind. In virtually every case, these
signals have a nite duration T common to both signals; this duration is known as the bit interval .
Exactly what signals we use ultimately a ects how well the bits can be received. Interestingly, baseband
and modulated signal sets can yield the same performance. Other considerations determine how signal set
choice a ects digital communication performance.
This content is available online at http://cnx.org/content/m0519/2.10/ .
