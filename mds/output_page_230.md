224 CHAPTER 6. INFORMATION COMMUNICATION

Figure 6.20: To correct errors that occur in the digital channel, a channel coder and decoder are added
to the communication system. Properly designed channel coding can greatly reduce the probability (from
the uncoded value of p ) that a data bit b ( n ) is received incorrectly even when the probability of c ( l )
be received in error remains p or becomes larger. This system forms the Fundamental Model of Digital
Communication.

Figure 6.21: The upper portion depicts the result of directly modulating the bit stream b ( n ) into a
transmitted signal x ( t ) using a baseband BPSK signal set. R is the datarate produced by the source
coder. If that bit stream passes through a (3,1) channel coder to yield the bit stream c ( l ) , the resulting
transmitted signal requires a bit interval T three times smaller than the uncoded version. This reduction
in the bit interval means that the transmitted energy/bit decreases by a factor of three, which results in
an increased error probability in the receiver.
that should not prevent us from studying commonly used error correcting codes that not only nd their way
into all digital communication systems, but also into CDs and bar codes used on merchandise.

### 6.25 Repetition Codes

Perhaps the simplest error correcting code is the repetition code . Here, the transmitter sends the data
bit several times, an odd number of times in fact. Because the error probability p is always less than
, we
know that more of the bits should be correct rather than in error. Simple majority voting of the received
bits (hence the reason for the odd number) determines the transmitted bit more accurately than sending
it alone. For example, let's consider the three-fold repetition code: for every bit b ( n ) emerging from the
source coder, the channel coder produces three. Thus, the bit stream emerging from the channel coder c ( l )
has a data rate three times higher than that of the original bit stream b ( n ) . The coding table (Table 1.1)
illustrates when the majority-vote decoder can correct errors and when it can't.
Thus, if one bit of the three bits is received in error, the receiver can correct the error; if more than
one error occurs, the channel decoder announces the bit is 1 instead of transmitted value of 0 . Using this
repetition code, the probability of
b
b ( n ) 6 = 0 equals 3 p (1 p ) + p . This probability of a decoding error
is always less than p , the uncoded value, so long as p <
.
This content is available online at http://cnx.org/content/m0071/2.22/ .
