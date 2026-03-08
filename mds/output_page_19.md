13

2.1.2 Euler's Formula

Surprisingly, the polar form of a complex number z can be expressed mathematically as

z = re (2.2)

To show this result, we use Euler's relations that express exponentials with imaginary arguments in terms
of trigonometric functions.
e = cos + j sin (2.3)

cos =
e + e
sin =
e e
j
(2.4)

The rst of these is easily derived from the Taylor's series for the exponential.

e = 1 +
x
+
x
+
x
+ : : :

Substituting j for x , we nd that

e = 1 + j

j

+ : : :

because j = 1 , j = j , and j = 1 . Grouping separately the real-valued terms and the imaginary-valued
ones,

e = 1

+ + j

+ : : :

The real-valued terms correspond to the Taylor's series for cos ( ) , the imaginary ones to sin ( ) , and Euler's
rst relation results. The remaining relations are easily derived from the rst. Because of the relationship
r =
p + b , we see that multiplying the exponential in (2.3) by a real constant corresponds to setting the
radius of the complex number by the constant.

2.1.3 Calculating with Complex Numbers

Adding and subtracting complex numbers expressed in Cartesian form is quite easy: You add (subtract) the
real parts and imaginary parts separately.

( z z ) = ( a a ) + j ( b b ) (2.5)

To multiply two complex numbers in Cartesian form is not quite as easy, but follows directly from following
the usual rules of arithmetic.
z z = ( a + jb ) ( a + jb )

= a a b b + j ( a b + a b )
(2.6)

Note that we are, in a sense, multiplying two vectors to obtain another vector. Complex arithmetic provides
a unique way of de ning vector multiplication.

Exercise 2.3 (Solution on p. 30.)
What is the product of a complex number and its conjugate?

Division requires mathematical manipulation. We convert the division problem into a multiplication problem
by multiplying both the numerator and denominator by the conjugate of the denominator.

z
=
a + jb + jb

=
a + jb + jb

a jb jb

=
( a + jb ) ( a jb ) + b

=
a a + b b + j ( a b a b ) + b

(2.7)
