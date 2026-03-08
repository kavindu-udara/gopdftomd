198 CHAPTER 6. INFORMATION COMMUNICATION

e
R =
1
e
C =
e
G =

e
L =

r
+ cosh
d r

(6.3)

The voltage between the two conductors and the current owing through them will depend on distance x
along the transmission line as well as time. We express this dependence as v ( x; t ) and i ( x; t ) . When we place
a sinusoidal source at one end of the transmission line, these voltages and currents will also be sinusoidal
because the transmission line model consists of linear circuit elements. As is customary in analyzing linear
circuits, we express voltages and currents as the real part of complex exponential signals, and write circuit
variables as a complex amplitude here dependent on distance times a complex exponential: v ( x; t ) =
Re V ( x ) e and i ( x; t ) = Re I ( x ) e . Using the transmission line circuit model, we nd from
KCL, KVL, and v-i relations the equations governing the complex amplitudes.

KCL at Center Node
I ( x ) = I ( x x ) V ( x )
e
G + j 2 f
e
C x (6.4)

V-I relation for RL series

V ( x ) V ( x + x ) = I ( x )
e
R + j 2 f L x (6.5)

Rearranging and taking the limit x ! 0 yields the so-called transmission line equations .

d
I ( x ) =
e
G + j 2 f
e
C V ( x )

d
V ( x ) =
e
R + j 2 f
e
L I ( x )

(6.6)

By combining these equations, we can obtain a single equation that governs how the voltage's or the current's
complex amplitude changes with position along the transmission line. Taking the derivative of the second
equation and plugging the rst equation into the result yields the equation governing the voltage.

d
V ( x ) =
e
G + j 2 f
e
C
e
R + j 2 f
e
L V ( x ) (6.7)

This equation's solution is
V ( x ) = V e + V e (6.8)

Calculating its second derivative and comparing the result with our equation for the voltage can check this
solution.
d
V ( x ) = V e + V e

= V ( x )

(6.9)

Our solution works so long as the quantity satis es

=

r
e
G + j 2 f
e
C
e
R + j 2 f
e
L

= a ( f ) + jb ( f )

(6.10)

Thus, depends on frequency, and we express it in terms of real and imaginary parts as indicated. The
quantities V and V are constants determined by the source and physical considerations. For example,
let the spatial origin be the middle of the transmission line model Figure 6.2. Because the circuit model
contains simple circuit elements, physically possible solutions for voltage amplitude cannot increase with
