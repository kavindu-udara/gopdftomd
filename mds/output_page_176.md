170 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Figure 5.18: The input to the simple example system, a unit sample, is shown at the top, with the
outputs for several system parameter values shown below.

Figure 5.19: The plot shows the unit-sample response of a length-5 boxcar lter.
population becomes extinct; if greater than one, the population ourishes. The same di erence
equation also describes the e ect of compound interest on deposits. Here, n indexes the times at
which compounding occurs (daily, monthly, etc.), a equals the compound interest rate plus one,
and b = 1 (the bank provides no gain). In signal processing applications, we typically require that
the output remain bounded for any input. For our example, that means that we restrict j a j < 1
and choose values for it and the gain according to the application.

Exercise 5.26 (Solution on p. 193.)
Note that the di erence equation (5.42),

y ( n ) = a y ( n 1) + + a y ( n p ) + b x ( n ) + b x ( n 1) + + b x ( n q )

does not involve terms like y ( n + 1) or x ( n + 1) on the equation's right side. Can such terms also
be included? Why or why not?

Example 5.4
A somewhat di erent system has no a coe cients. Consider the di erence equation

y ( n ) =
1
x ( n ) + + x ( n q + 1) (5.46)

Because this system's output depends only on current and previous input values, we need not

be concerned with initial conditions. When the input is a unit-sample, the output equals
1
for

n = f 0 ; : : : ; q 1 g , then equals zero thereafter. Such systems are said to be FIR ( F inite I mpulse
