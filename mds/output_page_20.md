14 CHAPTER 2. SIGNALS AND SYSTEMS

Because the nal result is so complicated, it's best to remember how to perform division multiplying
numerator and denominator by the complex conjugate of the denominator than trying to remember the
nal result.
The properties of the exponential make calculating the product and ratio of two complex numbers much
simpler when the numbers are expressed in polar form.

z z = r e r e = r r e

z
=
r e
e
=
r
e
(2.8)

To multiply, the radius equals the product of the radii and the angle the sum of the angles. To divide,
the radius equals the ratio of the radii and the angle the di erence of the angles. When the original
complex numbers are in Cartesian form, it's usually worth translating into polar form, then performing the
multiplication or division (especially in the case of the latter). Addition and subtraction of polar forms
amounts to converting to Cartesian form, performing the arithmetic operation, and converting back to polar
form.
Example 2.1
When we solve circuit problems, the crucial quantity, known as a transfer function, will always be
expressed as the ratio of polynomials in the variable s = j 2 f . What we'll need to understand the
circuit's e ect is the transfer function in polar form. For instance, suppose the transfer function
equals
s + 2 + s + 1
(2.9)

s = j 2 f (2.10)

Performing the required division is most easily accomplished by rst expressing the numerator and
denominator each in polar form, then calculating the ratio. Thus,

s + 2 + s + 1
=
j 2 f + 2 4 f + j 2 f + 1
(2.11)

=

p f e
4 f ) + 4 f e

(2.12)

=

s f
4 f + 16 f
e
( ( ( )))
(2.13)

### 2.2 Elemental Signals

Elemental signals are the building blocks with which we build complicated signals. By de nition,
elemental signals have a simple structure. Exactly what we mean by the structure of a signal will unfold in
this section of the course. Signals are nothing more than functions de ned with respect to some independent
variable, which we take to be time for the most part. Very interesting signals are not functions solely of
time; one great example of which is an image. For it, the independent variables are x and y (two-dimensional
space). Video signals are functions of three variables: two spatial dimensions and time. Fortunately, most
of the ideas underlying modern signal theory can be exempli ed with one-dimensional signals.
This content is available online at http://cnx.org/content/m0004/2.29/ .
