12 CHAPTER 2. SIGNALS AND SYSTEMS

Figure 2.1: A complex number is an ordered pair ( a , b ) that can be regarded as coordinates in the plane.
Complex numbers can also be expressed in polar coordinates as r \ .
z = Re [ z ] + j Im [ z ]

z = Re [ z ] j Im [ z ]
(2.1)

Using Cartesian notation, the following properties easily follow.

If we add two complex numbers, the real part of the result equals the sum of the real parts and the
imaginary part equals the sum of the imaginary parts. This property follows from the laws of vector
addition.
a + jb + a + jb = a + a + j ( b + b )

In this way, the real and imaginary parts remain separate.
The product of j and a real number is an imaginary number: ja . The product of j and an imaginary
number is a real number: j ( jb ) = b because j = 1 . Consequently, multiplying a complex number
by j rotates the number's position by 90 degrees.

Exercise 2.1 (Solution on p. 30.)
Use the de nition of addition to show that the real and imaginary parts can be expressed as a
sum/di erence of a complex number and its conjugate. Re [ z ] =
and Im [ z ] =
.

Complex numbers can also be expressed in an alternate form, polar form , which we will nd quite useful.
Polar form arises arises from the geometric interpretation of complex numbers. The Cartesian form of a
complex number can be re-written as

a + jb =
p + b
a + b
+ j
b + b

By forming a right triangle having sides a and b , we see that the real and imaginary parts correspond to the
cosine and sine of the triangle's base angle. We thus obtain the polar form for complex numbers.

z = a + jb = r \

r = j z j =
p + b

= arctan ( b=a )

a = r cos

b = r sin

The quantity r is known as the magnitude of the complex number z , and is frequently written as j z j . The
quantity is the complex number's angle . In using the arc-tangent formula to nd the angle, we must take
into account the quadrant in which the complex number lies.

Exercise 2.2 (Solution on p. 30.)
Convert 3 2 j to polar form.

![Im7](../output/images/test_018_Im7.png)
