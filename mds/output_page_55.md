49

same frequency, we have
X
v =
X
V e = 0 (3.12)

which means
X
V = 0 (3.13)

the complex amplitudes of the voltages obey KVL . We can easily imagine that the complex amplitudes
of the currents obey KCL.
What we have discovered is that source(s) equaling a complex exponential of the same frequency forces
all circuit variables to be complex exponentials of the same frequency. Consequently, the ratio of voltage to
current for each element equals the ratio of their complex amplitudes, which depends only on the source's
frequency and element values.
This situation occurs because the circuit elements are linear and time-invariant. For example, suppose we
had a circuit element where the voltage equaled the square of the current: v ( t ) = Ki ( t ) . If i ( t ) = Ie ,
v ( t ) = KI e , meaning that voltage and current no longer had the same frequency and that their ratio
was time-dependent.
Because for linear circuit elements the complex amplitude of voltage is proportional to the complex
amplitude of current V = ZI assuming complex exponential sources means circuit elements behave
as if they were resistors, where instead of resistance, we use impedance. Because complex amplitudes
for voltage and current also obey Kirchho 's laws, we can solve circuits using voltage and
current divider and the series and parallel combination rules by considering the elements to
be impedances.

### 3.10 Time and Frequency Domains

When we nd the di erential equation relating the source and the output, we are faced with solving the
circuit in what is known as the time domain . What we emphasize here is that it is often easier to nd
the output if we use impedances. Because impedances depend only on frequency, we nd ourselves in the
frequency domain . A common error in using impedances is keeping the time-dependent part, the complex
exponential, in the fray. The entire point of using impedances is to get rid of time and concentrate on
frequency. Only after we nd the result in the frequency domain do we go back to the time domain and put
things back together again.
To illustrate how the time domain, the frequency domain and impedances t together, consider the time
domain and frequency domain to be two work rooms. Since you can't be two places at the same time, you
are faced with solving your circuit problem in one of the two rooms at any point in time. Impedances and
complex exponentials are the way you get between the two rooms. Security guards make sure you don't try
to sneak time domain variables into the frequency domain room and vice versa. Figure 3.23 shows how this
works.
This content is available online at http://cnx.org/content/m10708/2.9/ .
