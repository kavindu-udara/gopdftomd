23

The gain can be positive or negative (if negative, we would say that the ampli er inverts its input) and
can be greater than one or less than one. If less than one, the ampli er actually attenuates . A real-world
example of an ampli er is your home stereo. You control the gain by turning the volume control.

2.6.3 Delay

A system serves as a time delay (Figure 2.14) when the output signal equals the input signal at an earlier
time.
y ( t ) = x ( t ) (2.31)

Figure 2.14: A delay.

Here, is the delay. The way to understand this system is to focus on the time origin: The output at time
t = equals the input at time t = 0 . Thus, if the delay is positive, the output emerges later than the input,
and plotting the output amounts to shifting the input plot to the right. The delay can be negative, in which
case we say the system advances its input. Such systems are di cult to build (they would have to produce
signal values derived from what the input will be ), but we will have occasion to advance signals in time.

2.6.4 Time Reversal

With a time-reversal system, the output signal equals the input signal ipped about the vertical axis (the
time origin).
y ( t ) = x ( t ) (2.32)

Figure 2.15: A time reversal system.

Again, such systems are di cult to build, but the notion of time reversal occurs frequently in communications
systems.
Exercise 2.5 (Solution on p. 30.)
Mentioned earlier was the issue of whether the ordering of systems mattered. In other words, if
we have two systems in cascade, does the output depend on which comes rst? Determine if the
ordering matters for the cascade of an ampli er and a delay and for the cascade of a time-reversal
system and a delay.

2.6.5 Derivative Systems and Integrators

Systems that perform calculus-like operations on their inputs can produce waveforms signi cantly di erent
than present in the input. Derivative systems operate in a straightforward way: A rst-derivative system
would have the input-output relationship y ( t ) =
x ( t ) . Integral systems have the complication that the
