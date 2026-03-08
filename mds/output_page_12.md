6 CHAPTER 1. INTRODUCTION

In the communications model, the source is a system having no input but producing an output; a sink has
an input and no output.
Understanding signal generation and how systems work amounts to understanding signals, the nature
of the information they represent, how information is transformed between analog and digital forms, and
how information can be processed by systems operating on information-bearing signals. This understanding
demands two di erent elds of knowledge. One is electrical science: How are signals represented and ma-
nipulated electrically? The second is signal science: What is the structure of signals, no matter what their
source, what is their information content, and what capabilities does this structure force upon communication
systems?

### 1.4 The Fundamental Signal: The Sinusoid

The most ubiquitous and important signal in electrical engineering is the sinusoid .

Sine De nition
s ( t ) = A cos (2 ft + ) or A cos ( !t + ) (1.1)

A is known as the sinusoid's amplitude , and determines the sinusoid's size. The amplitude conveys the
sinusoid's physical units (volts, lumens, etc). The frequency f has units of Hz (Hertz) or s , and determines
how rapidly the sinusoid oscillates per unit time. The temporal variable t always has units of seconds, and
thus the frequency determines how many oscillations/second the sinusoid has. AM radio stations have carrier
frequencies of about 1 MHz (one mega-hertz or 10 Hz), while FM stations have carrier frequencies of about
100 MHz. Frequency can also be expressed by the symbol ! , which has units of radians/second. Clearly,
! = 2 f . In communications, we most often express frequency in Hertz. Finally, is the phase , and
determines the sine wave's behavior at the origin ( t = 0 ). It has units of radians, but we can express it in
degrees, realizing that in computations we must convert from degrees to radians. Note that if =
, the
sinusoid corresponds to a sine function, having a zero value at the origin.

A sin (2 ft + ) = A cos 2 ft +

(1.2)

Thus, the only di erence between a sine and cosine signal is the phase; we term either a sinusoid.
We can also de ne a discrete-time variant of the sinusoid: A cos (2 fn + ) . Here, the independent
variable is n and represents the integers. Frequency now has no dimensions, and takes on values between 0
and 1.

Exercise 1.1 (Solution on p. 9.)
Show that cos (2 fn ) = cos (2 ( f + 1) n ) , which means that a sinusoid having a frequency larger
than one corresponds to a sinusoid having a frequency less than one.

note: Notice that we shall call either sinusoid an analog signal. Only when the discrete-time
signal takes on a nite set of values can it be considered a digital signal.

Exercise 1.2 (Solution on p. 9.)
Can you think of a simple signal that has a nite number of values but is de ned in continuous
time? Such a signal is also an analog signal.

1.4.2 Communicating Information with Signals

The basic idea of communication engineering is to use a signal's parameters to represent either real numbers or
other signals. The technical term is to modulate the carrier signal's parameters to transmit information
from one place to another. To explore the notion of modulation, we can send a real number (today's
temperature, for example) by changing a sinusoid's amplitude accordingly. If we wanted to send the daily
This content is available online at http://cnx.org/content/m0003/2.15/ .
