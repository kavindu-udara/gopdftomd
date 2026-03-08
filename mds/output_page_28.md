22 CHAPTER 2. SIGNALS AND SYSTEMS

2.5.3 Feedback Interconnection

Figure 2.12: The feedback con guration.

The subtlest interconnection con guration has a system's output also contributing to its input. Engineers
would say the output is fed back to the input through system 2, hence the terminology. The mathematical
statement of the feedback interconnection (Figure 2.12) is that the feed-forward system produces the output:
y ( t ) = S [ e ( t )] . The input e ( t ) equals the input signal minus the output of some other system's output to
y ( t ) : e ( t ) = x ( t ) S [ y ( t )] . Feedback systems are omnipresent in control problems, with the error signal
used to adjust the output to achieve some condition de ned by the input (controlling) signal. For example,
in a car's cruise control system, x ( t ) is a constant representing what speed you want, and y ( t ) is the car's
speed as measured by a speedometer. In this application, system 2 is the identity system (output equals
input).

### 2.6 Simple Systems

Systems manipulate signals, creating output signals derived from their inputs. Why the following are cate-
gorized as simple will only become evident towards the end of the course.

2.6.1 Sources

Sources produce signals without having input. We like to think of these as having controllable parameters,
like amplitude and frequency. Examples would be oscillators that produce periodic signals like sinusoids and
square waves and noise generators that yield signals with erratic waveforms (more about noise subsequently).
Simply writing an expression for the signals they produce speci es sources. A sine wave generator might
be speci ed by y ( t ) = A sin (2 f t ) u( t ) , which says that the source was turned on at t = 0 to produce a
sinusoid of amplitude A and frequency f .

2.6.2 Ampli ers

An ampli er (Figure 2.13) multiplies its input by a constant known as the ampli er gain .

y ( t ) = Gx ( t ) (2.30)

Figure 2.13: An ampli er.
This content is available online at http://cnx.org/content/m0006/2.24/ .
