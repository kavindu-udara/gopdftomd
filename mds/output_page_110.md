104 CHAPTER 4. FREQUENCY DOMAIN

the term involving a .

Z

s ( t ) cos
2 lt
dt =

Z

a cos
2 lt
dt +
X
a

Z

cos
2 kt
cos
2 lt
dt

+
X
b

Z

sin
2 kt
cos
2 lt
dt

(4.14)

The rst and third terms are zero; in the second, the only non-zero term in the sum results when the indices

k and l are equal (but not zero), in which case we obtain
a T
. If k = 0 = l , we obtain a T . Consequently,

a =
2

Z

s ( t ) cos
2 lt
dt ; l 6 = 0

All of the Fourier coe cients can be found similarly.

a =
1

Z

s ( t ) dt

a =
2

Z

s ( t ) cos
2 kt
dt ; k 6 = 0

b =
2

Z

s ( t ) sin
2 kt
dt

(4.15)

Exercise 4.4 (Solution on p. 139.)
The expression for a is referred to as the average value of s ( t ) . Why?

Exercise 4.5 (Solution on p. 139.)
What is the Fourier series for a unit-amplitude square wave?

Example 4.2
Let's nd the Fourier series representation for the half-wave recti ed sinusoid shown in Figure 4.5
(page 108).

s ( t ) =

8
>
<

>
:

sin
2 t
0 t <
T

0
T
t < T

(4.16)

We have a choice: we can nd the complex spectral coe cients c or we can nd the classic
series coe cients. Given one result, we can nd the other. In general, nding the complex-valued
coe cients is easier, but in this case let's nd the classic series coe cients because of the presence
of the sinusoid in the signal's expression. Begin with the sine terms in the series; to nd b we must
calculate the integral

b =
2

Z
sin
2 t
sin
2 kt
dt (4.17)

Using our trigonometric identities turns our integral of a product of sinusoids into a sum of integrals
of individual sinusoids, which are much easier to evaluate.

Z
sin
2 t
sin
2 kt
dt =
1

Z
cos
2 ( k 1) t
cos
2 ( k + 1) t
dt

=

(
k = 1

0 otherwise

(4.18)
