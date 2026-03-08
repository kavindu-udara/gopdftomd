199

distance along the transmission line. Expressing in terms of its real and imaginary parts in our solution
shows that such increases are a (mathematical) possibility. V ( x ) = V e + V e The voltage
cannot increase without limit; because a ( f ) is always positive, we must segregate the solution for negative
and positive x . The rst term will increase exponentially for x < 0 unless V = 0 in this region; a similar
result applies to V for x > 0 . These physical constraints give us a cleaner solution.

V ( x ) =

(
V e ; if x > 0

V e ; if x < 0
(6.11)

This solution suggests that voltages (and currents too) will decrease exponentially along a transmission
line. The space constant , also known as the attenuation constant , is the distance over which the voltage
decreases by a factor of
. It equals the reciprocal of a ( f ) , which depends on frequency, and is expressed by
manufacturers in units of dB/m.
The presence of the imaginary part of , b ( f ) , also provides insight into how transmission lines work.
Because the solution for x > 0 is proportional to e , we know that the voltage's complex amplitude will
vary sinusoidally in space . The complete solution for the voltage has the form

v ( x; t ) = Re
h
V e e
i
(6.12)

The complex exponential portion has the form of a propagating wave . If we could take a snapshot of the
voltage (take its picture at t = t ), we would see a sinusoidally varying waveform along the transmission
line. One period of this variation, known as the wavelength , equals = 2 =b . If we were to take a second
picture at some later time t = t , we would also see a sinusoidal voltage. Because

2 ft bx = 2 f ( t + t t ) bx = 2 ft b x
2 f
( t t )

the second waveform appears to be the rst one, but delayed shifted to the right in space. Thus, the
voltage appeared to move to the right with a speed equal to 2 f=b (assuming b > 0 ). We denote this
propagation speed by c , and it equals

c =
2 f
=
2 f

r
e
G + j 2 f
e
C
e
R + j 2 f
e
L

(6.13)

The characteristics of the voltage signal shown in equation (6.12) depend on the values of a and b , and

how they depend on frequency. The simplest results occur in the high-frequency region where j 2 f
e
L
e
R

and j 2 f
e
C
e
G . In this case, simpli es to

q 4 f
e
L
e
C , which seemingly makes it pure imaginary with

a = 0 and b = 2 f
p
L
e
C . Using this result, we nd the propagation speed to be

lim c =
1
L
e
C
(6.14)

For typical coaxial cable, this propagation speed is a fraction (one-third to two-thirds) of the speed of light.
While this high-frequency analysis shows that the dominant high-frequency component of is its imagi-
nary part, there could be (and is!) a smaller real part. Since the real part of is the attenuation factor a , a
more detailed analysis is required to determine if a = 0 (no attenuation) or is non-zero. One way of pursuing
a more detailed analysis is to exploit equation (6.10) by

= ( a + jb ) =
e
G + j 2 f
e
C
e
R + j 2 f
e
L
