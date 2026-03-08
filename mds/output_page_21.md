15

2.2.1 Sinusoids

Perhaps the most common real-valued signal is the sinusoid.

s ( t ) = A cos (2 f t + ) (2.14)

For this signal, A is its amplitude, f its frequency, and its phase.

2.2.2 Complex Exponentials

The most important signal is complex-valued, the complex exponential.

s ( t ) = Ae

= Ae e
(2.15)

Here, j denotes
p 1 . Ae is known as the signal's complex amplitude . Considering the complex
amplitude as a complex number in polar form, its magnitude is the amplitude A and its angle the signal phase.
The complex amplitude is also known as a phasor . The complex exponential cannot be further decomposed
into more elemental signals, and is the most important signal in electrical engineering ! Mathematical
manipulations at rst appear to be more di cult because complex-valued numbers are introduced. In fact,
early in the twentieth century, mathematicians thought engineers would not be su ciently sophisticated
to handle complex exponentials even though they greatly simpli ed solving circuit problems. Steinmetz
introduced complex exponentials to electrical engineering, and demonstrated that mere engineers could use
them to good e ect and even obtain right answers! See Section 2.1 for a review of complex numbers and
complex arithmetic.
The complex exponential de nes the notion of frequency: it is the only signal that contains only one
frequency component. The sinusoid consists of two frequency components: one at the frequency + f and
the other at f .

Euler relation: This decomposition of the sinusoid can be traced to Euler's relation.

cos (2 ft ) =
e + e
(2.16)

sin (2 ft ) =
e e
j
(2.17)

e = cos (2 ft ) + j sin (2 ft ) (2.18)

Decomposition: The complex exponential signal can thus be written in terms of its real and
imaginary parts using Euler's relation. Thus, sinusoidal signals can be expressed as either the real
or the imaginary part of a complex exponential signal, the choice depending on whether cosine or
sine phase is needed, or as the sum of two complex exponentials. These two decompositions are
mathematically equivalent to each other.

A cos (2 ft + ) = Re Ae e (2.19)

A sin (2 ft + ) = Im Ae e (2.20)

Using the complex plane, we can envision the complex exponential's temporal variations as seen in the
above gure (Figure 2.2). The magnitude of the complex exponential is A , and the initial value of the
complex exponential at t = 0 has an angle of . As time increases, the locus of points traced by the complex
exponential is a circle (it has constant magnitude of A ). The number of times per second we go around the
circle equals the frequency f . The time taken for the complex exponential to go around the circle once is
known as its period T , and equals
. The projections onto the real and imaginary axes of the rotating

vector representing the complex exponential signal are the cosine and sine signal of Euler's relation (2.16).
http://www.edisontechcenter.org/CharlesProteusSteinmetz.html
