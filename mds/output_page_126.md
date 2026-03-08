120 CHAPTER 4. FREQUENCY DOMAIN

The middle term in the expression for Y ( f ) consists of the di erence of two terms: the constant
1 and the complex exponential e . Because of the Fourier transform's linearity, we simply
subtract the results.

Y ( f ) ! 1 e u( t ) 1 e u( t ) (4.44)

Note that in delaying the signal how we carefully included the unit step. The second term in this
result does not begin until t = . Thus, the waveforms shown in the Filtering Periodic Signals
example (Figure 4.10) mentioned above are exponentials. We say that the time constant of an
exponentially decaying signal equals the time it takes to decrease by 1 =e of its original value. Thus,
the time-constant of the rising and falling portions of the output equal the product of the circuit's
resistance and capacitance.

Exercise 4.16 (Solution on p. 140.)
Derive the lter's output by considering the terms in (4.41) in the order given. Integrate last rather
than rst. You should get the same answer.

In this example, we used the table extensively to nd the inverse Fourier transform, relying mostly on

what multiplication by certain factors, like
1 2 f
and e , meant. We essentially treated multiplication

by these factors as if they were transfer functions of some ctitious circuit. The transfer function
1 2 f
corresponded to a circuit that integrated, and e to one that delayed. We even implicitly interpreted
the circuit's transfer function as the input's spectrum! This approach to nding inverse transforms breaking
down a complicated expression into products and sums of simple components is the engineer's way of
breaking down the problem into several subproblems that are much easier to solve and then gluing the results
together. Along the way we may make the system serve as the input, but in the rule Y ( f ) = X ( f ) H ( f ) ,
which term is the input and which is the transfer function is merely a notational matter (we labeled one
factor with an X and the other with an H ).

4.9.1 Transfer Functions

The notion of a transfer function applies well beyond linear circuits. Although we don't have all we need
to demonstrate the result as yet, all linear, time-invariant systems have a frequency-domain input-output
relation given by the product of the input's Fourier transform and the system's transfer function. Thus,
linear circuits are a special case of linear, time-invariant systems. As we tackle more sophisticated problems
in transmitting, manipulating, and receiving information, we will assume linear systems having certain
properties (transfer functions) without worrying about what circuit has the desired property. At this point,
you may be concerned that this approach is glib, and rightly so. Later we'll show that by involving software
that we really don't need to be concerned about constructing a transfer function from circuit elements and
op-amps.

4.9.2 Commutative Transfer Functions

Another interesting notion arises from the commutative property of multiplication (exploited in an example
above (Example 4.6)): We can arbitrarily choose an order in which to apply each product. Consider a
cascade of two linear, time-invariant systems. Because the Fourier transform of the rst system's output is
X ( f ) H ( f ) and it serves as the second system's input, the cascade's output spectrum is X ( f ) H ( f ) H ( f ) .
Because this product also equals X ( f ) H ( f ) H ( f ) , the cascade having the linear systems in the
opposite order yields the same result . Furthermore, the cascade acts like a single linear system,
having transfer function H ( f ) H ( f ) . This result applies to other con gurations of linear, time-invariant
systems as well; see this Frequency Domain Problem (Problem 4.13). Engineers exploit this property by
determining what transfer function they want, then breaking it down into components arranged according to
standard con gurations. Using the fact that op-amp circuits can be connected in cascade with the transfer
function equaling the product of its component's transfer function (see this analog signal processing problem
