# Chapter 6

# Information Communication

### 6.1 Information Communication

As far as a communications engineer is concerned, signals express information. Because systems manipulate
signals, they also a ect the information content. Information comes neatly packaged in both analog and
digital forms. Speech, for example, is clearly an analog signal, and computer les consist of a sequence
of bytes, a form of discrete-time signal despite the fact that the index sequences byte position, not time
sample. Communication systems endeavor not to manipulate information, but to transmit it from one
place to another, so-called point-to-point communication , from one place to many others, broadcast
communication , or from many to many, like a telephone conference call or a chat room. Communication
systems can be fundamentally analog, like radio, or digital, like computer networks.
This chapter develops a common theory that underlies how such systems work. We describe and analyze
several such systems, some old like AM radio, some new like computer networks. The question as to which is
better, analog or digital communication, has been answered, because of Claude Shannon's fundamental work
on a theory of information published in 1948, the development of cheap, high-performance computers, and
the creation of high-bandwidth communication systems. The answer is to use a digital communication
strategy . In most cases, you should convert all information-bearing signals into discrete-time, amplitude-
quantized signals. Fundamentally digital signals, like computer les (which are a special case of symbolic
signals), are in the proper form. Because of the Sampling Theorem, we know how to convert analog signals
into digital ones. Shannon showed that once in this form, a properly engineered system can communi-
cate digital information with no error despite the fact that the communication channel thrusts
noise onto all transmissions . This startling result has no counterpart in analog systems; AM radio
will remain noisy. The convergence of these theoretical and engineering results on communications systems
has had important consequences in other arenas. The audio compact disc (CD) and the digital videodisk
(DVD) are now considered digital communications systems, with communication design considerations used
throughout.
Go back to the fundamental model of communication (Figure 1.3). Communications design begins with
two fundamental considerations.

1. What is the nature of the information source, and to what extent can the receiver tolerate errors in
the received information?
2. What are the channel's characteristics and how do they a ect the transmitted signal?

In short, what are we going to send and how are we going to send it? Interestingly, digital as well as
analog transmission are accomplished using analog signals, like voltages in Ethernet (an example of wireline
communications) and electromagnetic radiation ( wireless ) in cellular telephone.
This content is available online at http://cnx.org/content/m0513/2.8/ .

195
