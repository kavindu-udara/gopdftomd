98 CHAPTER 4. FREQUENCY DOMAIN

as sum of harmonically related sine waves: sinusoids having frequencies that are integer multiples of
the fundamental frequency . Because the signal has period T , the fundamental frequency is 1 =T . The
Fourier series expresses the signal as a superposition of complex exponentials having frequencies k=T , k =
f : : :; 1 ; 0 ; 1 ; : : : g .

s ( t ) =
X
c e
(4.1)

The zeroth coe cient equals the signal's average value and is real-valued for real-valued signals: c =
1

R
s ( t ) dt . In general, the other Fourier coe cients are complex-valued; we will nd many special cases

wherein come coe cients are real-valued. The family of functions e
are called basis functions and

form the foundation of the Fourier series. No matter what the periodic signal might be, these functions are
always present and form the representation's building blocks. They depend on the signal period T , and are
indexed by k .

Key point: Assuming we know the period, knowing the Fourier coe cients is equivalent to
knowing the signal. Thus, it makes no di erence if we have a time-domain or a frequency-domain
characterization of the signal.

Exercise 4.1 (Solution on p. 139.)
What is a sinusoid's Fourier series? You can nd this expression by inspection!

To nd the Fourier coe cients, we note the orthogonality property

Z

e
e
dt =

(
T; if k = l

0 ; if k 6 = l
= T ( k l ) (4.2)

( n ) is the unit-sample de ned in equation (2.28). Using this de nition, ( k l ) equals one when k l = 0
( k = l ) and zero when k l 6 = 0 ( k 6 = l ). Assuming for the moment that the Fourier series works, we can
nd a signal's Fourier coe cients, its spectrum , by exploiting this orthogonality property. Simply multiply
each side of (4.1) by e and integrate over the interval [0 ; T ] .

c =
1

Z

s ( t ) e
dt (4.3)

Example 4.1
Finding the Fourier series coe cients for the square wave sq ( t ) is very simple. Mathematically,
this signal can be expressed over one period as

sq ( t ) =

(
1 ; 0 < t <
1 ;
< t < T

The expression for the Fourier coe cients has the form

c =
1

Z
e
dt
1

Z
e
dt (4.4)

note: When integrating an expression containing j , treat it just like any other constant.

The two integrals are very similar, one equaling the negative of the other. The nal expression
becomes

c =
2 2 k
( 1) 1

=

(
k odd

0 k even

(4.5)
