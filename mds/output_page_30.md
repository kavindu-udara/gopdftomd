24 CHAPTER 2. SIGNALS AND SYSTEMS

integral's limits must be de ned. It is a signal theory convention that the elementary integral operation have
a lower limit of 1 , and that the value of all signals at t = 1 equals zero. A simple integrator would
have input-output relation

y ( t ) =

Z

x ( ) d (2.33)

2.6.6 Linear Systems

Linear systems are a class of systems rather than having a speci c input-output relation. Linear systems
form the foundation of system theory, and are the most important class of systems in communications. They
have the property that when the input is expressed as a weighted sum of component signals, the output
equals the same weighted sum of the outputs produced by each component. When S [ ] is linear,

S [ G x ( t ) + G x ( t )] = G S [ x ( t )] + G S [ x ( t )] (2.34)

for all choices of signals and gains.
This general input-output relation property can be manipulated to indicate speci c properties shared by
all linear systems.

S [ Gx ( t )] = G S [ x ( t )] The colloquialism summarizing this property is Double the input, you double
the output. Note that this property is consistent with alternate ways of expressing gain changes:
Since 2 x ( t ) also equals x ( t ) + x ( t ) , the linear system de nition provides the same output no matter
which of these is used to express a given signal.
S [0] = 0 If the input is identically zero for all time , the output of a linear system must be zero.
This property follows from the simple derivation S [0] = S [ x ( t ) x ( t )] = S [ x ( t )] S [ x ( t )] = 0 .

Just why linear systems are so important is related not only to their properties, which are divulged throughout
this course, but also because they lend themselves to relatively simple mathematical analysis. Said another
way, They're the only systems we thoroughly understand!
We can nd the output of any linear system to a complicated input by decomposing the input into simple
signals. The equation above (2.34) says that when a system is linear, its output to a decomposed input is
the sum of outputs to each input. For example, if

x ( t ) = e + sin (2 f t )

the output S ( x ( t )) of any linear system equals

y ( t ) = S e + S [sin (2 f t )]

2.6.7 Time-Invariant Systems

Systems that don't change their input-output relation with time are said to be time-invariant. The mathe-
matical way of stating this property is to use the signal delay concept described in Section 2.6.3.

y ( t ) = S [ x ( t )] = ) y ( t ) = S [ x ( t )] (2.35)

If you delay (or advance) the input, the output is similarly delayed (advanced). Thus, a time-invariant
system responds to an input you may supply tomorrow the same way it responds to the same input applied
today; today's output is merely delayed to occur tomorrow.
The collection of linear, time-invariant systems are the most thoroughly understood systems. Much of
the signal processing and system theory discussed here concentrates on such systems. For example, electric
circuits are, for the most part, linear and time-invariant. Nonlinear ones abound, but characterizing them
so that you can predict their behavior for any input remains an unsolved problem.
