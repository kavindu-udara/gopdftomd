111

is has no e ect: Whatever you pick for a value has no practical relevance for either the signal's spectrum or
for how a system responds to the signal. The Fourier series value at the discontinuity is the average of the
values on either side of the jump.

### 4.6 Encoding Information in the Frequency Domain

To emphasize the fact that every periodic signal has both a time and frequency domain representation, we
can exploit both to encode information into a signal. Refer to the Fundamental Model of Communication
(Figure 1.3). We have an information source, and want to construct a transmitter that produces a signal
x ( t ) . For the source, let's assume we have information to encode every T seconds. For example, we want
to represent typed letters produced by an extremely good typist (a key is struck every T seconds). Let's
consider the Fourier series formula in the light of trying to encode information.

x ( t ) =
X
c e
(4.26)

We use a nite sum here merely for simplicity (fewer parameters to determine). An important aspect of
the spectrum is that each frequency component c can be manipulated separately: Instead of nding the
Fourier spectrum from a time-domain speci cation, let's construct it in the frequency domain by selecting
the c according to some rule that relates coe cient values to the alphabet. In de ning this rule, we want to
always create a real-valued signal x ( t ) . Because of the Fourier spectrum's properties (Property 4.1, p. 99),
the spectrum must have conjugate symmetry. This requirement means that we can only assign positive-
indexed coe cients (positive frequencies), with negative-indexed ones equaling the complex conjugate of the
corresponding positive-indexed ones.
Assume we have N letters to encode: f a ; : : : ; a g . One simple encoding rule could be to make a single
Fourier coe cient be non-zero and all others zero for each letter. For example, if a occurs, we make c = 1
and c = 0 , k 6 = n . In this way, the n harmonic of the frequency 1 =T is used to represent a letter. Note
that the bandwidth the range of frequencies required for the encoding equals N=T . Another possibility
is to consider the binary representation of the letter's index. For example, if the letter a occurs, converting
13 to its base-2 representation, we have 13 = 1101 . We can use the pattern of zeros and ones to represent
directly which Fourier coe cients we turn on (set equal to one) and which we turn o .

Exercise 4.9 (Solution on p. 140.)
Compare the bandwidth required for the direct encoding scheme (one nonzero Fourier coe cient
for each letter) to the binary number scheme. Compare the bandwidths for a 128-letter alphabet.
Since both schemes represent information without loss we can determine the typed letter uniquely
from the signal's spectrum both are viable. Which makes more e cient use of bandwidth and
thus might be preferred?

Exercise 4.10 (Solution on p. 140.)
Can you think of an information-encoding scheme that makes even more e cient use of the spec-
trum? In particular, can we use only one Fourier coe cient to represent N letters uniquely?

We can create an encoding scheme in the frequency domain (p. 111) to represent an alphabet of letters. But,
as this information-encoding scheme stands, we can represent one letter for all time. However, we note that
the Fourier coe cients depend only on the signal's characteristics over a single period. We could change
the signal's spectrum every T as each letter is typed. In this way, we turn spectral coe cients on and o as
letters are typed, thereby encoding the entire typed document. For the receiver in the Fundamental Model
of Communication (Figure 1.3) to retrieve the typed letter, it would simply use the Fourier formula for the
Fourier spectrum for each T -second interval to determine what each typed letter was. Figure 4.9 shows
such a signal in the time-domain.
This content is available online at http://cnx.org/content/m0043/2.17/ .
Fourier Series and Their Properties, (2) http://cnx.org/content/m0065/latest/\#complex
