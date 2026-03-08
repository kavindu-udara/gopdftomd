# Chapter 5

# Digital Signal Processing

### 5.1 Introduction to Digital Signal Processing

Not only do we have analog signals signals that are real- or complex-valued functions of a continuous
variable such as time or space we can de ne digital ones as well. Digital signals are sequences , functions
de ned only for the integers. We thus use the notation s ( n ) to denote a discrete-time one-dimensional signal
such as a digital music recording and s ( m; n ) for a discrete- time two-dimensional signal like a photo taken
with a digital camera. Sequences are fundamentally di erent than continuous-time signals. For example,
continuity has no meaning for sequences.
Despite such fundamental di erences, the theory underlying digital signal processing mirrors that for ana-
log signals: Fourier transforms, linear ltering, and linear systems parallel what previous chapters described.
These similarities make it easy to understand the de nitions and why we need them, but the similarities
should not be construed as analog wannabes. We will discover that digital signal processing is not an
approximation to analog processing. We must explicitly worry about the delity of converting analog signals
into digital ones. The music stored on CDs, the speech sent over digital cellular telephones, and the video
carried by digital television all evidence that analog signals can be accurately converted to digital ones and
back again.
The key reason why digital signal processing systems have a technological advantage today is the com-
puter : computations, like the Fourier transform, can be performed quickly enough to be calculated as the
signal is produced, and programmability means that the signal processing system can be easily changed.
This exibility has obvious appeal, and has been widely accepted in the marketplace. Programmability
means that we can perform signal processing operations impossible with analog systems (circuits). We will
also discover that digital systems enjoy an algorithmic advantage that contributes to rapid processing
speeds: Computations can be restructured in non-obvious ways to speed the processing. This exibility
comes at a price, a consequence of how computers work. How do computers perform signal processing?

### 5.2 Introduction to Computer Organization

5.2.1 Computer Architecture

To understand digital signal processing systems, we must understand a little about how computers compute.
The modern de nition of a computer is an electronic device that performs calculations on data, presenting
the results to humans or other computers in a variety of (hopefully useful) ways.
This content is available online at http://cnx.org/content/m10781/2.3/ .
Taking a systems viewpoint for the moment, a system that produces its output as rapidly as the input arises is said to
be a real-time system. All analog systems operate in real time; digital ones that depend on a computer to perform system
computations may or may not work in real time. Clearly, we need real-time signal processing systems. Only recently have
computers become fast enough to meet real-time requirements while performing non-trivial signal processing.
This content is available online at http://cnx.org/content/m10263/2.28/ .

143
