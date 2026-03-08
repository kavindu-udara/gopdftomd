# Chapter 2

# Signals and Systems

### 2.1 Complex Numbers

While the fundamental signal used in electrical engineering is the sinusoid, it can be expressed mathematically
in terms of an even more fundamental signal: the complex exponential . Representing sinusoids in terms of
complex exponentials is not a mathematical oddity. Fluency with complex numbers and rational functions
of complex variables is a critical skill all engineers master. Understanding information and power system
designs and developing new systems all hinge on using complex numbers. In short, they are critical to
modern electrical engineering, a realization made over a century ago.

2.1.1 De nitions

The notion of the square root of 1 originated with the quadratic formula: the solution of certain quadratic
equations mathematically exists only if the so-called imaginary quantity
p 1 could be de ned. Euler rst
used i for the imaginary unit but that notation did not take hold until roughly Ampère's time. Ampère
used the symbol i to denote current (intensité de current). It wasn't until the twentieth century that the
importance of complex numbers to circuit theory became evident. By then, using i for current was entrenched
and electrical engineers chose j for writing complex numbers.
An imaginary number has the form jb =
p b . A complex number , z , consists of the ordered pair
( a , b ), a is the real component and b is the imaginary component (the j is suppressed because the imaginary
component of the pair is always in the second position). The imaginary number jb equals ( 0 , b ). Note that
a and b are real-valued numbers.
Figure 2.1 shows that we can locate a complex number in what we call the complex plane . Here, a ,
the real part, is the x -coordinate and b , the imaginary part, is the y -coordinate. From analytic geometry, we
know that locations in the plane can be expressed as the sum of vectors, with the vectors corresponding to
the x and y directions. Consequently, a complex number z can be expressed as the (vector) sum z = a + jb
where j indicates the y -coordinate. This representation is known as the Cartesian form of z . An imaginary
number can't be numerically added to a real number; rather, this notation for a complex number represents
vector addition, but it provides a convenient notation when we perform arithmetic manipulations.
Some obvious terminology. The real part of the complex number z = a + jb , written as Re [ z ] , equals
a . We consider the real part as a function that works by selecting that component of a complex number
not multiplied by j . The imaginary part of z , Im [ z ] , equals b : that part of a complex number that is
multiplied by j . Again, both the real and imaginary parts of a complex number are real-valued.
The complex conjugate of z , written as z , has the same real part as z but an imaginary part of the
This content is available online at http://cnx.org/content/m0081/2.27/ .
http://www-groups.dcs.st-and.ac.uk/~history/Biographies/Euler.html
http://www-groups.dcs.st-and.ac.uk/~history/Biographies/Ampere.html

11
