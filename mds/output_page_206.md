200 CHAPTER 6. INFORMATION COMMUNICATION

Expanding the expressions, we nd that

a b + j 2 ab =
e
G
e
R 4 f
e
L
e
C + j 2 f (
e
R
e
C +
e
G
e
L )

Considering the high-frequency limit, the constant term on the right side can be ignored. Setting b =

2 f
p
L
e
C and equating j 2 ab with j 2 f (
e
R
e
C +
e
G
e
L ) , we nd that the real part indeed equals a non-zero
constant.

a =
1

e
R
+
e
GZ

!

Here, Z is de ned to be

q
L=
e
C and its importance is demonstrated below. The exponential attenuation of

high-quality coaxial cable predicted by this result is typically 1 dB (or less) per kilometer.

Exercise 6.1 (Solution on p. 255.)
Find the propagation speed in terms of physical parameters for both the coaxial cable and twisted
pair examples.

By using the second of the transmission line equation (6.6), we can solve for the current's complex
amplitude. Considering the spatial region x > 0 , for example, we nd that

d
V ( x ) = V ( x ) =
e
R + j 2 f
e
L I ( x )

which means that the ratio of voltage and current complex amplitudes does not depend on distance.

V ( x ) ( x )
=

s
R + j 2 f
e
L
G + j 2 f
e
C

= Z

(6.15)

The quantity Z is known as the transmission line's characteristic impedance . Note that when the signal
frequency is su ciently high, the characteristic impedance is real, which means the transmission line appears
resistive in this high-frequency regime.

lim Z =

s
L
C
(6.16)

Typical values for characteristic impedance are 50 and 75 .
A related transmission line is the optic ber. Here, the electromagnetic eld is light, and it propagates
down a cylinder of glass. In this situation, we don't have two conductors in fact we have none and the
energy is propagating in what corresponds to the dielectric material of the coaxial cable. Optic ber com-
munication has exactly the same properties as other transmission lines: Signal strength decays exponentially
according to the ber's space constant and propagates at some speed less than light would in free space.
From the encompassing view of Maxwell's equations, the only di erence is the electromagnetic signal's fre-
quency. Because no electric conductors are present and the ber is protected by an opaque insulator, optic
ber transmission is interference-free.

Exercise 6.2 (Solution on p. 255.)
From tables of physical constants, nd the frequency of a sinusoid in the middle of the visible light
range. Compare this frequency with that of a mid-frequency cable television signal.

To summarize, we use transmission lines for high-frequency wireline signal communication. In wireline
communication, we have a direct, physical connection a circuit between transmitter and receiver. When
we select the transmission line characteristics and the transmission frequency so that we operate in the high-
frequency regime, signals are not ltered as they propagate along the transmission line: The characteristic
impedance is real-valued the transmission line's equivalent impedance is a resistor and all the signal's
