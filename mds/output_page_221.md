215

### 6.17 Digital Communication in the Presence of Noise

When we incorporate additive noise into our channel model, so that r ( t ) = s ( t )+ n ( t ) , errors can creep in.
If the transmitter sent bit 0 using a BPSK signal set, the integrators' outputs in the matched lter receiver
(Figure 6.14) would be

Z

r ( t ) s ( t ) dt = A T +

Z

n ( t ) s ( t ) dt

Z

r ( t ) s ( t ) dt = A T +

Z

n ( t ) s ( t ) dt

(6.45)

The quantities containing noise terms cause errors in the receiver's decision-making process. Because they
involve noise, the values of these integrals are random quantities drawn from some probability distribution
that vary erratically from bit interval to bit interval. Because the noise has zero average value and has an
equal amount of power in all frequency bands, the values of the integrals will hover about zero. What is
important is how much they vary. If the noise is such that its integral term is more negative than A T ,
then the receiver will make an error, deciding that the transmitted zero-valued bit was indeed a one. The
probability that this situation occurs depends on three factors:

Signal Set Choice The di erence between the signal-dependent terms in the integrators' outputs
(equations (6.45)) de nes how large the noise term must be for an incorrect receiver decision to result.
What a ects the probability of such errors occurring is the energy in the di erence of the received
signals in comparison to the noise term's variability. The signal-di erence energy equals
Z

( s ( t ) s ( t )) dt

For our BPSK baseband signal set, the di erence-signal-energy term is 4 A T .
Variability of the Noise Term We quantify variability by the spectral height of the white noise
added by the channel.
Probability Distribution of the Noise Term The value of the noise terms relative to the
signal terms and the probability of their occurrence directly a ect the likelihood that a receiver error
will occur. For the white noise we have been considering, the underlying distributions are Gaussian.
Deriving the following expression for the probability the receiver makes an error on any bit transmission
is complicated but can be found in other modules: here and here .

p = Q

0

@

s
s ( t ) s ( t ) dt N

1

A

= Q

0

@

s A T

1

A
for the BPSK case

(6.46)

Here Q ( ) is the integral Q ( x ) =

R
e d . This integral has no closed form expression, but it

can be accurately computed. As Figure 6.15 illustrates, Q ( ) is a decreasing, very nonlinear function.

The term A T equals the energy expended by the transmitter in sending the bit; we label this term E . We
arrive at a concise expression for the probability the matched lter receiver makes a bit-reception error when
the BPSK signal set is used.

p = Q

0

@

s E

1

A
(6.47)
This content is available online at http://cnx.org/content/m0546/2.14/ .
"Detection of Signals in Noise" http://cnx.org/content/m16253/latest/
"Continuous-Time Detection Theory" http://cnx.org/content/m11406/latest/
