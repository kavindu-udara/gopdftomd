7

Figure 1.5
the amplitude at midnight. We could relate temperature to amplitude by the formula A = A (1 + kT ) ,
where A and k are constants that the transmitter and receiver must both know.
If we had two numbers we wanted to send at the same time, we could modulate the sinusoid's frequency
as well as its amplitude. This modulation scheme assumes we can estimate the sinusoid's amplitude and
frequency; we shall learn that this is indeed possible.
Now suppose we have a sequence of parameters to send. We have exploited all of the sinusoid's two
parameters. What we can do is modulate them for a limited time (say T seconds), and send two parameters
every T . This simple notion corresponds to how a modem works. Here, typed characters are encoded into
eight bits, and the individual bits are encoded into a sinusoid's amplitude and frequency. We'll learn how
this is done in subsequent modules, and more importantly, we'll learn what the limits are on such digital
communication schemes.

### 1.5 Introduction Problems

Problem 1.1: RMS Values
The rms (root-mean-square) value of a periodic signal is de ned to be

rms [ s ] =

s

Z

s ( t ) dt

where T is de ned to be the signal's period : the smallest positive number such that s ( t ) = s ( t + T ) .

(a) What is the period of s ( t ) = A sin (2 f t + ) ?
(b) What is the rms value of this signal? How is it related to the peak value?
(c) What is the period and rms value of the depicted (Figure 1.5) square wave , generically denoted by
sq ( t ) ?
(d) By inspecting any device you plug into a wall socket, you'll see that it is labeled 110 volts AC. What
is the expression for the voltage provided by a wall socket? What is its rms value?

Problem 1.2: Modems
The word modem is short for modulator-demodulator. Modems are used not only for connecting com-
puters to telephone lines, but also for connecting digital (discrete-valued) sources to generic channels. In this
problem, we explore a simple kind of modem, in which binary information is represented by the presence or
absence of a sinusoid (presence representing a 1 and absence a 0 ). Consequently, the modem's transmitted
signal that represents a single bit has the form

x ( t ) = A sin (2 f t ) , 0 t T

Within each bit interval T , the amplitude is either A or zero.
This content is available online at http://cnx.org/content/m10353/2.17/ .
