114 CHAPTER 4. FREQUENCY DOMAIN

calculations entirely in the frequency domain. Using Fourier series, we can calculate how any linear circuit
will respond to a periodic input.

### 4.8 Derivation of the Fourier Transform

Fourier series clearly open the frequency domain as an interesting and useful way of determining how circuits
and systems respond to periodic input signals. Can we use similar techniques for non-periodic signals? What
is the response of the lter to a single pulse? Addressing these issues requires us to nd the Fourier spectrum
of all signals, both periodic and non-periodic ones. We need a de nition for the Fourier spectrum of a signal,
periodic or not. This spectrum is calculated by what is known as the Fourier transform .
Let s ( t ) be a periodic signal having period T . We want to consider what happens to this signal's
spectrum as we let the period become longer and longer. We denote the spectrum for any assumed value of
the period by c ( T ) . We calculate the spectrum according to the familiar formula

c ( T ) =
1

Z

s ( t ) e
dt (4.29)

where we have used a symmetric placement of the integration interval about the origin for subsequent deriva-
tional convenience. Let f be a xed frequency equaling k=T ; we vary the frequency index k proportionally
as we increase the period. De ne

S ( f ) T c ( T ) =

Z

s ( t ) e dt (4.30)

making the corresponding Fourier series

s ( t ) =
X
S ( f ) e
1
(4.31)

As the period increases, the spectral lines become closer together, becoming a continuum. Therefore,

lim s ( t ) s ( t ) =

Z

S ( f ) e df (4.32)

with

S ( f ) =

Z

s ( t ) e dt (4.33)

S ( f ) is the Fourier transform of s ( t ) (the Fourier transform is symbolically denoted by the uppercase version
of the signal's symbol) and is de ned for any signal for which the integral ((4.33)) converges.

Example 4.4
Let's calculate the Fourier transform of the pulse signal p ( t ) .

P ( f ) =

Z

p ( t ) e dt =

Z

e dt =
1 j 2 f
e 1

= e
sin ( f )

Note how closely this result resembles the expression for Fourier series coe cients of the periodic
pulse signal (4.10).

Figure 4.11 shows how increasing the period does indeed lead to a continuum of coe cients, and that the

Fourier transform does correspond to what the continuum becomes. The quantity
sin ( t )
has a special name,
This content is available online at http://cnx.org/content/m0046/2.21/ .
